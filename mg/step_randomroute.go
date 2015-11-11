package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
)

type StepRandomroute struct {
	Step
}
func NewStepRandomroute(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepRandomroute {
	stepRandomroute := new(StepRandomroute)
	stepRandomroute.appManifest = appManifest
	stepRandomroute.cliConnection = cliConnection
	return stepRandomroute
}

func (s *StepRandomroute) Run() error {
	s.appManifest.RandomRoute(session.AppName,
		askYesOrNo("Do you want use a random route <%s> ? ", true),
	)

	return nil
}
