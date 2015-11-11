package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"strconv"
	"errors"
	"bufio"
	"os"
	"github.com/daviddengcn/go-colortext"
)

type StepDomain struct {
	Step
	domains []string
}

func NewStepDomain(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepDomain {
	stepDomain := new(StepDomain)
	stepDomain.appManifest = appManifest
	stepDomain.cliConnection = cliConnection
	return stepDomain
}

func (s *StepDomain) Run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Searching domains ...")
	domains, err := s.getDomains()
	fmt.Println("Available domains:")
	if err != nil {
		return err
	}
	for num, domain := range domains {
		fmt.Println(strconv.Itoa(num) + ". " + domain)
	}
	for true {
		fmt.Print(fmt.Sprintf("Which domain do you want <%s> ? ", domains[0]))
		domainBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		domain := string(domainBytes)
		if (domain == "") {
			domain = "0"
		}
		domain, err = s.findDomain(domain)
		if err != nil {
			showError(err)
			continue
		}
		fmt.Print(fmt.Sprintf("What is your host <%s> ? ", session.AppName))
		hostBytes, _, err := reader.ReadLine()
		host := string(hostBytes)
		if (host == "") {
			host = session.AppName
		}
		s.appManifest.Domain(session.AppName, host, domain)
		ct.Foreground(ct.Green, false)
		fmt.Println("Domain added.")
		ct.ResetColor()
		if askYesOrNo("Do you want add another domain <%s> ? ", true) {
			continue
		}
		break
	}
	return nil
}
func (s *StepDomain) findDomain(domainNameOrInt string) (string, error) {
	domains, err := s.getDomains()
	if err != nil {
		return "", err
	}
	domain, err := s.findDatabyNameOrId(domains, domainNameOrInt)
	if err != nil && IsNotValidData(err) {
		return "", errors.New("Not a valid domain.")
	}else if err != nil {
		return "", errors.New("Domain not found.")
	}
	return domain, nil
}
func (s *StepDomain) getDomains() ([]string, error) {
	if len(s.domains) > 0 {
		return s.domains, nil
	}
	domains, err := s.parseDataFromCli([]string{"domains"}, 2)
	if err != nil {
		return nil, err
	}
	s.domains = domains
	return domains, nil
}