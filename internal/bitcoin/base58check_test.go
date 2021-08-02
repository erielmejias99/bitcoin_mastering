package bitcoin

import (
	"github.com/bitcoin_mastering/internal/bitcoin/consts"
	"github.com/bitcoin_mastering/internal/bitcoin/key_generation"
	"math/big"
	"testing"
)

func TestPrivateKeyEncodeWif(t *testing.T){
	privateKeyHex := "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"
	var privateKey,_ = new( big.Int ).SetString( privateKeyHex, 16 )
	encodePrivateKey, err  := Encode( consts.PrivateKeyWif, privateKey )
	if err != nil {
		t.Error( err.Error() )
	}
	if  encodePrivateKey != "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" {
		t.FailNow()
	}
}

func TestPrivateKeyDecodeWif(t *testing.T){
	privateKeyWif := "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"
	decodedPrivateKey, err  := Decode( privateKeyWif )
	if err != nil {
		t.Error( err.Error() )
	}
	if  decodedPrivateKey.Text(16) != "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd" {
		t.FailNow()
	}
}
func TestPrivateKeyEncodeWifCompressed(t *testing.T){
	privateKeyHex := "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"
	var privateKey,_ = new( big.Int ).SetString( privateKeyHex, 16 )
	encodePrivateKey, err  := Encode( consts.PrivateKeyWifCompressed, privateKey )
	if err != nil {
		t.Error( err.Error() )
	}
	if  encodePrivateKey != "KxFC1jmwwCoACiCAWZ3eXa96mBM6tb3TYzGmf6YwgdGWZgawvrtJ" {
		t.FailNow()
	}
}

func TestPrivateKeyDecodeWifCompressed(t *testing.T){
	privateKeyWifCompressed := "KxFC1jmwwCoACiCAWZ3eXa96mBM6tb3TYzGmf6YwgdGWZgawvrtJ"
	decodedPrivateKey, err  := Decode( privateKeyWifCompressed )
	if err != nil {
		t.Error( err.Error() )
	}
	if  decodedPrivateKey.Text(16) != "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd" {
		t.FailNow()
	}
}

func TestPrivateKeyDecodeStringWif(t *testing.T){
	privateKeyWif := "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"
	decodedPrivateKey, err  := DecodeString( privateKeyWif )
	if err != nil {
		t.Error( err.Error() )
	}
	if  decodedPrivateKey!= "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd" {
		t.FailNow()
	}
}

func TestPrivateKeyEncodeStringWif(t *testing.T){
	privateKeyHex := "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"
	encodePrivateKey, err  := EncodeString( consts.PrivateKeyWif, privateKeyHex )
	if err != nil {
		t.Error( err.Error() )
	}
	if  encodePrivateKey != "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" {
		t.FailNow()
	}
}

func TestEncodeDecodeWithRandomKeys(t *testing.T){
	const total = 1000
	var errorCount = 0
	for totalTest := total;totalTest > 0; totalTest--{
		privateKey := key_generation.GeneratePrivateKey()
		encodedPrivateKey, err := Encode( consts.PrivateKeyWif, privateKey)
		if err != nil {
			errorCount++
			t.Errorf( "PrivateKey: %s  Error: %s",
				privateKey.Text(16), err.Error() )
			continue
		}
		decodedPrivateKey, err := DecodeString( encodedPrivateKey )
		if err != nil{
			errorCount++
			t.Errorf( "PrivateKey: %s | EncodedWif %s | Error: %s",
				privateKey.Text(16), encodedPrivateKey, err.Error() )
			continue
		}
		if  privateKey.Text(16) != decodedPrivateKey{
			errorCount++
			t.Errorf( "Diferent keys")
		}
	}
	if errorCount != 0{
		t.Errorf("Errors/TotalCases -> %d/%d", errorCount, total )
	}
}

func TestEncodeDecodeWithRandomKeysBigInt(t *testing.T){
	const total = 5
	var errorCount = 0
	for totalTest := total;totalTest > 0; totalTest--{
		privateKey := key_generation.GeneratePrivateKey()

		// encodes Wif
		encodedPrivateKeyWif, err := Encode( consts.PrivateKeyWif, privateKey)
		if err != nil {
			errorCount++
			t.Errorf( "PrivateKey: %s  Error: %s",
				privateKey.Text(16), err.Error() )
			continue
		}
		decodedPrivateKey, err := Decode( encodedPrivateKeyWif )
		if err != nil{
			errorCount++
			t.Errorf( "PrivateKey: %s | EncodedWif %s | Error: %s",
				privateKey.Text(16), encodedPrivateKeyWif, err.Error() )
			continue
		}
		if  privateKey.Cmp( decodedPrivateKey ) != 0{
			errorCount++
			t.Errorf( "Diferent keys")
		}

		// WifCompressed
		encodedPrivateKeyWifCompressed, err := Encode( consts.PrivateKeyWifCompressed, privateKey)
		if err != nil {
			errorCount++
			t.Errorf( "PrivateKey: %s  Error: %s",
				privateKey.Text(16), err.Error() )
			continue
		}
		decodedPrivateKeyFromWifCompressed, err := DecodeString( encodedPrivateKeyWifCompressed )
		if err != nil {
			errorCount++
			t.Errorf( "Error decoding compressed wif %s", err.Error())
			continue
		}
		if  privateKey.Text(16) != decodedPrivateKeyFromWifCompressed +"01"{
			errorCount++
			t.Errorf( "Diferent keys Compressed %s -> %s", privateKey.Text(16), decodedPrivateKeyFromWifCompressed)
		}
	}
	if errorCount != 0{
		t.Errorf("Errors/TotalCases -> %d/%d", errorCount, total )
	}
}