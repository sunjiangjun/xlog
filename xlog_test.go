package xlog

import (
	"testing"
	"time"
)

type M struct {
	Name string
	Id   int
}

func Test(t *testing.T) {
	xl := NewXLogger().BuildFile("access", 2*time.Second)

	ss := []*M{{Name: "vivi", Id: 1}, {Name: "wang", Id: 21}}
	xl.PrintlnSlice(ss)
	xl.BuildFormatter(FORMAT_JSON)
	time.Sleep(15 * time.Second)
}
