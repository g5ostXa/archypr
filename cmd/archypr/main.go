package main

import (
	"github.com/g5ostXa/archypr/internal/aurhelper"
	"github.com/g5ostXa/archypr/internal/checkdepends"
	"github.com/g5ostXa/archypr/internal/header"
	"github.com/g5ostXa/archypr/internal/pacmanconf"
)

func main() {

	// Show header and prompt to srart the install
	header.InstallStart()

	// Configure /etc/pacman.conf
	pacmanconf.Configure()

	// Install paru and all archypr dependendices
	aurhelper.Check()
	checkdepends.Validate()
}
