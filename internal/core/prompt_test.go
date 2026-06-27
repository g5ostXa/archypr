package core

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestConfirmAcceptsYes(t *testing.T) {

	var output bytes.Buffer
	silenceLogger(t, &output)

	confirmed := Confirm(strings.NewReader("y\n"), "Install?")

	if !confirmed {
		t.Fatal("expected y to confirm")
	}
}

func TestConfirmRejectsNoAndEmptyInput(t *testing.T) {

	tests := map[string]string{
		"no":    "n\n",
		"empty": "\n",
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			var output bytes.Buffer
			silenceLogger(t, &output)

			confirmed := Confirm(strings.NewReader(input), "Install?")

			if confirmed {
				t.Fatalf("expected %q to cancel", input)
			}
		})
	}
}

func TestConfirmRepromptsAfterInvalidInput(t *testing.T) {

	var output bytes.Buffer
	silenceLogger(t, &output)

	confirmed := Confirm(strings.NewReader("maybe\ny\n"), "Install?")

	if !confirmed {
		t.Fatal("expected y after invalid input to confirm")
	}
}

func silenceLogger(t *testing.T, output *bytes.Buffer) {

	t.Helper()

	Logger.SetOutput(output)
	t.Cleanup(func() {
		Logger.SetOutput(os.Stderr)
	})
}
