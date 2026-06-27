package header

import (
	"fmt"
	"os"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/g5ostXa/archypr/internal/core"
	"github.com/g5ostXa/archypr/internal/styles"
)

var (
	MasterTitle = "archypr"
	Version     = "0.1"
)

var installPromptMsg = "Do you want to start the installation now? (Yy/Nn)"

func InstallStart() {

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

	fmt.Println()

	core.CommonSeparator()
	time.Sleep(1 * time.Second)

	fmt.Println()

	core.TimeLogger.Info("Initialized installer...")
	time.Sleep(1 * time.Second)

	if !core.Confirm(os.Stdin, installPromptMsg) {
		core.Logger.Info("Installation cancelled by user.")
		os.Exit(0)
	}
}
