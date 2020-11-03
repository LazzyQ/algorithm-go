package ac

import "testing"

func TestAhoCorasickAutomation_FindAll(t *testing.T) {
	ac := NewAC([]string{"明星", "林俊杰"})

	result := ac.FindAll("今晚出场的明星有周杰伦，林俊杰，陈奕迅")
	t.Log(result)
}