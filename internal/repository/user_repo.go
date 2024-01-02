package repository

import (
	"DoctorWho/internal/domain"
	"database/sql"
)

type repo struct {
	db *sql.DB
}
type Repo interface {
	Register(user domain.User) (int, error)
	Exist(Number string, pass string) (bool, error)
}

func NewRepo(db *sql.DB) Repo {
	return &repo{db: db}
}
func (r repo) Register(user domain.User) (id int, err error) {
	query := `
	insert into users (phone_number,password,role,created_at,updated_at,deleted_at) values($1,$2,$3,$4,$5,$6) returning id
`
	row := r.db.QueryRow(query, user.Phone_number(), user.Password(), user.Role(), user.Created_at(), user.Updated_at(), user.Deleted_at())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r repo) Exist(Number string, pass string) (exist bool, err error) {
	query := `
	SELECT EXISTS (
			SELECT password
			FROM users
			WHERE phone_number = $1
		)
`
	row := r.db.QueryRow(query, Number)
	var password string
	if err := row.Scan(&password); err != nil {
		return false, err
	}
	if pass != password {
		return false, err
	}
	return true, nil
}
