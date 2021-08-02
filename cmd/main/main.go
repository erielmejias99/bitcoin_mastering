package main

import (
	"fmt"
	"github.com/bitcoin_mastering/internal/bitcoin/key_generation"
)

func main() {
	privateKey := key_generation.GeneratePrivateKey()
	fmt.Printf( "%s", privateKey.Text(10) )
}