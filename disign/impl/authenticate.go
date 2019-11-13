package impl

import (
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
