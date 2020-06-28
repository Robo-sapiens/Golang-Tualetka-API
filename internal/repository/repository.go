package repository

import (
	"github.com/jackc/pgx"
	"main/internal/models"
)

type Repository interface {
	Register(user *models.User) error

}

type DB struct {
	DBConnPool *pgx.ConnPool
}

func NewDBStore(db *pgx.ConnPool) Repository {
	return &DB{
		DBConnPool: db,
	}
}