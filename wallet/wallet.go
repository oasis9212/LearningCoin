package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"ralo/utils"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
}

var w *wallet

const (
	fileName      string = "nomadcoin.wallet"
	signature            = "213f11232963141fc01289441e95107a9788d16988ce152e23f2bb9c9df4c0638ab39117239d51522b2f91217337a44af8ba9af9bd1f5dab42f0d35e0c924ce2"
	privateKey           = "30770201010420b5134fb80d9e1e846e78786a7452fe54034b2311092e803210729dbb4d76cc04a00a06082a8648ce3d030107a1440342000499e88c044fb0c21b2b11bde2d88558e693a6245f186b3dc519bc40733c36e69651f493be14fbb4ffdb9dadbfce739529bd7d9f04e50b1e2f36b6fe0a20493a0e"
	hashedmessage        = "53a08f7f927728ca9bb30852cdd1b6636584c60942ab01680ed4242b3ad284ef"
)

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() (key *ecdsa.PrivateKey) {
	keyAsbytes, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err = x509.ParseECPrivateKey(keyAsbytes)
	utils.HandleErr(err)
	return
}

func aFromKey(key *ecdsa.PrivateKey) string {
	return encodeBigInts(key.X.Bytes(), key.Y.Bytes())
}

func encodeBigInts(a, b []byte) string {
	z := append(a, b...)

	return fmt.Sprintf("%x", z)
}

func Sign(payload string, w *wallet) string {
	payloadASbytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadASbytes)
	utils.HandleErr(err)
	return encodeBigInts(r.Bytes(), s.Bytes())

}

func restoreBigInts(signature string) (*big.Int, *big.Int, error) {
	Bytes, err := hex.DecodeString(signature)
	if err != nil {
		return nil, nil, err
	}

	firstHalfBytes := Bytes[:len(Bytes)/2]
	secondHalfBytes := Bytes[len(Bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalfBytes)
	bigB.SetBytes(secondHalfBytes)

	return &bigA, &bigB, nil
}

func Verify(signture, payload, address string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleErr(err)
	x, y, err := restoreBigInts(address)
	utils.HandleErr(err)

	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	payloadbytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	ok := ecdsa.Verify(&publicKey, payloadbytes, r, s)
	return ok
}

func Wallet() *wallet {

	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = aFromKey(w.privateKey)
	}
	return w
}
