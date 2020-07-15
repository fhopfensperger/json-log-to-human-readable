package cmd

import "fmt"

type LogMessage struct {
	Timestamp string `json:"timestamp"`
	Level string `json:"level"`
	Message string `json:"message"`
	Exception Exception `json:"exception,omitempty"`
	LoggerName string `json:"loggerName"`
}

type Exception struct {
	RefId int `json:"refId"`
	ExceptionType string `json:"exceptionType"`
	Message string `json:"message"`
	CausedBy CausedBy `json:"causedBy"`
	Frames *[]Frame `json:"frames"`
}

type CausedBy struct {
	Exception *Exception `json:"exception,omitempty"`
}

type Frame struct {
	Class string `json:"class"`
	Method string `json:"method"`
	Line int `json:"line"`
}

func (lm LogMessage) print() {
	fmt.Printf("%v %v\t %v\t%v\n",lm.Level, lm.Timestamp, lm.LoggerName, lm.Message)

	// log message contains an error error
	if lm.Exception != (Exception{}) {
		lm.Exception.print()
	}
}

func (ex Exception) print()  {
	fmt.Printf("Caused by: %v. %s:\n", ex.ExceptionType, ex.Message)
	for _, frame := range *ex.Frames {
		fmt.Printf("\t at %s(%s:%v)\n",frame.Method, frame.Class ,frame.Line)
	}
	if ex.CausedBy != (CausedBy{}) {
		ex.CausedBy.Exception.print()
	}
}