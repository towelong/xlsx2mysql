package main

import (
	"sync"

	"github.com/towelong/xlsx2mysql/app"
)

var (
	wg sync.WaitGroup
)

func main() {
	files := []string{"example.xlsx"}
	for _, file := range files {
		wg.Add(1)
		go worker(file)
	}
	wg.Wait()
}

func worker(file string) {
	app.Xlsx2MySQL(file, "Sheet1")
	wg.Done()
}
