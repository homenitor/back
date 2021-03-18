package adapters

import (
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:02/01/2006 15:04:05.000} ▶ %{level:.4s} %{color:reset}%{message}`,
)

type Logging struct {
	log *logging.Logger
}

func NewLogging() *Logging {
	loggingBackend := logging.NewLogBackend(os.Stderr, "", 0)

	formatterBackend := logging.NewBackendFormatter(loggingBackend, format)

	leveledBackend := logging.AddModuleLevel(formatterBackend)
	leveledBackend.SetLevel(logging.DEBUG, "")

	logging.SetBackend(leveledBackend)

	logger := logging.MustGetLogger("homenitor")

	return &Logging{
		log: logger,
	}
}

func (l *Logging) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *Logging) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *Logging) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *Logging) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *Logging) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *Logging) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}
