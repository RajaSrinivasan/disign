package impl

import (
	"log"
	"testing"
)

var pubKeyFile = "/Users/rajasrinivasan/.ssh/id_rsa.pub"

func TestVerifyExternalKey(t *testing.T) {
	Verify("sign.go", "sign.go.sig", pubKeyFile)
	Verify("sign.go.sig", "sign.go.2.sig", pubKeyFile)
	Verify("sign_test.go", "sign_test.go.sig", pubKeyFile)
	Verify("sign_test.go", "sign_test.go.2.sig", pubKeyFile)
}

func TestLoadPrublicKey(t *testing.T) {
	pubkey, _ := loadPublicKey(pubKeyFile)
	log.Printf("%v\n", pubkey)
}
