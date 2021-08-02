package decode

import (
	"testing"
)

func TestValidateChecksumOk(t *testing.T) {
	privateKey := "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"
	err := ValidateChecksum( privateKey )
	if err != nil{
		t.Error( err.Error() )
	}
}

func TestValidateChecksumWrong(t *testing.T) {
	privateKey := "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB15cn"
	err := ValidateChecksum( privateKey )
	if err == nil{
		t.Error( "Must be an error because de checksum is not valid." )
	}
}