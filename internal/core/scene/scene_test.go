package scene

import (
	"testing"
)

func TestLoadYamlFile(t *testing.T) {
	var scene Scene
	err := LoadYamlFile("./test_data.yaml", &scene)
	if err != nil {
		t.Errorf("load yaml file error: %s", err)
		return
	}
	t.Log(scene)
}
