package main

import (
	"context"
	"os"
	"os/exec"
	"runtime/pprof"
)

func runDAGProfile(ctx context.Context, cfg Config) (string, error) {
	tmp, err := os.CreateTemp("", "stupidflow-prof")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())
	if err := pprof.StartCPUProfile(tmp); err != nil {
		return "", err
	}
	err = runDAG(ctx, cfg)
	pprof.StopCPUProfile()
	if err != nil {
		return "", err
	}
	out, err := exec.CommandContext(ctx, "go", "tool", "pprof", "-top", "-nodecount=5", os.Args[0], tmp.Name()).CombinedOutput()
	return string(out), err
}
