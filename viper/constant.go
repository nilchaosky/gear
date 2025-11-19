package viper

type CfgMod string

const (
	configEnv          = "GEAR_CONFIG_PATH"
	debugMode   CfgMod = "debug"
	testMode    CfgMod = "test"
	releaseMode CfgMod = "release"
	defaultFile        = "config.yaml"
	testFile           = "config.test.yaml"
	debugFile          = "config.debug.yaml"
	releaseFile        = "config.release.yaml"
)
