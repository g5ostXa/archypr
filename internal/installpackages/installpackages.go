package installpackages

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/g5ostXa/archypr/internal/core"
)

func Needed(packages []string) {

	if len(packages) == 0 {
		core.Logger.Info("No missing dependencies to install.")
		return
	}

	args := append([]string{"-S", "--needed"}, packages...)
	cmd := exec.Command("paru", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		core.Logger.Error(fmt.Sprintf("Failed to install dependencies: %v", err))
		os.Exit(1)
	}

	core.TimeLogger.Info("Dependencies installed successfully!")
}
