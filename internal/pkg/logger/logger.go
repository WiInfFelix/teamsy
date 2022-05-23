package logger

import (
	"io"
	"log"
	"os"
)

func InitLogger() error {

	log.SetFlags(2 | 3)

	log.Println("Starting logger....")

	logFile, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		return err
	}

	multiW := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiW)

	return nil
}
