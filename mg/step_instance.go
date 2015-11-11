package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"strconv"
	"bufio"
	"os"
	"errors"
)

type StepInstance struct {
	Step
}
const DEFAULT_INSTANCE = "1"
func NewStepInstance(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepInstance {
	stepInstance := new(StepInstance)
	stepInstance.appManifest = appManifest
	stepInstance.cliConnection = cliConnection
	return stepInstance
}

func (s *StepInstance) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(fmt.Sprintf("How much instances do you want <%s> ? ", DEFAULT_INSTANCE))
		instancesBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		instancesString := string(instancesBytes)
		if (instancesString == "") {
			instancesString = DEFAULT_INSTANCE
		}
		instances, err := s.getInstance(instancesString)
		if err != nil {
			showError(err)
			continue
		}
		s.appManifest.Instances(session.AppName, instances)
		break
	}
	return nil
}
func (s *StepInstance) getInstance(instancesString string) (int, error) {
	var instances int
	var errToReturn error = errors.New("Not a valid number of instances.")
	var err error
	instances, err = strconv.Atoi(instancesString)
	if err != nil {
		return 0, errToReturn
	}
	if instances < 0 {
		return 0, errToReturn
	}
	return instances, nil
}