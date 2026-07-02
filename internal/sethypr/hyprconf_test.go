package sethypr

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCopiesFiles(t *testing.T) {

	t.Parallel()

	srcFS := fstest.MapFS{
		"btop": &fstest.MapFile{
			Mode: os.ModeDir | 0o555,
		},
		"btop/btop.conf": &fstest.MapFile{
			Data: []byte("theme = mocha\n"),
			Mode: 0o444,
		},
	}
	dest := t.TempDir()

	if err := copyEmbeddedDir(srcFS, ".", dest); err != nil {

		t.Fatalf("copyEmbeddedDir() error = %v", err)
	}

	contents, err := os.ReadFile(filepath.Join(dest, "btop", "btop.conf"))
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if string(contents) != "theme = mocha\n" {
		t.Fatalf("copied file contents = %q", contents)
	}

	if err := os.WriteFile(filepath.Join(dest, "btop", "another.conf"), []byte("ok\n"), 0o644); err != nil {
		t.Fatalf("copied directory should be writable: %v", err)
	}
}

func TestCopyEmbeddedDir(t *testing.T) {

	t.Parallel()

	srcFS := fstest.MapFS{
		"fish/config.fish": &fstest.MapFile{
			Data: []byte("set fish_greeting\n"),
			Mode: 0o444,
		},
	}
	dest := t.TempDir()
	destFile := filepath.Join(dest, "fish", "config.fish")

	if err := os.MkdirAll(filepath.Dir(destFile), 0o755); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	if err := os.WriteFile(destFile, []byte("old\n"), 0o444); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if err := copyEmbeddedDir(srcFS, ".", dest); err != nil {
		t.Fatalf("copyEmbeddedDir() error = %v", err)
	}

	contents, err := os.ReadFile(destFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if string(contents) != "set fish_greeting\n" {
		t.Fatalf("copied file contents = %q", contents)
	}
}
