package common

import (
	std "gitlab.wallstcn.com/wscnbackend/ivankastd"
)

var (
	GlobalConf *Config
)

type Config struct {
	Micro     std.ConfigService `yaml:"micro"`
	Log       std.ConfigLog     `yaml:"log"`
	Bind      string            `yaml:"bind"`
	CertPem   string            `yaml:"certpem"`
	KeyPem    string            `yaml:"keypem"`
}

func LoadConfig(filePath string) {
	println("loading config")
	GlobalConf = &Config{}
	std.LoadConf(GlobalConf, filePath)
}

func Initalise() {
	std.InitLog(GlobalConf.Log)

}

