package impl

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
)

var HashAlg string
var Verbose bool

func reportError(err error, ctx string) {
	log.Printf("Error (%v) for %s\n", err, ctx)
}

func fullName(f string) (string, error) {
	_, err := os.Stat(f)
	if err != nil {
		reportError(err, "stat of "+f)
		return "", err
	}
	ap, err := filepath.Abs(f)
	if err != nil {
		reportError(err, "filepath.Abs of "+f)
		return "", err
	}
	return ap, err
}

func hashFile(f string) (string, error) {
	file, err := os.Open(f)
	if err != nil {
		reportError(err, f)
		return "", err
	}
	defer file.Close()

	var hashBytes []byte
	switch HashAlg {
	case "sha512":
		hsha512 := sha512.New()
		_, err = io.Copy(hsha512, file)
		if err != nil {
			reportError(err, "hash sha512")
			return "", err
		}
		hashBytes = hsha512.Sum(nil)
	case "md5":
		hmd5 := md5.New()
		_, err = io.Copy(hmd5, file)
		if err != nil {
			reportError(err, "hash md5")
			return "", err
		}
		hashBytes = hmd5.Sum(nil)
	}
	hashhex := hex.EncodeToString(hashBytes)
	return hashhex, nil
}
