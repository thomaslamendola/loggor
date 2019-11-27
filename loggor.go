package loggor

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var _filename string
var _logBasePath string
var _machineName string
var _logger string
var _source string

type log struct {
	Time        time.Time `json:"time"`
	Message     string    `json:"message"`
	Level       string    `json:"level"`
	Logger      string    `json:"logger"`
	MachineName string    `json:"machineName"`
	Source      string    `json:"source"`
}

// Private methods

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generic(level string, message string) {
	var err error
	localfilename := _logBasePath + time.Now().Format("20060102") + "_" + _filename
	f, err := os.OpenFile(localfilename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	log := log{
		Time:        time.Now(),
		Level:       level,
		Logger:      _logger,
		MachineName: _machineName,
		Source:      _source,
		Message:     message,
	}
	stringlog, err := json.Marshal(log)
	check(err)
	_, err = f.Write(stringlog)
	check(err)
	_, err = f.WriteString("\n")
	check(err)
	_, err = fmt.Println(stringlog)
	check(err)
}

//Public methods

//Initialize a loggor instance
func Initialize(logBasePath string, filename string, source string, machineName string) {
	if !strings.Contains(filename, ".") {
		filename += ".json"
	}
	_logBasePath = logBasePath
	_filename = filename
	_machineName = machineName
	_source = source
	_logger = "Loggor"
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
