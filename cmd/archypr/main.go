package main

import (
	"github.com/g5ostXa/archypr/internal/aurhelper"
	"github.com/g5ostXa/archypr/internal/header"
	"github.com/g5ostXa/archypr/internal/pacmanconf"
)

func main() {

	header.InstallStart()
	pacmanconf.Configure()

	aurhelper.Check()
}
