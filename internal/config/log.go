package config

type LogConf struct {
	Level uint32 `yaml:"level"`
	Path  string `yaml:"path"`
}
