package configv1

type DomainList struct {
	Name    string   `yaml:"name"`
	Domains []string `yaml:"domains"`
}

type ConfigV1 struct {
	Version string       `yaml:"config_version"`
	Groups  []DomainList `yaml:"domain_groups"`
}
