package log

import (
	"github.com/Sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func (l *Logger) Log(args ...interface{}) error {
	var msg interface{}
	m := make(map[string]interface{})
	for i := 0; i < len(args); i += 2 {
		k := args[i].(string)
		v := args[i+1]
		if k == "msg" {
			msg = v
		} else {
			m[k] = v
		}
	}
	l.WithFields(logrus.Fields(m)).Debug(msg)
	return nil
}

func New() *Logger {
	return &Logger{Logger: logrus.New()}
}
