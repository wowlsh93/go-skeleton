/*
Copyright stick Corp. All Rights Reserved.

Written by HAMA
*/

package logging

import (
	"fmt"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/wowlsh93/go-skeleton/core/Aapp/config"
	"io/ioutil"
	"os"
)

var logger = logrus.New()

func GetLogger() *logrus.Logger {
	return logger
}

func loatefilehook(configpath string) logrus.Hook {

	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "Jan _2 15:04:05.000000000"
	Formatter.FullTimestamp = true
	Formatter.ForceColors = true

	rotateFileHook, err := NewRotateFileHook(RotateFileConfig{
		Filename:   configpath + "scanner.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     7,
		Level:      logrus.DebugLevel,
		Formatter:  logrus.Formatter(Formatter),
	})

	if err != nil {
		fmt.Println("rotateFileHook  error")
	}

	return rotateFileHook
}

func InitLog(conf *config.Configuration) {

	logger.AddHook(filename.NewHook())
	logger.SetFormatter(&logrus.JSONFormatter{})

	conf.Server.Logoutput = "both"

	if conf.Server.Logoutput == "both" {
		rotateFileHook := loatefilehook(conf.Server.Logpath)
		logger.AddHook(rotateFileHook)
		logger.SetOutput(os.Stdout)

	} else if conf.Server.Logoutput == "file" {
		rotateFileHook := loatefilehook(conf.Server.Logpath)
		logger.AddHook(rotateFileHook)
		logger.SetOutput(ioutil.Discard)

	} else {
		logger.SetOutput(os.Stdout)
	}

	logger.SetLevel(logrus.DebugLevel)

}
