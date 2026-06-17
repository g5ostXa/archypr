package core

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
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
		fmt.Println("\033[2J\033[H")
		return
	}
	fmt.Println(string(output))
}
