package usecase

import (
	"main/internal/models"
	"main/internal/repository"
)

type UseCase interface {
	Register(user *models.User) (*models.User, error)
	DeleteAccount(userID int) error
	GetWhoBuy(roomID int) (*models.Member, error)
	CreateRoom(room *models.Room) (*models.Room, error)
	DeleteRoom(roomID int) error
	GetInfoAboutRoom(roomID int) (*models.Room, error)
	AddMemberIntoRoom(member *models.Member) error
	DeleteMemberFromRoom(member *models.Member) error

	UpdatePaperToilet(paper *models.Paper) (int, error)
	ChangeValuable(valuable *models.Valuable) (bool, error)
	ChangePayAbility(ability *models.PayAbility) (bool, error)

}

type useCase struct {
	repository repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{
		repository: repo,
	}
}