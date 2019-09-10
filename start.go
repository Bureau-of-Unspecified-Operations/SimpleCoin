package main

import (
	"fmt"
	"hash"
)

type txIn struct {
	pubKey    *PublicKey
	signature []byte
	block     *block
	txID      int
}

type txOut struct {
	pubKey *PublicKey
	coins  int
}

type transaction struct {
	inputs  []txIn
	outputs []txOut
	change  txOut
}

type block struct {
	prevHash uint64
	nonce    uint64
}
