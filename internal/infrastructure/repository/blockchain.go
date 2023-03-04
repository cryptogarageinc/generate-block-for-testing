package repository

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/repository"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/infrastructure/entity"
	"github.com/pkg/errors"

	resty "github.com/go-resty/resty/v2"
)

type BlockchainRpcConfig struct {
	Host     string
	UserID   string
	Password string
}

type blockchainRpc struct {
	config *BlockchainRpcConfig
	cli    *resty.Client
}

func (b *blockchainRpc) GetNewBlockHex(
	ctx context.Context,
	minimumTxAge int,
	fedpegScript string,
	pakEntries []string,
	blockSignScriptPubkey string,
	maximumWitnessSize uint32,
) (blockHex string, err error) {
	dynafedData := entity.DynafedData{
		SignBlockScript: blockSignScriptPubkey,
		MaxBlockWitness: int64(maximumWitnessSize),
		FedpegScript:    fedpegScript,
		ExtensionSpace:  pakEntries,
	}
	result, _, err := b.post(ctx, "getnewblockhex", minimumTxAge, dynafedData)
	if err != nil {
		return "", err
	}
	blockHex = result.(string)
	return blockHex, nil
}

func (b *blockchainRpc) SubmitBlock(
	ctx context.Context,
	blockHex string,
) error {
	_, _, err := b.post(ctx, "submitblock", blockHex)
	return err
}

func (b *blockchainRpc) GenerateToAddress(
	ctx context.Context,
	blockCount int,
	address string,
) (blockHashes []string, err error) {
	result, _, err := b.post(ctx, "generatetoaddress", int64(blockCount), address)
	if err != nil {
		return nil, err
	}

	workBlockHashes := result.([]interface{})
	blockHashes = make([]string, len(workBlockHashes))
	for i := range workBlockHashes {
		blockHashes[i] = workBlockHashes[i].(string)
	}
	return blockHashes, nil
}

func (b *blockchainRpc) GetMempoolTXIDs(ctx context.Context) ([]string, error) {
	result, _, err := b.post(ctx, "getrawmempool", false, true)
	if err != nil {
		return nil, err
	}
	rawMempoolData := result.(map[string]interface{})
	mempoolTxIDs := rawMempoolData["txids"].([]interface{})
	txIDs := make([]string, len(mempoolTxIDs))
	for i := range mempoolTxIDs {
		txIDs[i] = mempoolTxIDs[i].(string)
	}
	return txIDs, nil
}

func (b *blockchainRpc) GetBlockChainInfoWithElements(ctx context.Context) (*model.BlockChainInfo, error) {
	result, _, err := b.post(ctx, "getblockchaininfo")
	if err != nil {
		return nil, err
	}
	blockchainInfo := &model.BlockChainInfo{}
	rawBlockchainInfo := result.(map[string]interface{})
	blockchainInfo.Blocks = uint64(rawBlockchainInfo["blocks"].(float64))
	blockchainInfo.BestBlockHash = rawBlockchainInfo["bestblockhash"].(string)
	blockchainInfo.CurrentFedpegScript = rawBlockchainInfo["current_fedpeg_script"].(string)
	extensionSpace := rawBlockchainInfo["extension_space"].([]interface{})
	blockchainInfo.ExtensionSpace = make([]string, len(extensionSpace))
	for i := range extensionSpace {
		blockchainInfo.ExtensionSpace[i] = extensionSpace[i].(string)
	}
	return blockchainInfo, nil
}

func (b *blockchainRpc) post(
	ctx context.Context,
	method string,
	params ...interface{},
) (
	result interface{},
	res *resty.Response,
	err error,
) {
	requestParam := entity.RequestData{
		JsonRPC: "1.0",
		ID:      "1",
		Method:  method,
		Params:  params,
	}
	reqJson, err := json.Marshal(requestParam)
	if err != nil {
		return result, nil, err
	}
	responseJson := entity.ResponseData{
		Error: make(map[string]interface{}),
	}

	req := b.cli.R().EnableTrace().
		SetHeader("Cache-Control", "no-cache, no-store").
		SetBody(reqJson).
		SetResult(&responseJson)
	if b.config.UserID != "" {
		req = req.SetBasicAuth(b.config.UserID, b.config.Password)
	}

	res, err = req.Post("")
	if err != nil {
		return result, nil, errors.Wrapf(err, "Request: %v", requestParam)
	} else if !res.IsSuccess() {
		return result, res, errors.Errorf(
			"Invalid Status Code %d, Request: %v, Error: %v",
			res.StatusCode(), requestParam, responseJson.Error)
	}
	result = responseJson.Result
	return result, res, nil
}

func NewBlockchainRpcConfig(
	host,
	userID,
	password string,
) *BlockchainRpcConfig {
	return &BlockchainRpcConfig{
		Host:     host,
		UserID:   userID,
		Password: password,
	}
}

func NewBlockchainRpc(
	config *BlockchainRpcConfig,
) repository.Blockchain {
	host := config.Host
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}
	config.Host = host
	return &blockchainRpc{
		config: config,
		cli:    resty.New().SetBaseURL(host).SetDisableWarn(true), // for http
	}
}
