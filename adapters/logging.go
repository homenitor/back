package adapters

import (
	"os"
	"strings"

	"github.com/homenitor/back/config"
	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:02/01/2006 15:04:05.000} ▶ %{level:.4s} %{color:reset}%{message}`,
)

type Logging struct {
	log *logging.Logger
}

func NewLogging() *Logging {
	configLogLevel := config.LogLevel()
	logLevel := getLogLevel(configLogLevel)
	loggingBackend := logging.NewLogBackend(os.Stderr, "", 0)

	formatterBackend := logging.NewBackendFormatter(loggingBackend, format)

	leveledBackend := logging.AddModuleLevel(formatterBackend)
	leveledBackend.SetLevel(logLevel, "")

	logging.SetBackend(leveledBackend)

	logger := logging.MustGetLogger("homenitor")

	return &Logging{
		log: logger,
	}
}

func getLogLevel(configLogLevel string) logging.Level {
	lowercasedConfigLogLevel := strings.ToLower(configLogLevel)
	levels := make(map[string]logging.Level, 0)
	levels["debug"] = logging.DEBUG
	levels["info"] = logging.INFO
	levels["error"] = logging.ERROR

	level, ok := levels[lowercasedConfigLogLevel]
	if !ok {
		panic(ErrUnknownLogLevel)
	}

	return level
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
