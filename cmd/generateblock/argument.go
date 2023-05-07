package main

import "time"

type ArgError string

const (
	ErrHostName ArgError = "Connection host name not set"
)

type argument struct {
	Host                 string        `help:"connection host & port"`
	FedpegScript         string        `arg:"-s" help:"fedpeg script on dynafed"`
	Pak                  []string      `arg:"-p,separate" help:"pak entries"`
	Network              string        `arg:"-n" help:"network. (bitcoin:mainnet/testnet/regtest, liquid:liquidv1/liquidregtest/elementsregtest)"`
	Address              string        `arg:"-a" help:"bitcoin address for generatetoaddress"`
	RpcUserID            string        `help:"connection rpc userID"`
	RpcPassword          string        `help:"connection rpc password"`
	Logging              bool          `arg:"-l" help:"log output"`
	PollingTime          time.Duration `arg:"-t" help:"polling duration time"`
	GenerateCount        uint          `arg:"-c" help:"generate count"`
	IgnoreEmptyMempool   bool          `arg:"-m" help:"ignore empty mempool"`
	HasCheckInitialBlkDl bool          `arg:"-i" help:"check initial block download flag"`
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
	if a.GenerateCount == 0 {
		a.GenerateCount = env.GenerateCount
	}
	if !a.IgnoreEmptyMempool {
		a.IgnoreEmptyMempool = env.IgnoreEmptyMempool
	}
	if !a.HasCheckInitialBlkDl {
		a.HasCheckInitialBlkDl = env.HasCheckInitialBlkDl
	}
}

func (a *argument) Validate() error {
	if a.Host == "" {
		return ErrHostName
	}
	return nil
}
