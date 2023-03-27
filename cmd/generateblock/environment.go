package main

type environment struct {
	Host               string   `env:"GENERATE_BLOCK_CONNECTION_HOST"`
	FedpegScript       string   `env:"DYNAFED_FEDPEG_SCRIPT"`
	PakEntries         []string `env:"DYNAFED_PAK" envSeparator:","`
	Network            string   `env:"GENERATE_BLOCK_CONNECTION_NETWORK"`
	Address            string   `env:"GENERATE_BLOCK_GENERATETOADDRESS"`
	RpcUserID          string   `env:"CONNECTION_PRC_USERID"`
	RpcPassword        string   `env:"CONNECTION_PRC_PASSWORD"`
	GenerateCount      uint     `env:"GENERATE_BLOCK_COUNT" envDefault:"1"`
	IgnoreEmptyMempool bool     `env:"IGNORE_EMPTY_MEMPOOL"`
}
