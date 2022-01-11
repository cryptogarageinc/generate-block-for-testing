# generate-block-for-testing

This application is used to perform block generation in a testing network for the bitcoin or the liquid network.

## function

- generatetoaddress
  - use address
  - (in the future) generate address by bip32
  - (in the future) show the generated block's hash.
- generate block for dynamic federation (liquid network)
  - set fedpeg script
  - set pak entry
  - (in the future) block sign
  - (in the future) maximum witness size
  - (in the future) show the generated block's hash.

## usage

### command-line argument

```sh
Usage: generateblock.exe [--host HOST] [--fedpegscript FEDPEGSCRIPT] [--pak PAK] [--network NETWORK] [--address ADDRESS] [--rpcuserid RPCUSERID] [--rpcpassword RPCPASSWORD] [--logging]

Options:
  --host HOST            connection host & port
  --fedpegscript FEDPEGSCRIPT, -s FEDPEGSCRIPT
                         fedpeg script on dynafed
  --pak PAK, -p PAK      pak entries
  --network NETWORK, -n NETWORK
                         network. (bitcoin:mainnet/testnet/regtest, liquid:liquidv1/liquidregtest/elementsregtest)
  --address ADDRESS, -a ADDRESS
                         bitcoin address for generatetoaddress
  --rpcuserid RPCUSERID
                         connection rpc userID
  --rpcpassword RPCPASSWORD
                         connection rpc password
  --logging, -l          log output
  --help, -h             display this help and exit
```

### environment variable

- GENERATE_BLOCK_CONNECTION_HOST: host & port.
- GENERATE_BLOCK_CONNECTION_NETWORK: network type.
  - bitcoin: mainnet, testnet, regtest
  - liquid network: liquidv1, liquidregtest, (elementsregtest)
- CONNECTION_PRC_USERID: connection rpc userID
- CONNECTION_PRC_PASSWORD: connection rpc password
- (for generatetoaddress)
  - GENERATE_BLOCK_GENERATETOADDRESS: bitcoin address
- (liquid network parameter)
  - DYNAFED_FEDPEG_SCRIPT: fedpeg script.
  - DYNAFED_PAK: pak entry. To set multiple items, separate them with commas.

If both environment variable and command-line argument are set, the value set for command-line argument will take precedence.

### example

bitcoin:

```sh
generateblock -l --host localhost:18443 --rpcuserid bitcoinrpc --rpcpassword password -n regtest -a bcrt1qpaujknvwumkwplvpdlh6gtsv7hrl60a37fc9tx
```

liquid network:

```sh
generateblock -l --host localhost:18447 --rpcuserid elementsrpc --rpcpassword password -n elementsregtest -s 5121024241bff4d20f2e616bef2f6e5c25145c068d45a78da3ddba433b3101bbe9a37d51ae -p 02b6991705d4b343ba192c2d1b10e7b8785202f51679f26a1f2cdbe9c069f8dceb024fb0908ea9263bedb5327da23ff914ce1883f851337d71b3ca09b32701003d05 -p 030e07d4f657c0c169e04fac5d5a8096adb099874834be59ad1e681e22d952ccda0214156e4ae9168289b4d0c034da94025121d33ad8643663454885032d77640e3d
```

## build

```go
go build ./cmd/generateblock/
```

```sh
make build
```

### format

```sh
make gettools format
```
