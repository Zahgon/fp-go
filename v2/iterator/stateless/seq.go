package stateless

// ToSeq converts the stateless [Iterator] to an idiomatic go iterator
func ToSeq[T any](it Iterator[T]) Seq[T] { _ = "STUB: not implemented"; return nil }

// ToSeq2 converts the stateless [Iterator] to an idiomatic go iterator
func ToSeq2[K, V any](it Iterator[Pair[K, V]]) Seq2[K, V] { _ = "STUB: not implemented"; return nil }
