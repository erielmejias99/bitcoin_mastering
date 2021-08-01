package bitcoin

import "encoding/hex"

type EncodePrefix string

//HEX string
const (
	PrivateKeyWif  EncodePrefix = "80" //5 K L
	BitcoinAddress EncodePrefix = "00" // 1
	PSHAddress     EncodePrefix = "05" //3
	BitcoinTestNet EncodePrefix = "6f"
	BIP38EncryptedPrivateKey EncodePrefix = "0142" // 6P
	BIP32ExtendedPublicKey EncodePrefix = "0488b21e" // xpub
)

func (p EncodePrefix) Value() string {
	return string( p )
}

func (p EncodePrefix) Bytes() []byte {
	b, err := hex.DecodeString( p.Value() )
	if err != nil {
		panic( "Invalid prefix, error decoding to bytes" )
	}
	return b
}

func (p EncodePrefix ) byteSize() int{
	return len( p.Bytes() )
}