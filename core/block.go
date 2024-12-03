package core

import (
	"io"

	"github.com/JoaoRafa19/crypto-go/crypto"
	"github.com/JoaoRafa19/crypto-go/types"
)

// To save space we just hash the header
type Header struct {
	Version      uint32
	PrevBlocHash types.Hash
	Timestamp    uint64
	Height       uint32
	DataHash     uint32
}

// Hold the transactions and the header information
type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    crypto.Signature

	// cached version of the header hash
	hash types.Hash
}

func NewBlock(h *Header, tsc []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: tsc,
	}
}


func (b *Block) Sign(privKey crypto.PrivateKey) *crypto.Signature {
	sig,err := privKey.Sign()
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}


func (b*Block) hashableData() []byte {
	
}