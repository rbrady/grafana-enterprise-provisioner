package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Tenants []TenantConfig `yaml:"tenants"`
}

type TenantConfig struct {
	ServiceURL   string `yaml:"service_url"`
	Tenant       string `yaml:"tenant"`
	AccessPolicy string `yaml:"access-policy"`
	Token        string `yaml:"token"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
