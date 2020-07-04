package usecase

import "main/internal/models"

func (u *useCase) Register(user *models.User) (*models.User,error) {
	return u.repository.Register(user)
}

func (u *useCase) DeleteAccount(userID int) error {
	return u.repository.DeleteAccount(userID)
}

