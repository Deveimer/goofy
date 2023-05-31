package log

type Logger interface {
	Log(args ...interface{})
	Logf(format string, a ...interface{})
	Debug(args ...interface{})
	Debugf(format string, a ...interface{})
	Info(args ...interface{})
	Infof(format string, a ...interface{})
	Warn(args ...interface{})
	Warnf(format string, a ...interface{})
	Error(args ...interface{})
	Errorf(format string, a ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, a ...interface{})
	AddData(key string, value interface{})
}
