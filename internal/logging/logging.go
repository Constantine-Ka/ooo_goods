package logging

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

// CreateLogger Инициализация Логера
func CreateLogger() *zap.Logger {
	//Устанавливаем папку, где будут храниться логи.
	logFolder := "logs"
	if _, err := os.Stat(logFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(logFolder, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	//Конфигурируем логгер
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stdout",
			"logs/info.log",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
	//Создаем экземпляр логера
	return zap.Must(config.Build())
}
