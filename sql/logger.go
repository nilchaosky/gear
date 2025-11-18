package sql

import (
	"fmt"

	"github.com/nilchaosky/gear/logz"
	"gorm.io/gorm/logger"
)

type Writer struct {
	config GeneralDB
	writer logger.Writer
}

func NewWriter(config GeneralDB) *Writer {
	return &Writer{config: config}
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...any) {

	// 当有日志时候均需要输出到控制台
	fmt.Printf(message, data...)

	// 当开启了zap的情况，会打印到日志记录
	if logz.Print != nil && c.config.LogZap {
		switch c.config.LogLevel() {
		case logger.Silent:
			logz.Print.Debug(fmt.Sprintf(message, data...))
		case logger.Error:
			logz.Print.Error(fmt.Sprintf(message, data...))
		case logger.Warn:
			logz.Print.Warn(fmt.Sprintf(message, data...))
		case logger.Info:
			logz.Print.Info(fmt.Sprintf(message, data...))
		default:
			logz.Print.Info(fmt.Sprintf(message, data...))
		}
		return
	}
}
