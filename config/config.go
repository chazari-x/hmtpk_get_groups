package config

type Schedule struct {
	Groups []struct {
		ID   int    `yaml:"id"`
		Name string `yaml:"name"`
	} `yaml:"groups"`

	Teachers []struct {
		Name string `yaml:"name"`
	} `yaml:"teachers"`
}
