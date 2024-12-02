package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/JoaoRafa19/teste-blockchain/types"
)

// To save space we just hash the header
type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBynary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &h.Nonce)

}
func (h *Header) DecodeBynary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

// Hold the transactions and the header information
type Block struct {
	Header
	Transactions []Transaction

	// cached version of the header hash
	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	b.Header.EncodeBynary(buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	return b.hash
}

func (b *Block) EncodeBynary(w io.Writer) error {
	if err := b.Header.EncodeBynary(w); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil

}
func (b *Block) DecodeBynary(r io.Reader) error {
	if err := b.Header.DecodeBynary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.DecodeBynary(r); err != nil {
			return err
		}
	}
	return nil
}
