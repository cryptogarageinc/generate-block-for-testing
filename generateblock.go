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

func GenerateBlock(
	ctx context.Context,
	nodeInfo *Connection,
	network string,
	address string,
) error {
	handle := newHandler(nodeInfo.Host, nodeInfo.RpcUserID, nodeInfo.RpcPassword)
	return handle.GenerateBlock(ctx, network, "", []string{}, address)
}

func GenerateElementsDynafedBlock(
	ctx context.Context,
	nodeInfo *Connection,
	network string,
	fedpegScript string,
	pakEntries []string,
) error {
	handle := newHandler(nodeInfo.Host, nodeInfo.RpcUserID, nodeInfo.RpcPassword)
	return handle.GenerateBlock(ctx, network, fedpegScript, pakEntries, "")
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
