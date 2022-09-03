package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"iot-master-gateway/config"
	"os"
)

func Open(opts *config.Log) error {
	if opts.Text {
		logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: opts.Format})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: opts.Format})
	}

	// LogOutput to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	if opts.Output.Filename == "" {
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetOutput(&lumberjack.Logger{
			Filename:   opts.Output.Filename,
			MaxSize:    opts.Output.MaxSize,
			MaxAge:     opts.Output.MaxAge,
			MaxBackups: opts.Output.MaxBackups,
			LocalTime:  true,
		})
	}

	// Only log the warning severity or above.
	level, err := logrus.ParseLevel(opts.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)
	return nil
}

type Fields = logrus.Fields

var WithFields = logrus.WithFields

var Trace = logrus.Trace

var Tracef = logrus.Tracef

var Warn = logrus.Warn

var Warnf = logrus.Warnf

var Info = logrus.Info

var Infof = logrus.Infof

var Error = logrus.Error

var Errorf = logrus.Errorf

var Fatal = logrus.Fatal

var Fatalf = logrus.Fatalf

var Println = logrus.Println

var Print = logrus.Print

var Printf = logrus.Printf
