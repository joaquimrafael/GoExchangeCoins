package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	in := strings.NewReader("2\nUSD\nBRL\nclose\n")
	var out strings.Builder

	run(in, &out)

	got := out.String()
	if !strings.Contains(got, "Result: BRL 10.20") {
		t.Fatalf("expected conversion result in output, got:\n%s", got)
	}
}

func TestRunClose(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{name: "AtValuePrompt", input: "close\n"},
		{name: "AtOriginPrompt", input: "2\nclose\n"},
		{name: "AtTargetPrompt", input: "2\nUSD\nclose\n"},
		{name: "Uppercase", input: "CLOSE\n"},
		{name: "MixedCase", input: "Close\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out strings.Builder

			run(strings.NewReader(tc.input), &out)

			got := out.String()
			if strings.Contains(got, "Result:") {
				t.Fatalf("close should exit before any conversion, got result in output:\n%s", got)
			}
			if strings.Contains(got, "value must be a number") {
				t.Fatalf("close should not be parsed as a value, got error in output:\n%s", got)
			}
		})
	}
}
