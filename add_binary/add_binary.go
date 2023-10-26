package add_binary

import (
	"log"
	"math/big"
)

func addBinary(a string, b string) string {
	abi := big.NewInt(0)
	_, ok := abi.SetString(a, 2)
	if !ok {
		log.Fatalf("can't set string")
	}
	bbi := big.NewInt(0)
	_, ok = bbi.SetString(b, 2)
	if !ok {
		log.Fatalf("can't set string")
	}
	abi.Add(abi, bbi)

	return abi.Text(2)
}
