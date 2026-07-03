package sethypr

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/g5ostXa/archypr/internal/core"
	"github.com/g5ostXa/archypr/version"
)

var versionDir = "version"

func VersionCopy() {

	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		core.Logger.Fatal("Failed to get user's home directory...", err)
	}

	// Run this if successfully determined user's home dir
	versionDestDir := filepath.Join(homeDir, versionDir)

	core.Logger.Info("Copying version...")

	if err := copyEmbeddedversionDir(version.FS, ".", versionDestDir); err != nil {
		core.Logger.Fatal("Failed to copy version...", err)
	}
	core.Logger.Info("version copied successfully")
}

func copyEmbeddedversionDir(srcFS fs.FS, srcRoot, dest string) error {

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
			return os.MkdirAll(targetPath, writableversionDirMode(info.Mode()))
		}
		return copyEmbeddedversionFile(srcFS, path, targetPath, writableversionFileMode(info.Mode()))
	})
}

func copyEmbeddedversionFile(srcFS fs.FS, src, dest string, mode os.FileMode) error {

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}

	srcFile, err := srcFS.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := makeWritableversion(dest); err != nil {
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

func makeWritableversion(path string) error {

	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return nil
		}
		return os.Chmod(path, writableversionFileMode(info.Mode()))
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func writableversionDirMode(mode os.FileMode) os.FileMode {

	return mode.Perm() | 0o700
}

func writableversionFileMode(mode os.FileMode) os.FileMode {

	mode = mode.Perm() | 0o600
	if mode&0o111 != 0 {
		mode |= 0o100
	}
	return mode
}
