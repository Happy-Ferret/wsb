package cmd

import (
	"log"

	"github.com/qneyrat/wsb/wsbd/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start wsb server",
	Run: func(cmd *cobra.Command, args []string) {
		wbd := server.NewServer()

		log.Println("http server started on :4000")
		err := wbd.Start()
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
