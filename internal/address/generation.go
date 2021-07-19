package address

import (
	"fmt"
	"github.com/bitcoin_mastering/internal/bitcoin"
	"math/big"
)

func GeneratePrivateKey() * big.Int {
	var validPrivateKey bool
	var privateKey string
	decodedPrivateKey := new( big.Int )

	for !validPrivateKey{
		privateKey := bitcoin.RandomPrivateKey()
		decodedPrivateKey, _ = decodedPrivateKey.SetString( privateKey ,16)

		if decodedPrivateKey.Cmp( big.NewInt(0) ) == 1 &&
			decodedPrivateKey.Cmp( bitcoin.GetN() ) == -1{
			validPrivateKey = true
		}
	}

	fmt.Printf( "HEX: %s \nDEC: %s ", decodedPrivateKey.Text(16), privateKey )

	return decodedPrivateKey
}