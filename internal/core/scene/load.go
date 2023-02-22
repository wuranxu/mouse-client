package scene

import (
	json "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func LoadYamlFile(file string, scene *Scene) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, scene)
}

func LoadJSONFile(file string, scene *Scene) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, scene)
}

// Load load scene data
func Load(data []byte, scene *Scene) error {
	return yaml.Unmarshal(data, scene)
}
