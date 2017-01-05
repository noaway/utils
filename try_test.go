package utils

import (
	"testing"
)

type student struct {
	Id   int
	Name string
	Age  int
}

func TestTry(t *testing.T) {
	var stu interface{}
	Trycatch(func() {
		s := stu.(student).Name
		t.Log(s)
	}, func(e interface{}) {
		t.Log(e)
	})
}
