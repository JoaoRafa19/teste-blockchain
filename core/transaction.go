package core

import "io"

type Transaction struct {
	Data []byte
	
}

func (tx *Transaction) DecodeBynary(r io.Reader) error {
	return nil
}

func (tx *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}
