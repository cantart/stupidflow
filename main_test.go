package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestExecute_UsesLocalWhenDockerImageEmpty(t *testing.T) {
	temp := t.TempDir()
	cfgPath := filepath.Join(temp, "dag.yaml")
	os.WriteFile(cfgPath, []byte("tasks:\n- id: a\n  type: sleep\n"), 0o600)

	called := ""
	origLocal := runDAGWithProfilesFunc
	origDocker := runDAGInDockerFunc
	runDAGWithProfilesFunc = func(ctx context.Context, cfg Config, cpuFile, memFile string) error {
		called = "local"
		return nil
	}
	runDAGInDockerFunc = func(ctx context.Context, cfgPath, image, cpuFile, memFile string) error {
		called = "docker"
		return nil
	}
	defer func() {
		runDAGWithProfilesFunc = origLocal
		runDAGInDockerFunc = origDocker
	}()

	err := execute(context.Background(), cfgPath, "", "cpu.pprof", "mem.pprof")
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if called != "local" {
		t.Fatalf("expected local, got %s", called)
	}
}
