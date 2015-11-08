package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/cf/manifest"
	"fmt"
	"bufio"
	"os"
	"path"
)

type StepAppname struct {
	Step
}

func NewStepAppname(cliConnection plugin.CliConnection, appManifest manifest.AppManifest, manifestPath string) *StepAppname {
	stepAppname := new(StepAppname)
	stepAppname.appManifest = appManifest
	stepAppname.cliConnection = cliConnection
	stepAppname.manifestPath = manifestPath
	return stepAppname
}

func (s *StepAppname) Run() error {
	reader := bufio.NewReader(os.Stdin)
	defaultName, err := s.defaultName()
	if err != nil {
		return err
	}
	fmt.Print(fmt.Sprintf("What is the name of your app <%s> ? ", defaultName))
	appNameBytes, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	appName := string(appNameBytes)
	if (appName == "") {
		appName = defaultName
	}
	s.appName = appName
	s.appManifest.Memory(appName, 1024)
	s.appManifest.Instances(appName, 1)
	return nil
}
func (s *StepAppname) GetAppName() string {
	return s.appName
}
func (s *StepAppname) defaultName() (string, error) {
	defaultName := path.Base(path.Dir(s.manifestPath))
	if defaultName != "." {
		return defaultName, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Base(wd), nil
}