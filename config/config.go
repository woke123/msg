package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var C *Config
var config_path = "config/config.yaml"

type EmailData struct {
	Email    string `json:"email"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
}

type Msg struct {
	Email_sent string `json:"email_sent"`
	Email_fail string `json:"email_fail"`
}

type Email struct {
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Config struct {
	Redis  Redis  `yaml:"redis"`
	Email  Email  `yaml:"email"`
	Msg    Msg    `yaml:"msg"`
}

type Redis struct {
	Host  string `yaml:"host"`
	Auth  string `yaml:"auth"`
}

func InitMsg() error {
	C = new(Config)
	yamlFile, err := ioutil.ReadFile(config_path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, C)
	if err != nil {
		return err
	}
	return nil
}