package generic

import (
	"github.com/IBM/fp-go/v2/internal/functor"
)

func FromLens[S, A, HKTES, HKTA any](
	fmap functor.MapType[A, Endomorphism[S], HKTA, HKTES],
) func(Lens[S, A]) Traversal[S, A, HKTES, HKTA] {
	_ = "STUB: not implemented"
	return nil
}
