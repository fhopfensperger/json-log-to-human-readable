package cmd

import (
	"fmt"
	"math"
	"time"
)

type LogMessage struct {
	Timestamp  string    `json:"timestamp"`
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	Exception  Exception `json:"exception,omitempty"`
	LoggerName string    `json:"loggerName"`
}

type SpringBootLogMessage struct {
	Timestamp  string `json:"@timestamp"`
	Level      string `json:"level"`
	Message    string `json:"message"`
	Exception  string `json:"stack_trace,omitempty"`
	LoggerName string `json:"logger_name"`
}

type GoZapLogMessage struct {
	Level      string  `json:"level"`
	Timestamp  float64 `json:"ts"`
	Logger     string  `json:"logger"`
	Message    string  `json:"msg"`
	Controller string  `json:"controller,omitempty"`
	Request    string  `json:"request,omitempty"`
	Error      string  `json:"error,omitempty"`
	Stacktrace string  `json:"stacktrace,omitempty"`
}

type CommonLogMessage interface {
	print()
}

type Exception struct {
	RefId         int      `json:"refId"`
	ExceptionType string   `json:"exceptionType"`
	Message       string   `json:"message"`
	CausedBy      CausedBy `json:"causedBy"`
	Frames        *[]Frame `json:"frames"`
}

type CausedBy struct {
	Exception *Exception `json:"exception,omitempty"`
}

type Frame struct {
	Class  string `json:"class"`
	Method string `json:"method"`
	Line   int    `json:"line"`
}

func (lm *LogMessage) print() {
	fmt.Printf("%v %v\t %v\t%v\n", lm.Level, lm.Timestamp, lm.LoggerName, lm.Message)

	// log message contains an error error
	if lm.Exception != (Exception{}) {
		lm.Exception.print()
	}
}

func (ex *Exception) print() {
	fmt.Printf("Caused by: %v. %s:\n", ex.ExceptionType, ex.Message)
	for _, frame := range *ex.Frames {
		fmt.Printf("\t at %s(%s:%v)\n", frame.Method, frame.Class, frame.Line)
	}
	if ex.CausedBy != (CausedBy{}) {
		ex.CausedBy.Exception.print()
	}
}

func (alm *SpringBootLogMessage) print() {
	fmt.Printf("%v %v\t %v\t%v\n", alm.Level, alm.Timestamp, alm.LoggerName, alm.Message)
	// log message contains an error error
	if alm.Exception != "" {
		fmt.Printf("Exception: %s", alm.Exception)
	}
}

func (glm *GoZapLogMessage) print() {
	sec, dec := math.Modf(glm.Timestamp)
	timestamp := time.Unix(int64(sec), int64(dec*(1e9)))
	fmt.Printf("%v %v\t %v\tmsg: %v\tcontroller: %v\t request: %v\n", glm.Level, timestamp, glm.Logger, glm.Message, glm.Controller, glm.Request)
	// log message contains an error error
	if glm.Error != "" {
		fmt.Printf("error: %s", glm.Error)
		fmt.Printf("stacktrace: %s\n", glm.Stacktrace)
	}
}
