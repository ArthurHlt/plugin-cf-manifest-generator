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

type StepStack struct {
	Step
	stacks []string
}
const DEFAULT_STACK = "0"
func NewStepStack(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepStack {
	stepStack := new(StepStack)
	stepStack.appManifest = appManifest
	stepStack.cliConnection = cliConnection
	return stepStack
}

func (s *StepStack) Run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Searching stacks ...")
	stacks, err := s.getStacks()
	if err != nil {
		return err
	}
	fmt.Println("Available stacks:")
	for num, stack := range stacks {
		fmt.Println(strconv.Itoa(num) + ". " + stack)
	}
	for true {
		fmt.Print(fmt.Sprintf("Which stack do you want <%s> ? ", stacks[0]))
		stackBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		stack := string(stackBytes)
		if (stack == "") {
			stack = DEFAULT_STACK
		}
		stack, err = s.findStack(stack)
		if err != nil {
			showError(err)
			continue
		}
		s.appManifest.Stack(session.AppName, stack)
		break
	}
	return nil
}
func (s *StepStack) findStack(stackNameOrInt string) (string, error) {
	stacks, err := s.getStacks()
	if err != nil {
		return stackNameOrInt, err
	}
	stack, err := s.findDatabyNameOrId(stacks, stackNameOrInt)
	if err != nil && IsNotValidData(err) {
		return stackNameOrInt, errors.New("Not a valid stack.")
	}else if err != nil {
		return stackNameOrInt, errors.New("Stack not found.")
	}
	return stack, nil
}
func (s *StepStack) getStacks() ([]string, error) {
	if len(s.stacks) > 0 {
		return s.stacks, nil
	}
	services, err := s.parseDataFromCli([]string{"stacks"}, 4)
	if err != nil {
		return nil, err
	}
	s.stacks = services
	return services, nil
}