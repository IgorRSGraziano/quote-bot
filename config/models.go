package config

type Config struct {
	Group   string   `yaml:"group"`
	Command string   `yaml:"command"`
	Quotes  []string `yaml:"quotes"`
}
