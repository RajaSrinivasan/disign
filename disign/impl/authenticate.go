package impl

import (
	"log"
)

func init() {
	log.Println("Authenticae implementation init")
}

func authenticatefile(file string) {
	log.Printf("Authenticate file %s\n", file)
}

func Authenticate(pub string, files []string) {
	log.Printf("Authenticating using %s of %d files\n", pub, len(files))
	for _, f := range files {
		authenticatefile(f)
	}
}
