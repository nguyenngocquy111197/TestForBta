package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

/*
NewModel public
*/
func NewModel(path string) *logrus.Entry {

	var (
		log    *logrus.Logger = logrus.New()
		output io.Writer      = os.Stdout
	)

	if len(path) > 0 {
		path, err := filepath.Abs(path)
		if err != nil {
			log.Error(err)
		} else {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				if err := os.MkdirAll(path, os.ModePerm); err != nil {
					log.Error(err)
				}
			}

			output = newRotateWriter(log, fmt.Sprintf("%s/activity.log", path))
			output = io.MultiWriter(output, os.Stdout)
		}
	}

	switch os.Getenv("LOG_MODE") {
	case "info":
		log.SetLevel(logrus.InfoLevel)
		break
	case "debug":
		log.SetLevel(logrus.DebugLevel)
		break
	case "trace":
		log.SetLevel(logrus.TraceLevel)
		break
	default:
		log.SetLevel(logrus.DebugLevel)
	}

	if os.Getenv("LOG_FORMAT") == "json" {
		formatter := new(logrus.JSONFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		log.SetFormatter(formatter)
	} else {
		formatter := new(logrus.TextFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		log.SetFormatter(formatter)
	}

	log.SetOutput(output)
	return logrus.NewEntry(log)
}

/*
NewRawLog ...
*/
func NewRawLog(path string) *logrus.Entry {
	var (
		formatter           = new(logrus.JSONFormatter)
		log                 = logrus.New()
		output    io.Writer = os.Stdout
	)

	formatter.TimestampFormat = "2006-01-02 15:04:05"

	if len(path) > 0 {
		path, err := filepath.Abs(path)
		if err != nil {
			//
		} else {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.MkdirAll(path, os.ModePerm)
			}
			output = newRotateWriter(log, fmt.Sprintf("%s/data-raw.log", path))
		}
	} else {
		logrus.Println("Do not log raw data")
		return logrus.NewEntry(log)
	}

	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(formatter)
	log.SetOutput(output)

	return logrus.NewEntry(log)
}

/*
NewRotateWriter create new write
*/
func newRotateWriter(log *logrus.Logger, output string) io.Writer {
	lumb := &lumberjack.Logger{
		Filename: output,
		MaxSize:  50, // megabytes
		// MaxBackups: 7,
		// MaxAge:    1, //days
		LocalTime: true,
		// Compress:  true, // disabled by default
	}

	go func() {
		for {
			if err := lumb.Rotate(); err != nil {
				log.Error(err)
			}
			var toNewDay time.Duration
			{
				yyyy, mm, dd := time.Now().Add(24 * time.Hour).Date()
				toNewDay = -time.Since(time.Date(yyyy, mm, dd, 0, 0, 0, 0, time.Local))
			}
			<-time.After(toNewDay)
			log.Warningln("Rotate newfile")
		}
	}()

	return lumb
}
