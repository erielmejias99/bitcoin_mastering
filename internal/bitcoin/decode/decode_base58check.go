package decode

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bitcoin_mastering/internal/bitcoin/consts"
	"github.com/bitcoin_mastering/internal/bitcoin/encode"
	"math/big"
)

func discoverEncodePrefix( key string ) (encode.EncodePrefix, error){
	if key[ 0 ] == '5'{
		return encode.PrivateKeyWif, nil
	}
	if  key[ 0 ] == 'K' || key[ 0 ] == 'L'{
		return encode.PrivateKeyWifCompressed, nil
	}
	if key[ 0 ] == '1'{
		return encode.BitcoinAddress, nil
	}
	return "",errors.New("PrefixFormat not found")
}

func discoverEncodeFormat( key string ) (consts.EncodeFormat, error){
	if key[ 0 ] == '5'{
		return consts.PrivateKeyWif, nil
	}
	if  key[ 0 ] == 'K' || key[ 0 ] == 'L'{
		return consts.PrivateKeyWifCompressed, nil
	}
	if key[ 0 ] == '1'{
		return consts.Address, nil
	}
	return "",errors.New("EncodedFormat not found" )
}

func findAlphabetIndex( char uint8 ) (index int, err error) {

	var left, right, middle = 0, len(consts.Base58Alphabet)-1, 0
	for left <= right{
		middle = ( left + right )/2
		if consts.Base58Alphabet[ middle ] == char {
			return middle, nil
		}else{
			if char > consts.Base58Alphabet[middle] {
				left = middle + 1
			}else{
				right = middle - 1
			}
		}
	}
	return -1, errors.New("character not found in alphabet" )
}

// takes the entire base58check string and change base10
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

func ValidateChecksum( key string ) error{

	keyInt, err := decodeBase58CheckToInt( key )
	if err != nil {
		return errors.New( fmt.Sprintf( "Invalid key %s", err.Error() ) )
	}

	bytes := keyInt.Bytes()

	// prefix + payload without the checksum
	payload := bytes[ 0: len(bytes) - 4 ]

	// sha256( sha256( payload ) ) to get the right checksum
	firstSha := sha256.Sum256( payload )
	secondSha := sha256.Sum256( firstSha[:] )

	// compare right checksum and actualChecksum are equals
	hexChecksum := hex.EncodeToString( bytes[ len( bytes ) - 4: ] )
	hexCalculatedChecksum := hex.EncodeToString( secondSha[0:4] ) 
	if hexCalculatedChecksum != hexChecksum{
		return InvalidChecksum{ contains: hexChecksum, mustContains: hexCalculatedChecksum }
	}

	return nil
}

func ValidateEncodedKeySize( encodePrefix consts.EncodeFormat, key string ) error{
	switch encodePrefix {
	case consts.PrivateKeyWif:
		if len( key ) != 51 {
			if len( key ) > 51{
				return errors.New("too long key")
			}
			return errors.New("too short key")
		}
	case consts.PrivateKeyWifCompressed:
		if len( key ) != 52{
			return errors.New("invalid private key wif compressed, size didn't match" )
		}
	case consts.Address:
		panic("Not implemented" )
		return nil
	}

	return nil
}

// Entrypoint, decode base58Check
func Base58Check( key string ) (*big.Int, error){

	// validate all chars in the key belongs to base58Check alphabet
	for _, char := range key {
		index, err := findAlphabetIndex( uint8(char) )
		if err != nil || index == -1 {
			return nil, errors.New("invalid key, contains invalid characters")
		}
	}

	//determine encode format
	format, err := discoverEncodeFormat( key )
	prefix, err := discoverEncodePrefix( key )

	//validate key size
	err = ValidateEncodedKeySize( format, key )
	if err != nil {
		return nil, err
	}

	//validate checksum
	err = ValidateChecksum( key )
	if err != nil {
		return nil, err
	}

	// decode to big.Int
	keyInt, err := decodeBase58CheckToInt( key )
	if err != nil{
		return nil, err
	}

	// Some encodes contains suffixes
	extraSuffixBytes := 0
	if format == consts.PrivateKeyWifCompressed{
		extraSuffixBytes = 1
	}

	//remove firsts version byte
	bytesKey := keyInt.Bytes()[ prefix.ByteSize() :]

	//remove 4 checksum bytes
	bytesKey = bytesKey[0: ( len( bytesKey ) - 4 ) - extraSuffixBytes ]

	// convert to big.Int
	keyInt.SetBytes( bytesKey )

	return keyInt, nil
}


func Base58CheckString( key string ) (string, error){
	bigInt, err := Base58Check(key)
	if err != nil || bigInt == nil {
		return "", err
	}
	return bigInt.Text( 16 ), nil
}