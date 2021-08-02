package encode

import (
	"crypto/sha256"
	"github.com/bitcoin_mastering/internal/bitcoin/consts"
	"math/big"
)

func EncodeBase58Check( prefix EncodePrefix, key * big.Int ) string{
	checkSum := checksum( prefix, key.Bytes() )

	// add prefix
	fullPayload := prefix.Bytes()
	// add payload
	fullPayload = append(fullPayload, key.Bytes()... )
	//add checksum
	fullPayload = append(fullPayload, checkSum... )

	// Convert to big.Int to be able to change the base
	number := new( big.Int ).SetBytes( fullPayload  )

	var digit = new(big.Int)
	base := big.NewInt(58)

	invertedFormattedKey := ""
	zero := big.NewInt(0)
	for number.Cmp( zero ) != 0{
		digit.Mod( number, base )
		invertedFormattedKey += string(consts.Base58Alphabet[ digit.Int64() ])
		number.Div( number, base )
	}

	formattedKey := ""
	for i := len(invertedFormattedKey) - 1; i >= 0; i--{
		formattedKey += string(invertedFormattedKey[ i ])
	}
	return formattedKey
}

func checksum( prefix EncodePrefix, payload []byte) []byte  {
	fullPayload := prefix.Bytes()
	fullPayload = append(fullPayload,  payload... )

	fistSha256 := sha256.Sum256( fullPayload )
	secondSha256 := sha256.Sum256( fistSha256[:] )

	return secondSha256[:4]
}