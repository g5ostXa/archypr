package main

import (
	"github.com/g5ostXa/archypr/internal/aurhelper"
	"github.com/g5ostXa/archypr/internal/header"
)

func main() {

	header.InstallStart()
	aurhelper.Check()
}
