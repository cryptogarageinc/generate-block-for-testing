package repository

import (
	"context"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
)

type Blockchain interface {
	GetNewBlockHex(
		ctx context.Context,
		minimumTxAge int,
		fedpegScript string,
		pakEntries []string,
		blockSignScriptPubkey string,
		maximumWitnessSize uint32,
	) (blockHex string, err error)
	SubmitBlock(
		ctx context.Context,
		blockHex string,
	) error
	GenerateToAddress(
		ctx context.Context,
		blockCount int,
		address string,
	) (blockHashes []string, err error)
	GetMempoolTXIDs(ctx context.Context) ([]string, error)
	GetBlockChainInfo(
		ctx context.Context,
	) (*model.BlockChainInfo, error)
}
