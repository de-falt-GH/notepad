package user

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Storage interface {
	DetailUser(ctx context.Context, req *DetailUserRequest) (res DetailUserResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest) (err error)

	AddNote(ctx context.Context, req *AddNoteRequest) (res AddNoteResponse, err error)
	UpdateNote(ctx context.Context, req *UpdateNoteRequest) (err error)
	DetailNote(ctx context.Context, req *DetailNoteRequest) (res Note, err error)
	DeleteNote(ctx context.Context, req *DeleteNoteRequest) (err error)
	ListPrivateNotes(ctx context.Context, req *ListPrivateNotesRequest) (res []Note, err error)
}

type storage struct {
	conn *pgxpool.Conn
	log  *zap.SugaredLogger
}

func (s storage) DetailUser(ctx context.Context, req *DetailUserRequest) (res DetailUserResponse, err error) {
	query := `SELECT login, password_hash, email, name, info FROM "user" WHERE 1=1`
	args := []any{}
	cnt := 1

	if req.Id != 0 {
		query += " AND id=$" + strconv.Itoa(cnt)
		args = append(args, req.Id)
		cnt++
	}

	if req.Login != "" {
		query += " AND login=$" + strconv.Itoa(cnt)
		args = append(args, req.Login)
		cnt++
	}

	row := s.conn.QueryRow(ctx, query, args...)
	err = row.Scan(&res.Login, &res.PasswordHash, &res.Email, &res.Name, &res.Info)

	return
}

func (s storage) UpdateUser(ctx context.Context, req *UpdateUserRequest) (err error) {
	query := `UPDATE "user" SET login=$1, password_hash=$2, email=$3, name=$4, info=$5 WHERE id=$6`
	args := []any{req.Login, req.PasswordHash, req.Email, req.Name, req.Info, req.Id}

	s.log.Info(req.Login)
	_, err = s.conn.Exec(ctx, query, args...)

	return
}

func (s storage) AddNote(ctx context.Context, req *AddNoteRequest) (res AddNoteResponse, err error) {
	query := `INSERT INTO note (user_id, name, data, public) VALUES ($1, $2, $3, $4) RETURNING id`
	args := []any{req.UserId, req.Name, req.Data, req.Public}

	row := s.conn.QueryRow(ctx, query, args...)
	err = row.Scan(&res.NoteId)

	return
}

func (s storage) UpdateNote(ctx context.Context, req *UpdateNoteRequest) (err error) {
	query := `UPDATE note SET name=$1, data=$2, public=$3 WHERE id=$4`
	args := []any{req.Name, req.Data, req.Public, req.Id}

	_, err = s.conn.Exec(ctx, query, args...)

	return
}

func (s storage) DetailNote(ctx context.Context, req *DetailNoteRequest) (res Note, err error) {
	query := `SELECT id, name, data, public FROM note WHERE id=$1`
	args := []any{req.Id}

	row := s.conn.QueryRow(ctx, query, args...)
	err = row.Scan(&res.Id, &res.Name, &res.Data, &res.Public)

	return
}

func (s storage) DeleteNote(ctx context.Context, req *DeleteNoteRequest) (err error) {
	query := `DELETE FROM note WHERE id=$1`
	args := []any{req.Id}

	_, err = s.conn.Exec(ctx, query, args...)

	return
}

func (s storage) ListPrivateNotes(ctx context.Context, req *ListPrivateNotesRequest) (res []Note, err error) {
	query := `SELECT id, name, data, public FROM note WHERE user_id=$1`
	args := []any{req.UserId}
	cnt := 2

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
		err = rows.Scan(&note.Id, &note.Name, &note.Data, &note.Public)
		res = append(res, note)
	}

	return
}

func NewStorage(conn *pgxpool.Conn, log *zap.SugaredLogger) Storage {
	return &storage{conn: conn, log: log}
}
