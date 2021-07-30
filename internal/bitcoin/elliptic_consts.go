package bitcoin

import "math/big"

type EllipticParam string

const(
	P  EllipticParam = "115792089237316195423570985008687907853269984665640564039457584007908834671663"
	A  EllipticParam = "0"
	B  EllipticParam = "7"
	N  EllipticParam = "115792089237316195423570985008687907852837564279074904382605163141518161494337"
	Gx EllipticParam = "55066263022277343669578718895168534326250603453777594175500187360389116729240"
	Gy EllipticParam = "32670510020758816978083085130507043184471273380659243275938904335757337482424"
)


type EllipticPoint struct {
	X EllipticParam
	Y EllipticParam
}

var n * big.Int
func GetN() * big.Int{
	if n == nil || n.Text(10) != N.Value() {
		n, _ = new( big.Int ).SetString( "115792089237316195423570985008687907852837564279074904382605163141518161494337" ,10)
	}
	return n
}

func ( e EllipticParam ) Value() string{
	return string(e)
}

type EncodeFormat string

const(
	Hex 			EncodeFormat = "hex"
	Wif  			EncodeFormat = "wif"
	HexCompressed 	EncodeFormat = "hex_compressed"
	WifCompressed 	EncodeFormat = "wif_compressed"
)

func (e EncodeFormat) Value() string  {
	return string( e )
}