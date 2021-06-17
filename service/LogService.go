package service

import (
	"encoding/json"
	"os"
	"regexp"
	"time"
	"web-read/enum"
)

type logService struct {
	Log func(level string, message string, data map[string]interface{})
}

type logLine struct {
	Time    int64
	Level   string
	Message string
	Data    map[string]interface{}
}

// 注册日志
func (l logService) LogRegist() func(level string, message string, data map[string]interface{}) {
	logFile := os.Getenv("ROOT_DIR") + "/log/log-" + time.Now().Format(enum.DataZone) + ".log"
	// 文件不存在则创建
	_, fileExistErr := os.Stat(logFile)
	if fileExistErr != nil {
		f, _ := os.Create(logFile)
		_ = f.Close()
	}
	src, _ := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	return func(level string, message string, data map[string]interface{}) {
		content := logLine{
			Time:    time.Now().Unix(),
			Level:   level,
			Message: message,
			Data:    data,
		}
		contentBytes, _ := json.Marshal(content)
		// 去掉\n
		re, _ := regexp.Compile("\n")
		contentBytes = re.ReplaceAll(contentBytes, []byte(" "))
		contentBytes = append(contentBytes, []byte("\n")...)
		_, _ = src.Write(contentBytes)
	}
}

var LogService logService
