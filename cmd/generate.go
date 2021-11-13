package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	app "github.com/towelong/xlsx2mysql/app"
)

var wg sync.WaitGroup

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate SQL from Excel",
	Long:  "generate SQL from Excel",
	Run: func(cmd *cobra.Command, args []string) {
		generate(args)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generate(files []string) {
	for _, file := range files {
		wg.Add(1)
		go worker(file)
	}
	wg.Wait()
}

func worker(file string) {
	defer wg.Done()
	_, err := app.Xlsx2MySQL(file, "Sheet1")
	if err != nil {
		fmt.Println(file + "=============>" + "Convert Error...")
	} else {
		fmt.Println(file + "=============>" + "Convert Successfully!")
	}
}
