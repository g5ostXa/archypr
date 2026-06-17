package pacmanconf

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/g5ostXa/archypr/internal/core"
)

const PacmanConfPath = "/etc/pacman.conf"

func Configure() {
	if err := ConfigureFile(PacmanConfPath); err != nil {
		core.Logger.Warn("Failed to configure pacman", "error", err)
	}
}

func ConfigureFile(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat %s: %w", path, err)
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	updated, changes := updatePacmanConf(contents)
	if len(changes) == 0 {
		core.Logger.Info("Pacman is already configured !")
		return nil
	}

	if err := writeConfigFile(path, updated, fileInfo.Mode().Perm()); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	for _, change := range changes {
		core.Logger.Info(change)
	}
	core.Logger.Info("Pacman has been configured !")
	return nil
}

func writeConfigFile(path string, contents []byte, perm os.FileMode) error {
	if os.Geteuid() == 0 {
		return os.WriteFile(path, contents, perm)
	}

	if _, err := exec.LookPath("sudo"); err != nil {
		return fmt.Errorf("sudo is required to write %s when archypr is not run as root: %w", path, err)
	}

	cmd := exec.Command("sudo", "tee", path)
	cmd.Stdin = bytes.NewReader(contents)
	cmd.Stdout = io.Discard
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sudo tee %s: %w", path, err)
	}

	return nil
}

func updatePacmanConf(contents []byte) ([]byte, []string) {
	text := strings.ReplaceAll(string(contents), "\r\n", "\n")
	hasTrailingNewline := strings.HasSuffix(text, "\n")
	lines := strings.Split(text, "\n")
	if hasTrailingNewline {
		lines = lines[:len(lines)-1]
	}

	changes := make([]string, 0, 4)
	hasILoveCandy := false
	parallelDownloadsLine := -1

	for index, line := range lines {
		trimmed := strings.TrimSpace(line)

		switch {
		case trimmed == "ILoveCandy":
			hasILoveCandy = true
		case trimmed == "#Color":
			lines[index] = uncommentLine(line)
			changes = append(changes, "Enabled pacman colors")
		case trimmed == "#VerbosePkgLists":
			lines[index] = uncommentLine(line)
			changes = append(changes, "Enabled pacman verbose package lists")
		case strings.HasPrefix(trimmed, "#ParallelDownloads"):
			lines[index] = uncommentLine(line)
			line = lines[index]
			trimmed = strings.TrimSpace(line)
			changes = append(changes, "Enabled pacman parallel downloads")
		}

		if strings.HasPrefix(trimmed, "ParallelDownloads") {
			parallelDownloadsLine = index
		}
	}

	if !hasILoveCandy && parallelDownloadsLine >= 0 {
		lines = insertLine(lines, parallelDownloadsLine+1, "ILoveCandy")
		changes = append(changes, "Enabled pacman ILoveCandy")
	}

	result := strings.Join(lines, "\n")
	if hasTrailingNewline {
		result += "\n"
	}

	return []byte(result), changes
}

func insertLine(lines []string, index int, line string) []string {
	lines = append(lines, "")
	copy(lines[index+1:], lines[index:])
	lines[index] = line
	return lines
}

func uncommentLine(line string) string {
	commentIndex := strings.Index(line, "#")
	if commentIndex == -1 {
		return line
	}

	return line[:commentIndex] + line[commentIndex+1:]
}
