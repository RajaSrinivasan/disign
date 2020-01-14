package impl

import (
	"testing"
)

var pvtKeyFile = "/Users/rajasrinivasan/.ssh/id_rsa"

func TestSign(t *testing.T) {

	Sign("sign.go", "sign.go.sig", pvtKeyFile)
	Sign("sign_test.go", "sign_test.go.sig", pvtKeyFile)

	Sign("sign.go", "sign.go.2.sig", pvtKeyFile)
	Sign("sign_test.go", "sign_test.go.2.sig", pvtKeyFile)
}

func TestSignFiles(t *testing.T) {
	files := []string{"sign.go", "sign_test.go"}
	SignFiles(files, pvtKeyFile)
	files = []string{"authenticate.go", "authenticate_test.go", "sign.go", "sign_test.go"}
	SignFiles(files, pvtKeyFile)
}

func TestLoadPrivateKey(t *testing.T) {
	pvtkey, _ := loadPrivateKey(pvtKeyFile)
	showPrivateKey(pvtkey)
}
