// Package option provides a small, strongly typed set of tools for building with the options pattern.
package option

// Option provides an interface for options to satisfy.
// Concrete usages of Option are typically Funcs returned from a builder function.
type Option[T any] interface {
	Apply(T) (T, error)
}

// Apply applies a list of options to t.
func Apply[T any](t T, options ...Option[T]) (T, error) {
	return NewOptions(options...).Apply(t)
}
