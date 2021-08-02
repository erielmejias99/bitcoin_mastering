package bitcoin

import (
	"encoding/hex"
	"errors"
	"github.com/bitcoin_mastering/internal/bitcoin/consts"
	"github.com/bitcoin_mastering/internal/bitcoin/decode"
	"github.com/bitcoin_mastering/internal/bitcoin/encode"
	"math/big"
	"strings"
)

func Encode( format consts.EncodeFormat, key *big.Int ) (string, error){

	switch format {
	case consts.Hex:
		encoded := key.Text(16)
		return encoded, nil
	case consts.HexCompressed:
		return key.Text(16) + "01", nil
	case consts.Wif:
		return encode.EncodeBase58Check( encode.PrivateKeyWif, key ), nil
	case consts.WifCompressed:
		suffix,_ := hex.DecodeString( "01" )
		key.Append( suffix, 16 )
		return encode.EncodeBase58Check(encode.PrivateKeyWif, key ), nil
	default:
		panic("Invalid format")
	}
}

func EncodeString( format consts.EncodeFormat, key string ) (string, error){

	// In golang hex digit are in lowercase
	key = strings.ToLower( key )

	keyInt, ok := new(big.Int).SetString( key, 16 )
	if !ok {
		return "", errors.New("error decoding key to big.Int")
	}
	return Encode( format, keyInt )
}


func Decode( key string ) (*big.Int, error ){
	return decode.Base58Check( key )
}

func DecodeString( key string ) (string,error){
	return decode.Base58CheckString( key )
}
