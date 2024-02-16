package option

// Func is a function that satisfies the Option interface.
type Func[T any] func(T) (T, error)

// Apply applies the Func to t.
func (f Func[T]) Apply(t T) (T, error) {
	return f(t)
}
