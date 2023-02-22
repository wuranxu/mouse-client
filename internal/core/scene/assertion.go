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

func Assert(check []*Check) error {
	for _, ck := range check {
		if ck.Disabled {
			continue
		}
		switch ck.AssertType {
		case Equal:
			if ck.Actually != ck.Expected {
				return fmt.Errorf("【%s】 %s: %v != %v", ck.Name, ck.ErrorMsg, ck.Expected, ck.Actually)
			}
		}
	}
	return nil
}
