package kubeinteraction

import (
	"strconv"

	"github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/keys"
	apipac "github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/v1alpha1"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/formatting"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/params/info"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/params/version"
	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

const (
	StateStarted   = "started"
	StateQueued    = "queued"
	StateCompleted = "completed"
	StateFailed    = "failed"
)

func AddLabelsAndAnnotations(event *info.Event, pipelineRun *tektonv1.PipelineRun, repo *apipac.Repository, providerinfo *info.ProviderConfig) {
	// Add labels on the soon to be created pipelinerun so UI/CLI can easily
	// query them.
	labels := map[string]string{
		// These keys are used in LabelSelector query so we are keeping in Labels as it is.
		// But adding same keys to Annotations so UI/CLI can fetch the actual value instead of modified value
		"app.kubernetes.io/managed-by": pipelinesascode.GroupName,
		"app.kubernetes.io/version":    formatting.CleanValueKubernetes(version.Version),
		keys.URLOrg:                    formatting.CleanValueKubernetes(event.Organization),
		keys.URLRepository:             formatting.CleanValueKubernetes(event.Repository),
		keys.SHA:                       formatting.CleanValueKubernetes(event.SHA),
		keys.Repository:                formatting.CleanValueKubernetes(repo.GetName()),
		keys.State:                     StateStarted,
		keys.EventType:                 formatting.CleanValueKubernetes(event.EventType),

		// We are deprecating these below keys from labels and adding it to Annotations
		// In PAC v0.20.x releases we will remove these keys from Labels
		keys.Sender:      formatting.CleanValueKubernetes(event.Sender),
		keys.Branch:      formatting.CleanValueKubernetes(event.BaseBranch),
		keys.GitProvider: providerinfo.Name,
	}

	annotations := map[string]string{
		keys.ShaTitle:      event.SHATitle,
		keys.ShaURL:        event.SHAURL,
		keys.RepoURL:       event.URL,
		keys.URLOrg:        event.Organization,
		keys.URLRepository: event.Repository,
		keys.SHA:           event.SHA,
		keys.Sender:        event.Sender,
		keys.EventType:     event.EventType,
		keys.Branch:        event.BaseBranch,
		keys.Repository:    repo.GetName(),
		keys.GitProvider:   providerinfo.Name,
		keys.State:         StateStarted,
	}

	if event.PullRequestNumber != 0 {
		labels[keys.PullRequest] = strconv.Itoa(event.PullRequestNumber)
		annotations[keys.PullRequest] = strconv.Itoa(event.PullRequestNumber)
	}

	// TODO: move to provider specific function
	if providerinfo.Name == "github" || providerinfo.Name == "github-enterprise" {
		if event.InstallationID != -1 {
			annotations[keys.InstallationID] = strconv.FormatInt(event.InstallationID, 10)
		}
		if event.GHEURL != "" {
			annotations[keys.GHEURL] = event.GHEURL
		}
	}

	// GitLab
	if event.SourceProjectID != 0 {
		annotations[keys.SourceProjectID] = strconv.Itoa(event.SourceProjectID)
	}
	if event.TargetProjectID != 0 {
		annotations[keys.TargetProjectID] = strconv.Itoa(event.TargetProjectID)
	}

	// Don't overwrite the labels if there is some who already exist set by the user in repo
	if pipelineRun.GetLabels() == nil {
		pipelineRun.Labels = map[string]string{}
	}
	// Don't overwrite the annotation if there is some who already exist set by the user in repo
	if pipelineRun.GetAnnotations() == nil {
		pipelineRun.Annotations = map[string]string{}
	}

	for k, v := range labels {
		pipelineRun.Labels[k] = v
	}
	for k, v := range annotations {
		pipelineRun.Annotations[k] = v
	}
}
