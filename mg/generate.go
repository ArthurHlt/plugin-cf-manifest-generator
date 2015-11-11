package mg

import (
	"os"
	"path"
	"errors"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"github.com/cloudfoundry/cli/plugin"
	"fmt"
	"github.com/daviddengcn/go-colortext"
)
const DEFAULT_MANIFEST = "manifest.yml"

type Session struct {
	AppName      string
	ManifestPath string
}
type Generator struct {
	cliConnection plugin.CliConnection
	appManifest   *manifest.Manifest
	Session       *Session
	steps         []StepInterface
	stepsDetails  []StepInterface
}
var session *Session = &Session{}
func NewGenerator(manifestPath string, cliConnection plugin.CliConnection) *Generator {
	session.ManifestPath = manifestPath
	return &Generator{
		cliConnection: cliConnection,
		appManifest: manifest.NewManifest(),
		Session: session,
	}
}
func (g *Generator) Generate() error {
	manifestPath, err := g.getFinalManifestPath(session.ManifestPath, false)
	session.ManifestPath = manifestPath
	if err != nil {
		return err
	}
	g.appManifest.FileSavePath(manifestPath)
	for true {
		for _, step := range g.getSteps() {
			err = step.Run()
			if err != nil {
				return err
			}
			fmt.Print("\n")
		}
		if askYesOrNo("Do you want set more detailed informations <%s> ? ", true) {
			for _, step := range g.getStepsDetails() {
				err = step.Run()
				if err != nil {
					return err
				}
				fmt.Print("\n")
			}
		}
		err = g.appManifest.Save()
		if err != nil {
			return err
		}
		ct.Foreground(ct.Green, false)
		fmt.Println(fmt.Sprintf("App %s has been added.\n", session.AppName))
		ct.ResetColor()
		if !askYesOrNo("Do you want to add an other app <%s> ? ", true) {
			break;
		}
		fmt.Print("\n")
	}
	return nil
}
func (g *Generator) getStepsDetails() []StepInterface {
	if len(g.stepsDetails) > 0 {
		return g.stepsDetails
	}
	stepsDetails := []StepInterface{
		NewStepStack(g.cliConnection, g.appManifest),
		NewStepTimeout(g.cliConnection, g.appManifest),
		NewStepRandomroute(g.cliConnection, g.appManifest),
		NewStepNoRoute(g.cliConnection, g.appManifest),
		NewStepNoHostname(g.cliConnection, g.appManifest),
	}
	g.stepsDetails = stepsDetails
	return stepsDetails
}
func (g *Generator) getSteps() []StepInterface {
	if len(g.steps) > 0 {
		return g.steps
	}
	steps := []StepInterface{
		NewStepAppname(g.cliConnection, g.appManifest),
		NewStepInstance(g.cliConnection, g.appManifest),
		NewStepMemory(g.cliConnection, g.appManifest),
		NewStepPath(g.cliConnection, g.appManifest),
		NewStepDiskQuota(g.cliConnection, g.appManifest),
		NewStepDomain(g.cliConnection, g.appManifest),
		NewStepBuildpack(g.cliConnection, g.appManifest),
		NewStepService(g.cliConnection, g.appManifest),
		NewStepEnv(g.cliConnection, g.appManifest),
		NewStepCommand(g.cliConnection, g.appManifest),
		NewStepInherit(g.cliConnection, g.appManifest),
	}
	g.steps = steps
	return steps
}
func (g *Generator) getFinalManifestPath(manifestPath string, finalFind bool) (string, error) {
	fileInfo, err := os.Stat(manifestPath)
	if err == nil && fileInfo.IsDir() {
		return path.Join(manifestPath, DEFAULT_MANIFEST), nil
	}
	dir := path.Dir(manifestPath)
	fileInfo, err = os.Stat(dir)
	if err == nil && dir != "." && fileInfo.IsDir() {
		return manifestPath, nil
	}
	if finalFind {
		return "", errors.New("Cannot find directory " + dir)
	}
	wd, err := os.Getwd()
	if (err != nil) {
		return "", err
	}
	manifestPath, err = g.getFinalManifestPath(path.Join(wd, manifestPath), true)
	return manifestPath, err
}
