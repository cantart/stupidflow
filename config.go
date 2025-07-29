package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// TaskType enumerates possible workload types.
type TaskType string

const (
	TaskSleep       TaskType = "sleep"
	TaskCPUBurn     TaskType = "cpu_burn"
	TaskMemorySpike TaskType = "memory_spike"
	TaskFakeIO      TaskType = "fake_io"
)

type Task struct {
	ID        string   `yaml:"id"`
	Type      TaskType `yaml:"type"`
	Duration  string   `yaml:"duration,omitempty"`
	SizeMB    int      `yaml:"size_mb,omitempty"`
	Image     string   `yaml:"image,omitempty"`
	DependsOn []string `yaml:"depends_on,omitempty"`
}

type Config struct {
	Tasks []Task `yaml:"tasks"`
}

func loadConfig(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, validateConfig(cfg)
}

func validateConfig(cfg Config) error {
	if len(cfg.Tasks) != 1 {
		return fmt.Errorf("config must contain exactly one task")
	}
	return nil
}
