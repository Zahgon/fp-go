package generic

import (
	M "github.com/IBM/fp-go/v2/monoid"
)

//go:inline
func ApplicativeMonoid[GA ~func(R) A, R, A any](m M.Monoid[A]) M.Monoid[GA] {
	_ = "STUB: not implemented"
	return nil
}
