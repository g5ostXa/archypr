package aurhelper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/g5ostXa/archypr/internal/core"
	"github.com/g5ostXa/archypr/internal/styles"
)

var helperInstallPromptMsg = "→ Do you want to install paru now? (Yy/Nn):"

func installAurHelper() {

	fmt.Println()
	core.Logger.Info("Paru was not found on your system.")

	reader := bufio.NewReader(os.Stdin)

	for {
		lipgloss.Print(styles.CommonPromptStyle.Render(helperInstallPromptMsg))
		fmt.Println()

		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "n" || input == "" {
			fmt.Println()
			core.Logger.Info("Installation cancelled by user.")
			os.Exit(0)
		} else if input == "y" {
			break
		}

		core.Logger.Warn("Invalid input. Please type y or n ...")
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		core.TimeLogger.Fatal("Could not determine home directory...")
	}

	cacheDir := filepath.Join(homeDir, ".cache")
	paruDir := filepath.Join(cacheDir, "paru")

	core.Logger.Info("Building paru...")

	gitCmd := exec.Command("git", "clone", "https://aur.archlinux.org/paru.git", paruDir)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr

	if err := gitCmd.Run(); err != nil {
		core.TimeLogger.Fatal("Failed to clone paru")
	}

	buildCmd := exec.Command("makepkg", "-si", "--noconfirm")
	buildCmd.Dir = paruDir
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	if err := buildCmd.Run(); err != nil {
		fmt.Println()
		core.TimeLogger.Fatal("Failed to build and install paru...")
	}

	fmt.Println()
	core.Logger.Info("Paru has been successfully built and installed!")
}

func Check() {

	if _, err := exec.LookPath("paru"); err == nil {
		// If aur helper is already installed, run this...
		core.Logger.Info("Paru is already installed !")
		return
	}
	// If aur helper is NOT installed, run this...
	installAurHelper()
}
