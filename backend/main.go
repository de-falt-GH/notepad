package main

import (
	"context"
	"kursarbeit/api"
	"kursarbeit/storage"
	common_storage "kursarbeit/storage/common"
	user_storage "kursarbeit/storage/user"
	"os"

	"go.uber.org/zap"
)

func main() {
	os.Setenv("POSTGRES_HOST", "0.0.0.0")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "kursarbeit")
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "postgres")
	os.Setenv("POSTGRES_SSLMODE", "disable")

	os.Setenv("BACKEND_HOST", "127.0.0.1")
	os.Setenv("BACKEND_PORT", "8080")

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
