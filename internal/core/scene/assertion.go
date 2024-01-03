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
		if s, ok := params[ck.Actually]; ok {
			ck.Actually = ToString(s)
		}
		if s, ok := params[ck.Expected]; ok {
			ck.Expected = ToString(s)
		}
		switch ck.AssertType {
		case Equal:
			if ck.Actually != ck.Expected {
				return fmt.Errorf("【%s】 %s: %v is not equal to %v", ck.Name, ck.ErrorMsg, ck.Actually, ck.Expected)
			}
		case NotEqual:
			if ck.Actually == ck.Expected {
				return fmt.Errorf("【%s】 %s: %v is equal to %v", ck.Name, ck.ErrorMsg, ck.Actually, ck.Expected)
			}
		}
	}
	return nil
}
