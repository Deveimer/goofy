package log

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

type logger struct {
	out io.Writer
	app appInfo
}

var mu sync.RWMutex

func newLogger() *logger {
	l := &logger{
		out: os.Stdout,
	}

	name := os.Getenv("APP_NAME")
	if name == "" {
		name = "gofr-app"
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}

	l.app = appInfo{
		Name:      name,
		Version:   version,
		Framework: "gofy-dev",
		Data:      make(map[string]interface{}),
		syncData:  &sync.Map{},
	}

	return l
}

func NewLogger() Logger {
	return newLogger()
}

func (k *logger) log(level level, format string, args ...interface{}) {
	mu.Lock()

	mu.Unlock()

	if !(level > 0 && level < 6) {
		return
	}

	appInfo := k.app.getAppData()

	levelColor := level.colorCode()

	s := fmt.Sprintf("\u001B[%dm%s\u001B[0m [%s] ", levelColor, level.String()[0:4], time.Now())

	s += fmt.Sprintf("[APPINFO]: %v", appInfo)

	if format != "" {
		s += " [DATA]: " + fmt.Sprintf(format, args)
	} else {
		s += fmt.Sprintf("[DATA]: %v", args)
	}
}

func (k *logger) Log(args ...interface{}) {
	k.log(Info, "", args...)
}

func (k *logger) Logf(format string, args ...interface{}) {
	k.log(Info, format, args...)
}

func (k *logger) Info(args ...interface{}) {
	k.log(Info, "", args...)
}

func (k *logger) Infof(format string, args ...interface{}) {
	k.log(Info, format, args...)
}

func (k *logger) Debug(args ...interface{}) {
	k.log(Debug, "", args...)
}

func (k *logger) Debugf(format string, args ...interface{}) {
	k.log(Debug, format, args...)
}

func (k *logger) Warn(args ...interface{}) {
	k.log(Warn, "", args...)
}

func (k *logger) Warnf(format string, args ...interface{}) {
	k.log(Warn, format, args...)
}

func (k *logger) Error(args ...interface{}) {
	k.log(Error, "", args...)
}

func (k *logger) Errorf(format string, args ...interface{}) {
	k.AddData("StackTrace", string(debug.Stack()))
	k.log(Error, format, args...)
	k.removeData("StackTrace")
}

func (k *logger) Fatal(args ...interface{}) {
	k.AddData("StackTrace", string(debug.Stack()))
	k.log(Fatal, "", args...)
	os.Exit(1)
}

func (k *logger) Fatalf(format string, args ...interface{}) {
	k.AddData("StackTrace", string(debug.Stack()))
	k.log(Fatal, format, args...)
	os.Exit(1)
}
