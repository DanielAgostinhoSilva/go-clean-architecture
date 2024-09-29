package usecase

import "context"

type Input interface{}
type Output interface{}

type UseCase interface {
	Execute(ctx context.Context, input Input) (Output, error)
}

type UnitUseCase interface {
	Execute(ctx context.Context, input Input) error
}

type NullaryUseCase interface {
	Execute(ctx context.Context) (Output, error)
}
