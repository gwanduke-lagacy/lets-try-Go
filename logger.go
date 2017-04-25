package main

import (
	"io"
	"log"
	"os"
)

var myStdLogger *log.Logger
var myFileLogger *log.Logger

// PrepareFileToLogging 로깅할 파일을 준비합니다
func PrepareFileToLogging(filepath string) (fp *os.File, err error) {
	fpLog, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	return fpLog, err
}

// WriteDoubleLogging 지정한 파일과, 표준 출력으로 로깅 진행
func WriteDoubleLogging(fp *os.File, logText string) {
	multiWriter := io.MultiWriter(fp, os.Stdout)
	log.SetOutput(multiWriter)
	log.Println(logText)
	log.SetOutput(os.Stdout)
}
