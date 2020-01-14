package impl

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"io/ioutil"
	"log"
)

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
		log.Printf("%s\n", err)
		return err
	}
	return nil
}

func Verify(file string, sigfile string, pubkeyfile string) error {
	log.Printf("Verifying %s signature %s using %s\n", file, sigfile, pubkeyfile)
	pubbytes, err := ioutil.ReadFile(pubkeyfile)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	log.Printf("Loaded %s %d bytes\n", pubkeyfile, len(pubbytes))
	pubkey, err := x509.ParsePKIXPublicKey(pubbytes)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	rsapubkey := pubkey.(*rsa.PublicKey)
	err = verify(file, sigfile, rsapubkey)

	return nil
}
