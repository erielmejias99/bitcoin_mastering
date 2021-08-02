package decode

import (
	"fmt"
	"github.com/bitcoin_mastering/internal/bitcoin/consts"
)

type LetterNotFound struct {
	char uint8
}

func (l LetterNotFound) Error()string{
	return fmt.Sprintf( "invalid key, character '%s' doesn't belong to Base58 alphabet: %s", string( l.char ), consts.Base58Alphabet)
}


type InvalidChecksum struct {
	contains string
	mustContains string
}

func (i InvalidChecksum) Error() string {
	return fmt.Sprintf( "Invalid checksum, contains: %s, Must be: %s",
		i.contains, i.mustContains )
}