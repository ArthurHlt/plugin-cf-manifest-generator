package manifest

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)
func NewManifest() *Manifest {
	manifest := &Manifest{}
	manifestContent := &ManifestContent{}
	application := make([]Application, 0)
	manifestContent.Applications = application
	manifest.Content = manifestContent
	return manifest
}

func (m *Manifest) FileSavePath(manifestPath string) {
	m.ManifestPath = manifestPath
}

func (m *Manifest) Save() error {
	var data []byte
	var err error
	m.clean()
	if len(m.Content.Applications) == 1 {
		data, err = yaml.Marshal(m.Content.Applications[0])
	}else {
		data, err = yaml.Marshal(m.Content)
	}
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(m.ManifestPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (m *Manifest) clean() {
	if len(m.Content.Applications) <= 1 {
		return
	}
	for _, app := range m.Content.Applications {
		if app.Inherit != "" {
			m.Inherit(app.Name, app.Inherit)
		}
	}
}
func (m *Manifest) Instances(appName string, instances int) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Instances = instances
}
func (m *Manifest) Env(appName string, key, value string) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Env[key] = value
}
func (m *Manifest) Domains(appName string, host string, domain string) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Hosts = append(m.Content.Applications[index].Hosts, host)
	m.Content.Applications[index].Domains = append(m.Content.Applications[index].Domains, domain)
}
func (m *Manifest) Domain(appName string, host string, domain string) {
	index := m.FindOrCreateApp(appName)
	if m.Content.Applications[index].Host != "" && m.Content.Applications[index].Domain != "" {
		m.Domains(appName, host, domain)
	}
	m.Content.Applications[index].Host = host
	m.Content.Applications[index].Domain = domain
}
func (m *Manifest) Path(appName, path string) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Path = path
}
func (m *Manifest) Command(appName, command string) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Command = command
}
func (m *Manifest) Buildpack(appName, buildpack string) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Buildpack = buildpack
}
func (m *Manifest) Service(appName, service string) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Services = append(m.Content.Applications[index].Services, service)
}
func (m *Manifest) Memory(appName string, memory int) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Memory = strconv.Itoa(memory) + "M"
}
func (m *Manifest) NoRoute(appName string, noRoute bool) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].NoRoute = noRoute
}
func (m *Manifest) RandomRoute(appName string, randomRoute bool) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].RandomRoute = randomRoute
}
func (m *Manifest) NoHostname(appName string, noHostname bool) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].NoHostname = noHostname
}
func (m *Manifest) Inherit(appName string, inherit string) {
	index := m.FindOrCreateApp(appName)
	if len(m.Content.Applications) == 2 {
		m.Content.Applications[0].Inherit = ""
	}
	if len(m.Content.Applications) > 1 {
		m.Content.Inherit = inherit
	}else {
		m.Content.Applications[index].Inherit = inherit
	}

}
func (m *Manifest) Stack(appName string, stack string) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Stack = stack
}
func (m *Manifest) Timeout(appName string, timeout int) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].Timeout = timeout
}
func (m *Manifest) DiskQuota(appName string, diskQuota int) {
	index := m.FindOrCreateApp(appName)
	m.Content.Applications[index].DiskQuota = strconv.Itoa(diskQuota) + "M"
}
func (m *Manifest) FindOrCreateApp(appName string) int {
	var app Application
	var index int
	for index, app = range m.Content.Applications {
		if app.Name == appName {
			return index
		}
	}
	app = Application{
		Name: appName,
		Env: make(map[string]string),
	}
	m.Content.Applications = append(m.Content.Applications, app)
	return len(m.Content.Applications) - 1
}