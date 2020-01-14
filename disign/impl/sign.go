package impl

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

const signatureFileType string = ".sig"
const rsaKeySize = 2048

func showPrivateKey(pvt *rsa.PrivateKey) {
	log.Printf("%s\n", pvt.D.Text(16))
	pvtkeybytes, _ := x509.MarshalPKCS8PrivateKey(pvt)

	log.Printf("PvtKeyBytes: %x\n", pvtkeybytes)

	pub := pvt.Public()
	rsapub := pub.(*rsa.PublicKey)

	log.Printf("Public Key Size %d, Exponent %d\n", rsapub.Size(), rsapub.E)
	log.Printf("PublicKey: %s\n", rsapub.N.Text(16))
}

func loadPrivateKey(keyfile string) (*rsa.PrivateKey, error) {
	keybytes, _ := ioutil.ReadFile(keyfile)
	key, err := ssh.ParseRawPrivateKey(keybytes)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	if Verbose {
		log.Printf("%q\n", key)
	}
	return key.(*rsa.PrivateKey), nil
}

func sign(file string, sigfile string, pvt *rsa.PrivateKey) error {

	log.Printf("Signing %s creating %s\n", file, sigfile)

	databytes, _ := ioutil.ReadFile(file)
	h := sha256.New()
	h.Write(databytes)
	datahash := h.Sum(nil)
	if Verbose {
		log.Printf("Datahash: %x\n", datahash)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, pvt, crypto.SHA256, datahash[:])
	if err != nil {
		log.Printf("Error from signing: %s\n", err)
		return err
	}

	sigf, _ := os.Create(sigfile)
	defer sigf.Close()
	sigf.Write(signature)
	if Verbose {
		fmt.Printf("Signature: %x\n", signature)
	}
	return nil
}

func Sign(file string, sigfile string, pvtkeyfile string) error {
	rsapvtkey, err := loadPrivateKey(pvtkeyfile)
	if err != nil {
		return err
	}
	err = sign(file, sigfile, rsapvtkey)
	return err
}

func SignFiles(files []string, pvtkeyfile string) error {
	log.Printf("Signing using %s of %d files\n", pvtkeyfile, len(files))

	rsapvtkey, err := loadPrivateKey(pvtkeyfile)
	if err != nil {
		return err
	}
	for _, f := range files {
		sigfname := f + signatureFileType
		err = sign(f, sigfname, rsapvtkey)
		if err != nil {
			return err
		}
	}
	return nil
}
