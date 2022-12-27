package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/spinute/redis-inventory/cmd/app"
)

func main() {
	err := doc.GenMarkdownTree(app.RootCmd, "./docs/cobra")
	if err != nil {
		log.Fatal(err)
	}
}
