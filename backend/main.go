package main

import (
	"context"

	"github.com/jackc/pgx"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/database_name")
	if err != nil {
		// Handle error
	}
	defer conn.Close()

	ctx := context.Background() // In real applications, use a more specific context as needed
	conn.Exec(ctx, "INSERT INTO users(name, email) VALUES($1, $2)", "John Doe", "john.doe@example.com")

}
