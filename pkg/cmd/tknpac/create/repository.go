package create

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	apipac "github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/v1alpha1"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/cli"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/cli/prompt"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/cli/webhook"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/cmd/tknpac/generate"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/formatting"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/git"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/params"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/params/info"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	noColorFlag = "no-color"
)

type repoOptions struct {
	event        *info.Event
	repository   *apipac.Repository
	run          *params.Run
	gitInfo      *git.Info
	pacNamespace string

	ioStreams *cli.IOStreams
	cliOpts   *cli.PacCliOpts
}

func repositoryCommand(run *params.Run, ioStreams *cli.IOStreams) *cobra.Command {
	var githubURLForWebhook, gitlabURLForWebhook string
	var onlyWebhook, githubWebhook bool
	createOpts := &repoOptions{
		event:      info.NewEvent(),
		repository: &apipac.Repository{},
		run:        run,
	}
	cmd := &cobra.Command{
		Use:     "repository",
		Aliases: []string{"repo"},
		Short:   "Create a repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			ctx := context.Background()
			createOpts.ioStreams = ioStreams
			createOpts.cliOpts = cli.NewCliOptions(cmd)
			createOpts.ioStreams.SetColorEnabled(!createOpts.cliOpts.NoColoring)

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			createOpts.gitInfo = git.GetGitInfo(cwd)
			if err := run.Clients.NewClients(ctx, &run.Info); err != nil {
				return err
			}

			if !onlyWebhook {
				if err := getRepoURL(createOpts); err != nil {
					return err
				}

				if err := getOrCreateNamespace(ctx, createOpts); err != nil {
					return err
				}

				if err := createRepoCRD(ctx, createOpts); err != nil {
					return err
				}

				gopt := generate.MakeOpts()
				gopt.GitInfo = createOpts.gitInfo
				gopt.IOStreams = createOpts.ioStreams
				gopt.CLIOpts = createOpts.cliOpts

				// defaulting the values for repo create command
				gopt.Event.EventType = "[pull_request, push]"
				gopt.Event.BaseBranch = "main"

				if err := generate.Generate(gopt); err != nil {
					return err
				}
			}

			config := &webhook.Options{
				Run:                 run,
				RepositoryURL:       createOpts.gitInfo.URL,
				PACNamespace:        createOpts.pacNamespace,
				RepositoryName:      createOpts.repository.Name,
				RepositoryNamespace: createOpts.repository.Namespace,
				GitHubWebhook:       githubWebhook,
				IOStreams:           ioStreams,
			}

			if githubURLForWebhook != "" {
				config.ProviderAPIURL = githubURLForWebhook
			} else if gitlabURLForWebhook != "" {
				config.ProviderAPIURL = gitlabURLForWebhook
			}

			return config.Install(ctx)
		},
		Annotations: map[string]string{
			"commandType": "main",
		},
	}

	cmd.PersistentFlags().BoolP(noColorFlag, "C", !ioStreams.ColorEnabled(), "Disable coloring")
	cmd.PersistentFlags().StringVar(&createOpts.repository.Name, "name", "", "Repository name")
	cmd.PersistentFlags().StringVar(&createOpts.event.URL, "url", "", "Repository URL")
	cmd.PersistentFlags().StringVarP(&createOpts.repository.Namespace, "namespace", "n", "",
		"The target namespace where the runs will be created")
	cmd.PersistentFlags().StringVarP(&createOpts.pacNamespace, "pac-namespace",
		"", "", "The namespace where pac is installed")
	cmd.PersistentFlags().StringVarP(&githubURLForWebhook, "github-api-url", "", "", "GitHub Enterprise API URL")
	cmd.PersistentFlags().StringVarP(&gitlabURLForWebhook, "gitlab-api-url", "", "", "GitLab Hosted API URL")
	cmd.PersistentFlags().BoolVar(&onlyWebhook, "webhook", false, "Skip repository creation, proceed with configuring webhook")
	cmd.PersistentFlags().BoolVar(&githubWebhook, "github-webhook", false, "Allows configuring webhook if GitHub App is already configured")

	return cmd
}

// getOrCreateNamespace ask and create namespace or use the default one
func getOrCreateNamespace(ctx context.Context, opts *repoOptions) error {
	if opts.repository.Namespace != "" {
		return nil
	}

	// by default, use the current namespace unless it's default or
	// pipelines-as-code and then propose some meaningful namespace based on the
	// git url.
	autoNS := opts.run.Info.Kube.Namespace

	if (autoNS == "default" || autoNS == "pipelines-as-code") &&
		opts.gitInfo.URL != "" {
		autoNS = filepath.Base(opts.gitInfo.URL) + "-pipelines"
	}

	var chosenNS string
	msg := fmt.Sprintf("Please enter the namespace where the pipeline should run (default: %s):", autoNS)
	if err := prompt.SurveyAskOne(&survey.Input{Message: msg}, &chosenNS); err != nil {
		return err
	}

	// set the namespace as the default one
	if chosenNS == "" {
		chosenNS = autoNS
	}
	// check if the namespace exists if it does just exit
	_, err := opts.run.Clients.Kube.CoreV1().Namespaces().Get(ctx, chosenNS, metav1.GetOptions{})
	if err == nil {
		opts.repository.Namespace = chosenNS
		return nil
	}

	fmt.Fprintf(opts.ioStreams.Out, "%s Namespace %s is not found\n",
		opts.ioStreams.ColorScheme().WarningIcon(),
		chosenNS,
	)
	msg = fmt.Sprintf("Would you like me to create the namespace %s?", chosenNS)
	var createNamespace bool
	if err := prompt.SurveyAskOne(&survey.Confirm{Message: msg, Default: true}, &createNamespace); err != nil {
		return err
	}
	if !createNamespace {
		return fmt.Errorf("you need to create the target namespace first")
	}

	_, err = opts.run.Clients.Kube.CoreV1().Namespaces().Create(ctx,
		&v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: chosenNS,
			},
		},
		metav1.CreateOptions{})
	opts.repository.Namespace = chosenNS
	return err
}

// getRepoURL get the repository URL from the user using the git url as default.
func getRepoURL(opts *repoOptions) error {
	if opts.event.URL != "" {
		return nil
	}

	q := "Enter the Git repository url containing the pipelines "
	var err error
	if opts.gitInfo.URL != "" {
		opts.gitInfo.URL, err = cleanupGitURL(opts.gitInfo.URL)
		if err != nil {
			return err
		}
		q += fmt.Sprintf("(default: %s)", opts.gitInfo.URL)
	}
	q += ": "
	if err := prompt.SurveyAskOne(&survey.Input{Message: q}, &opts.event.URL); err != nil {
		return err
	}
	if opts.event.URL != "" {
		return nil
	}
	if opts.event.URL == "" && opts.gitInfo.URL != "" {
		opts.event.URL = opts.gitInfo.URL
		return nil
	}

	return fmt.Errorf("no url has been provided")
}

func cleanupGitURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path), nil
}

func createRepoCRD(ctx context.Context, opts *repoOptions) error {
	repoOwner, err := formatting.GetRepoOwnerFromURL(opts.event.URL)
	if err != nil {
		return fmt.Errorf("invalid git URL: %s, it should be of format: https://gitprovider/project/repository", opts.event.URL)
	}
	repositoryName := strings.ReplaceAll(repoOwner, "/", "-")
	opts.repository, err = opts.run.Clients.PipelineAsCode.PipelinesascodeV1alpha1().Repositories(opts.repository.Namespace).Create(
		ctx,
		&apipac.Repository{
			ObjectMeta: metav1.ObjectMeta{
				Name: repositoryName,
			},
			Spec: apipac.RepositorySpec{
				URL: opts.event.URL,
			},
		},
		metav1.CreateOptions{})
	if err != nil {
		return err
	}
	cs := opts.ioStreams.ColorScheme()
	fmt.Fprintf(opts.ioStreams.Out, "%s Repository %s has been created in %s namespace\n",
		cs.SuccessIconWithColor(cs.Green),
		repositoryName,
		opts.repository.Namespace,
	)
	return nil
}
