package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xlsx2mysql",
	Short: "An tool of helping your fastly generate SQL from Excel",
	Long: `An tool of helping your fastly generate SQL from Excel.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}
