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
	if hexFormat !=
}
