package sethypr

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/g5ostXa/archypr/assets"
	"github.com/g5ostXa/archypr/internal/core"
)

const AssetsDir = "assets"

func AssetsCopy() {

	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		core.Logger.Fatal("Failed to get user's home directory...", err)
	}

	// Run this if successfully determined user's home dir
	assetsDestDir := filepath.Join(homeDir, AssetsDir)

	core.Logger.Info("Copying assets...")

	if err := copyEmbeddedAssetsDir(assets.FS, ".", assetsDestDir); err != nil {
		core.Logger.Fatal("Failed to copy assets...", err)
	}
	core.Logger.Info("Assets copied successfully")
}

func copyEmbeddedAssetsDir(srcFS fs.FS, srcRoot, dest string) error {

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
			return os.MkdirAll(targetPath, writableAssetsDirMode(info.Mode()))
		}
		return copyEmbeddedAssetsFile(srcFS, path, targetPath, writableAssetsFileMode(info.Mode()))
	})
}

func copyEmbeddedAssetsFile(srcFS fs.FS, src, dest string, mode os.FileMode) error {

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}

	srcFile, err := srcFS.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := makeWritableAssets(dest); err != nil {
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

func makeWritableAssets(path string) error {

	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return nil
		}
		return os.Chmod(path, writableAssetsFileMode(info.Mode()))
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func writableAssetsDirMode(mode os.FileMode) os.FileMode {

	return mode.Perm() | 0o700
}

func writableAssetsFileMode(mode os.FileMode) os.FileMode {

	mode = mode.Perm() | 0o600
	if mode&0o111 != 0 {
		mode |= 0o100
	}
	return mode
}
