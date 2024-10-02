package logging

import (
	"log"
	"os"
	"strings"
)

var DebugMode bool
var LogBuffer strings.Builder

const (
	LogPrefix  = "#--- "
	LogEnabled = true
)

func LogInfo(format string, v ...interface{}) {
	if LogEnabled {
		log.Printf(LogPrefix+format, v...)
	}
}

func LogError(err error) {
	if LogEnabled {
		log.Printf(LogPrefix+"Error: %v", err)
	}
}

func LogFatal(err error) {
	log.Fatal(LogPrefix+"Fatal error: ", err)
}

type LogWriter struct{}

func (w LogWriter) Write(p []byte) (n int, err error) {
	LogBuffer.Write(p)
	return os.Stdout.Write(p)
}

func Init() {
	DebugMode = os.Getenv("DEBUG") == "true"

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(LogWriter{})

	if DebugMode {
		LogInfo("Debug mode is enabled")
	}
}
