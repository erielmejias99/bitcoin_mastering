package bitcoin

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"time"
)


func RandomPrivateKey() string {
	privateKey, _ := RandomString(32)

	maxRand := new(big.Int)
	maxRand.SetBit( big.NewInt(0) , 255,1)
	randInt, err := rand.Int ( rand.Reader, maxRand )
	if err != nil {
		return ""
	}

	privateKey = append(privateKey, randInt.Bytes()...)
	privateKey = append(privateKey, []byte( strconv.Itoa( int(time.Now().UnixNano()) ) )... )

	return fmt.Sprintf("%X", sha256.Sum256( privateKey ))
}


func RandomString( n int ) ([]byte, error )  {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return b, err
	}
	return b, nil
}


func EncodePrivateKey( key *big.Int, format EncodeFormat ) string{
	switch format {
	case Hex:
		encoded := key.Text(16)
		return encoded
	case HexCompressed:
		return key.Text(16) + "01"
	case Wif:
		return EncodeBase58Check(PrivateKeyWif, key )
	case WifCompressed:
		return ""
		//return Encode( []byte(key.Text(10) + "01" ) )
	default:
		panic("Invalid format")
	}
}