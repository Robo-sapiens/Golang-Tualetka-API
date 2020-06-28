package repository

import (
	"main/internal/models"
	"main/internal/tools/errors"
)


func (db *DB) Register(user *models.User) error {
	result, err := db.DBConnPool.Exec("INSERT INTO users (nickname, name, phone, status) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING",
		user.Nickname,
		user.Name,
		user.Phone,
		"common")
	if result.RowsAffected() == 0 {
		return errors.UserAlreadyExists
	}
	return err
}
