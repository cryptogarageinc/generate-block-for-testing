package handler

import (
	"context"
	"errors"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/usecase"
)

type Handler interface {
	GenerateBlock(
		ctx context.Context,
		networkType string,
		fedpegScript string,
		pak []string,
		address string,
		ignoreEmptyMempool bool,
	) error
}

type handler struct {
	usecase usecase.GenerateBlock
}

func (h *handler) GenerateBlock(
	ctx context.Context,
	networkType string,
	fedpegScript string,
	pak []string,
	address string,
	ignoreEmptyMempool bool,
) error {
	if networkType == "" {
		return errors.New("networkType is empty")
	}
	if err := model.ValidateNetworkType(networkType); err != nil {
		return err
	}
	return h.usecase.GenerateBlock(ctx, &model.Configuration{
		Network:            model.NewNetworkType(networkType),
		FedpegScript:       fedpegScript,
		PakEntries:         pak,
		Address:            address,
		IgnoreEmptyMempool: ignoreEmptyMempool,
	})
}

func NewHandler(usecase usecase.GenerateBlock) Handler {
	return &handler{
		usecase: usecase,
	}
}
