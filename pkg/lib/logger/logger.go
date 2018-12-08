package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Info(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info(args...)
}

func Error(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Error(args...)
}

func Trace(index string, source string, traceid string, err string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"index":   index,
		"source":  source,
		"traceid": traceid,
		"err":     err,
	}).Info(args...)
}
