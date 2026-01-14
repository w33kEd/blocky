package main

import (
	"fmt"
	"log"

	"github.com/w33ked/go-blocky/wallet"
)

// helper functions
func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
}
