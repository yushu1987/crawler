package lib

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"time"
)

// export outside
var Log = &logrus.Entry{
	Logger: loger,
}

// inside
var loger = logrus.New()

func InitLog() {
	// 自带切分日志
	InitLoggerLogrotate(
		Config.GetString("log.logpath"),
		Config.GetString("log.logfile"),
		Config.GetString("log.level"),
		Config.GetString("log.format"),
		Config.GetDuration("log.loglifetime")*time.Hour,
		Config.GetDuration("log.logrotation")*time.Hour,
	)
	// ignore terminal stdout
	loger.SetOutput(ioutil.Discard)

	// 需要linux logrotate来切分
	// Log.Logger = NewLogger(
	// 	Config.GetString("log.logpath"),
	// 	Config.GetString("log.logfile"),
	// 	Config.GetString("log.level"),
	// 	Config.GetString("log.format"),
	// )
}

func InitLoggerLogrotate(logPath, logFileName, level, format string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(logPath+logFileName), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime),    // 日志切割时间间隔
	)
	if err != nil {
		Log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
		os.Exit(99)
	}

	switch level {
	case "debug":
		loger.SetLevel(logrus.DebugLevel)
	case "info":
		loger.SetLevel(logrus.InfoLevel)
	case "warn":
		loger.SetLevel(logrus.WarnLevel)
	case "error":
		loger.SetLevel(logrus.ErrorLevel)
	default:
		loger.SetLevel(logrus.InfoLevel)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})

	if format == "json" {
		lfHook.SetFormatter(&logrus.JSONFormatter{})
	} else {
		lfHook.SetFormatter(&logrus.TextFormatter{})
	}

	loger.AddHook(lfHook)
}

// need linux logrotate process split log
func NewLogger(logPath, fileName, level, typeof string) *logrus.Logger {
	file := logPath + fileName
	cLog := logrus.New()
	fileFd, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	// Lshortfile
	if err != nil {
		fmt.Printf("config local file system logger error. %+v \n", errors.WithStack(err))
		os.Exit(99)
	}
	cLog.Out = fileFd

	switch typeof {
	case "text":
		// Text
		cLog.Formatter = &logrus.TextFormatter{}
	default:
		// Json
		cLog.Formatter = &logrus.JSONFormatter{}
	}

	levelFlag, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("parse log level faild, error: %+v \n", errors.WithStack(err))
		levelFlag = logrus.InfoLevel
	}
	cLog.Level = levelFlag
	return cLog
}
