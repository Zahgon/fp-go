package codec

func onTypeError(expType string) func(any) error { _ = "STUB: not implemented"; return nil }

// Is checks if a value can be converted to type T.
// Returns Some(value) if the conversion succeeds, None otherwise.
// This is a type-safe cast operation.
func Is[T any]() ReaderResult[any, T] { _ = "STUB: not implemented"; return nil }
