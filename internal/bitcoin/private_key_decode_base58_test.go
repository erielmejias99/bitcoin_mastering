package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestDecodeStringToInt(t *testing.T){
	privateKeyBase58CheckWif := "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ"

	resp := decodeBase58CheckToInt( privateKeyBase58CheckWif )
	if resp == nil{
		t.Error( "Error decoding base58CheckString" )
		return
	}

	hexFormat := hex.EncodeToString( resp.Bytes() )
	if hexFormat != "800c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d507a5b8d"{
		t.Errorf( "Incorrect response from the decoder, Must be 800c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d507a5b8d and was %s", hexFormat)
	}
}

func TestFindCharIndexInAphabet(t *testing.T){

	for i, char := range alphabet{
		if i == 57{
			print( "sd")
		}
		index, err := findAlphabetIndex( uint8(char) )
		if err != nil{
			t.Errorf("Error looking for the index %s", err.Error() )
		}
		if index != i {
			t.Errorf( "Letter not found in base58 alphabet looking for %d char %s", i, string(alphabet[i]) )
		}
	}
}