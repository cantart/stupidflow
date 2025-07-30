package main

import (
	"context"
	"testing"
)

type dagTestCase struct {
	name    string
	cfg     Config
	wantErr bool
}

func TestRunDAG(t *testing.T) {
	cases := []dagTestCase{
		{
			name:    "single task",
			cfg:     Config{Tasks: []Task{{ID: "a", Type: TaskSleep, Duration: "1ms"}}},
			wantErr: false,
		},
		{
			name: "multiple tasks",
			cfg: Config{Tasks: []Task{
				{ID: "a", Type: TaskSleep},
				{ID: "b", Type: TaskSleep},
			}},
			wantErr: true,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := runDAG(context.Background(), tc.cfg)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestRunDAGWithRunner(t *testing.T) {
	calls := []string{}
	runner := func(ctx context.Context, t Task) error {
		calls = append(calls, t.ID)
		return nil
	}
	cfg := Config{Tasks: []Task{{ID: "a", Type: TaskSleep}}}
	if err := runDAGWithRunner(context.Background(), cfg, runner); err != nil {
		t.Fatalf("run: %v", err)
	}
	if len(calls) != 1 || calls[0] != "a" {
		t.Fatalf("unexpected order: %v", calls)
	}
}
