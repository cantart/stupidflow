package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildDockerArgs(t *testing.T) {
	args := buildDockerArgs("/cfg.yaml", "/bin/app", "/out", "alpine", "cpu.pprof", "mem.pprof")
	want := []string{"run", "--rm", "-v", "/cfg.yaml:/app/dag.yaml", "-v", "/bin/app:/app/stupidflow", "-v", "/out:/out", "alpine", "/app/stupidflow", "--config", "/app/dag.yaml", "--profile", "/out/cpu.pprof", "--mem-profile", "/out/mem.pprof"}
	if len(args) != len(want) {
		t.Fatalf("unexpected len %d", len(args))
	}
	for i := range want {
		if args[i] != want[i] {
			t.Fatalf("want %s got %s at %d", want[i], args[i], i)
		}
	}
}

func TestRunDAGInDocker_RemovesTempFiles(t *testing.T) {
	temp := t.TempDir()
	os.Setenv("TMPDIR", temp)
	defer os.Unsetenv("TMPDIR")

	execCommandContext = func(ctx context.Context, name string, args ...string) *exec.Cmd {
		if name == "docker" {
			for i := range args {
				if args[i] == "-v" && i+1 < len(args) {
					m := args[i+1]
					if strings.HasSuffix(m, ":/out") {
						dir := strings.TrimSuffix(m, ":/out")
						os.WriteFile(filepath.Join(dir, "cpu.pprof"), []byte("cpu"), 0o600)
						os.WriteFile(filepath.Join(dir, "mem.pprof"), []byte("mem"), 0o600)
						break
					}
				}
			}
		}
		return exec.CommandContext(ctx, "true")
	}
	defer func() { execCommandContext = exec.CommandContext }()

	cfgPath := filepath.Join(temp, "dag.yaml")
	os.WriteFile(cfgPath, []byte("tasks:\n- id: a\n  type: sleep\n"), 0o600)

	cpu := filepath.Join(temp, "cpu.pprof")
	mem := filepath.Join(temp, "mem.pprof")
	if err := runDAGInDocker(context.Background(), cfgPath, "alpine", cpu, mem); err != nil {
		t.Fatalf("run: %v", err)
	}
	if _, err := os.Stat(filepath.Join(temp, "stupidflow-bin")); !os.IsNotExist(err) {
		t.Fatalf("binary not removed")
	}
	matches, err := filepath.Glob(filepath.Join(temp, "stupidflow-out*"))
	if err != nil {
		t.Fatalf("glob: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("out dir not removed")
	}
}
