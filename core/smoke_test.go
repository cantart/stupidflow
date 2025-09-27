package core

import (
	"context"
	"testing"
)

func TestSmoke(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		ctx     context.Context
		wantErr bool
	}{
		{
			name:    "valid context",
			ctx:     context.Background(),
			wantErr: false,
		},
		{
			name:    "nil context",
			ctx:     nil,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := Smoke(tc.ctx)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Smoke() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
