package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestDecodeStringToInt(t *testing.T){
	privateKeyBase58CheckWif := "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ"

	resp, err := decodeBase58CheckToInt( privateKeyBase58CheckWif )
	if err != nil{
		t.Errorf( "Error decoding base58CheckString" )
		return
	}

	hexFormat := hex.EncodeToString( resp.Bytes() )
	if hexFormat != "800c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d507a5b8d"{
		t.Errorf( "Incorrect response from the decoder, Must be 800c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d507a5b8d and was %s", hexFormat)
	}
}

func TestFindCharIndexInAlphabet(t *testing.T){

	for i, char := range alphabet{
		index, err := findAlphabetIndex( uint8(char) )
		if err != nil{
			t.Errorf("Error looking for the index %s", err.Error() )
		}
		if index != i {
			t.Errorf( "Letter not found in base58 alphabet looking for %d char %s", i, string(alphabet[i]) )
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

func TestDecodeBase58CheckWifKey(t *testing.T) {
	base58Wif, err := DecodeBase58Check( "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil || base58Wif == nil {
		t.Error( err )
	}
	decodedWif := hex.EncodeToString( base58Wif.Bytes() )
	if decodedWif != "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"{
		t.Fail()
	}
}

func TestDecodeBase58CheckStringWifKey(t *testing.T) {
	base58Wif, err := DecodeBase58CheckString( "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil || base58Wif == "" {
		t.Error( err )
	}
	if base58Wif != "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"{
		t.Fail()
	}
}

func TestDecodeBase58CheckStringBadWifKey(t *testing.T) {
	_, err := DecodeBase58CheckString( "J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil {
		t.Log("Wrong prefix passed!")
	}
	_, err = DecodeBase58CheckString( "5J3mbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn" )
	if err != nil {
		t.Log("Missing character passed!")
	}
	_, err = DecodeBase58CheckString( "5J3mbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2Jpbnkeyhfs45YB1Jcn" )
	if err != nil {
		t.Log("Extra character passed!")
	}
}