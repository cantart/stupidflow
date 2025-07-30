package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func checkFileNotEmpty(t *testing.T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	if info.Size() == 0 {
		t.Fatalf("%s is empty", path)
	}
}

func TestRunDAGWithProfiles(t *testing.T) {
	cfg := Config{Tasks: []Task{{ID: "a", Type: TaskSleep, Duration: "1ms"}}}
	cases := []struct {
		name string
		cpu  bool
		mem  bool
	}{
		{name: "cpu", cpu: true},
		{name: "mem", mem: true},
		{name: "both", cpu: true, mem: true},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cpuFile := ""
			memFile := ""
			if tc.cpu {
				cpuFile = filepath.Join(t.TempDir(), "cpu.pprof")
			}
			if tc.mem {
				memFile = filepath.Join(t.TempDir(), "mem.pprof")
			}
			if err := runDAGWithProfiles(context.Background(), cfg, cpuFile, memFile); err != nil {
				t.Fatalf("run: %v", err)
			}
			if tc.cpu {
				checkFileNotEmpty(t, cpuFile)
			}
			if tc.mem {
				checkFileNotEmpty(t, memFile)
			}
		})
	}
}
