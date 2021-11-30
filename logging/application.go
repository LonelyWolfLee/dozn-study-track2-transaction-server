package logging

func Trace(message string) {
	logger.trace.Println(message)
}

func Info(message string) {
	logger.info.Println(message)
}

func Warn(message string) {
	logger.warn.Println(message)
}

func Error(message string) {
	logger.error.Fatalln(message)
}
