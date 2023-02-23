package scene

import (
	"fmt"
)

type Assertion string

const (
	Equal    Assertion = "equal"
	NotEqual Assertion = "not_equal"
	Contain  Assertion = "contain"
)

func Assert(check []*Check, params map[string][]byte) error {
	for _, ck := range check {
		if ck.Disabled {
			continue
		}
		switch ck.AssertType {
		case Equal:
			if s, ok := params[ck.Actually]; ok {
				ck.Actually = ToString(s)
			}
			if s, ok := params[ck.Expected]; ok {
				ck.Expected = ToString(s)
			}
			if ck.Actually != ck.Expected {
				return fmt.Errorf("【%s】 %s: %v != %v", ck.Name, ck.ErrorMsg, ck.Expected, ck.Actually)
			}
		}
	}
	return nil
}
