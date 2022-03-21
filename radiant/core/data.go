package core

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Server map[string]string `yaml:"Server"`
	Error  map[string]string `yaml:"Error"`
	Asset  map[string]string `yaml:"Asset"`
	Database map[string]string `yaml:"Database"`
	Sites map[string][]string `yaml:"Sites"`
	Siteallowance map[string]string `yaml:"SiteAllowance"`
	Acceptedlanguage map[string][]string `yaml:"Language"`
}

type DbConfiguration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

type SiteInfo struct {
	Host string
	Name string
}

var Site SiteInfo

var Language []string

var Configure Conf

var Database DbConfiguration

func init() {
	var err error
	// Reading Config file
	configYml, err := ioutil.ReadFile("radiant/core/config/conf.yaml")
	Configure = Conf{}
	err = yaml.Unmarshal(configYml, &Configure)
	if err != nil {
		log.Fatal(err)
	}
	Language = Configure.Acceptedlanguage["acceptedlanguage"]
}

func GetDbConfig() DbConfiguration {
	Database.DB_HOST=Configure.Database["DB_HOST"]
	Database.DB_NAME=Configure.Database["DB_NAME"]
	Database.DB_PORT=Configure.Database["DB_PORT"]
	Database.DB_USERNAME=Configure.Database["DB_USERNAME"]
	Database.DB_PASSWORD=Configure.Database["DB_PASSWORD"]
	return Database
}
