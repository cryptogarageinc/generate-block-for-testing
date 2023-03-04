package model

type BlockChainInfo struct {
	Blocks              uint64
	BestBlockHash       string
	CurrentFedpegScript string
	ExtensionSpace      []string
}
