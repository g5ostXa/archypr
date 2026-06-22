package core

import (
	"bufio"
	"io"
	"strings"

	"github.com/g5ostXa/archypr/internal/styles"
)

func Confirm(input io.Reader, prompt string) bool {
	reader := bufio.NewReader(input)

	for {
		Logger.Print(styles.CommonPromptStyle.Render(prompt))

		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))

		switch answer {
		case "y":
			return true
		case "n", "":
			return false
		default:
			Logger.Warn("Invalid input. Please type y or n ...")
		}
	}
}
