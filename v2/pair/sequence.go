package pair

func MonadSequence[L, A, HKTA, HKTPA any](
	mmap func(HKTA, Kleisli[L, A, A]) HKTPA,
	fas Pair[L, HKTA],
) HKTPA {
	_ = "STUB: not implemented"
	return *new(HKTPA)
}

func MonadTraverse[L, A, HKTA, HKTPA any](
	mmap func(HKTA, Kleisli[L, A, A]) HKTPA,
	f func(A) HKTA,
	fas Pair[L, A],
) HKTPA {
	_ = "STUB: not implemented"
	return *new(HKTPA)
}

func Sequence[L, A, HKTA, HKTPA any](
	mmap func(Kleisli[L, A, A]) func(HKTA) HKTPA,
) func(Pair[L, HKTA]) HKTPA {
	_ = "STUB: not implemented"
	return nil
}

func Traverse[L, A, HKTA, HKTPA any](
	mmap func(Kleisli[L, A, A]) func(HKTA) HKTPA,
) func(func(A) HKTA) func(Pair[L, A]) HKTPA {
	_ = "STUB: not implemented"
	return nil
}
