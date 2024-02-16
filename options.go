package option

import (
	"errors"
	"fmt"
)

// Options is a list of Option, and satisfies the Option interface.
type Options[T any] []Option[T]

// NewOptions converts a []Option to an Options.
func NewOptions[T any](options ...Option[T]) Options[T] {
	opts := make(Options[T], len(options))
	for i, option := range options {
		opts[i] = option
	}

	return opts
}

// Apply applies a list of options to t.
func (o Options[T]) Apply(t T) (T, error) {
	var errs []error
	for i, option := range o {
		var err error
		if t, err = option.Apply(t); err != nil {
			errs = append(errs, fmt.Errorf("failed to apply option %d: %w", i, err))
		}
	}

	return t, errors.Join(errs...)
}
