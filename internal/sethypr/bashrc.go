package sethypr

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/g5ostXa/archypr/bashrc"
	"github.com/g5ostXa/archypr/internal/core"
)

// Since .bashrc is at the root of the embedded package, walk from "."
var bashrcPath = "."

func BashrcCopy() {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		core.Logger.Fatal("Failed to get user's home directory...", err)
	}

	// Points cleanly to ~/
	BashrcDestDir := filepath.Join(homeDir)

	core.Logger.Info("Copying Bashrc...")

	// Pass the embedded file system from the bashrc package
	if err := copyEmbeddedbashrcPath(bashrc.FS, bashrcPath, BashrcDestDir); err != nil {
		core.Logger.Fatal("Failed to copy Bashrc...", err)
	}
	core.Logger.Info("Bashrc copied successfully")
}

func copyEmbeddedbashrcPath(srcFS fs.FS, srcRoot, dest string) error {
	return fs.WalkDir(srcFS, srcRoot, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Example: path is ".bashrc", srcRoot is "." -> relPath becomes ".bashrc"
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
		// targetPath joins (~/) with (".bashrc") -> (~/.bashrc)
		targetPath := filepath.Join(dest, filepath.FromSlash(relPath))
		if entry.IsDir() {
			return os.MkdirAll(targetPath, writablebashrcPathMode(info.Mode()))
		}
		return copyEmbeddedBashrcFile(srcFS, path, targetPath, writableBashrcFileMode(info.Mode()))
	})
}

func copyEmbeddedBashrcFile(srcFS fs.FS, src, dest string, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	srcFile, err := srcFS.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := makeWritableBashrc(dest); err != nil {
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

func makeWritableBashrc(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return nil
		}
		return os.Chmod(path, writableBashrcFileMode(info.Mode()))
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func writablebashrcPathMode(mode os.FileMode) os.FileMode {
	return mode.Perm() | 0o700
}

func writableBashrcFileMode(mode os.FileMode) os.FileMode {
	mode = mode.Perm() | 0o600
	if mode&0o111 != 0 {
		mode |= 0o100
	}
	return mode
}
