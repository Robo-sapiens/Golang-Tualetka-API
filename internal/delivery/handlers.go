package delivery

import useCase "main/internal/usecase"


type Handlers struct {
	usecase useCase.UseCase
}

func NewHandlers(ucases useCase.UseCase) *Handlers {
	return &Handlers{
		usecase: ucases,
	}
}