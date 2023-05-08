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
	ExistMempool(ctx context.Context) (bool, error)
	IsInitialBlockDownload(ctx context.Context) (bool, error)
	CompareDynafed(
		ctx context.Context,
		fedpegScript string,
		pakEntries []string,
	) (bool, bool, error)
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

func (g *generateBlock) ExistMempool(
	ctx context.Context,
) (bool, error) {
	txIDs, err := g.blockchainRepo.GetMempoolTXIDs(ctx)
	if err != nil {
		return false, err
	}
	return len(txIDs) > 0, nil
}

func (g *generateBlock) IsInitialBlockDownload(
	ctx context.Context,
) (bool, error) {
	bi, err := g.blockchainRepo.GetBlockChainInfo(ctx)
	if err != nil {
		return false, err
	}
	return bi.IsInitialBlockDownload, nil
}

func (g *generateBlock) CompareDynafed(
	ctx context.Context,
	fedpegScript string,
	pakEntries []string,
) (bool, bool, error) {
	bi, err := g.blockchainRepo.GetBlockChainInfo(ctx)
	if err != nil {
		return false, false, err
	}
	if bi.CurrentFedpegScript != fedpegScript || !g.compareSlice(bi.ExtensionSpace, pakEntries) {
		return false, bi.IsInitialBlockDownload, nil
	}
	return true, bi.IsInitialBlockDownload, nil
}

func (g *generateBlock) compareSlice(src []string, dst []string) bool {
	srcMap := make(map[string]bool, len(src))
	for _, str := range src {
		srcMap[str] = false
	}
	cnt := 0
	for _, str := range dst {
		if exist, ok := srcMap[str]; ok && !exist {
			srcMap[str] = true
			cnt++
		}
	}
	return cnt == len(src)
}

func NewGenerateBlock(blockchainRepo repository.Blockchain) GenerateBlock {
	return &generateBlock{
		blockchainRepo: blockchainRepo,
	}
}
