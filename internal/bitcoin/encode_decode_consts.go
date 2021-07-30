package bitcoin

import "encoding/hex"

type PrivatePrefix string

//HEX string
const (
	WifPrefix PrivatePrefix = "80"
)

func (p PrivatePrefix) Value() string {
	return string( p )
}

func (p PrivatePrefix) Bytes() []byte {
	b, err := hex.DecodeString( p.Value() )
	if err != nil {
		panic( "Invalid prefix, error decoding to bytes" )
	}
	return b
}