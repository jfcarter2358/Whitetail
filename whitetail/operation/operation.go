package operation

import (
	"io/ioutil"
	"os"
	"whitetail/analyst"
	"whitetail/config"
	"whitetail/dashboard"
	"whitetail/observer"
	"whitetail/probe"

	"gopkg.in/yaml.v2"
)

type Operation struct {
	Probes     map[string]probe.Probe         `yaml:"probes"`
	Analysts   map[string]analyst.Analyst     `yaml:"analysts"`
	Dashboards map[string]dashboard.Dashboard `yaml:"dashboards"`
	Observers  map[string]observer.Observer   `yaml:"observers"`
}

var Operations Operation

func LoadOperation() error {
	yamlFile, err := ioutil.ReadFile(config.Config.OperationPath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Operations)
	if err != nil {
		return err
	}

	return nil
}

func SaveOperation() error {
	data, err := yaml.Marshal(Operations)
	if err != nil {
		return err
	}

	err = os.WriteFile(config.Config.OperationPath, data, 0777)
	if err != nil {
		return err
	}

	return nil
}
