package aiven

// ref is a helper function to return a pointer to a value.
func ref[T any](v T) *T {
	return &v
}
