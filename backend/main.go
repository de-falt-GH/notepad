package main

import (
	"context"
	"kursarbeit/api"
	"kursarbeit/storage"
	common_storage "kursarbeit/storage/common"
	user_storage "kursarbeit/storage/user"

	"go.uber.org/zap"
)

func main() {
	log := zap.Must(zap.NewProduction()).Sugar()
	defer log.Sync()

	pool, err := storage.Connect()
	if err != nil {
		log.Error("connecting to db failed: ", err)
		return
	}
	defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Error("acquiring conn from pool failed: ", err)
		return
	}

	commonStorage := common_storage.NewStorage(conn, log)
	userStorage := user_storage.NewStorage(conn, log)

	api.Serve(commonStorage, userStorage, log)
}
