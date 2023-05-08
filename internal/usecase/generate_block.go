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
			if config.HasCheckInitialBlkDl {
				isInitialBlkDl, err := g.generateBlockService.IsInitialBlockDownload(ctx)
				if err != nil {
					return err
				} else if isInitialBlkDl {
					// need to generate block
				} else {
					return pkgerror.ErrEmptyMempoolTx
				}
			} else {
				return pkgerror.ErrEmptyMempoolTx
			}
		case !exist && config.CanDynafed():
			compare, isInitialBlkDl, err := g.generateBlockService.CompareDynafed(
				ctx, config.FedpegScript, config.PakEntries)
			if err != nil {
				return err
			} else if compare {
				if config.HasCheckInitialBlkDl && isInitialBlkDl {
					// need to generate block
				} else {
					return pkgerror.ErrEmptyMempoolTx
				}
			}
			// need to generate block.
		}
	}

	var err error
	for i := uint(0); i < config.GetGenerateCount(); i++ {
		switch {
		case config.CanDynafed():
			err = g.generateBlockService.GenerateDynafedBlock(
				ctx,
				config.FedpegScript,
				config.PakEntries)
		case config.Address == "":
			err = errors.Errorf("address is not set")
		default:
			err = g.generateBlockService.GenerateToAddress(ctx, config.Address)
		}
		if err != nil {
			break
		}
	}
	return err
}

func NewGenerateBlock(
	generateBlockService service.GenerateBlock,
) GenerateBlock {
	return &generateBlock{
		generateBlockService: generateBlockService,
	}
}
