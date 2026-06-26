package sethypr

import (
	"io"
	"os"
	"path/filepath"

	"github.com/g5ostXa/archypr/internal/core"
)

var hyprDots = "dotfiles"

func SourceCopy() {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		core.Logger.Fatal("Failed to get user's home directory...", err)
	}

	// Run this if successfully deternined user's home dir
	dotfilesSourceDir := filepath.Join(".", hyprDots)
	dotfilesDestDir := filepath.Join(homeDir, hyprDots)

	core.Logger.Info("Copying dotfiles...")

	if err := copyDir(dotfilesSourceDir, dotfilesDestDir); err != nil {
		core.Logger.Fatal("Failed to copy dotfiles...", err)
	}
	core.Logger.Info("Dotfiles copied successfully")
}

func copyDir(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}
		return copyFile(path, targetPath, info.Mode())
	})
}

func copyFile(src, dest string, mode os.FileMode) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
