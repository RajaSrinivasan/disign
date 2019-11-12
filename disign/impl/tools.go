package impl

import (
	"log"
	"os"
	"path/filepath"
)

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
