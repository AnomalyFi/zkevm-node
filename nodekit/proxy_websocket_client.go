package nodekit

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	c             *websocket.Conn
	pendingBlocks chan RollupBlock
}

func NewWSClient(addr string) (*WSClient, error) {
	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}

	ws := &WSClient{
		c:             c,
		pendingBlocks: make(chan RollupBlock),
	}

	go func() {
		for {
			msgType, rawMsg, err := c.ReadMessage()
			if err != nil {
				log.Printf("failed to read message, err: %s\n", err)
			}

			// respond Pong
			if msgType == websocket.PingMessage {
				c.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(writeWait))
				continue
			}

			msg := new(Message)
			err = json.Unmarshal(rawMsg, msg)
			if err != nil {
				log.Printf("unable to unmarshal msg, err: %s\n", err)
				continue
			}

			blk := new(RollupBlock)
			err = json.Unmarshal([]byte(msg.Message), blk)
			if err != nil {
				log.Printf("unable to unmarshal blk, err :%s\n", err)
				continue
			}

			log.Printf("receiving block %+v\n", blk)
			ws.pendingBlocks <- *blk
		}
	}()

	return ws, nil
}

func (cli *WSClient) SubscribeBlock() error {
	return cli.c.WriteJSON(BlockSubscriptionMessage)
}

func (cli *WSClient) ListenBlock(ctx context.Context) (*RollupBlock, error) {
	select {
	case blk := <-cli.pendingBlocks:
		return &blk, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
