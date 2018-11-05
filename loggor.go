package loggor

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

var _filename string

type log struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Level     string    `json:"level"`
}

// Private methods

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generic(level string, message string) {
	var err error
	localfilename := time.Now().Format("20060102") + "_" + _filename
	f, err := os.OpenFile(localfilename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	log := log{
		Message:   message,
		Level:     level,
		Timestamp: time.Now(),
	}
	stringlog, err := json.Marshal(log)
	check(err)
	_, err = f.Write(stringlog)
	check(err)
	_, err = f.WriteString("\n")
	check(err)
}

//Public methods

//Initialize a loggor instance
func Initialize(filename string) {
	if !strings.Contains(filename, ".") {
		filename += ".log"
	}
	_filename = filename
}

//Info prints an info log
func Info(message string) {
	generic("INFO", message)
}

//Warn prints a warn log
func Warn(message string) {
	generic("WARN", message)
}

//Error prints an error log
func Error(message string) {
	generic("ERROR", message)
}
