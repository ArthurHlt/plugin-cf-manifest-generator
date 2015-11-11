package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
)

type StepTimeout struct {
	Step
}
const DEFAULT_TIMEOUT = "80"
func NewStepTimeout(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepTimeout {
	stepTimeout := new(StepTimeout)
	stepTimeout.appManifest = appManifest
	stepTimeout.cliConnection = cliConnection
	return stepTimeout
}

func (s *StepTimeout) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(fmt.Sprintf("How much time your app have to start <%s> (in seconds) ? ", DEFAULT_TIMEOUT))
		timeoutBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		timeoutString := string(timeoutBytes)
		if (timeoutString == "") {
			timeoutString = DEFAULT_TIMEOUT
		}
		timeout, err := s.getTimeout(timeoutString)
		if err != nil {
			showError(err)
			continue
		}
		s.appManifest.Timeout(session.AppName, timeout)
		break
	}
	return nil
}
func (s *StepTimeout) getTimeout(timeoutString string) (int, error) {
	timeout, err := strconv.Atoi(timeoutString)
	if err != nil {
		return 0, errors.New("It's not a valid number.")
	}
	if timeout <= 0 {
		return 0, errors.New("You can't pass a number under or equal to 0.")
	}
	return timeout, nil
}