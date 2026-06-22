package pacmanconf

import "testing"

func TestUpdatePacmanConf(t *testing.T) {

	input := []byte(`#Color
#VerbosePkgLists
#ParallelDownloads = 5
`)

	got, changes := updatePacmanConf(input)
	if len(changes) == 0 {
		t.Fatal("expected pacman.conf to change")
	}

	want := []byte(`Color
VerbosePkgLists
ParallelDownloads = 5
ILoveCandy
`)
	if string(got) != string(want) {
		t.Fatalf("unexpected pacman.conf output\nwant:\n%s\ngot:\n%s", want, got)
	}
}

func TestUpdatePacmanConfIsIdempotent(t *testing.T) {

	input := []byte(`Color
VerbosePkgLists
ParallelDownloads = 5
ILoveCandy
`)

	got, changes := updatePacmanConf(input)
	if len(changes) != 0 {
		t.Fatal("expected already configured pacman.conf to be unchanged")
	}

	if string(got) != string(input) {
		t.Fatalf("expected output to match input\nwant:\n%s\ngot:\n%s", input, got)
	}
}
