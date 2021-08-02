package encode

import "fmt"

type InvalidPrivateKey struct {
	Hint string
}

func (i *InvalidPrivateKey) Error()string {
	return fmt.Sprintf( "Invalid error key hint: %s", i.Hint )
}