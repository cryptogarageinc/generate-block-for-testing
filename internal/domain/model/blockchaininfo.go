package model

type BlockChainInfo struct {
	Blocks                 uint64
	BestBlockHash          string
	IsInitialBlockDownload bool

	// elements only parameters
	CurrentFedpegScript string
	ExtensionSpace      []string
}
