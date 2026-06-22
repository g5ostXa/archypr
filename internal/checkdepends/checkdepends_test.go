package checkdepends

import (
	"reflect"
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
