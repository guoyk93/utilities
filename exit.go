package utilities

import (
	"log"
	"os"
)

func Exit(err *error) {
	if *err == nil {
		return
	}
	log.Println("exit with error:", (*err).Error())
	os.Exit(1)
}
