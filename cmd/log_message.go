package cmd

import "fmt"

type LogMessage struct {
	Timestamp  string    `json:"timestamp"`
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	Exception  Exception `json:"exception,omitempty"`
	LoggerName string    `json:"loggerName"`
}

type AlternativeLogMessage struct {
	Timestamp  string `json:"@timestamp"`
	Level      string `json:"level"`
	Message    string `json:"message"`
	Exception  string `json:"stack_trace,omitempty"`
	LoggerName string `json:"logger_name"`
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

func (alm *AlternativeLogMessage) print() {
	fmt.Printf("%v %v\t %v\t%v\n", alm.Level, alm.Timestamp, alm.LoggerName, alm.Message)
	// log message contains an error error
	if alm.Exception != "" {
		fmt.Printf("Exception: %s", alm.Exception)
	}
}
