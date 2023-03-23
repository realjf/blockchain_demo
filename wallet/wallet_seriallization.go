package wallet

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/realjf/blockchain_demo/helper"
)

type WalletIterator struct {
	PrivateKey    string
	PublicKey     string
	PubKey        []byte
	WalletAddress string
}

type WalletSerializer struct {
	Wallets map[string]WalletIterator
}

func SerializeWallets(wallets *Wallets) {
	serializer := WalletSerializer{
		Wallets: map[string]WalletIterator{},
	}
	for _, w := range wallets.Wallets {
		address := string(w.Address()[:])
		private, public := helper.Encode(w.PrivateKey, w.PublicKey)
		serializer.Wallets[address] = WalletIterator{
			PrivateKey:    private,
			PublicKey:     public,
			PubKey:        w.PubKey,
			WalletAddress: address,
		}
	}
	var content bytes.Buffer

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(serializer)
	helper.Handle(err)

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	helper.Handle(err)
	fmt.Println("save wallets file!")
}

func DeserializeWallets() (*Wallets, error) {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return nil, err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return nil, err
		}
		fmt.Println("read EOF: wallets file!")
	}

	serializer := WalletSerializer{
		Wallets: map[string]WalletIterator{},
	}

	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&serializer)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return nil, err
		}
		fmt.Println("wallets file is empty!")
	}

	wallets := &Wallets{
		Wallets: make(map[string]*Wallet),
	}

	for address, w := range serializer.Wallets {
		private, public := helper.Decode(w.PrivateKey, w.PublicKey)

		wallets.Wallets[address] = &Wallet{
			PrivateKey: private,
			PublicKey:  public,
			PubKey:     w.PubKey,
		}
	}

	return wallets, nil
}
