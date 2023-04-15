package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vinoMamba/go-mangosteen/internal/database"
	"github.com/vinoMamba/go-mangosteen/internal/router"
)

func Run() {
	rootCmd := &cobra.Command{
		Use:   "mangosteen",
		Short: "Mangosteen App Cli",
	}
	srvCmd := &cobra.Command{
		Use:   "server",
		Short: "Run Server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "handle the database",
	}
	lib1Cmd := &cobra.Command{
		Use:   "sql",
		Short: "use database/sql",
		Run: func(cmd *cobra.Command, args []string) {
			database.PgConnect()
			database.PgCreateTables()
			defer database.Close()
		},
	}
	dbCmd.AddCommand(lib1Cmd)

	rootCmd.AddCommand(srvCmd)
	rootCmd.AddCommand(dbCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunServer() {
	r := router.New()
	err := r.Run(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
