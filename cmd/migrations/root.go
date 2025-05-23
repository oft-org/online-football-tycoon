package migrations

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/cmd/migrations"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oft",
	Short: "Online Football Tycoon CLI",
}

func Execute() {
	rootCmd.AddCommand(migrations.MigrationsCmd)
	rootCmd.AddCommand(ServerCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
