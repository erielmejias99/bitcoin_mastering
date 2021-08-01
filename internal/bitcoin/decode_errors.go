package bitcoin

import "fmt"

type LetterNotFound struct {
	char uint8
}

func (l LetterNotFound) Error()string{
	return fmt.Sprintf( "invalid key, character '%s' doesn't belong to Base58 alphabet: %s", string( l.char ), alphabet )
}
