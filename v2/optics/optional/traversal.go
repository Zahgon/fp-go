package optional

import (
	"github.com/IBM/fp-go/v2/internal/functor"
	"github.com/IBM/fp-go/v2/internal/pointed"
)

func AsTraversal[R ~func(func(A) HKTA) func(S) HKTS, S, A, HKTS, HKTA any](
	fof pointed.OfType[S, HKTS],
	fmap functor.MapType[A, S, HKTA, HKTS],
) func(Optional[S, A]) R {
	_ = "STUB: not implemented"
	return nil
}
