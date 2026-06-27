package sethypr

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/g5ostXa/archypr/dotfiles"
	"github.com/g5ostXa/archypr/internal/core"
)

const hyprDots = "dotfiles"

func SourceCopy() {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		core.Logger.Fatal("Failed to get user's home directory...", err)
	}

	// Run this if successfully determined user's home dir
	dotfilesDestDir := filepath.Join(homeDir, hyprDots)

	core.Logger.Info("Copying dotfiles...")

	if err := copyEmbeddedDir(dotfiles.FS, ".", dotfilesDestDir); err != nil {
		core.Logger.Fatal("Failed to copy dotfiles...", err)
	}
	core.Logger.Info("Dotfiles copied successfully")
}

func copyEmbeddedDir(srcFS fs.FS, srcRoot, dest string) error {
	return fs.WalkDir(srcFS, srcRoot, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcRoot, path)
		if err != nil {
			return err
		}
		if relPath == "." {
			return nil
		}

		info, err := entry.Info()
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dest, filepath.FromSlash(relPath))

		if entry.IsDir() {
			return os.MkdirAll(targetPath, writableDirMode(info.Mode()))
		}
		return copyEmbeddedFile(srcFS, path, targetPath, writableFileMode(info.Mode()))
	})
}

func copyEmbeddedFile(srcFS fs.FS, src, dest string, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}

	srcFile, err := srcFS.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := makeExistingFileWritable(dest); err != nil {
		return err
	}

	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	return destFile.Chmod(mode)
}

func makeExistingFileWritable(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return nil
		}
		return os.Chmod(path, writableFileMode(info.Mode()))
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func writableDirMode(mode os.FileMode) os.FileMode {
	return mode.Perm() | 0o700
}

func writableFileMode(mode os.FileMode) os.FileMode {
	mode = mode.Perm() | 0o600
	if mode&0o111 != 0 {
		mode |= 0o100
	}
	return mode
}
