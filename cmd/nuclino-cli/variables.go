package main

import "github.com/jakoubek/nuclino-cli/internal/vcs"

var (
	version   = vcs.Version()
	buildTime string
)
