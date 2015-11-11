package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
)

type StepCommand struct {
	Step
}
func NewStepCommand(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepCommand {
	stepCommand := new(StepCommand)
	stepCommand.appManifest = appManifest
	stepCommand.cliConnection = cliConnection
	return stepCommand
}

func (s *StepCommand) Run() error {
	reader := bufio.NewReader(os.Stdin)
	if !askYesOrNo("Do you want set a command <%s> ? ", true) {
		return nil
	}
	fmt.Print(fmt.Sprintf("What is your command ? "))
	commandBytes, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	command := string(commandBytes)
	s.appManifest.Command(session.AppName, command)
	return nil
}