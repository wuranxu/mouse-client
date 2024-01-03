package scene

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateStringRender(t *testing.T) {
	result, err := Render(`fggrfgrf{{uuid .honey.ff}}`, map[string]any{"honey": map[string]any{"ff": "lixiaoyao"}})
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "lixiaoyao", result, "should be rendered as 【lixiaoyao】")
}
