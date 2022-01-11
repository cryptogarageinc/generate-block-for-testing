package main

import (
	"fmt"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
)

type ArgError string

const (
	ErrHostName ArgError = "Connection host name not set"
)

type argument struct {
	Host         string   `help:"connection host & port"`
	FedpegScript string   `arg:"-s" help:"fedpeg script on dynafed"`
	Pak          []string `arg:"-p,separate" help:"pak entries"`
	Network      string   `arg:"-n" help:"network. (bitcoin:mainnet/testnet/regtest, liquid:liquidv1/liquidregtest/elementsregtest)"`
	Address      string   `arg:"-a" help:"bitcoin address for generatetoaddress"`
	RpcUserID    string   `help:"connection rpc userID"`
	RpcPassword  string   `help:"connection rpc password"`
	Logging      bool     `arg:"-l" help:"log output"`
}

// Error returns the error string.
func (e ArgError) Error() string {
	return string(e)
}

func (a *argument) setValueFromEnvironment(env *environment) {
	if a.Host == "" {
		a.Host = env.Host
	}
	if a.Network == "" {
		a.Network = env.Network
	}
	if a.FedpegScript == "" {
		a.FedpegScript = env.FedpegScript
	}
	if len(a.Pak) == 0 {
		a.Pak = env.PakEntries
	}
	if a.Address == "" {
		a.Address = env.Address
	}
	if a.RpcUserID == "" {
		a.RpcUserID = env.RpcUserID
	}
	if a.RpcPassword == "" {
		a.RpcPassword = env.RpcPassword
	}
}

func (a *argument) Validate() error {
	if a.Host == "" {
		return ErrHostName
	}
	if err := model.ValidateNetworkType(a.Network); err != nil {
		return err
	}
	return nil
}

func (a *argument) ToConfigurationModel() *model.Configuration {
	var network model.NetworkType
	if a.Network != "" {
		network = model.NewNetworkType(a.Network)
	} else {
		network = model.ElementsRegtest
		fmt.Println("set: default network elementsRegTest")
	}

	config := &model.Configuration{
		Network:      network,
		FedpegScript: a.FedpegScript,
		PakEntries:   a.Pak,
		Address:      a.Address,
	}
	return config
}
