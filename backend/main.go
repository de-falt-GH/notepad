package main

import (
	"context"

	"kursarbeit/storage"

	"go.uber.org/zap"
)

func main() {
	log := zap.Must(zap.NewProduction()).Sugar()
	defer log.Sync()

	conn, err := storage.Connect()
	if err != nil {
		log.Error("connecting to db failed: ", err)
		return
	}
	defer conn.Close(context.Background())

}
