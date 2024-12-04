package core

import (
	"testing"

	"github.com/JoaoRafa19/crypto-go/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()

	tx := &Transaction{
		Data: []byte("foo bar baz"),
	}

	assert.Nil(t, tx.Sign(privateKey))
	assert.NotNil(t, tx.Signature)

}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo bar baz"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo bar baz"),
	}

	assert.Nil(t, tx.Sign(privKey))
	return tx
}
