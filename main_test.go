package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	wantVersion := strings.TrimSpace(versionString)
	tests := []struct {
		name     string
		args     []string
		wantCode int
		wantOut  string
	}{
		{"no args prints hello", []string{}, 0, "hello world\n"},
		{"-v prints version", []string{"-v"}, 0, wantVersion + "\n"},
		{"-version prints version", []string{"-version"}, 0, wantVersion + "\n"},
		{"--version prints version", []string{"--version"}, 0, wantVersion + "\n"},
		{"invalid flag returns non-zero", []string{"--nope"}, 2, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out, errOut bytes.Buffer
			code := run(tt.args, &out, &errOut)
			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d (stderr=%q)", code, tt.wantCode, errOut.String())
			}
			if tt.wantOut != "" && out.String() != tt.wantOut {
				t.Errorf("stdout = %q, want %q", out.String(), tt.wantOut)
			}
		})
	}
}
