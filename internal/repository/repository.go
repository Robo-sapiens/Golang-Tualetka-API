package repository

import "github.com/jackc/pgx"

type Repository interface {


}

type DB struct {
	DBConnPool *pgx.ConnPool
}

func NewDBStore(db *pgx.ConnPool) Repository {
	return &DB{
		DBConnPool: db,
	}
}