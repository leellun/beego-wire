package logz

import (
	"log"
	"os"
)

// Logger 是一个日志记录器结构体，用于包装标准库的日志功能
type Logger struct {
	logger *log.Logger
}

var (
	logger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
)

// Log 打印一条错误日志
func Log(typeStr, err string) {
	logger.Println("["+typeStr+"]", err)
}

// Info 打印一条信息日志
func Info(message string) {
	Log("INFO", message)
}

// Error 打印一条错误日志
func Error(message string) {
	Log("ERROR", message)
}

// Warn 打印一条错误日志
func Warn(message string) {
	Log("WARN", message)
}
