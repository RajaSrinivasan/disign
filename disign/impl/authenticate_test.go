package impl

import (
	"log"
	"testing"
)

var pubKeyFile = "/Users/rajasrinivasan/.ssh/id_rsa.pub"

func TestAuthenticate(t *testing.T) {
	//t.Println("Testing authentication of digital signatures")
	Authenticate("sign.go", "sign.go.sig", pubKeyFile)
	Authenticate("sign.go.sig", "sign.go.2.sig", pubKeyFile)
	Authenticate("sign_test.go", "sign_test.go.sig", pubKeyFile)
	Authenticate("sign_test.go", "sign_test.go.2.sig", pubKeyFile)
}

func TestAuthenticateFiles(t *testing.T) {
	files := []string{"sign.go", "sign_test.go"}
	AuthenticateFiles(files, pubKeyFile)
	files = []string{"authenticate.go", "authenticate_test.go", "sign.go", "sign_test.go"}
	AuthenticateFiles(files, pubKeyFile)
}

func TestLoadPrublicKey(t *testing.T) {
	pubkey, _ := loadPublicKey(pubKeyFile)
	log.Printf("%v\n", pubkey)
}
