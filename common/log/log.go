package log

// NOTES:
// - Quitar el flag y usar el common.IsArgDefined()

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/debeando/lightflow/config"

	"github.com/sirupsen/logrus"
)

var debug bool

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		DisableQuote:    true,
	})

	if flag.Lookup("test.v") != nil {
		logrus.SetOutput(ioutil.Discard)
	} else {
		logrus.SetOutput(os.Stdout)
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
