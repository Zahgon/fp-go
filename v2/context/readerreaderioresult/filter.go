package readerreaderioresult

import (
	"github.com/IBM/fp-go/v2/option"
)

//go:inline
func Filter[C, HKTA, A any](
	filter func(Predicate[A]) Endomorphism[HKTA],
) func(Predicate[A]) Operator[C, HKTA, HKTA] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func FilterArray[C, A any](p Predicate[A]) Operator[C, []A, []A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func FilterIter[C, A any](p Predicate[A]) Operator[C, Seq[A], Seq[A]] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func FilterMap[C, HKTA, HKTB, A, B any](
	filter func(option.Kleisli[A, B]) Reader[HKTA, HKTB],
) func(option.Kleisli[A, B]) Operator[C, HKTA, HKTB] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func FilterMapArray[C, A, B any](p option.Kleisli[A, B]) Operator[C, []A, []B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func FilterMapIter[C, A, B any](p option.Kleisli[A, B]) Operator[C, Seq[A], Seq[B]] {
	_ = "STUB: not implemented"
	return nil
}
