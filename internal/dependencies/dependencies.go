package dependencies

import (
	"os/exec"

	"github.com/g5ostXa/archypr/internal/core"
)

func Check() {
	packages := []string{
		"lolcat",
		"figlet",
	}

	for _, pkg := range packages {
		cmd := exec.Command("paru", "-Qi", pkg)

		if err := cmd.Run(); err != nil {
			// Run something here if not all dependencies are installed...
			core.TimeLogger.Fatal("Some dependencies are missing...")
		}
	}
	core.Logger.Info("All dependencies are already installed !")
}
