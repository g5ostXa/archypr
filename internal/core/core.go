package core

import (
	"fmt"
	"os"
	"os/exec"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/log"
	"github.com/g5ostXa/archypr/internal/styles"
)

var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          " 󰣇 Log",
})

func ClearScreen() {

	cmd := exec.Command("clear")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("\033[2J\033[H")
		return
	}
	fmt.Println(string(output))
}

func Separator() {

	var separator = "....................................................."

	fmt.Println()
	lipgloss.Println(styles.CommonSeparatorStyle.Render(separator))
}
