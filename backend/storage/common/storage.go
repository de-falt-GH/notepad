package common

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Storage interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (id int, err error)
	DetailUser(ctx context.Context, req *DetailUserRequest) (res *DetailUserResponse, err error)
}

type storage struct {
	Storage
	conn *pgxpool.Conn
	log  *zap.SugaredLogger
}

func (s storage) CreateUser(ctx context.Context, req *CreateUserRequest) (id int, err error) {
	query := `INSERT INTO "user"(login, password_hash, email, name, info)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	args := []any{
		req.Login,
		req.PasswordHash,
		req.Email,
		req.Name,
		req.Info,
	}

	row := s.conn.QueryRow(ctx, query, args)
	err = row.Scan(&id)

	return
}

func (s storage) DetailUser(ctx context.Context, req *DetailUserRequest) (res *DetailUserResponse, err error) {
	res = &DetailUserResponse{}

	query := `SELECT id, login, password_hash, email, name, info FROM "user" WHERE 1=1`
	args := []any{}
	cnt := 1

	if req.Login != "" {
		query += " AND login=$" + strconv.Itoa(cnt)
		args = append(args, req.Login)
		cnt++
	}

	row := s.conn.QueryRow(ctx, query, args)
	row.Scan(&res.Id, &res.Login, &res.PasswordHash, &res.Email, &res.Name, &res.Info)

	return
}

func NewStorage(conn *pgxpool.Conn, log *zap.SugaredLogger) *storage {
	return &storage{conn: conn, log: log}
}
