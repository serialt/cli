package t

import (
	"github.com/serialt/cli/config"
	"go.uber.org/zap"
)

// 存放一些公共的变量

var (
	// Listen   = ":9879"
	// Host     = ""
	// Username = ""
	// Password = ""

	// 其他配置文件
	ConfigFile = "~/.git-mirror.yaml"

	Config *config.MyConfig
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)
