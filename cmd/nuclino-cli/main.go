package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type config struct {
	env            string
	nuclinoApiKey  string
	listWorkspaces bool
	nuclinoApi struct{
		baseUrl  = 'https://api.nuclino.com'
	}
}

func main() {
	var cfg config

	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.nuclinoApiKey, "apiKey", "", "API key for the Nuclino API")

	flag.BoolVar(&cfg.listWorkspaces, "workspaces", false, "List workspaces")

	displayVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	// If the version flag value is true, then print out the version number and
	// immediately exit.
	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Built  :\t%s\n", buildTime)
		os.Exit(0)
	}

	log.Println("List workspaces?", cfg.listWorkspaces)

}
