package option

import (
	"fmt"
)

// Options is a list of Option, and satisfies the Option interface.
type Options[T any] []Option[T]

// NewOptions wraps the provided Option values into an Options slice.
func NewOptions[T any](options ...Option[T]) Options[T] {
	opts := make(Options[T], len(options))
	for i, option := range options {
		opts[i] = option
	}

	return opts
}

// Apply applies a list of options to t.
func (o Options[T]) Apply(t T) (T, error) {
	for i, option := range o {
		var err error
		if t, err = option.Apply(t); err != nil {
			return t, fmt.Errorf("failed to apply option %d: %w", i, err)
		}
	}

	return t, nil
}
