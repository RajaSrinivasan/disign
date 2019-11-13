package impl

import (
	"log"
)

var pvtKeyFilename string

const signatureFileType string = ".sig"

func signFile(file string) {

	fp, err := fullName(file)
	if err != nil {
		log.Printf("Cannot Sign %s\n", file)
		return
	}
	sigfile := fp + signatureFileType
	log.Printf("Signing file %s generating %s\n", fp, sigfile)
	filehash, err := hashFile(fp)
	if err != nil {
		return
	}
	log.Printf("hash is %s\n", filehash)

}

func Sign(pvt string, files []string) {
	log.Printf("Signing using %s of %d files\n", pvt, len(files))
	pvtKeyFilename, _ = fullName(pvt)
	for _, f := range files {
		signFile(f)
	}
}
