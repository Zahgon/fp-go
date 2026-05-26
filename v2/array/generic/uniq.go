package generic

// StrictUniq converts an array of arbitrary items into an array or unique items
// where uniqueness is determined by the built-in uniqueness constraint
func StrictUniq[AS ~[]A, A comparable](as AS) AS { _ = "STUB: not implemented"; return *new(AS) }

// uniquePredUnsafe returns a predicate on a map for uniqueness
func uniquePredUnsafe[PRED ~func(A) K, A any, K comparable](f PRED) func(int, A) bool {
	_ = "STUB: not implemented"
	return nil
}

// Uniq converts an array of arbitrary items into an array or unique items
// where uniqueness is determined based on a key extractor function
func Uniq[AS ~[]A, PRED ~func(A) K, A any, K comparable](f PRED) func(as AS) AS {
	_ = "STUB: not implemented"
	return nil

	// we need to create a new predicate for each iteration
}
