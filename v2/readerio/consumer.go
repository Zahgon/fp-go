package readerio

//go:inline
func ChainConsumer[R, A any](c Consumer[A]) Operator[R, A, Void] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainFirstConsumer[R, A any](c Consumer[A]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}
