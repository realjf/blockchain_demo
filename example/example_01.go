package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "Wallets.dat"

type Wallets struct {
	WalletsMap map[string]*Wallet
}

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func (w *Wallets) SaveWallets() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(&w)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}

}

func NewWallets() (*Wallets, error) {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		wallets := &Wallets{}
		wallets.WalletsMap = make(map[string]*Wallet)
		return wallets, err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			log.Panic(err)
		}
	}

	wallets := Wallets{
		WalletsMap: make(map[string]*Wallet),
	}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			log.Panic(err)
		}
	}

	return &wallets, nil
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{
		PrivateKey: private,
		PublicKey:  public,
	}

	return &wallet
}

func main() {
	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	newWallet := MakeWallet()
	wallets.WalletsMap["1"] = newWallet
	wallets.SaveWallets()
}
