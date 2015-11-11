package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"errors"
)

type StepMemory struct {
	Step
}
const DEFAULT_MEMORY = "512M"
func NewStepMemory(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepMemory {
	stepMemory := new(StepMemory)
	stepMemory.appManifest = appManifest
	stepMemory.cliConnection = cliConnection
	return stepMemory
}

func (s *StepMemory) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(fmt.Sprintf("How much memory do you want <%s> ? ", DEFAULT_MEMORY))
		memoryBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		memoryString := string(memoryBytes)
		if (memoryString == "") {
			memoryString = DEFAULT_MEMORY
		}
		memory, err := s.getMemory(memoryString)
		if err != nil {
			showError(err)
			continue
		}
		s.appManifest.Memory(session.AppName, memory)
		break
	}
	return nil
}
func (s *StepMemory) getMemory(memoryString string) (int, error) {
	var errToReturn error = errors.New("Not a valid memory.")
	size, err := sizeMachineParsing(memoryString)
	if err != nil {
		return 0, errToReturn
	}
	return size, nil
}