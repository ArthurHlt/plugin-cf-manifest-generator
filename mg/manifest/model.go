package manifest

type Manifest struct {
	ManifestPath string
	Content      *ManifestContent
}

type ManifestContent struct {
	Applications []Application `yaml:"applications"`
	Inherit      string `yaml:"inherit,omitempty"`
}
type Application struct {
	Name        string `yaml:"name"`
	Memory      string `yaml:"memory"`
	Instances   int `yaml:"instances"`
	Path        string `yaml:"path"`
	Buildpack   string `yaml:"buildpack,omitempty"`
	Command     string `yaml:"command,omitempty"`
	Domains     []string `yaml:"domains,omitempty"`
	Hosts       []string `yaml:"hosts,omitempty"`
	RandomRoute bool `yaml:"random-route,omitempty"`
	Timeout     int `yaml:"timeout,omitempty"`
	Env         map[string]string  `yaml:"env,omitempty"`
	NoRoute     bool `yaml:"no-route,omitempty"`
	Services    []string `yaml:"services,omitempty"`
	DiskQuota   string `yaml:"disk_quota,omitempty"`
	Stack       string `yaml:"stack,omitempty"`
	NoHostname  bool `yaml:"no-hostname,omitempty"`
	Host        string `yaml:"host"`
	Domain      string `yaml:"domain,omitempty"`
	Inherit     string `yaml:"inherit,omitempty"`
}