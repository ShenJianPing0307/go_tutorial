package abc

import (
	"testing"
)

func TestCalAdd(t *testing.T) {
	res := calAdd(1, 2)
	if res != 3 {
		t.Fatalf("calAdd函数错误，期望值是%v，返回值是%v", 3, res)
	} else {
		t.Logf("calAdd函数正确，期望值是%v，返回值是%v", 3, res)

	}
}
