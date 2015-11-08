package mg

import (
	"os"
	"path"
	"errors"
	"fmt"
	"github.com/cloudfoundry/cli/cf/manifest"
	"github.com/cloudfoundry/cli/plugin"
)
const DEFAULT_MANIFEST = "manifest.yml"

type Generator struct {
	ManifestPath  string
	cliConnection plugin.CliConnection
	appManifest   manifest.AppManifest
	stepAppname   *StepAppname
}
func NewGenerator(manifestPath string, cliConnection plugin.CliConnection) *Generator {
	return &Generator{
		ManifestPath: manifestPath,
		cliConnection: cliConnection,
		appManifest: manifest.NewGenerator(),
	}
}
func (g *Generator) Generate() error {
	manifestPath, err := g.getFinalManifestPath(g.ManifestPath, false)
	g.ManifestPath = manifestPath
	if err != nil {
		return err
	}
	g.appManifest.FileSavePath(manifestPath)

	g.stepAppname = NewStepAppname(g.cliConnection, g.appManifest, g.ManifestPath)
	err = g.stepAppname.Run()

	if err != nil {
		return err
	}
	for _, step := range g.steps() {
		err = step.Run()
		if err != nil {
			return err
		}
	}
	err = g.appManifest.Save()
	return err
}
func (g *Generator) steps() []StepInterface {
	fmt.Println(g.stepAppname.GetAppName())
	return []StepInterface{
		NewStepDomain(g.cliConnection, g.appManifest, g.ManifestPath, g.stepAppname.GetAppName()),
	}
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