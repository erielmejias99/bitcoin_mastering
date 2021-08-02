package main

import (
	"fmt"
	"github.com/bitcoin_mastering/internal/bitcoin"
)

func main() {
	privateKey := bitcoin.GeneratePrivateKey()
	fmt.Printf( "%s", privateKey.Text(10) )
}