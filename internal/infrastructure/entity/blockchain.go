package entity

type RequestData struct {
	JsonRPC string        `json:"jsonrpc,"`
	ID      string        `json:"id,"`
	Method  string        `json:"method,"`
	Params  []interface{} `json:"params,"`
}

type ResponseData struct {
	Result interface{}            `json:"result,"`
	Error  map[string]interface{} `json:"error,"`
	ID     string                 `json:"id,"`
}

type DynafedData struct {
	SignBlockScript string   `json:"signblockscript"`
	MaxBlockWitness int64    `json:"max_block_witness"`
	FedpegScript    string   `json:"fedpegscript"`
	ExtensionSpace  []string `json:"extension_space"`
}
