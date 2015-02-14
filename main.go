package main

import (
	"fmt"
	"os"

	"github.com/oremj/cdn-purge/cdns/edgecast"
	"github.com/spf13/cobra"
)

func purgeEdgecastCmd() *cobra.Command {
	edgecastAPI := new(edgecast.API)

	purgeEdgecastCmd := &cobra.Command{
		Use:   "edgecast [url]",
		Short: "purges urls from edgecast",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Usage()
				fmt.Println("url must be provided")
				os.Exit(1)
			}

			id, err := edgecastAPI.Purge(args[0])
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			fmt.Print(id)
		},
	}
	purgeEdgecastCmd.Flags().StringVarP(&edgecastAPI.Token, "token", "t", "", "Access token.")
	purgeEdgecastCmd.Flags().StringVarP(&edgecastAPI.AccountId, "account-id", "", "", "Access token.")

	return purgeEdgecastCmd
}

func main() {

	mainCmd := &cobra.Command{
		Use: "cdn-purge",
	}

	mainCmd.AddCommand(purgeEdgecastCmd())
	mainCmd.Execute()
}
