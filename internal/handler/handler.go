package handler

import (
	"context"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/usecase"
)

type Handler interface {
	GenerateBlock(
		ctx context.Context,
		config *model.Configuration,
	) error
}

type handler struct {
	usecase usecase.GenerateBlock
}

func (h *handler) GenerateBlock(
	ctx context.Context,
	config *model.Configuration,
) error {
	return h.usecase.GenerateBlock(ctx, config)
}

func NewHandler(usecase usecase.GenerateBlock) Handler {
	return &handler{
		usecase: usecase,
	}
}
