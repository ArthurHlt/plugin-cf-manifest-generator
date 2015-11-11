package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"errors"
	"path"
)

type StepInherit struct {
	Step
}
const DEFAULT_INHERIT = "manifest.yml"
func NewStepInherit(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepInherit {
	stepInherit := new(StepInherit)
	stepInherit.appManifest = appManifest
	stepInherit.cliConnection = cliConnection
	return stepInherit}

func (s *StepInherit) Run() error {
	reader := bufio.NewReader(os.Stdin)
	if !askYesOrNo("Do you want herit from another manifest <%s> ? ", true) {
		return nil
	}
	for true {
		fmt.Print(fmt.Sprintf("What is the path of this manifest file <%s> ? ", DEFAULT_INHERIT))
		inheritBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		inherit := string(inheritBytes)
		if (inherit == "") {
			inherit = DEFAULT_INHERIT}
		inherit, err = s.getInherit(inherit)
		if err != nil {
			showError(err)
			if !askYesOrNo("Do you want to force it <%s> ? ", false) {
				continue
			}
		}
		s.appManifest.Inherit(inherit)
		break
	}
	return nil
}
func (s *StepInherit) getInherit(inherit string) (string, error) {
	var errToReturn = errors.New("This manifest doesn't exist.")
	if path.IsAbs(inherit) {
		_, err := os.Stat(inherit)
		if err != nil {
			return inherit, errToReturn
		}
	}
	finalPath := path.Join(path.Dir(session.ManifestPath), inherit)
	_, err := os.Stat(finalPath)
	if err != nil {
		return inherit, errToReturn
	}
	return inherit, nil
}