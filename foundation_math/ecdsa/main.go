package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	// elliptic curve
	secp256k1 := btcec.S256()
	// Alice generate key pair
	privateKey, _ := btcec.NewPrivateKey(secp256k1)
	publicKey := privateKey.PubKey()
	fmt.Printf("private: \t%v\n", hex.EncodeToString(privateKey.Serialize()))
	fmt.Printf("public: \t%v\n", hex.EncodeToString(publicKey.SerializeUncompressed()))
	// msg
	msg := "hello"
	if len(os.Args) >= 2 {
		msg = os.Args[1]
	}
	fmt.Printf("message: \t%v\n", msg)
	// hash of the msg
	hash := sha256.Sum256([]byte(msg))
	fmt.Printf("sha256: \t%v\n", hex.EncodeToString(hash[:]))
	// sign hash with private key
	sign, _ := privateKey.Sign(hash[:])
	fmt.Printf("signature: \t%v\n", hex.EncodeToString(sign.Serialize()))
	// Alice send msg and signature to Bob
	sign2, _ := btcec.ParseSignature(sign.Serialize(), secp256k1)
	// Bob decrypted the hash and verify the msg
	fmt.Printf("verify: \t%v\n", sign2.Verify(hash[:], publicKey))
}
