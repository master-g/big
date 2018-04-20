package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type KeyInfo struct {
	privateKey   []byte   // 0
	publicKey    []byte   // 1
	sha256hash1  [32]byte // 2
	ripemd160of2 []byte   // 3
	version3     []byte   // 4
	sha256hash4  [32]byte // 5
	sha256hash5  [32]byte // 6
	first4bytes6 []byte   // 7
	step8        []byte   // 8
	Address      string
}

var (
	AddressTestNet byte = 0x6F
	AddressRealNet byte = 0x00
)

func NewKeyInfo() (keyInfo *KeyInfo, err error) {
	keyInfo = &KeyInfo{}

	// T("step 0. create an ECDSA secp256k1 private key")
	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		keyInfo = nil
		return
	}
	keyInfo.privateKey = privateKey.Serialize()

	// T("step 1. public key")
	keyInfo.publicKey = privateKey.PubKey().SerializeUncompressed()

	// T("step 2. SHA-256 hash of 2")
	keyInfo.sha256hash1 = sha256.Sum256(keyInfo.publicKey)

	// T("step 3. RIPEMD-160 hash of 3")
	hash := ripemd160.New()
	hash.Write(keyInfo.sha256hash1[:])
	keyInfo.ripemd160of2 = hash.Sum(nil)

	// T("step 4. address version to the begining of 4")
	keyInfo.version3 = append([]byte{AddressRealNet}, keyInfo.ripemd160of2...)

	// T("step 5. SHA-256 hash of 4")
	keyInfo.sha256hash4 = sha256.Sum256(keyInfo.version3)

	// T("step 6. SHA-256 hash of 5")
	keyInfo.sha256hash5 = sha256.Sum256(keyInfo.sha256hash4[:])

	// T("step 7. first 4 bytes of 6")
	keyInfo.first4bytes6 = keyInfo.sha256hash5[:4]

	// T("step 8. add 7 to the end of 4")
	keyInfo.step8 = append(keyInfo.version3, keyInfo.first4bytes6...)

	// T("step 9. Base58 encoding of 8")
	keyInfo.Address = base58.Encode(keyInfo.step8)

	return keyInfo, nil
}

func (ki *KeyInfo) Detail() {
	T("step 0. create an ECDSA secp256k1 private key")
	P(hexify(ki.privateKey))

	T("step 1. public key")
	P(hexify(ki.publicKey))

	T("step 2. SHA-256 hash of 2")
	P(hexify(ki.sha256hash1[:]))

	T("step 3. RIPEMD-160 hash of 3")
	P(hexify(ki.ripemd160of2))

	T("step 4. address version to the begining of 4")
	P(hexify(ki.version3))

	T("step 5. SHA-256 hash of 4")
	P(hexify(ki.sha256hash4[:]))

	T("step 6. SHA-256 hash of 5")
	P(hexify(ki.sha256hash5[:]))

	T("step 7. first 4 bytes of 6")
	P(hexify(ki.first4bytes6))

	T("step 8. add 7 to the end of 4")
	P(hexify(ki.step8))

	T("step 9. Base58 encoding of 8")
	P(ki.Address)
}

func T(s string) {
	fmt.Println(s)
}

func P(s string) {
	fmt.Println("    " + s)
}

func hexify(b []byte) string {
	return strings.ToUpper(hex.EncodeToString(b))
}

func Test() error {
	count := 0
	start := time.Now()
	for {
		keyInfo, err := NewKeyInfo()
		if err != nil {
			return err
		}

		count++

		if strings.HasPrefix(strings.ToLower(keyInfo.Address), "1") {
			keyInfo.Detail()
			fmt.Printf("cost %v runs\n", count)
			elapse := time.Now().Sub(start)
			fmt.Printf("took %v\n", elapse)
			break
		}
	}

	return nil
}

func main() {
	if err := Test(); err != nil {
		fmt.Errorf("error %v", err)
		os.Exit(1)
	}
}
