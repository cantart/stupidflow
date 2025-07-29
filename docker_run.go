package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
)

func buildDockerArgs(cfgPath, binPath, outDir, image string, profile bool) []string {
	args := []string{"run", "--rm", "-v", cfgPath + ":/app/dag.yaml", "-v", binPath + ":/app/stupidflow", "-v", outDir + ":/out", image, "/app/stupidflow", "--config", "/app/dag.yaml"}
	if profile {
		args = append(args, "--profile", "/out/profile.pb.gz")
	}
	return args
}

func runDAGInDocker(ctx context.Context, cfgPath, image string, profile bool) (string, error) {
	binPath := filepath.Join(os.TempDir(), "stupidflow-bin")
	if err := exec.CommandContext(ctx, "go", "build", "-o", binPath, ".").Run(); err != nil {
		return "", err
	}
	outDir, err := os.MkdirTemp("", "stupidflow-out")
	if err != nil {
		return "", err
	}
	absCfg, err := filepath.Abs(cfgPath)
	if err != nil {
		return "", err
	}
	args := buildDockerArgs(absCfg, binPath, outDir, image, profile)
	c := exec.CommandContext(ctx, "docker", args...)
	if err := c.Run(); err != nil {
		return "", err
	}
	if !profile {
		return "", nil
	}
	prof := filepath.Join(outDir, "profile.pb.gz")
	out, err := exec.CommandContext(ctx, "go", "tool", "pprof", "-top", "-nodecount=5", binPath, prof).CombinedOutput()
	return string(out), err
}
