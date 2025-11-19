package viper

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(mod CfgMod, cfg any) (*viper.Viper, string) {
	config := getConfigPath(mod)

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//配置文件热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&cfg); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	root, _ := filepath.Abs("..")
	return v, root
}

// getConfigPath 获取配置文件路径, 优先级: 命令行 > 环境变量 > 默认值
func getConfigPath(mod CfgMod) (config string) {
	// `-c` flag parse
	flag.StringVar(&config, "c", "", "选择的配置文件地址")
	flag.Parse()
	if config != "" { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", config)
		return
	}
	if env := os.Getenv(configEnv); env != "" {
		config = env
		fmt.Printf("您正在使用 %s 环境变量, config 的路径为 %s\n", configEnv, config)
		return
	}

	switch mod {
	case debugMode:
		config = debugFile
	case testMode:
		config = testFile
	case releaseMode:
		config = releaseFile
	}
	fmt.Printf("您正在使用 %s 模式运行, config 的路径为 %s\n", mod, config)

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = defaultFile
		fmt.Printf("配置文件路径不存在, 使用默认配置文件路径: %s\n", config)
	}

	return
}
