package logger

var logLevel = 0

func SetLogLevel(level int) {
	logLevel = level
}

/*
Log lowest level without REST calls
*/
func LogLevel1() bool {
	return logLevel >= 1
}

// Log REST Call URLs
func LogLevel2() bool {
	return logLevel >= 2
}
