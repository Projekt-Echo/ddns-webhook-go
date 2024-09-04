package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	zapCfg := zap.NewProductionConfig()
	zapCfg.OutputPaths = []string{"stdout", "./logs.log"}
	zapCfg.DisableStacktrace = true
	Logger, _ = zapCfg.Build()
}
