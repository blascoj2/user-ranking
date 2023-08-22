package flags

import "flag"

func ProvideFlags() Flags {
	var traceLevelName = flag.String("trace_level", "INFO", "Trace level (ERROR, WARN, INFO, DEBUG, TRACE)")
	flag.Parse()

	loggingFlags := LoggingFlags{TraceLevel: *traceLevelName}

	return Flags{
		LoggingFlags: loggingFlags,
	}
}

type Flags struct {
	LoggingFlags LoggingFlags
}

type LoggingFlags struct {
	TraceLevel string
}
