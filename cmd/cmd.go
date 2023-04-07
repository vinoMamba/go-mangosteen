package cmd

import (
	"log"

	"github.com/vinoMamba/go-mangosteen/internal/database"
	"github.com/vinoMamba/go-mangosteen/internal/router"
)

func RunServer() {
	database.PgConnect()
	database.PgCreateTables()
	defer database.Close()

	r := router.New()
	err := r.Run(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
