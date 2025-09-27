package core

import (
	"errors"
	"strings"
)

// BaseNode contains the common metadata required by all nodes.
type BaseNode struct {
	ID     string         `json:"id"`
	Config map[string]any `json:"config,omitempty"`
}

var (
	errNilBaseNode   = errors.New("base node must not be nil")
	errEmptyNodeID   = errors.New("node id must not be empty")
	errNilNodeConfig = errors.New("node config must not be nil")
)

// Validate returns a slice of validation errors for the base node.
func (n *BaseNode) Validate() []error {
	if n == nil {
		return []error{errNilBaseNode}
	}

	errs := make([]error, 0, 2)

	if strings.TrimSpace(n.ID) == "" {
		errs = append(errs, errEmptyNodeID)
	}

	if n.Config == nil {
		errs = append(errs, errNilNodeConfig)
	}

	return errs
}
