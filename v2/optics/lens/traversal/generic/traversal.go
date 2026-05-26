package generic

import (
	"github.com/IBM/fp-go/v2/internal/functor"
)

func Compose[B, HKTB, S, A, HKTS, HKTA any](
	fmap functor.MapType[A, S, HKTA, HKTS],
) func(Traversal[A, B, HKTA, HKTB]) func(Lens[S, A]) Traversal[S, B, HKTS, HKTB] {
	_ = "STUB: not implemented"
	return nil
}
