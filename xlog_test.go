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

	ss := []*M{&M{Name: "sunhongtao", Id: 1}, &M{Name: "wang", Id: 20}}
	xl.PrintlnSlice(ss)
	xl.BuildFormatter(FORMAT_JSON)
	time.Sleep(10 * time.Second)
}
