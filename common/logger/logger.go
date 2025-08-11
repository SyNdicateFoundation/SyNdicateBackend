package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func SetupLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&customFormat{})
	Logger.SetLevel(logrus.DebugLevel)
}

type customFormat struct{}

func (F *customFormat) Format(ent *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[SyNdicateBackend] [%s] %s\n", ent.Level.String(), ent.Message)), nil
}
