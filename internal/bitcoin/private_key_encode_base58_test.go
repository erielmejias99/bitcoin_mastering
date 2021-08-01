package bitcoin

import (
	"encoding/hex"
	"math/big"
	"testing"
)

const privateKeyHex = "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"

func TestPrivateKeyEncodeHex(t *testing.T) {
	var privateKey,_ = new( big.Int ).SetString(privateKeyHex, 16 )

	if EncodePrivateKey( privateKey, Hex ) != "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd" {
		t.FailNow()
	}
}

func TestPrivateKeyEncodeWif(t *testing.T){
	var privateKey,_ = new( big.Int ).SetString(privateKeyHex, 16 )
	encodePrivateKey := EncodePrivateKey( privateKey, Wif )
	if  encodePrivateKey != "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" {
		t.FailNow()
	}
}

func TestWifCheckSum(t *testing.T) {
	payloadByte, _ := hex.DecodeString( "0c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d")
	base58Checksum := hex.EncodeToString(  checksum ( WifPrefix, payloadByte ) )
	if base58Checksum != "507a5b8d"{
		t.FailNow()
	}
}

func TestWifPrivateKeyEncode(t *testing.T) {
	//test case 1
	privateKey, _ := new(big.Int).SetString("0c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d", 16 )
	privateKeyWifFmt := EncodePrivateKey( privateKey, Wif )
	if privateKeyWifFmt != "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ"{
		t.Logf("Returned %s must be %s", privateKeyWifFmt, "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ" )
		t.Fail()
	}

	//test case 2
	privateKey, _ = new(big.Int).SetString("1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd", 16 )
	privateKeyWifFmt = EncodePrivateKey( privateKey, Wif )
	if privateKeyWifFmt != "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"{
		t.Logf("Returned %s must be %s", privateKeyWifFmt, "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
		t.FailNow()
	}
}