package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
)

type StepNoRoute struct {
	Step
}

func NewStepNoRoute(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepNoRoute {
	stepNoRoute := new(StepNoRoute)
	stepNoRoute.appManifest = appManifest
	stepNoRoute.cliConnection = cliConnection
	return stepNoRoute
}

func (s *StepNoRoute) Run() error {
	s.appManifest.NoRoute(session.AppName,
		askYesOrNo("Do you want no route <%s> ? ", true),
	)
	return nil
}
