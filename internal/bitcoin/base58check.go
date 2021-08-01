package bitcoin

import (
	"crypto/sha256"
	"errors"
	"math/big"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func EncodeBase58Check( prefix PrivatePrefix, key * big.Int ) string{
	checkSum := checksum( prefix, key.Bytes() )

	// add prefix
	fullPayload := prefix.Bytes()
	// add payload
	fullPayload = append(fullPayload, key.Bytes()... )
	//add checksum
	fullPayload = append(fullPayload, checkSum... )

	// Convert to big.Int to be able to change the base
	number := new( big.Int ).SetBytes( fullPayload  )

	var digit = new(big.Int)
	base := big.NewInt(58)

	invertedFormattedKey := ""
	zero := big.NewInt(0)
	for number.Cmp( zero ) != 0{
		digit.Mod( number, base )
		invertedFormattedKey += string(alphabet[ digit.Int64() ])
		number.Div( number, base )
	}

	formattedKey := ""
	for i := len(invertedFormattedKey) - 1; i >= 0; i--{
		formattedKey += string(invertedFormattedKey[ i ])
	}
	return formattedKey
}

func discoverEncode( key string ) (PrivatePrefix, error){
	if key[ 0 ] == '5'{
		return WifPrefix, nil
	}
	return "",errors.New("PrefixFormat not found")
}

func findAlphabetIndex( char uint8 ) (index int, err error) {

	var left, right, middle = 0, len( alphabet )-1, 0
	for left <= right{
		middle = ( left + right )/2
		if alphabet[ middle ] == char {
			return middle, nil
		}else{
			if char > alphabet[ middle ] {
				left = middle + 1
			}else{
				right = middle - 1
			}
		}
	}

	return -1, errors.New("character not found in alphabet" )
}

func decodeBase58CheckToInt( key string ) (*big.Int, error){

	//calculate last value first
	index, err := findAlphabetIndex( key[ len(key)-1 ] )
	if err !=nil{
		return nil, LetterNotFound{ char: key[ len(key)-1 ] }
	}
	keyInt := big.NewInt( int64(index) )
	base58 := big.NewInt(58 )

	for i := len( key ) - 2; i >= 0; i--{
		index, err := findAlphabetIndex( key[ i ] )
		if err != nil{
			return nil,err
		}
		//calculate pow
		baseExp := big.NewInt( 58 )
		for expIndex := len( key ) - i - 1; expIndex > 1; expIndex --{
			baseExp.Mul( baseExp, base58 )
		}
		//multiply by the index of the letter
		baseExp.Mul( baseExp, big.NewInt( int64( index ) ) )

		// sum with the total kwy calculation
		keyInt.Add( keyInt, baseExp )
	}

	return keyInt, nil
}

func DecodeBase58Check( key string ) (string, error){

	// Todo create regex just with de alphabet of Base58

	// Todo validate size depending on Type

	// decode tu Hex

	//determine encode format
	//format, err := discoverEncode( key )
	//if err != nil{
	//	return "", &InvalidPrivateKey{hint: "Unknown prefix format." }
	//}

	//remove checksum
	//5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ

	return "", nil
}

//func EncodeBase58HexString( prefix PrivatePrefix, key * big.Int ) string{
//	base58Byte := EncodeBeforeBaseBase58( prefix, key )
//	return hex.EncodeToString( base58Byte )
//}

func checksum( prefix PrivatePrefix, payload []byte) []byte  {
	fullPayload := prefix.Bytes()
	fullPayload = append(fullPayload,  payload... )

	fistSha256 := sha256.Sum256( fullPayload )
	secondSha256 := sha256.Sum256( fistSha256[:] )

	return secondSha256[:4]
}