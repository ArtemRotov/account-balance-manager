package main

import (
	"flag"

	"github.com/ArtemRotov/account-balance-manager/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to config file")
}

func main() {
	flag.Parse()

	app.Run(configPath)
}
