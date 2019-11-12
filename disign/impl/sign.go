package impl

import (
	"log"
)

var pvtKeyFilename string

func signFile(file string) {
	log.Printf("Signing file %s\n", file)
}

func Sign(pvt string, files []string) {
	log.Printf("Signing using %s of %d files\n", pvt, len(files))
	pvtKeyFilename, _ = fullName(pvt)
	for _, f := range files {
		signFile(f)
	}
}
