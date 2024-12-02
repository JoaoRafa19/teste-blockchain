package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/JoaoRafa19/teste-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     123123123,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBynary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBynary(buf))
	assert.Equal(t, h, hDecode)
}

func Test_Encoding_Decoding_Block(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     123123123,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBynary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBynary(buf))

	assert.Equal(t, b, bDecode, "Encoded and decoded block mut be equal")
}

func Test_Block_Hash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
		},
		Transactions: []Transaction{},
	}

	h := b.Hash()
	fmt.Println(h)
	assert.False(t, h.IsZero())
}
