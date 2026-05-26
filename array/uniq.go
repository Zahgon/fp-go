package array

// StrictUniq converts an array of arbitrary items into an array or unique items
// where uniqueness is determined by the built-in uniqueness constraint
func StrictUniq[A comparable](as []A) []A { _ = "STUB: not implemented"; return nil }

// Uniq converts an array of arbitrary items into an array or unique items
// where uniqueness is determined based on a key extractor function
func Uniq[A any, K comparable](f func(A) K) func(as []A) []A { _ = "STUB: not implemented"; return nil }
