package cmd

import (
	"log"

	"github.com/vinoMamba/go-mangosteen/internal/router"
)

func RunServer() {

	r := router.New()
	err := r.Run(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
