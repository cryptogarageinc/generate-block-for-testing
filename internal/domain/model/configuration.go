package model

import (
	"strings"

	"github.com/pkg/errors"
)

type NetworkType string

const (
	Mainnet         NetworkType = "mainnet"
	Testnet         NetworkType = "testnet"
	Regtest         NetworkType = "regtest"
	LiquidV1        NetworkType = "liquidv1"
	ElementsRegtest NetworkType = "elementsregtest"
)

func (n NetworkType) String() string {
	return string(n)
}

func (n NetworkType) IsElements() bool {
	switch n {
	case LiquidV1, ElementsRegtest:
		return true
	default:
		return false
	}
}

func NewNetworkType(network string) NetworkType {
	switch strings.ToLower(network) {
	case Mainnet.String(), Testnet.String(), Regtest.String(), LiquidV1.String():
		return NetworkType(network)
	case ElementsRegtest.String(), "liquidregtest":
		return ElementsRegtest
	default:
		panic("no match network type")
	}
}

func ValidateNetworkType(network string) error {
	switch strings.ToLower(network) {
	case Mainnet.String(), Testnet.String(), Regtest.String(), LiquidV1.String():
	case ElementsRegtest.String(), "liquidregtest":
	default:
		return errors.Errorf("no match network type, %s", network)
	}
	return nil
}

type Configuration struct {
	Network            NetworkType
	FedpegScript       string
	PakEntries         []string
	Address            string
	IgnoreEmptyMempool bool
}

func (c *Configuration) CanDynafed() bool {
	if c.Network.IsElements() && c.FedpegScript != "" && len(c.PakEntries) > 0 {
		return true
	}
	return false
}

func (c *Configuration) IsBitcoinNetwork() bool {
	switch c.Network {
	case Mainnet, Testnet, Regtest:
		return true
	default:
		return false
	}
}
