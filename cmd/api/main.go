package main

import (
	"log"

	"github.com/manuhdez/golang-api-hex/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
