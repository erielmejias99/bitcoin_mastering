package key_generation

import (
	"github.com/bitcoin_mastering/internal/bitcoin/consts"
	"math/big"
)

func GeneratePrivateKey() * big.Int {
	var validPrivateKey bool
	//var privateKey string
	decodedPrivateKey := new( big.Int )

	for !validPrivateKey{
		privateKey := RandomPrivateKey()
		decodedPrivateKey, _ = decodedPrivateKey.SetString( privateKey ,16)

		if decodedPrivateKey.Cmp( big.NewInt(0) ) == 1 &&
			decodedPrivateKey.Cmp( consts.GetN() ) == -1 &&
			len( decodedPrivateKey.Text(16) ) == 64{
			validPrivateKey = true
		}
	}

	//fmt.Printf( "HEX: %s \nDEC: %s ", decodedPrivateKey.Text(16), privateKey )

	return decodedPrivateKey
}