package key_generation

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
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

	privateKeySha := sha256.Sum256( privateKey )
	return hex.EncodeToString( privateKeySha[:] )
}


func RandomString( n int ) ([]byte, error )  {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return b, err
	}
	return b, nil
}
