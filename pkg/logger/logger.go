package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type Logger struct {
	logFile *os.File
}

// Config 日志配置
type Config struct {
	LogDir    string // 日志目录，默认为 "logs"
	EnableLog bool   // 是否启用日志，默认为 true
}

// NewLogger 创建新的日志实例
func NewLogger(config *Config) (*Logger, error) {
	if config == nil {
		config = &Config{
			LogDir:    "logs",
			EnableLog: true,
		}
	}

	if !config.EnableLog {
		return &Logger{}, nil
	}

	// 创建日志目录
	if err := os.MkdirAll(config.LogDir, 0755); err != nil {
		return nil, fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 生成日志文件名
	now := time.Now()
	logFileName := fmt.Sprintf("%04d_%02d_%02d_%02d_%02d_%02d.log",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	logFilePath := filepath.Join(config.LogDir, logFileName)

	// 创建日志文件
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("创建日志文件失败: %w", err)
	}

	logger := &Logger{
		logFile: logFile,
	}

	// 设置日志输出
	logger.setupLogger()

	return logger, nil
}

// setupLogger 设置日志输出格式和目标
func (l *Logger) setupLogger() {
	if l.logFile == nil {
		return
	}

	// 设置日志输出到文件和控制台
	multiWriter := io.MultiWriter(l.logFile, os.Stdout)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 设置Gin的日志输出
	gin.DefaultWriter = multiWriter
}

// Close 关闭日志文件
func (l *Logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

// Info 记录信息日志
func (l *Logger) Info(v ...interface{}) {
	log.Println(v...)
}

// Infof 格式化记录信息日志
func (l *Logger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Error 记录错误日志
func (l *Logger) Error(v ...interface{}) {
	log.Println(v...)
}

// Errorf 格式化记录错误日志
func (l *Logger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Fatal 记录致命错误日志并退出程序
func (l *Logger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

// Fatalf 格式化记录致命错误日志并退出程序
func (l *Logger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
