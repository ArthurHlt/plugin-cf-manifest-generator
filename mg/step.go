package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/cf/manifest"
)

type StepInterface interface {
	Run() error
}
type Step struct {
	StepInterface
	cliConnection plugin.CliConnection
	appManifest   manifest.AppManifest
	manifestPath  string
	appName       string
}