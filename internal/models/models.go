package models

import (
	"database/sql"
	"github.com/koopa0/go-api/internal/repository/dbrepo"
)

type Models struct {
	DB dbrepo.DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: dbrepo.DBModel{DB: db},
	}
}

type User struct {
	ID       int
	Email    string
	Password string
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}
