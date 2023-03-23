package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"

	"github.com/realjf/blockchain_demo/helper"
)

const (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	PubKey     []byte
}

func (w *Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PubKey)

	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := helper.Base58Encode(fullHash)

	// fmt.Printf("pub key: %x\n", w.PublicKey)
	// fmt.Printf("pub hash: %x\n", pubHash)
	// fmt.Printf("address: %x\n", address)

	return address
}

func ValidateAddress(address string) bool {
	pubKeyHash := helper.Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-checksumLength:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-checksumLength]
	targetChecksum := Checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Equal(actualChecksum, targetChecksum)
}

func NewKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, []byte) {
	curve := elliptic.P521()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	helper.Handle(err)

	pubKey := elliptic.Marshal(curve, private.PublicKey.X, private.PublicKey.Y)
	return private, &private.PublicKey, pubKey
}

func MakeWallet() *Wallet {
	private, public, pubKey := NewKeyPair()
	return &Wallet{
		PrivateKey: private,
		PublicKey:  public,
		PubKey:     pubKey,
	}
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	helper.Handle(err)

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}
