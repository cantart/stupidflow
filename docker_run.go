package main

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

var execCommandContext = exec.CommandContext

func buildDockerArgs(cfgPath, binPath, outDir, image, cpuFile, memFile string) []string {
	args := []string{"run", "--rm",
		"-v", cfgPath + ":/app/dag.yaml",
		"-v", binPath + ":/app/stupidflow",
		"-v", outDir + ":/out",
		image, "/app/stupidflow",
		"--config", "/app/dag.yaml",
		"--profile", "/out/" + filepath.Base(cpuFile),
		"--mem-profile", "/out/" + filepath.Base(memFile),
	}
	return args
}

func runDAGInDocker(ctx context.Context, cfgPath, image, cpuFile, memFile string) error {
	binPath := filepath.Join(os.TempDir(), "stupidflow-bin")
	if err := execCommandContext(ctx, "go", "build", "-o", binPath, ".").Run(); err != nil {
		return err
	}
	outDir, err := os.MkdirTemp("", "stupidflow-out")
	if err != nil {
		return err
	}
	absCfg, err := filepath.Abs(cfgPath)
	if err != nil {
		return err
	}
	args := buildDockerArgs(absCfg, binPath, outDir, image, cpuFile, memFile)
	c := execCommandContext(ctx, "docker", args...)
	if err := c.Run(); err != nil {
		return err
	}
	if err := copyFile(filepath.Join(outDir, filepath.Base(cpuFile)), cpuFile); err != nil {
		return err
	}
	if err := copyFile(filepath.Join(outDir, filepath.Base(memFile)), memFile); err != nil {
		return err
	}
	os.Remove(binPath)
	os.RemoveAll(outDir)
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
