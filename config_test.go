package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	yamlData := "tasks:\n- id: a\n  type: sleep\n  duration: 1s\n  image: busybox\n"
	tmp := filepath.Join(t.TempDir(), "dag.yaml")
	if err := os.WriteFile(tmp, []byte(yamlData), 0o600); err != nil {
		t.Fatalf("write tmp: %v", err)
	}
	cfg, err := loadConfig(tmp)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if len(cfg.Tasks) != 1 || cfg.Tasks[0].ID != "a" || cfg.Tasks[0].Image != "busybox" || cfg.Tasks[0].Type != TaskSleep {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}
