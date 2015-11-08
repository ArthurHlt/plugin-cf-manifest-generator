package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/cf/manifest"
	"fmt"
	"strconv"
	"strings"
	"errors"
	"bufio"
	"os"
)

type StepDomain struct {
	Step
	domains []string
}

func NewStepDomain(cliConnection plugin.CliConnection, appManifest manifest.AppManifest, manifestPath, appName string) *StepDomain {
	stepDomain := new(StepDomain)
	stepDomain.appManifest = appManifest
	stepDomain.cliConnection = cliConnection
	stepDomain.manifestPath = manifestPath
	stepDomain.appName = appName
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
		fmt.Print(fmt.Sprintf("What is your domain <%s> ? ", domains[0]))
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
			fmt.Println(fmt.Sprintf("%v", err))
			continue
		}
		fmt.Print(fmt.Sprintf("What is your host <%s> ? ", s.appName))
		hostBytes, _, err := reader.ReadLine()
		host := string(hostBytes)
		if (host == "") {
			host = s.appName
		}
		s.appManifest.Domain(s.appName, host, domain)
		break
	}
	return nil
}
func (s *StepDomain) findDomain(domainNameOrInt string) (string, error) {
	domains, err := s.getDomains()
	if err != nil {
		return "", err
	}
	numDomain, err := strconv.Atoi(domainNameOrInt)
	if err == nil {
		if numDomain < 0 || numDomain > len(domains) - 1 {
			return "", errors.New("Not a valid domain.")
		}
		return domains[numDomain], nil
	}
	for _, domain := range domains {
		if domain == domainNameOrInt {
			return domain, nil
		}
	}
	return "", errors.New("Domain not found.")
}
func (s *StepDomain) getDomains() ([]string, error) {
	if len(s.domains) > 0 {
		return s.domains, nil
	}
	domainsTerminal, err := s.cliConnection.CliCommandWithoutTerminalOutput("domains")
	if err != nil {
		return nil, err
	}
	if len(domainsTerminal) <= 2 {
		return nil, errors.New(strings.Join(domainsTerminal, "\n"))
	}
	domainsUnparsed := domainsTerminal[2:]
	var domainsParsed []string = make([]string, len(domainsUnparsed))
	for index, domainUnparsed := range domainsUnparsed {
		domainsParsed[index] = strings.Split(domainUnparsed, " ")[0]
	}
	s.domains = domainsParsed
	return domainsParsed, nil
}