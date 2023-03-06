package usecase

import (
	"context"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/service"
	pkgerror "github.com/cryptogarageinc/generate-block-for-testing/internal/pkg/errors"
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
	if config.IgnoreEmptyMempool {
		exist, err := g.generateBlockService.ExistMempool(ctx)
		switch {
		case err != nil:
			return err
		case !exist && !config.CanDynafed():
			return pkgerror.ErrEmptyMempoolTx
		case !exist && config.CanDynafed():
			compare, err := g.generateBlockService.CompareDynafed(
				ctx, config.FedpegScript, config.PakEntries)
			if err != nil {
				return err
			} else if compare {
				return pkgerror.ErrEmptyMempoolTx
			}
			// need to generate block.
		}
	}
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
