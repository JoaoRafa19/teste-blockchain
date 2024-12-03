package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/JoaoRafa19/crypto-go/types"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:      1,
		PrevBlocHash: types.RandomHash(),
		Height:       height,
		Timestamp:    uint64(time.Now().UnixNano()),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestHashBlock(t *testing.T) {
	// b := new(Block)
	b := randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}
