package usecase

import "main/internal/models"

func (u *useCase) CreateRoom(room *models.Room) (*models.Room,error){
	return u.repository.CreateRoom(room)
}

