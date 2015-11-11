package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"errors"
)

type StepDiskQuota struct {
	Step
}
const DEFAULT_DISKQUOTA = "1G"
func NewStepDiskQuota(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepDiskQuota {
	stepDiskQuota := new(StepDiskQuota)
	stepDiskQuota.appManifest = appManifest
	stepDiskQuota.cliConnection = cliConnection
	return stepDiskQuota
}

func (s *StepDiskQuota) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(fmt.Sprintf("What is the size of your disk <%s> ? ", DEFAULT_DISKQUOTA))
		diskQuotaBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		diskQuotaString := string(diskQuotaBytes)
		if (diskQuotaString == "") {
			diskQuotaString = DEFAULT_DISKQUOTA
		}
		diskQuota, err := s.getDiskQuota(diskQuotaString)
		if err != nil {
			showError(err)
			continue
		}
		s.appManifest.DiskQuota(session.AppName, diskQuota)
		break
	}
	return nil
}
func (s *StepDiskQuota) getDiskQuota(diskQuota string) (int, error) {
	var errToReturn error = errors.New("Not a valid disk quota.")
	size, err := sizeMachineParsing(diskQuota)
	if err != nil {
		return 0, errToReturn
	}
	return size, nil
}