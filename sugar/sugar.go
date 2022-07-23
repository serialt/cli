package sugar

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

	Config *config.MyConfig
	// Logger *zap.Logger
	Log *zap.SugaredLogger
)
