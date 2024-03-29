package v1_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	server "github.com/carbonfive/go-filecoin-rest-api"
	v1 "github.com/carbonfive/go-filecoin-rest-api/handlers/v1"
	"github.com/carbonfive/go-filecoin-rest-api/test"
	"github.com/carbonfive/go-filecoin-rest-api/types"
)

func TestBlockHeaderHandler_ServeHTTP(t *testing.T) {

	t.Run("all fields pass through", func(t *testing.T) {
		tbh := types.BlockHeader{
			Miner:                 "someaddress",
			Tickets:               [][]byte{[]byte("ticket1")},
			ElectionProof:         []byte("electionproof"),
			Parents:               []cid.Cid{test.RequireTestCID(t, []byte("parent1"))},
			ParentWeight:          big.NewInt(1234),
			Height:                34343,
			ParentStateRoot:       test.RequireTestCID(t, []byte("stateroot")),
			ParentMessageReceipts: test.RequireTestCID(t, []byte("receipts")),
			Messages:              test.RequireTestCID(t, []byte("messages")),
			BLSAggregate:          []byte("blsa"),
			Timestamp:             939393,
			BlockSig:              []byte("blocksig"),
		}
		tb := types.Block{
			ID:     test.RequireTestCID(t, []byte("block")),
			Header: tbh,
		}

		bhh := v1.BlockHandler{Callback: func(blockId string) (*types.Block, error) {
			return &tb, nil
		}}

		cbs := &server.V1Callbacks{GetBlockByID: bhh.Callback}
		s := test.CreateTestServer(t, cbs, false)
		s.Run()
		defer func() {
			assert.NoError(t, s.Shutdown())
		}()

		body := test.RequireGetResponseBody(t, s.Config().Port, "chain/blocks/1111")
		var actual types.Block
		require.NoError(t, json.Unmarshal(body, &actual))
		assert.True(t, actual.ID.Equals(tb.ID))
		assert.Equal(t, "block", actual.Kind)
		assert.Equal(t, "blockHeader", actual.Header.Kind)
		assert.Equal(t, tb.Header.Miner, actual.Header.Miner)
		assert.Equal(t, tb.Header.Tickets[0], actual.Header.Tickets[0])
		assert.Equal(t, tb.Header.ElectionProof, actual.Header.ElectionProof)
		assert.True(t, tb.Header.Parents[0].Equals(actual.Header.Parents[0]))
		assert.Equal(t, tb.Header.ParentWeight, actual.Header.ParentWeight)
		assert.Equal(t, tb.Header.Height, actual.Header.Height)
		assert.True(t, tb.Header.ParentStateRoot.Equals(actual.Header.ParentStateRoot))
		assert.True(t, tb.Header.ParentMessageReceipts.Equals(actual.Header.ParentMessageReceipts))
		assert.True(t, tb.Header.Messages.Equals(actual.Header.Messages))
		assert.Equal(t, tb.Header.BLSAggregate, actual.Header.BLSAggregate)
		assert.Equal(t, tb.Header.Timestamp, actual.Header.Timestamp)
		assert.Equal(t, tb.Header.BlockSig, actual.Header.BlockSig)
	})

}
