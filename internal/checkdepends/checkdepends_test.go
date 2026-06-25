package checkdepends

import (
	"reflect"
	"testing"
)

func TestMissingDependencies(t *testing.T) {
	expected := []string{packages[0], packages[len(packages)/2], packages[len(packages)-1]}
	missingPackages := make(map[string]bool, len(expected))
	for _, pkg := range expected {
		missingPackages[pkg] = true
	}

	missing := missingDependencies(packages, func(pkg string) bool {
		return !missingPackages[pkg]
	})

	if !reflect.DeepEqual(missing, expected) {
		t.Fatalf("expected missing packages %v, got %v", expected, missing)
	}
}

func TestMissingDependenciesAllInstalled(t *testing.T) {
	missing := missingDependencies(packages, func(string) bool {
		return true
	})

	if len(missing) != 0 {
		t.Fatalf("expected no missing packages, got %v", missing)
	}
}

func TestMissingDependenciesChecksAll(t *testing.T) {
	checked := make(map[string]int, len(packages))

	missingDependencies(packages, func(pkg string) bool {
		checked[pkg]++
		return true
	})

	for _, pkg := range packages {
		if checked[pkg] != 1 {
			t.Fatalf("expected %s to be checked once, got %d", pkg, checked[pkg])
		}
	}
}
