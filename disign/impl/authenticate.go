package impl

import (
	"log"
)

func authenticateFile(file string) {
	log.Printf("Authenticate file %s\n", file)
}

func Authenticate(pub string, files []string) {
	log.Printf("Authenticating using %s of %d files\n", pub, len(files))
	for _, f := range files {
		authenticateFile(f)
	}
}
