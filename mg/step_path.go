package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"path"
	"errors"
)

type StepPath struct {
	Step
}
const DEFAULT_PATH = "./"
func NewStepPath(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepPath {
	stepPath := new(StepPath)
	stepPath.appManifest = appManifest
	stepPath.cliConnection = cliConnection
	return stepPath
}

func (s *StepPath) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(fmt.Sprintf("What is the path of your app <%s> ? ", DEFAULT_PATH))
		pathBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		pathName := string(pathBytes)
		if (pathName == "") {
			pathName = DEFAULT_PATH
		}
		pathName, err = s.getPath(pathName)
		if err != nil {
			showError(err)
			if !askYesOrNo("Do you want to force it <%s> ? ", false) {
				continue
			}
		}
		s.appManifest.Path(session.AppName, pathName)
		break
	}
	return nil
}
func (s *StepPath) getPath(pathName string) (string, error) {
	var errToReturn = errors.New("Path doesn't exist.")
	if path.IsAbs(pathName) {
		_, err := os.Stat(pathName)
		if err != nil {
			return pathName, errToReturn
		}
	}
	finalPath := path.Join(path.Dir(session.ManifestPath), pathName)
	_, err := os.Stat(finalPath)
	if err != nil {
		return pathName, errToReturn
	}
	return pathName, nil
}