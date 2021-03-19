package log

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/debeando/lightflow/config"
	"github.com/sirupsen/logrus"
)

var debug bool

type myFormatter struct {
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var data string

	if len(entry.Data) > 0 {
		for k, v := range entry.Data {
			var level string

			switch entry.Level.String() {
			case "info":
				level = "I"
			case "debug":
				level = "D"
			case "warning":
				level = "W"
			case "error":
				level = "E"
			}

			data += fmt.Sprintf(
				"%s %s %s%s\n",
				entry.Time.Format("2006-01-02 15:04:05"),
				level,
				entry.Message,
				fmt.Sprintf(" %s: %#v", k, v),
			)
		}
	} else {
		var level string

		switch entry.Level.String() {
		case "info":
			level = "I"
		case "debug":
			level = "D"
		case "warning":
			level = "W"
		case "error":
			level = "E"
		}

		data = fmt.Sprintf(
			"%s %s %s\n",
			entry.Time.Format("2006-01-02 15:04:05"),
			level,
			entry.Message,
		)
	}

	return []byte(data), nil
}

func init() {
	logrus.SetFormatter(new(myFormatter))

	logrus.SetLevel(logrus.ErrorLevel)

	if flag.Lookup("test.v") != nil {
		logrus.SetOutput(ioutil.Discard)
	} else {
		logrus.SetOutput(os.Stderr)
	}

	logrus.SetLevel(logrus.InfoLevel)

	if config.Load().General.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Info(m string, f map[string]interface{}) {
	logrus.WithFields(f).Info(m)
}

func Warning(m string, f map[string]interface{}) {
	logrus.WithFields(f).Warning(m)
}

func Error(m string, f map[string]interface{}) {
	logrus.WithFields(f).Error(m)
}

func Debug(m string, f map[string]interface{}) {
	if flag.Lookup("debug") != nil && flag.Lookup("debug").Value.(flag.Getter).Get().(bool) {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.WithFields(f).Debug(m)
}
