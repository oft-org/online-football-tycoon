package main

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/cmd/migrations"
	serverCmd "github.com/robertobouses/online-football-tycoon/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oft",
	Short: "Online Football Tycoon CLI",
}

func main() {
	rootCmd.AddCommand(migrations.MigrationsCmd)
	rootCmd.AddCommand(serverCmd.ServerCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
