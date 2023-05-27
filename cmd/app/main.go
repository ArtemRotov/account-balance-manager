package main

import (
	"flag"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to config file")
}

func main() {

	flag.Parse()

}
