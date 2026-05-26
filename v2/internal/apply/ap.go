// Copyright (c) 2023 - 2025 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apply

func MonadAp[HKTGA, HKTGB, HKTGAB, HKTFGAB, HKTFGGAB, HKTFGA, HKTFGB any](
	fap func(HKTFGGAB, HKTFGA) HKTFGB,
	fmap func(HKTFGAB, func(HKTGAB) func(HKTGA) HKTGB) HKTFGGAB,
	gap func(HKTGAB, HKTGA) HKTGB,

	fab HKTFGAB,
	fa HKTFGA) HKTFGB {
	_ = "STUB: not implemented"
	return *new(HKTFGB)
}

func Ap[HKTGA, HKTGB, HKTGAB, HKTFGAB, HKTFGGAB, HKTFGA, HKTFGB any](
	fap func(HKTFGA) func(HKTFGGAB) HKTFGB,
	fmap func(func(HKTGAB) func(HKTGA) HKTGB) func(HKTFGAB) HKTFGGAB,
	gap func(HKTGA) func(HKTGAB) HKTGB,

	fa HKTFGA) func(HKTFGAB) HKTFGB {
	_ = "STUB: not implemented"
	return nil
}

// func Ap[HKTGA, HKTGB, HKTGAB, HKTFGAB, HKTFGGAB, HKTFGA, HKTFGB any](
// 	fap func(HKTFGA) func(HKTFGGAB) HKTFGB,
// 	fmap func(func(HKTGAB) func(HKTGA) HKTGB) func(HKTFGAB) HKTFGGAB,
// 	gap func(HKTGA) func(HKTGAB) HKTGB,

// 	fa HKTFGA) func(HKTFGAB) HKTFGB {

// 	return fap(fmap(F.Bind1st(F.Bind1st[HKTGAB, HKTGA, HKTGB], gap)), fa)
// }

// export function ap<F, G>(
// 	F: Apply<F>,
// 	G: Apply<G>
//   ): <A>(fa: HKT<F, HKT<G, A>>) => <B>(fab: HKT<F, HKT<G, (a: A) => B>>) => HKT<F, HKT<G, B>> {
// 	return <A>(fa: HKT<F, HKT<G, A>>) => <B>(fab: HKT<F, HKT<G, (a: A) => B>>): HKT<F, HKT<G, B>> =>
// 	  F.ap(
// 		F.map(fab, (gab) => (ga: HKT<G, A>) => G.ap(gab, ga)),
// 		fa
// 	  )
//   }

//  function apFirst<F>(A: Apply<F>): <B>(second: HKT<F, B>) => <A>(first: HKT<F, A>) => HKT<F, A> {
// 	return (second) => (first) =>
// 	  A.ap(
// 		A.map(first, (a) => () => a),
// 		second
// 	  )
//   }

// Functor<F>.map: <A, () => A>(fa: HKT<F, A>, f: (a: A) => () => A) => HKT<F, () => A>

// Apply<F>.ap: <B, A>(fab: HKT<F, (a: B) => A>, fa: HKT<F, B>) => HKT<F, A>

func MonadApFirst[HKTGA, HKTGB, HKTGBA, A, B any](
	fap func(HKTGBA, HKTGB) HKTGA,
	fmap func(HKTGA, func(A) func(B) A) HKTGBA,

	first HKTGA,
	second HKTGB,
) HKTGA {
	_ = "STUB: not implemented"
	return *new(HKTGA)
}

func ApFirst[HKTGA, HKTGB, HKTGBA, A, B any](
	fap func(HKTGB) func(HKTGBA) HKTGA,
	fmap func(func(A) func(B) A) func(HKTGA) HKTGBA,

	second HKTGB,
) func(HKTGA) HKTGA {
	_ = "STUB: not implemented"
	return nil
}

func MonadApSecond[HKTGA, HKTGB, HKTGBB, A, B any](
	fap func(HKTGBB, HKTGB) HKTGB,
	fmap func(HKTGA, func(A) func(B) B) HKTGBB,

	first HKTGA,
	second HKTGB,
) HKTGB {
	_ = "STUB: not implemented"
	return *new(HKTGB)
}

func ApSecond[HKTGA, HKTGB, HKTGBB, A, B any](
	fap func(HKTGB) func(HKTGBB) HKTGB,
	fmap func(func(A) func(B) B) func(HKTGA) HKTGBB,

	second HKTGB,
) func(HKTGA) HKTGB {
	_ = "STUB: not implemented"
	return nil
}

func MonadApS[S1, S2, B, HKTBGBS2, HKTS1, HKTS2, HKTB any](
	fap func(HKTBGBS2, HKTB) HKTS2,
	fmap func(HKTS1, func(S1) func(B) S2) HKTBGBS2,
	fa HKTS1,
	key func(B) func(S1) S2,
	fb HKTB,
) HKTS2 {
	_ = "STUB: not implemented"
	return *new(HKTS2)
}

func ApS[S1, S2, B, HKTBGBS2, HKTS1, HKTS2, HKTB any](
	fap func(HKTB) func(HKTBGBS2) HKTS2,
	fmap func(func(S1) func(B) S2) func(HKTS1) HKTBGBS2,
	key func(B) func(S1) S2,
	fb HKTB,
) func(HKTS1) HKTS2 {
	_ = "STUB: not implemented"
	return nil
}
