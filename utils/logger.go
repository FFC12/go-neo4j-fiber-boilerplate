package utils

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLogger(toFile bool) {
	if toFile {
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
		WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
		ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	}
}
