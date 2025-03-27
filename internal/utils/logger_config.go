package utils

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerPath string = "./logs/"
var logFileName string = "log-" + time.Now().Format("2006-01-02") + ".log"

func Log() *zap.Logger {
	// Asegurar que el directorio de logs exista
	err := os.MkdirAll(loggerPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Configurar archivo de logs para diferentes niveles
	infoLogFile, _ := os.OpenFile(filepath.Join(loggerPath, logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errorLogFile, _ := os.OpenFile(filepath.Join(loggerPath, logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Priority
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	fileInfoWriting := zapcore.Lock(infoLogFile)
	fileErrorWriting := zapcore.Lock(errorLogFile)

	// Encoders
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// Cores
	core := zapcore.NewTee(
		// Consola
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),

		// Archivos
		zapcore.NewCore(fileEncoder, fileErrorWriting, highPriority),
		zapcore.NewCore(fileEncoder, fileInfoWriting, lowPriority),
	)

	log := zap.New(core)
	defer log.Sync()

	return log
}
