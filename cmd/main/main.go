package main

import (
	"fmt"
	"github.com/bitcoin_mastering/internal/address"
)

func main() {
	privateKey := address.GeneratePrivateKey()
	fmt.Printf( "%s", privateKey.Text(10) )
}