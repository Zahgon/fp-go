package readerioeither

// ChainConsumer chains a consumer (side-effect function) into a ReaderIOEither computation,
// replacing the success value with an empty struct.
//
// This is useful for performing side effects (like logging or printing) where you don't
// need to preserve the original value.
//
// Type parameters:
//   - R: The context type
//   - E: The error type
//   - A: The value type to consume
//
// Parameters:
//   - c: A consumer function that performs a side effect
//
// Returns:
//
//	An Operator that executes the consumer and returns struct{}
//
//go:inline
func ChainConsumer[R, E, A any](c Consumer[A]) Operator[R, E, A, struct{}] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstConsumer chains a consumer into a ReaderIOEither computation while preserving
// the original value.
//
// This is useful for performing side effects (like logging or printing) where you want
// to keep the original value for further processing.
//
// Type parameters:
//   - R: The context type
//   - E: The error type
//   - A: The value type to consume
//
// Parameters:
//   - c: A consumer function that performs a side effect
//
// Returns:
//
//	An Operator that executes the consumer and returns the original value
//
//go:inline
func ChainFirstConsumer[R, E, A any](c Consumer[A]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}
