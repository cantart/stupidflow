package main

import "testing"

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
