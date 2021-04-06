package libraries

type Logging interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}
