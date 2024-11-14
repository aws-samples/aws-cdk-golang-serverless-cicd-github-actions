package config

import (
	"cdk-infra/packages/types"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)
func Conf() types.Config {

	env := os.Getenv("ENV")
	var data []byte
	var err error
	switch env {
	case "sample":
		data, err = os.ReadFile("input/input.yaml")
	default:
		log.Fatal("no ENV set")
	}

	if err != nil {
		log.Fatalf("error reading yaml file:%v", err)
	}
	var config types.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("error unmarshal yaml:%v", err)
	}
	return config
}
