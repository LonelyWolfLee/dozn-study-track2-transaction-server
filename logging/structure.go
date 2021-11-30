package logging

import (
	"io/ioutil"
	"log"
	"os"
)

type logStruct struct {
	trace *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func get() logStruct {
	l := logStruct{
		log.New(ioutil.Discard, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return l
}

var logger = get()
