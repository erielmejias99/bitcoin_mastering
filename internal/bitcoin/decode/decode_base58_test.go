package decode

import (
	"encoding/hex"
	"github.com/bitcoin_mastering/internal/bitcoin/consts"
	"github.com/bitcoin_mastering/internal/bitcoin/encode"
	"github.com/bitcoin_mastering/internal/bitcoin/key_generation"
	"math/big"
	"testing"
)

func TestDecodeStringToInt(t *testing.T){
	privateKeyBase58CheckWif := "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ"

	resp, err := decodeBase58CheckToInt( privateKeyBase58CheckWif )
	if err != nil{
		t.Errorf( "Error decoding base58CheckString" )
		t.FailNow()
	}

	hexFormat := hex.EncodeToString( resp.Bytes() )
	if hexFormat != "800c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d507a5b8d"{
		t.Errorf( "Incorrect response from the decoder, Must be 800c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d507a5b8d and was %s", hexFormat)
	}
}

func TestFindCharIndexInAlphabet(t *testing.T){

	for i, char := range consts.Base58Alphabet {
		index, err := findAlphabetIndex( uint8(char) )
		if err != nil{
			t.Errorf("Error looking for the index %s", err.Error() )
		}
		if index != i {
			t.Errorf( "Letter not found in base58 alphabet looking for %d char %s", i, string(consts.Base58Alphabet[i]) )
		}
	}
}

func TestFindNotCorrectLetterInAlphabet(t *testing.T) {

	//Base58 is Base64 without the 0 (number zero), O (capital o), l (lower L), I (capital i),
	//and the symbols “+” and “/”. Or, more simply, it is a set of lowercase and capital letters
	//and numbers without the four (0, O, l, I)
	//123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
	_, err := findAlphabetIndex( uint8('0') )
	if err == nil{
		t.Errorf("The char 0 doesn't belongs to base58 alphabet, must be an error here.")
	}
	_, err = findAlphabetIndex( uint8('O') )
	if err == nil{
		t.Errorf("The char 0 doesn't belongs to base58 alphabet, must be an error here.")
	}
	_, err = findAlphabetIndex( uint8('l') )
	if err == nil{
		t.Errorf("The char 0 doesn't belongs to base58 alphabet, must be an error here.")
	}
	_, err = findAlphabetIndex( uint8('I') )
	if err == nil{
		t.Errorf("The char 0 doesn't belongs to base58 alphabet, must be an error here.")
	}
}

func TestBase58CheckWifKey(t *testing.T) {
	base58Wif, err := Base58Check( "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil || base58Wif == nil {
		t.Error( err )
	}
	decodedWif := hex.EncodeToString( base58Wif.Bytes() )
	if decodedWif != "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"{
		t.Fail()
	}
}

func TestBase58CheckStringWifKey(t *testing.T) {
	base58Wif, err := Base58CheckString( "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil || base58Wif == "" {
		t.Error( err )
	}
	if base58Wif != "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"{
		t.Fail()
	}
}

func TestBase58CheckStringBadWifKey(t *testing.T) {
	_, err := Base58CheckString( "J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil {
		t.Log("Wrong prefix passed!")
	}
	_, err = Base58CheckString( "5J3mbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil {
		t.Log("Missing character passed!")
	}
	_, err = Base58CheckString( "5J3mbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2Jpbnkeyhfs45YB1Jcn" )
	if err != nil {
		t.Log("Extra character passed!")
	}
}

func TestEncodeDecodeWithRandomKeys(t *testing.T){
	const total = 10000
	var errorCount = 0
	for totalTest := total;totalTest > 0; totalTest--{
		privateKey := key_generation.GeneratePrivateKey()
		encodedPrivateKey := encode.EncodeBase58Check(encode.PrivateKeyWif, privateKey)
		decodedPrivateKey, err := Base58CheckString( encodedPrivateKey )
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

func TestEncodeDecodeWif(t *testing.T) {
	key := "69fa7023a22b383bc9d778ac1dbfee1b2e3d96a4f6aefa115873b64fb7923312"
	privateKey := new(big.Int)
	privateKey.SetString( key, 16 )
	encoded := encode.EncodeBase58Check(encode.PrivateKeyWif, privateKey )
	decoded, err := Base58CheckString( encoded )
	if err != nil{
		t.Errorf( "Encoded %s | Error: %s", encoded, err.Error() )
	}
	if key != decoded{
		t.FailNow();
	}
}