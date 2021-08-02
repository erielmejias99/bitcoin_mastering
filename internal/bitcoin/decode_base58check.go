package bitcoin

import (
	"errors"
	"math/big"
)

func discoverEncode( key string ) (EncodePrefix, error){
	if key[ 0 ] == '5' || key[ 0 ] == 'K' || key[ 0 ] == 'L'{
		return PrivateKeyWif, nil
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

		//sum with the total kwy calculation
		keyInt.Add( keyInt, baseExp )
	}

	return keyInt, nil
}

func DecodeBase58Check( key string ) (*big.Int, error){

	// validate all chars in the key belongs to base58Check alphabet
	for i, char := range alphabet{
		index, err := findAlphabetIndex( uint8(char) )
		if err != nil || index != i {
			return nil, errors.New("invalid key, contains invalid characters")
		}
	}

	//determine encode format
	format, err := discoverEncode( key )
	if err != nil{
		return nil, &InvalidPrivateKey{hint: "Unknown prefix format." }
	}

	// Validate size depending on Key Format
	switch format {
	case PrivateKeyWif:
		if len( key ) != 51 {
			if len( key ) > 51{
				return nil, errors.New("too long key")
			}
			return nil, errors.New("too short key")
		}
	}

	// decode to big.Int
	keyInt, err := decodeBase58CheckToInt( key )
	if err != nil{
		return nil, err
	}

	//remove firsts version byte
	bytesKey := keyInt.Bytes()[ format.byteSize() :]

	//remove 4 checksum bytes
	bytesKey = bytesKey[0:len( bytesKey ) - 4 ]

	// convert to big.Int
	keyInt.SetBytes( bytesKey )

	return keyInt, nil
}


func DecodeBase58CheckString( key string ) (string, error){
	bigInt, err := DecodeBase58Check(key)
	if err != nil || bigInt == nil {
		return "", err
	}
	return bigInt.Text( 16 ), nil
}