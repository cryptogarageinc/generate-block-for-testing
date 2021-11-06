package service

import (
	"context"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/repository"
)

type GenerateBlock interface {
	GenerateDynafedBlock(
		ctx context.Context,
		fedpegScript string,
		pakEntries []string,
	) error
	GenerateToAddress(ctx context.Context, address string) error
}

const (
	wshOpTrueLockingScript  = "00204ae81572f06e1b88fd5ced7a1a000945432e83e1551e6f721ee9c00b8cc33260"
	maximumBlockWitnessSize = 520
)

type generateBlock struct {
	blockchainRepo repository.Blockchain
}

func (g *generateBlock) GenerateDynafedBlock(
	ctx context.Context,
	fedpegScript string,
	pakEntries []string,
) error {
	blockHex, err := g.blockchainRepo.GetNewBlockHex(
		ctx, 0, fedpegScript, pakEntries,
		wshOpTrueLockingScript, maximumBlockWitnessSize)
	if err != nil {
		return err
	}
	return g.blockchainRepo.SubmitBlock(ctx, blockHex)
}

func (g *generateBlock) GenerateToAddress(
	ctx context.Context,
	address string,
) error {
	_, err := g.blockchainRepo.GenerateToAddress(ctx, 1, address)
	// TODO(k-matsuzawa): In the future, I will make it return blockID.
	return err
}

func NewGenerateBlock(blockchainRepo repository.Blockchain) GenerateBlock {
	return &generateBlock{
		blockchainRepo: blockchainRepo,
	}
}
