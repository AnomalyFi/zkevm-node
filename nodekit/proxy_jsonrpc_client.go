package nodekit

import (
	"context"
	"strings"
)

const (
	Name            = "proxy"
	JSONRPCEndpoint = "/rpc"
)

type JSONRPCClient struct {
	requester *EndpointRequester
}

func NewJSONRPCClient(uri string) *JSONRPCClient {
	uri = strings.TrimSuffix(uri, "/")
	uri += JSONRPCEndpoint
	req := NewRequester(uri, Name)
	return &JSONRPCClient{requester: req}
}

type SubmitMsgTxArgs struct {
	RollupID []byte `json:"RollupID"`
	Data     []byte `json:"Data"`
}

type SubmitMsgTxReply struct {
	TxID string `json:"txId"`
}

func (j *JSONRPCClient) SubmitMsgTx(ctx context.Context, rollupID []byte, data []byte) (string, error) {
	resp := new(SubmitMsgTxReply)

	err := j.requester.SendRequest(ctx,
		"submitMsgTx",
		&SubmitMsgTxArgs{
			RollupID: rollupID,
			Data:     data,
		},
		resp,
	)

	if err != nil {
		return "", err
	}

	return resp.TxID, nil
}

type ConsumeBlockArgs struct{}
type ConsumeBlockReply struct {
	Block RollupBlock `json:"block"`
}

func (j *JSONRPCClient) ConsumeBlock(ctx context.Context) (*RollupBlock, error) {
	resp := new(ConsumeBlockReply)

	err := j.requester.SendRequest(ctx,
		"consumeBlock",
		&ConsumeBlockArgs{},
		resp,
	)

	if err != nil {
		return nil, err
	}

	return &resp.Block, nil
}
