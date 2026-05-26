package witherable

import (
	"github.com/IBM/fp-go/v2/internal/filterable"
	"github.com/IBM/fp-go/v2/internal/functor"
)

func Filter[A, HKT_G_A, HKT_F_HKT_G_A any](
	fmap functor.MapType[HKT_G_A, HKT_G_A, HKT_F_HKT_G_A, HKT_F_HKT_G_A],
	ffilter filterable.FilterType[A, HKT_G_A],
) func(func(A) bool) func(HKT_F_HKT_G_A) HKT_F_HKT_G_A {
	_ = "STUB: not implemented"
	return nil
}

func FilterMap[A, B, HKT_G_A, HKT_G_B, HKT_F_HKT_G_A, HKT_F_HKT_G_B any](
	fmap functor.MapType[HKT_G_A, HKT_G_B, HKT_F_HKT_G_A, HKT_F_HKT_G_B],
	ffilter filterable.FilterMapType[A, B, HKT_G_A, HKT_G_B],
) func(func(A) Option[B]) func(HKT_F_HKT_G_A) HKT_F_HKT_G_B {
	_ = "STUB: not implemented"
	return nil
}

// func Wither[A, HKTWA, HKTFOB, HKTFWB any](

// ) WitherType[A, HKTWA, HKTFOB, HKTFWB] {

// 	return func(f func(A) HKTFOB) func(HKTWA) HKTFWB {

// 	}
// }
