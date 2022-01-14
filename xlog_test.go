package xlog

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	xl := NewXLogger().BuildFile("access", 2*time.Second)
	xl.Info("name ,j,,,,,")
	xl.BuildFormatter(FORMAT_JSON)
	time.Sleep(10 * time.Second)
}
