package logger

import "log"

func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args)
}

func ErrorOf(format string, args ...interface{}) {
	log.Println(format, args)
}

func InfoOf(format string, args ...interface{}) {
	log.Println(format, args)
}
