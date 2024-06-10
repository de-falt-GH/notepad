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
	ListPublicNotes(ctx context.Context, req *ListPublicNotesRequest) (res []Note, err error)
}

type storage struct {
	Storage
	conn *pgxpool.Pool
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

	row := s.conn.QueryRow(ctx, query, args...)
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

	row := s.conn.QueryRow(ctx, query, args...)
	row.Scan(&res.Id, &res.Login, &res.PasswordHash, &res.Email, &res.Name, &res.Info)

	return
}

func (s storage) ListPublicNotes(ctx context.Context, req *ListPublicNotesRequest) (res []Note, err error) {
	query := `SELECT n.id, u.id, n.name, n.data, n.public, n.updated, u.name FROM note n LEFT JOIN "user" u ON n.user_id=u.id WHERE n.public=true ORDER BY n.updated DESC`
	args := []any{}
	cnt := 1

	if req.Search != "" {
		query += " AND name LIKE '%' || $" + strconv.Itoa(cnt) + " || '%'"
		args = append(args, req.Search)
		cnt++
	}

	if req.Skip != 0 {
		query += " SKIP $" + strconv.Itoa(cnt)
		args = append(args, req.Skip)
		cnt++
	}

	if req.Limit != 0 {
		query += " LIMIT $" + strconv.Itoa(cnt)
		args = append(args, req.Limit)
		cnt++
	}

	rows, err := s.conn.Query(ctx, query, args...)
	for rows.Next() {
		note := Note{}
		err = rows.Scan(&note.Id, &note.UserId, &note.Name, &note.Data, &note.Public, &note.Updated, &note.AuthorName)
		res = append(res, note)
	}

	return
}

func NewStorage(conn *pgxpool.Pool, log *zap.SugaredLogger) *storage {
	return &storage{conn: conn, log: log}
}
