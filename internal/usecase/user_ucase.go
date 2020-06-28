package usecase

import "main/internal/models"

func (u *useCase) Register(user *models.User) error {
	return u.repository.Register(user)
}
