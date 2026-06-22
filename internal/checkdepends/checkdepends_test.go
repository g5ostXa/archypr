package checkdepends

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestMissingDependenciesReturnsOnlyMissingPackages(t *testing.T) {
	packages := []string{"hyprland", "waybar", "ghostty"}
	installedPackages := map[string]bool{
		"hyprland": true,
		"ghostty":  true,
	}

	missing := missingDependencies(packages, func(pkg string) bool {
		return installedPackages[pkg]
	})

	expected := []string{"waybar"}
	if !reflect.DeepEqual(missing, expected) {
		t.Fatalf("expected missing packages %v, got %v", expected, missing)
	}
}

func TestMissingDependenciesReturnsEmptyWhenAllInstalled(t *testing.T) {
	packages := []string{"hyprland", "waybar"}

	missing := missingDependencies(packages, func(string) bool {
		return true
	})

	if len(missing) != 0 {
		t.Fatalf("expected no missing packages, got %v", missing)
	}
}

func TestConfirmDependencyInstallAcceptsYes(t *testing.T) {
	var output bytes.Buffer

	confirmed := confirmDependencyInstall(strings.NewReader("y\n"), &output)

	if !confirmed {
		t.Fatal("expected y to confirm dependency installation")
	}
}

func TestConfirmDependencyInstallRejectsNoAndEmptyInput(t *testing.T) {
	tests := map[string]string{
		"no":    "n\n",
		"empty": "\n",
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			var output bytes.Buffer

			confirmed := confirmDependencyInstall(strings.NewReader(input), &output)

			if confirmed {
				t.Fatalf("expected %q to cancel dependency installation", input)
			}
		})
	}
}

func TestConfirmDependencyInstallRepromptsAfterInvalidInput(t *testing.T) {
	var output bytes.Buffer

	confirmed := confirmDependencyInstall(strings.NewReader("maybe\ny\n"), &output)

	if !confirmed {
		t.Fatal("expected y after invalid input to confirm dependency installation")
	}
}
