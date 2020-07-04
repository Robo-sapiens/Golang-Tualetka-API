package repository

import (
	"github.com/jackc/pgx"
	"main/internal/models"
)

type Repository interface {
	Register(user *models.User) (*models.User, error)
	DeleteAccount(userID int) error
	CreateRoom(room *models.Room) (*models.Room, error)
	DeleteRoom(roomID int) error
	GetInfoAboutRoom(roomID int) (*models.Room, error)
	AddMemberIntoRoom(member *models.Member) error
	DeleteMemberFromRoom(member *models.Member) error

	UpdatePaperToilet(paper *models.Paper) (int, error)
	ChangeValuable(valuable *models.Valuable) (bool, error)
	ChangePayAbility(ability *models.PayAbility) (bool, error)
}

type DB struct {
	DBConnPool *pgx.ConnPool
}

func NewDBStore(db *pgx.ConnPool) Repository {
	return &DB{
		DBConnPool: db,
	}
}