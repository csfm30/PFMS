package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// https://grafana.com/docs/loki/latest/api/#examples-8
func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}

func ErrorLogin(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}

func ErrorHook(message interface{}, fields ...zap.Field) {

	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
