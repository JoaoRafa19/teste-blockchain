package core

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type BlockChain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

func NewBlockChain(genesis *Block) (*BlockChain, error) {
	bc := &BlockChain{
		headers: []*Header{},
		store:   NewMemStore(),
	}
	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)

	return bc, err

}

func (bc *BlockChain) SetValidator(v Validator) {
	bc.validator = v
}
func (bc *BlockChain) AddBlock(b *Block) error {
	//validate
	err := bc.validator.ValidateBlock(b)
	if err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

func (bc *BlockChain) HasBlock(heigh uint32) bool {
	return heigh <= bc.Height()
}

// [g, 1, 2, 3] = len 4 ; heigh = 3
func (bc *BlockChain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *BlockChain) addBlockWithoutValidation(b *Block) error {

	logrus.WithField("Adding New Block", logrus.Fields{
		"height": b.Height,
		"hash":   b.Hash(BlockHasher{}),
	}).Info("adding new block")

	bc.headers = append(bc.headers, b.Header)

	return bc.store.Put(b)
}

func (bc *BlockChain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("height (%+v) is too high", height)
	}

	return bc.headers[height], nil
}
