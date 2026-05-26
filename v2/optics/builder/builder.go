package builder

func MakeBuilder[S, A any](get func(S) Option[A], set func(A) Endomorphism[S], name string) Builder[S, A] {
	_ = "STUB: not implemented"
	return nil
}

func ComposeLensPrism[S, A, B any](r Prism[A, B]) func(Lens[S, A]) Builder[S, B] {
	_ = "STUB: not implemented"
	return nil
}
