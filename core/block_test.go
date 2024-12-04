package core

import (
	"testing"
	"time"

	"github.com/JoaoRafa19/crypto-go/crypto"
	"github.com/JoaoRafa19/crypto-go/types"
	"github.com/stretchr/testify/assert"
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

func randomBlockWithSignature(t *testing.T, heigh uint32) *Block {
	privkey := crypto.GeneratePrivateKey()

	bc := randomBlock(heigh)

	assert.Nil(t, bc.Sign(privkey))
	return bc
}

func TestSignBlock(t *testing.T) {
	priv := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(priv))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	priv := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(priv))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())
	b.Height = 100
	assert.NotNil(t, b.Verify())
}
