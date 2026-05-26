package iooption

//go:inline
func ChainConsumer[A any](c Consumer[A]) Operator[A, struct{}] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainFirstConsumer[A any](c Consumer[A]) Operator[A, A] { _ = "STUB: not implemented"; return nil }
