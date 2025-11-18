package logz

import (
	"fmt"
	"os"

	"github.com/nilchaosky/gear/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Config *Logger
	Print  *zap.Logger
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(Config.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", Config.Director)
		_ = os.Mkdir(Config.Director, os.ModePerm)
	}
	levels := Config.Levels()
	length := len(Config.Levels())
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if Config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	Print = logger
	return logger
}
