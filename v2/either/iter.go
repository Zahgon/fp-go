package either

func TraverseIterG[GA ~func(yield func(A) bool), GB ~func(yield func(B) bool), E, A, B any](f Kleisli[E, A, B]) Kleisli[E, GA, GB] {
	_ = "STUB: not implemented"
	return nil
}

func TraversableIter[E, A, B any]() Traversable[E, A, B, Seq[A], Seq[B]] {
	_ = "STUB: not implemented"
	return nil
}
