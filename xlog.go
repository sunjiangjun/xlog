package xlog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

const (
	FORMAT_JSON = 1
	FORMAT_TXT  = 2
)
const (
	STD  = 1
	FILE = 2
	HTTP = 3
)

type Level int

//FormatLog
/**
FORMAT_JSON = 1
FORMAT_TXT  = 2
*/
type FormatLog int

//OutType
/**
STD=1
FILE=2
HTTP=3
*/
type OutType int

type XLogger interface {
	PrintlnSlice(args interface{})
}

type XLog struct {
	XLogger
	*logrus.Logger
	outType OutType
}

func (log *XLog) PrintlnSlice(args interface{}) {
	body, err := json.Marshal(args)
	if err != nil {
		log.Error(err)
	}
	log.Println(string(body))
}

func NewXLogger() *XLog {
	lgs := logrus.New()
	lgs.SetFormatter(&logrus.JSONFormatter{})
	lgs.SetOutput(os.Stdout)
	return &XLog{Logger: lgs}
}
func (log *XLog) BuildOutType(out OutType) *XLog {
	log.outType = out
	return log
}
func (log *XLog) BuildFile(prefix string, RotationTime time.Duration) *XLog {
	if log.outType == FILE {
		fileName := fmt.Sprintf("%v_log_%v", prefix, "%Y%m%d%H%M")
		rl, err := rotate.New(fileName, rotate.WithRotationTime(RotationTime))
		if err != nil {
			panic(err)
		}
		log.SetOutput(rl)
	}
	return log
}

func (log *XLog) BuildFormatter(format FormatLog) *XLog {
	if format == FORMAT_JSON {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	if format == FORMAT_TXT {
		log.SetFormatter(&logrus.TextFormatter{})
	}
	return log
}

func (log *XLog) BuildLevel(l Level) *XLog {
	log.SetLevel(logrus.Level(l))
	return log
}
