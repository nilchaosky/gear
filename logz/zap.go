package logz

import (
	"fmt"
	"os"

	"github.com/nilchaosky/gear/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	config Config
	Print  = defaultLogger()
)

func defaultLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap(c Config) (logger *zap.Logger) {
	config = c
	if config.Director == "" {
		config.Director = "logs"
	}
	if ok, _ := utils.PathExists(config.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.Director)
		_ = os.Mkdir(config.Director, os.ModePerm)
	}
	levels := config.levels()
	length := len(config.levels())
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := newZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	Print = logger
	return logger
}
