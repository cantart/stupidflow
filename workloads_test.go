package main

import (
	"context"
	"testing"
)

func TestRunFakeIO(t *testing.T) {
	cases := []struct {
		name    string
		cancel  bool
		wantErr bool
	}{
		{name: "ok", cancel: false, wantErr: false},
		{name: "canceled", cancel: true, wantErr: true},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			if tc.cancel {
				cancel()
			} else {
				defer cancel()
			}
			err := runFakeIO(ctx, 1)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
