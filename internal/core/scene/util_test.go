package scene

import (
	"fmt"
	"testing"
)

func TestToBytes(t *testing.T) {
	a := ToBytes("hahaha")
	fmt.Println(a)
	result := ToString(a)
	fmt.Println(result)
}
