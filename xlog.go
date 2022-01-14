package xlog

import (
	"fmt"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
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

/**
FORMAT_JSON = 1
FORMAT_TXT  = 2
*/
type FORMAT_LOG int

/**
STD=1
FILE=2
HTTP=3
*/
type OUT_TYPE int

type XLogger interface {
}

type XLog struct {
	*logrus.Logger
	outType OUT_TYPE
}

func NewXLogger() *XLog {
	lgs := logrus.New()
	lgs.SetFormatter(&logrus.JSONFormatter{})
	lgs.SetOutput(os.Stdout)
	return &XLog{Logger: lgs}
}
func (log *XLog) BuildOutType(out OUT_TYPE) {
	log.outType = out
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

func (log *XLog) BuildFormatter(format FORMAT_LOG) *XLog {
	if format == FORMAT_JSON {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	if format == FORMAT_TXT {
		log.SetFormatter(&logrus.TextFormatter{})
	}
	return log
}

/**
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.WarnLevel)
*/
