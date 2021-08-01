package bitcoin

import "fmt"

type InvalidPrivateKey struct {
	hint string
}

func (i *InvalidPrivateKey) Error()string {
	return fmt.Sprintf( "Invalid error key hint: %s", i.hint )
}