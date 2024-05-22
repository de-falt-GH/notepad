package main

import (
	"kursarbeit/api"
	"kursarbeit/storage"
	common_storage "kursarbeit/storage/common"
	user_storage "kursarbeit/storage/user"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	log := zap.New(core).Sugar()
	defer log.Sync()

	pool, err := storage.Connect()
	if err != nil {
		log.Error("connecting to db failed: ", err)
		return
	}
	defer pool.Close()

	commonStorage := common_storage.NewStorage(pool, log)
	userStorage := user_storage.NewStorage(pool, log)

	api.Serve(commonStorage, userStorage, log)
}
