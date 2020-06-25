package usecase

import "main/internal/repository"

type UseCase interface {

}

type useCase struct {
	repository repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{
		repository: repo,
	}
}