package nodekit

import (
	"github.com/0xPolygonHermez/zkevm-node/hex"
	"github.com/0xPolygonHermez/zkevm-node/log"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type RollupBlock struct {
	// seq block height
	Height  uint64   `json:"height"`
	ChainID []byte   `json:"chainID"`
	Txs     [][]byte `json:"txs"`

	txs []*ethTypes.Transaction
}

func (b *RollupBlock) UnmarshalTxs() {
	for _, rawTx := range b.Txs {
		tx, err := hexToTx(string(rawTx))
		if err != nil {
			log.Errorf("err unmarshal tx, err: %+v\n", err)
			continue
		}
		b.txs = append(b.txs, tx)
	}
}

func hexToTx(str string) (*ethTypes.Transaction, error) {
	tx := new(ethTypes.Transaction)

	b, err := hex.DecodeHex(str)
	if err != nil {
		return nil, err
	}

	if err := tx.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return tx, nil
}

func (b *RollupBlock) GetTxs() []*ethTypes.Transaction {
	return b.txs
}
