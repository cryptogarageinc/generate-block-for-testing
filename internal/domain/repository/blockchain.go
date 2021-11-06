package repository

import "context"

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
}
