package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log2 *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	log2, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log2.Info(message, fields...)

}

func Fatal(message string, fields ...zap.Field) {
	log2.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log2.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log2.Error(message, fields...)
}

func LogPersonal(message string) {
	// open file and create if non-existent
	file, err := os.OpenFile("producer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "message ", log.LstdFlags|log.Lshortfile)

	logger.Output(2, message)

}
