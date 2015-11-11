package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
)

type StepNoHostname struct {
	Step
}
func NewStepNoHostname(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepNoHostname {
	stepNoHostname := new(StepNoHostname)
	stepNoHostname.appManifest = appManifest
	stepNoHostname.cliConnection = cliConnection
	return stepNoHostname}

func (s *StepNoHostname) Run() error {
	s.appManifest.NoHostname(session.AppName,
		askYesOrNo("Do you want to use no hostname <%s> ? ", true),
	)
	return nil
}