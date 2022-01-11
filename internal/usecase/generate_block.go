package usecase

import (
	"context"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/service"
	"github.com/pkg/errors"
)

type GenerateBlock interface {
	GenerateBlock(
		ctx context.Context,
		config *model.Configuration,
	) error
}

type generateBlock struct {
	generateBlockService service.GenerateBlock
}

func (g *generateBlock) GenerateBlock(
	ctx context.Context,
	config *model.Configuration,
) error {
	if config.CanDynafed() {
		return g.generateBlockService.GenerateDynafedBlock(
			ctx,
			config.FedpegScript,
			config.PakEntries)
	}
	if config.Address == "" {
		return errors.Errorf("address is not set")
	}
	return g.generateBlockService.GenerateToAddress(ctx, config.Address)
}

func NewGenerateBlock(
	generateBlockService service.GenerateBlock,
) GenerateBlock {
	return &generateBlock{
		generateBlockService: generateBlockService,
	}
}
