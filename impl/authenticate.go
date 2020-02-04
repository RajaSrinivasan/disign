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

func authenticate(file string, sigfile string, pubkey *rsa.PublicKey) error {
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

func Authenticate(file string, sigfile string, pubkeyfile string) error {

	rsapubkey, err := loadPublicKey(pubkeyfile)
	if err != nil {
		return err
	}
	log.Printf("Loaded public key %s\n", pubkeyfile)
	err = authenticate(file, sigfile, rsapubkey)
	return err
}

func AuthenticateFiles(files []string, pub string) error {
	log.Printf("Authenticating using %s of %d files\n", pub, len(files))
	pubkeyfile, err := loadPublicKey(pub)
	if err != nil {
		log.Printf("Cannot authenticate files. Unable to load Public Key file\n")
		return err
	}
	for _, f := range files {
		sigfile := f + signatureFileType
		err = authenticate(f, sigfile, pubkeyfile)
		if err != nil {
			return err
		}
	}
	return nil
}
