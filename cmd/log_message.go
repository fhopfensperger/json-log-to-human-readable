package cmd

type LogMessage struct {
	Timestamp string `json:"timestamp"`
	Level string `json:"level"`
	Message string `json:"message"`
}