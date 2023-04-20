package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "handle the database",
	}
	lib1Cmd := &cobra.Command{
		Use:   "sql",
		Short: "use database/sql",
		Run: func(cmd *cobra.Command, args []string) {
			// database.PgConnect()
			// database.PgCreateTables()
			// defer database.Close()
		},
	}
	dbCmd.AddCommand(lib1Cmd)
	rootCmd.AddCommand(dbCmd)
}
