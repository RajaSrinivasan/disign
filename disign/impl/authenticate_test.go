package impl

import (
	"testing"
)

var pubKeyFile = "/Users/rajasrinivasan/.ssh/id_rsa.pub"

func TestVerifyExternalKey(t *testing.T) {
	Verify("sign.go", "sign.go.sig", pubKeyFile)
	Verify("sign.go", "sign.go.2.sig", pubKeyFile)
	Verify("sign_test.go", "sign_test.go.sig", pubKeyFile)
	Verify("sign_test.go", "sign_test.go.2.sig", pubKeyFile)
}
