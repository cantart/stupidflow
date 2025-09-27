package core

import "testing"

func TestBaseNodeValidate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		node    *BaseNode
		wantLen int
	}{
		{
			name: "valid node",
			node: &BaseNode{
				ID:     "start",
				Config: map[string]any{},
			},
			wantLen: 0,
		},
		{
			name: "missing id",
			node: &BaseNode{
				ID:     "",
				Config: map[string]any{},
			},
			wantLen: 1,
		},
		{
			name: "nil config",
			node: &BaseNode{
				ID: "task",
			},
			wantLen: 1,
		},
		{
			name:    "missing id and config",
			node:    &BaseNode{},
			wantLen: 2,
		},
		{
			name:    "nil base node",
			node:    nil,
			wantLen: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.node.Validate()
			if len(err) != tc.wantLen {
				t.Fatalf("Validate() error count = %d, want %d", len(err), tc.wantLen)
			}
		})
	}
}
