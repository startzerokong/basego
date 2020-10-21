package log

import (
	"fmt"
	"github.com/startzerokong/basego/util"
	"os"
	"time"
)

func AllLog() {
	config := util.GetLogConfig()
	allLogPath := config.AllLogPath
	allLogName := config.AllLogName
	fullPath := allLogPath + "/" + allLogName

	timeStr := time.Now().Format("2006-01-02 15:04:05")

	file, _ := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	_, _ = fmt.Fprintln(file, )
	_ = file.Close()
}
