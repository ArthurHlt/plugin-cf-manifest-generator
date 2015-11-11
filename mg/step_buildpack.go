package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
	"net/url"
	"net/http"
	"regexp"
)

type StepBuildpack struct {
	Step
	buildpacks []string
}
func NewStepBuildpack(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepBuildpack {
	stepBuildpack := new(StepBuildpack)
	stepBuildpack.appManifest = appManifest
	stepBuildpack.cliConnection = cliConnection
	return stepBuildpack}

func (s *StepBuildpack) Run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Searching buildpacks ...")
	buildpacks, err := s.getBuildpacks()
	fmt.Println("Available buildpacks:")
	if err != nil {
		return err
	}
	for num, buildpack := range buildpacks {
		fmt.Println(strconv.Itoa(num) + ". " + buildpack)
	}
	for true {
		fmt.Print("Which buildpack do you want <none> (tip: you can pass an url) ? ")
		buildpackBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		buildpack := string(buildpackBytes)
		if (buildpack == "") {
			break
		}
		buildpack, err = s.findBuildpacks(buildpack)
		if err != nil {
			showError(err)
			continue
		}
		s.appManifest.Buildpack(session.AppName, buildpack)
		break
	}
	return nil
}
func (s *StepBuildpack) findBuildpacks(buildpackNameOrInt string) (string, error) {
	if s.isUrl(buildpackNameOrInt) && s.isValidUrl(buildpackNameOrInt) {
		return buildpackNameOrInt, nil
	}else if s.isUrl(buildpackNameOrInt) && !s.isValidUrl(buildpackNameOrInt) {
		return "", errors.New("Not a valid URL.")
	}
	buildpacks, err := s.getBuildpacks()
	if err != nil {
		return "", err
	}
	buildpack, err := s.findDatabyNameOrId(buildpacks, buildpackNameOrInt)
	if err != nil && IsNotValidData(err) {
		return "", errors.New("Not a valid buildpack.")
	}else if err != nil {
		return "", errors.New("Buildpack not found.")
	}
	return buildpack, nil
}
func (s *StepBuildpack) isValidUrl(buildpackUrl string) bool {
	if !s.isUrl(buildpackUrl) {
		return false
	}
	resp, err := http.Get(buildpackUrl)
	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
func (s *StepBuildpack) isUrl(buildpackUrl string) bool {
	_, err := url.Parse(buildpackUrl)
	if err != nil {
		return false
	}
	match, err := regexp.MatchString("^http(s?)://", buildpackUrl)
	if err != nil {
		return false
	}
	return match
}
func (s *StepBuildpack) getBuildpacks() ([]string, error) {
	if len(s.buildpacks) > 0 {
		return s.buildpacks, nil
	}
	buildpacks, err := s.parseDataFromCli([]string{"buildpacks"}, 2)
	if err != nil {
		return nil, err
	}
	s.buildpacks = buildpacks
	return buildpacks, nil
}