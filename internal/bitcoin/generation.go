package bitcoin

import (
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
			decodedPrivateKey.Cmp( GetN() ) == -1 &&
			len( decodedPrivateKey.Text(16) ) == 64{
			validPrivateKey = true
		}
	}

	//fmt.Printf( "HEX: %s \nDEC: %s ", decodedPrivateKey.Text(16), privateKey )

	return decodedPrivateKey
}