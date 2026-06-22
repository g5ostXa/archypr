package core

import (
	"fmt"
	"os"
	"os/exec"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/log"
	"github.com/g5ostXa/archypr/internal/styles"
)

// Common logger
var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: false,
	Prefix:          ":",
})

// Time-stamped logger:
var TimeLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})

func ClearScreen() {

	cmd := exec.Command("clear")

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("\033[2J\033[H")
	}
	fmt.Print(string(output))
}

func CommonSeparator() {

	var commonSeparator = "⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯ ⋯"
	lipgloss.Println(styles.ComSeparatorStyle.Render(commonSeparator))
}
