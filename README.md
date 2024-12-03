# Teste Aprendendo Conceitos de Blockchain

Generic propose blockchain for uses like office files, criptocurrency

[![.github/workflows/go.yml](https://github.com/JoaoRafa19/crypto-go/actions/workflows/go.yml/badge.svg)](https://github.com/JoaoRafa19/crypto-go/actions/workflows/go.yml)


## Features

- [-] Server
- [X] Block
    - [X] Block's hash
    - [x] Test
- [X] Transaction
    - [x] Transaction list Hash
    - [x] Test
- [ ] Key
- [X] Transport => tcp, udp, 
    - [X] Local transport layer
- [ ] Crypto Keypairs and signature

## Todos
Improvements and fixes that can be implemented

## Types 

- Hash

```go
type Hash [32]uint8
```

### Mistakes to remember 

On the struct Transaction on Signing the transaction the object was missing the value of the transaction's `Signature`, returnning a null value 

```go
func TestSignTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()

	tx := &Transaction{
		Data: []byte("foo bar baz"),
	}

	assert.Nil(t, tx.Sign(privateKey))
	assert.NotNil(t, tx.Signature) // FAIL

}
```
Beacuse of the method signatue (tx T) insted of (tx *T), the object was missing the reference

**`old`** :
```go
func (tx Transaction) Sign(privKey crypto.PrivateKey) error {
```

**`fixed`**:
 ```go
func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
````
