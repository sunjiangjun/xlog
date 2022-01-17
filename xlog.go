package xlog

import (
	"encoding/json"
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

const (

	//FatalLevel = // level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = iota + 1
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type Level int

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
	PrintlnSlice(args interface{})
}

type XLog struct {
	XLogger
	*logrus.Logger
	outType OUT_TYPE
}

func (x *XLog) PrintlnSlice(args interface{}) {
	body, err := json.Marshal(args)
	if err != nil {
		x.Error(err)
	}
	x.Println(string(body))
}

func NewXLogger() *XLog {
	lgs := logrus.New()
	lgs.SetFormatter(&logrus.JSONFormatter{})
	lgs.SetOutput(os.Stdout)
	return &XLog{Logger: lgs}
}
func (log *XLog) BuildOutType(out OUT_TYPE) *XLog {
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

func (log *XLog) BuildFormatter(format FORMAT_LOG) *XLog {
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

/**
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.WarnLevel)
*/
