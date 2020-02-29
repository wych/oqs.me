package logger

import (
	"log"
	"os"
	"path"

	"oqs.me/config"
)

var (
	NormalLogger *log.Logger
	ErrorLogger  *log.Logger
)

func LoggerInit() {
	nDir := path.Dir(config.Conf.Logger.Normal)
	if !isExist(nDir) {
		os.Mkdir(nDir, 755)
	}
	eDir := path.Dir(config.Conf.Logger.Error)
	if !isExist(eDir) {
		os.Mkdir(eDir, 755)
	}
}

func initNormalLogger() {
	f, err := os.OpenFile(config.Conf.Logger.Normal, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Normal logger init error: %s", err.Error())
	}
	NormalLogger = log.New(f, "Log| ", 0)
}

func initErrorLogger() {
	f, err := os.OpenFile(config.Conf.Logger.Error, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error logger init error: %s", err.Error())
	}
	ErrorLogger = log.New(f, "Error| ", 0)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
