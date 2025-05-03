package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var matchID string

var matchesCmd = &cobra.Command{
	Use:   "matches",
	Short: "Play a match by ID (POST to /match/play)",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()

		url := "http://localhost:8080/match/play"

		payload := map[string]interface{}{
			"id": matchID,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(payload).
			Post(url)

		if err != nil {
			fmt.Println("Error al hacer la solicitud:", err)
			return
		}

		fmt.Println("Código de estado:", resp.StatusCode())
		fmt.Println("Respuesta:", resp.String())
	},
}

func init() {
	matchesCmd.Flags().StringVar(&matchID, "id", "", "Match ID to play (required)")
	matchesCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(matchesCmd)
}
