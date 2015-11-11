package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"github.com/daviddengcn/go-colortext"
	"fmt"
	"strings"
	"errors"
	"strconv"
	"bufio"
	"os"
)

type StepInterface interface {
	Run() error
}
type Step struct {
	StepInterface
	cliConnection plugin.CliConnection
	appManifest   *manifest.Manifest
}
func showError(err error) {
	if err != nil {
		ct.Foreground(ct.Yellow, true)
		fmt.Println(fmt.Sprintf("%v", err))
		ct.ResetColor()
	}
}
func parseSizeToInt(sizeString, suffix string) (int, error) {
	if suffix != "" {
		sizeString = strings.TrimSuffix(sizeString, suffix)
	}
	size, err := strconv.Atoi(sizeString)
	if err != nil {
		return 0, err
	}
	if size < 0 {
		return 0, errors.New("Should be superior to 0")
	}
	return size, nil
}
func sizeMachineParsing(sizeString string) (int, error) {
	var size int
	var suffix string
	var errToReturn error = errors.New("Not valid.")
	var err error
	var multiplicator int = 1
	sizeString = strings.ToLower(sizeString)
	if (strings.HasSuffix(sizeString, "mb")) {
		suffix = "mb"
	}else if (strings.HasSuffix(sizeString, "m")) {
		suffix = "m"
	} else if (strings.HasSuffix(sizeString, "gb")) {
		suffix = "gb"
		multiplicator = 1024
	}else if (strings.HasSuffix(sizeString, "g")) {
		suffix = "g"
		multiplicator = 1024
	}
	size, err = parseSizeToInt(sizeString, suffix)
	if err != nil {
		return 0, errToReturn
	}
	return size * multiplicator, nil
}
func IsErrorFromCli(err error) bool {
	errString := strings.ToLower(fmt.Sprintf("%v", err))
	if strings.Contains(errString, "error executing cli") {
		return true
	}
	return false
}
func askYesOrNo(message string, noByDefault bool) bool {
	var defaultMessage string
	var defaultValue string
	var defaultReturn bool
	if noByDefault {
		defaultMessage = "y/N"
		defaultValue = "y"
		defaultReturn = true
	}else {
		defaultMessage = "Y/n"
		defaultValue = "n"
		defaultReturn = false
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(fmt.Sprintf(message, defaultMessage))
	forceBytes, _, err := reader.ReadLine()
	if err != nil {
		return !defaultReturn
	}
	force := string(forceBytes)
	if (strings.Contains(strings.ToLower(force), defaultValue)) {
		return defaultReturn
	}
	return !defaultReturn
}
func (s *Step) parseDataFromCli(command []string, startLine int) ([]string, error) {
	output, err := s.cliConnection.CliCommandWithoutTerminalOutput(command...)
	if err != nil {
		return nil, errors.New(strings.Join(output, "\n"))
	}
	if len(output) < startLine {
		return nil, errors.New(strings.Join(output, "\n"))
	}
	datasUnparsed := output[startLine:]
	var datasParsed []string = make([]string, len(datasUnparsed))
	for index, dataUnparsed := range datasUnparsed {
		datasParsed[index] = strings.Split(dataUnparsed, " ")[0]
	}
	return datasParsed, nil
}
func IsNotValidData(err error) bool {
	errString := fmt.Sprintf("%v", err)
	if (errString == "Not valid.") {
		return true
	}
	return false
}
func IsNotFoundData(err error) bool {
	errString := fmt.Sprintf("%v", err)
	if (errString == "Not found.") {
		return true
	}
	return false
}
func (s *Step) findDatabyNameOrId(datas []string, dataNameOrInt string) (string, error) {
	index, err := strconv.Atoi(dataNameOrInt)
	if err == nil {
		if index < 0 || index > len(datas) - 1 {
			return dataNameOrInt, errors.New("Not valid.")
		}
		return datas[index], nil
	}
	for _, data := range datas {
		if data == dataNameOrInt {
			return data, nil
		}
	}
	return dataNameOrInt, errors.New("Not found.")
}