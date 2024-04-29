package log

import (
	"fmt"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	stderr = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	stdout = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger().Level(zerolog.InfoLevel)
)

func SetDebugMode() {
	stderr = stderr.Level(zerolog.DebugLevel)
	stdout = stdout.Level(zerolog.DebugLevel)
}

var (
	prefix      string // 日志前缀
	prefixMutex *sync.Mutex
)

func SetPrefix(str string) {
	prefix = str
}

func Prefix() string {
	return prefix
}

func logf(level zerolog.Level, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)

	if prefix != "" {
		msg = prefix + " " + msg
	}

	switch level {
	case zerolog.DebugLevel:
		stderr.Debug().Msg(msg)
	case zerolog.InfoLevel:
		stdout.Info().Msg(msg)
	case zerolog.ErrorLevel:
		stderr.Error().Msg(msg)
	default:
		stderr.Debug().Msg(msg)
	}
}

func Errorf(format string, v ...interface{}) {
	logf(zerolog.ErrorLevel, format, v...)
}

func Infof(format string, v ...interface{}) {
	logf(zerolog.InfoLevel, format, v...)
}

func Debugf(format string, v ...interface{}) {
	logf(zerolog.DebugLevel, format, v...)
}

func Err(err error) {
	message := err.Error()
	Errorf(message)
}
