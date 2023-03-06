package generateblock

import (
	"context"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/service"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/handler"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/infrastructure/repository"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/usecase"
)

type Connection struct {
	Host        string
	RpcUserID   string
	RpcPassword string
}

type Generator struct {
	handle  handler.Handler
	network string

	ignoreEmptyMempool bool
}

func NewGenerator(
	nodeInfo *Connection,
	network string,
) *Generator {
	return &Generator{
		handle:  newHandler(nodeInfo.Host, nodeInfo.RpcUserID, nodeInfo.RpcPassword),
		network: network,
	}
}

func (g *Generator) WithIgnoreEmptyMempool(ignoreEmptyMempool bool) *Generator {
	g.ignoreEmptyMempool = ignoreEmptyMempool
	return g
}

func (g *Generator) GenerateBlock(
	ctx context.Context,
	address string,
) error {
	return g.handle.GenerateBlock(ctx, g.network, "", []string{}, address, g.ignoreEmptyMempool)
}

func (g *Generator) GenerateElementsDynafedBlock(
	ctx context.Context,
	fedpegScript string,
	pakEntries []string,
) error {
	return g.handle.GenerateBlock(ctx, g.network, fedpegScript, pakEntries, "", g.ignoreEmptyMempool)
}

func GenerateBlock(
	ctx context.Context,
	nodeInfo *Connection,
	network string,
	address string,
) error {
	handle := newHandler(nodeInfo.Host, nodeInfo.RpcUserID, nodeInfo.RpcPassword)
	return handle.GenerateBlock(ctx, network, "", []string{}, address, false)
}

func GenerateElementsDynafedBlock(
	ctx context.Context,
	nodeInfo *Connection,
	network string,
	fedpegScript string,
	pakEntries []string,
) error {
	handle := newHandler(nodeInfo.Host, nodeInfo.RpcUserID, nodeInfo.RpcPassword)
	return handle.GenerateBlock(ctx, network, fedpegScript, pakEntries, "", false)
}

func newHandler(
	host, rpcUserID, rpcPassword string,
) handler.Handler {
	blockchainConfig := repository.NewBlockchainRpcConfig(
		host, rpcUserID, rpcPassword)
	blockchainRepo := repository.NewBlockchainRpc(blockchainConfig)

	genBlockService := service.NewGenerateBlock(blockchainRepo)
	genBlockUsecase := usecase.NewGenerateBlock(genBlockService)

	return handler.NewHandler(genBlockUsecase)
}
