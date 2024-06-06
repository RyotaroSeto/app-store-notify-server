package main

import (
	"os"

	"app-store-notify-server/pkg/cmd/api"
)

func main() {
	os.Exit(run())
}

func run() int {
	api.Run()
	return 0
}
