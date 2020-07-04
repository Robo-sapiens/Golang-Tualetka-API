package repository

import (
	"main/internal/models"
	"main/internal/tools/errors"
)


func (db *DB) Register(user *models.User) (*models.User, error) {
	err := db.DBConnPool.QueryRow(`INSERT INTO users (nickname, name, phone, status)
    VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING RETURNING id, status`, user.Nickname, user.Name, user.Phone, "common").
		Scan(&user.ID, &user.Status)
	if err != nil {
		return nil, errors.UserAlreadyExists
	}
	return user, nil
}


func (db *DB) DeleteAccount(userID int) error {
	result, err := db.DBConnPool.Exec("DELETE FROM members WHERE user_id = $1", userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.MemberNotFound
	}
	result, err = db.DBConnPool.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.UserNotFound
	}
	return nil
}
