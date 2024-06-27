package system

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// GeneralLogger exported
var Logger *log.Logger

// ErrorLogger exported
var ErrorLog *log.Logger

func init() {
	absPath, err := filepath.Abs("./log")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/general-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	mw := io.MultiWriter(os.Stdout, generalLog)
	Logger = log.New(mw, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	//Logger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
