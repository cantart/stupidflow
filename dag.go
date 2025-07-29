package main

import (
	"context"
	"errors"
	"time"
)

func runTask(ctx context.Context, t Task) error {
	d, _ := time.ParseDuration(t.Duration)
	switch t.Type {
	case TaskSleep:
		return runSleep(ctx, d)
	case TaskCPUBurn:
		return runCPUBurn(ctx, d)
	case TaskMemorySpike:
		return runMemorySpike(ctx, t.SizeMB, d)
	case TaskFakeIO:
		return runFakeIO(ctx, t.SizeMB)
	default:
		return errors.New("unknown task type")
	}
}

func runDAGWithRunner(ctx context.Context, cfg Config, exec func(context.Context, Task) error) error {
	remaining := make(map[string]Task)
	for _, t := range cfg.Tasks {
		remaining[t.ID] = t
	}
	completed := make(map[string]bool)

	for len(remaining) > 0 {
		progress := false
		for id, t := range remaining {
			ready := true
			for _, dep := range t.DependsOn {
				if !completed[dep] {
					ready = false
					break
				}
			}
			if !ready {
				continue
			}
			if err := exec(ctx, t); err != nil {
				return err
			}
			completed[id] = true
			delete(remaining, id)
			progress = true
		}
		if !progress {
			return errors.New("cyclic dependencies detected")
		}
	}
	return nil
}

func runDAG(ctx context.Context, cfg Config) error {
	return runDAGWithRunner(ctx, cfg, runTask)
}
