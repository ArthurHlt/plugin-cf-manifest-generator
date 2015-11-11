package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"github.com/daviddengcn/go-colortext"
)

type StepEnv struct {
	Step
}
func NewStepEnv(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepEnv {
	stepEnv := new(StepEnv)
	stepEnv.appManifest = appManifest
	stepEnv.cliConnection = cliConnection
	return stepEnv
}

func (s *StepEnv) Run() error {
	reader := bufio.NewReader(os.Stdin)
	if !askYesOrNo("Do you want add environment variables <%s> ? ", true) {
		return nil
	}
	for true {
		fmt.Print("Key ? ")
		keyBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		key := string(keyBytes)
		if key == "" {
			break
		}
		fmt.Print("Value ? ")
		valueBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		value := string(valueBytes)
		s.appManifest.Env(session.AppName, key, value)
		ct.Foreground(ct.Green, false)
		fmt.Println("Env var added. Type enter when finish.")
		ct.ResetColor()
	}

	return nil
}