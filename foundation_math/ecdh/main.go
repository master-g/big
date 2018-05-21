package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	// Elliptic curve
	secp256k1 := btcec.S256()
	// Alice
	alicePrivate, _ := btcec.NewPrivateKey(secp256k1)
	alicePublic := alicePrivate.PubKey()
	// Bob
	bobPrivate, _ := btcec.NewPrivateKey(secp256k1)
	bobPublic := bobPrivate.PubKey()
	// Shared secret
	aliceS := btcec.GenerateSharedSecret(alicePrivate, bobPublic)
	bobS := btcec.GenerateSharedSecret(bobPrivate, alicePublic)
	// Results
	fmt.Printf("%v: \t%v\n", "Alice's private key",
		hex.EncodeToString(alicePrivate.Serialize()))
	fmt.Printf("%v: \t%v\n", "Alice's public key",
		hex.EncodeToString(alicePublic.SerializeUncompressed()))
	fmt.Printf("%v: \t%v\n", "Bob's private key",
		hex.EncodeToString(bobPrivate.Serialize()))
	fmt.Printf("%v: \t%v\n", "Bob's public key",
		hex.EncodeToString(bobPublic.SerializeUncompressed()))
	fmt.Printf("%v: \t%v\n", "Alice's shared secret",
		hex.EncodeToString(aliceS))
	fmt.Printf("%v: \t%v\n", "Bob's shared secret",
		hex.EncodeToString(bobS))
}
