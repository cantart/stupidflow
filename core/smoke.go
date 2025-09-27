package core

import (
	"context"
	"errors"
)

var errNilContext = errors.New("context must not be nil")

// Smoke validates that basic core preconditions are satisfied before tests run.
func Smoke(ctx context.Context) error {
	if ctx == nil {
		return errNilContext
	}

	return nil
}
