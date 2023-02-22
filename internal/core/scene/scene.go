// Package scene include test file(yaml) and how to parse yaml
package scene

// Step test step
type Step struct {
	Name        string            `yaml:"name" json:"name"`
	Url         string            `yaml:"url" json:"url"`
	Headers     map[string]string `json:"headers" yaml:"headers"`
	Method      string            `yaml:"method" json:"method"`
	Body        string            `json:"body" yaml:"body"`
	StatusCheck bool              `json:"status_check" yaml:"status_check"`
	Out         []*Out            `json:"out" yaml:"out"`
	Check       []*Check          `json:"check" yaml:"check"`
	Query       map[string]string `json:"query" yaml:"query"`
	Timeout     int               `json:"timeout"`
}

// Out extract parameters
type Out struct {
	Name string `json:"name" yaml:"name"`
	// parameters from
	From        FromType    `json:"from" yaml:"from"`
	ExtractType ExtractType `json:"extractType" yaml:"extract_type"`
	Expression  string      `json:"expression" yaml:"expression"`
	Variable    string      `json:"variable" yaml:"variable"`
}

// Check assert for the step
type Check struct {
	Name       string    `json:"name" yaml:"name"`
	AssertType Assertion `json:"assert_type" yaml:"assert_type"`
	Expected   string    `json:"expected" yaml:"expected"`
	Actually   string    `json:"actually" yaml:"actually"`
	ErrorMsg   string    `json:"error_msg" yaml:"error_msg"`
	Disabled   bool      `json:"disabled" yaml:"disabled"`
}

// Scene a testing scene
type Scene struct {
	Name  string  `json:"name" yaml:"name"`
	Steps []*Step `json:"steps" yaml:"steps"`
}
