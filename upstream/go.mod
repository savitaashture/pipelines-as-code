module github.com/openshift-pipelines/pipelines-as-code

go 1.22.0

toolchain go1.23.0

require (
	code.gitea.io/gitea v1.22.1
	code.gitea.io/sdk/gitea v0.19.0
	github.com/AlecAivazis/survey/v2 v2.3.7
	github.com/bradleyfalzon/ghinstallation/v2 v2.11.0
	github.com/cloudevents/sdk-go/v2 v2.15.2
	github.com/fvbommel/sortorder v1.1.0
	github.com/gfleury/go-bitbucket-v1 v0.0.0-20240822132758-a3031aa024b4
	github.com/gobwas/glob v0.2.3
	github.com/google/cel-go v0.21.0
	github.com/google/go-cmp v0.6.0
	github.com/google/go-github/scrape v0.0.0-20240903181346-00caddcf95be
	github.com/google/go-github/v62 v62.0.0
	github.com/google/go-github/v64 v64.0.0
	github.com/hako/durafmt v0.0.0-20210608085754-5c1018a4e16b
	github.com/jonboulle/clockwork v0.4.0
	github.com/juju/ansiterm v1.0.0
	github.com/ktrysmt/go-bitbucket v0.9.80
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.20
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d
	github.com/mitchellh/mapstructure v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.9.0
	github.com/tektoncd/pipeline v0.63.0
	github.com/xanzy/go-gitlab v0.108.0
	go.opencensus.io v0.24.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240823005443-9b4947da3948
	golang.org/x/oauth2 v0.22.0
	golang.org/x/sync v0.8.0
	golang.org/x/text v0.17.0
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools/v3 v3.5.1
	k8s.io/api v0.31.0
	k8s.io/apimachinery v0.31.0
	k8s.io/client-go v0.31.0
	k8s.io/utils v0.0.0-20240902221715-702e33fdd3c3
	knative.dev/eventing v0.42.1
	knative.dev/pkg v0.0.0-20240815051656-89743d9bbf7c
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/coreos/go-oidc/v3 v3.9.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
)

replace (
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20230515203736-54b630e78af5
	knative.dev/eventing => knative.dev/eventing v0.39.0
)

require (
	contrib.go.opencensus.io/exporter/ocagent v0.7.1-0.20200907061046-05415f1de66d // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.4.2 // indirect
	contrib.go.opencensus.io/exporter/zipkin v0.1.2 // indirect
	github.com/PuerkitoBio/goquery v1.9.2 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudevents/sdk-go/observability/opencensus/v2 v2.15.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/davidmz/go-pageant v1.0.2 // indirect
	github.com/emicklei/go-restful/v3 v3.12.1 // indirect
	github.com/evanphx/json-patch/v5 v5.9.0 // indirect
	github.com/go-fed/httpsig v1.1.1-0.20201223112313-55836744818e // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.1
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic v0.7.0 // indirect
	github.com/google/gnostic-models v0.6.9-0.20230804172637-c7be7c783f49 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.20.2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.58.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/prometheus/statsd_exporter v0.27.1 // indirect
	github.com/rickb777/date v1.20.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/xlzd/gotp v0.1.0 // indirect
	go.uber.org/automaxprocs v1.5.3 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/term v0.23.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/api v0.195.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/grpc v1.66.0 // indirect
	google.golang.org/protobuf v1.34.2
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.31.0 // indirect
	k8s.io/klog/v2 v2.130.1
	k8s.io/kube-openapi v0.0.0-20240903163716-9e1beecbcb38 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)
