package logging

import (
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

type TraceLevel int

const (
	TRACE_LEVEL TraceLevel = iota
	DEBUG_LEVEL TraceLevel = iota
	INFO_LEVEL  TraceLevel = iota
	WARN_LEVEL  TraceLevel = iota
	ERROR_LEVEL TraceLevel = iota
)

func GetTraceLevelWithName(string string) TraceLevel {
	switch string {
	case "TRACE":
		return TRACE_LEVEL
	case "DEBUG":
		return DEBUG_LEVEL
	case "INFO":
		return INFO_LEVEL
	case "WARN":
		return WARN_LEVEL
	case "ERROR":
		return ERROR_LEVEL
	default:
		return INFO_LEVEL
	}
}

func InitLogger(traceLevel TraceLevel) {

	var logLevel log.Level
	switch traceLevel {
	case DEBUG_LEVEL:
		logLevel = log.DebugLevel
	case WARN_LEVEL:
		logLevel = log.WarnLevel
	case INFO_LEVEL:
		logLevel = log.InfoLevel
	case TRACE_LEVEL:
		logLevel = log.TraceLevel
	case ERROR_LEVEL:
		logLevel = log.ErrorLevel
	}

	log.SetFormatter(&formatter.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"service", "file"},
		ShowFullLevel:   true,
		NoFieldsColors:  true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)
	log.SetReportCaller(false)
}

func Log(level log.Level, msg string, fields log.Fields) {
	log.WithFields(fields).Log(level, msg)
}