package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vinoMamba/go-mangosteen/internal/router"
)

func init() {
	var srvCmd = &cobra.Command{
		Use:   "server",
		Short: "Run Server",
		Run: func(cmd *cobra.Command, args []string) {
			r := router.New()
			err := r.Run(":3000")
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	rootCmd.AddCommand(srvCmd)
}
