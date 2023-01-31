package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	// 生成一个全局的logger
	zap.ReplaceGlobals(logger)
}
