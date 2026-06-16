package header

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/g5ostXa/archypr/internal/core"
	"github.com/g5ostXa/archypr/internal/styles"
)

var (
	MasterTitle = "archypr"
	Version     = "0.0-1"
)

var installPromptMsg = "→ Do you want to start the installation now? (Yy/Nn):"

func InstallStart() {

	core.ClearScreen()

	core.Logger.Info("Initializing installer...")
	time.Sleep(1 * time.Second)

	core.ClearScreen()

	lipgloss.Println(
		styles.MasterTileStyle.Render(
			fmt.Sprintf("Welcome to %s !", MasterTitle),
		),
	)

	lipgloss.Println(
		styles.MasterVersionStyle.Render(
			fmt.Sprintf("Version: %s", Version),
		),
	)

	core.Separator()
	reader := bufio.NewReader(os.Stdin)

	for {
		lipgloss.Println(styles.CommonPromptStyle.Render(installPromptMsg))

		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "n" || input == "" {
			core.Logger.Info("Installation cancelled by user.")
			os.Exit(0)
		} else if input == "y" {
			break
		}

		lipgloss.Println(styles.CommonWarningStyle.Render("-> Invalid input. Please type 'y' or 'n'."))
	}
}
