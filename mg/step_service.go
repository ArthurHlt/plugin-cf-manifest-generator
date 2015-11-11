package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"errors"
	"strings"
	"strconv"
	"github.com/daviddengcn/go-colortext"
)

type StepService struct {
	Step
	services    []string
	marketplace []string
}
const DEFAULT_SERVICE = "0"
func NewStepService(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepService {
	stepService := new(StepService)
	stepService.appManifest = appManifest
	stepService.cliConnection = cliConnection
	return stepService
}

func (s *StepService) Run() error {
	reader := bufio.NewReader(os.Stdin)
	if !askYesOrNo("Do you want add services <%s> ? ", true) {
		return nil
	}
	for true {
		fmt.Println("Searching services ...")
		services, err := s.getServices()
		fmt.Println("Available services:")
		if err != nil {
			return err
		}
		for num, buildpack := range services {
			fmt.Println(strconv.Itoa(num) + ". " + buildpack)
		}
		fmt.Print(fmt.Sprintf("Which service do you want <%s> (tip: if enter a new service you could create it) ? ", services[0]))
		serviceBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		service := string(serviceBytes)
		if (service == "") {
			service = DEFAULT_SERVICE
		}
		service, err = s.findService(service)
		if err != nil {
			showError(err)
			err = s.runCreateService(service)
			if err != nil {
				showError(err)
				continue
			}
		}
		s.appManifest.Service(session.AppName, service)
		ct.Foreground(ct.Green, false)
		fmt.Println("Service added.")
		ct.ResetColor()
		if askYesOrNo("Do you want add another service <%s> ? ", true) {
			continue
		}
		break
	}
	return nil
}
func (s *StepService) createServiceFromMarketplace(serviceName, serviceType, plan string) ([]string, error) {
	return s.cliConnection.CliCommand("create-service", serviceType, plan, serviceName)
}
func (s *StepService) createUserProvidedService(serviceName, data string) ([]string, error) {
	return s.cliConnection.CliCommand("create-user-provided-service", serviceName, "-p", data)
}
func (s *StepService) runCreateService(service string) error {
	if !askYesOrNo("Do you want to create service " + service + " <%s> ?", false) {
		return errors.New("So, service not found.")
	}
	reader := bufio.NewReader(os.Stdin)
	if askYesOrNo("Is it an user-provided service <%s> ? ", true) {
		fmt.Print("Data for your service ? ")
		dataBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		data := string(dataBytes)
		if err != nil {
			return err
		}
		createMessage, err := s.createUserProvidedService(service, data)
		fmt.Println(strings.Join(createMessage, "\n"))
		return nil
	}
	fmt.Println("Searching in marketplace ...")
	_, err := s.getMarketplace()
	if err != nil {
		return err
	}

	for true {
		fmt.Print("Service type ? ")
		serviceTypeBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		serviceType := string(serviceTypeBytes)

		fmt.Print("Service plan ? ")
		servicePlanBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		servicePlan := string(servicePlanBytes)

		createMessage, err := s.createServiceFromMarketplace(service, serviceType, servicePlan)
		if err != nil && !IsErrorFromCli(err) {
			return err
		}else if err != nil && IsErrorFromCli(err) {
			showError(errors.New(strings.Join(createMessage, "\n")))
			continue
		}
		fmt.Println(strings.Join(createMessage, "\n"))
		break
	}
	return nil
}
func (s *StepService) getMarketplace() ([]string, error) {
	if len(s.marketplace) > 0 {
		return s.marketplace, nil
	}
	marketplace, err := s.cliConnection.CliCommand("marketplace")
	if err != nil {
		return nil, err
	}
	s.marketplace = marketplace
	return marketplace, nil
}
func (s *StepService) findService(serviceNameOrInt string) (string, error) {
	services, err := s.getServices()
	if err != nil {
		return serviceNameOrInt, err
	}
	service, err := s.findDatabyNameOrId(services, serviceNameOrInt)
	if err != nil && IsNotValidData(err) {
		return serviceNameOrInt, errors.New("Not a valid service.")
	}else if err != nil {
		return serviceNameOrInt, errors.New("Service not found.")
	}
	return service, nil
}

func (s *StepService) getServices() ([]string, error) {
	if len(s.services) > 0 {
		return s.services, nil
	}
	services, err := s.parseDataFromCli([]string{"services"}, 4)
	if err != nil {
		return nil, err
	}
	s.services = services
	return services, nil
}