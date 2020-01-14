package impl

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"io/ioutil"
	"log"

	ssh "github.com/ianmcmahon/encoding_ssh"
)

func loadPublicKey(pubkey string) (*rsa.PublicKey, error) {

	pubbytes, err := ioutil.ReadFile(pubkey)
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	// decode string ssh-rsa format to native type
	pub_key, err := ssh.DecodePublicKey(string(pubbytes))
	if err != nil {
		log.Printf("%v\n", err)
	}

	rsapubkey := pub_key.(*rsa.PublicKey)
	return rsapubkey, nil
}

func authenticateFile(file string) {

	fp, err := fullName(file)
	if err != nil {
		log.Printf("File %s does not exist\n", file)
		return
	}
	sfp, err := fullName(file + signatureFileType)
	if err != nil {
		log.Printf("Signature file missing. Cannot authenticate %s\n", fp)
		return
	}
	log.Printf("Authenticate file %s using %s\n", fp, sfp)
	filehash, err := hashFile(fp)
	if err != nil {
		return
	}
	log.Printf("hash is %s\n", filehash)

}

func Authenticate(pub string, files []string) {
	log.Printf("Authenticating using %s of %d files\n", pub, len(files))
	for _, f := range files {
		authenticateFile(f)
	}
}

func verify(file string, sigfile string, pubkey *rsa.PublicKey) error {
	databytes, _ := ioutil.ReadFile(file)
	h := sha256.New()
	h.Write(databytes)
	hashed := h.Sum(nil)

	sigbytes, _ := ioutil.ReadFile(sigfile)
	err := rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed, sigbytes)
	if err != nil {
		log.Printf("Verifying %s using signature: %s - %s\n", file, sigfile, err)
		return err
	}
	log.Printf("Verified the signature %s of file %s\n", sigfile, file)
	return nil
}

func Verify(file string, sigfile string, pubkeyfile string) error {

	rsapubkey, err := loadPublicKey(pubkeyfile)
	if err != nil {
		return err
	}
	log.Printf("Loaded public key %s\n", pubkeyfile)
	err = verify(file, sigfile, rsapubkey)
	return err
}
