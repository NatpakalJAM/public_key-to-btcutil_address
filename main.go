package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

// Keys = input public key
var Keys = []string{
	"0x040dc6785daab4f31d390cf53a9c04a39088380ac9d4828da48f5e90273d51bb0ceaac56d963074282d87907d98487ba9c29e5719409c4cc42c8d646ae7c190b42",
	"0x020dc6785daab4f31d390cf53a9c04a39088380ac9d4828da48f5e90273d51bb0c",
	"0250863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352",
	"0450863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b23522cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6",
}

func main() {
	for _, key := range Keys {
		findAddress(key)
	}
}

func findAddress(pubKey string) {
	key := pubKey
	key = strings.TrimPrefix(key, "0x")

	// hex string -> []byte
	bKey, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}

	// PubKey -> Address
	mainNetAddr, err := btcutil.NewAddressPubKey(bKey, &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}

	// result
	fmt.Printf("/*--------------------------------------------------------------*/\n")
	fmt.Printf("Input Public Key: %s\n", pubKey)
	fmt.Printf("Format: %v\n", getPubKeyFormat(mainNetAddr.Format()))
	fmt.Printf("Address: %s\n", mainNetAddr.EncodeAddress())
	fmt.Printf("/*--------------------------------------------------------------*/\n\n\n")
}

func getPubKeyFormat(format btcutil.PubKeyFormat) (strPubKeyFormat string) {
	switch format {
	case 0:
		strPubKeyFormat = "Uncompressed"
	case 1:
		strPubKeyFormat = "Compressed"
	default:
		strPubKeyFormat = "Unknown"
	}
	return
}
