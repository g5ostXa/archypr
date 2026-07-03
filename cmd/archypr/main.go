package main

import (
	"github.com/g5ostXa/archypr/internal/aurhelper"
	"github.com/g5ostXa/archypr/internal/checkdepends"
	"github.com/g5ostXa/archypr/internal/header"
	"github.com/g5ostXa/archypr/internal/pacmanconf"
	"github.com/g5ostXa/archypr/internal/sethypr"
)

func main() {

	// Show header and prompt to start the install
	header.InstallStart()

	// Configure /etc/pacman.conf
	pacmanconf.Configure()

	// Install paru and all archypr dependencies
	aurhelper.Check()
	checkdepends.Validate()

	// Apply config files
	sethypr.SourceCopy()
	sethypr.AssetsCopy()
	sethypr.BashrcCopy()
	sethypr.VersionCopy()
}
