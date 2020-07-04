package usecase

import "main/internal/models"

func (u *useCase) CreateRoom(room *models.Room) (*models.Room,error){
	return u.repository.CreateRoom(room)
}

func (u *useCase) DeleteRoom(roomID int) error {
	return u.repository.DeleteRoom(roomID)
}

func (u *useCase) GetInfoAboutRoom(roomID int) (*models.Room, error) {
	return u.repository.GetInfoAboutRoom(roomID)
}