package nodekit

type RollupBlock struct {
	// seq block height
	Height  uint64   `json:"height"`
	ChainID []byte   `json:"chainID"`
	Txs     [][]byte `json:"txs"`
}
