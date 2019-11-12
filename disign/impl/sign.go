package impl

import (
	"log"
)

func init() {
	log.Println("Sign implementation init")
}

func signfile(file string) {
	log.Printf("Signing file %s\n", file)
}
func Sign(pvt string, files []string) {
	log.Printf("Signing using %s of %d files\n", pvt, len(files))
}
