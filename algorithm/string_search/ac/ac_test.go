package ac

import (
	"strings"
	"testing"
)

func TestAhoCorasickAutomation_FindAll(t *testing.T) {

	kw := "有什么区别,么区别,区别"

	ac := NewAC(strings.Split(kw, ","))

	result := ac.FindAll("有什么区别")
	t.Log(result)
}