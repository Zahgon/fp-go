// Copyright (c) 2023 IBM Corp.
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

package functor

// HKTFGA = HKT[F, HKT[G, A]]
// HKTFGB = HKT[F, HKT[G, B]]
func MonadMap[A, B, HKTGA, HKTGB, HKTFGA, HKTFGB any](
	fmap func(HKTFGA, func(HKTGA) HKTGB) HKTFGB,
	gmap func(HKTGA, func(A) B) HKTGB,
	fa HKTFGA,
	f func(A) B) HKTFGB {
	_ = "STUB: not implemented"
	return *new(HKTFGB)
}

func Map[A, B, HKTGA, HKTGB, HKTFGA, HKTFGB any](
	fmap func(func(HKTGA) HKTGB) func(HKTFGA) HKTFGB,
	gmap func(func(A) B) func(HKTGA) HKTGB,
	f func(A) B) func(HKTFGA) HKTFGB {
	_ = "STUB: not implemented"
	return nil
}

func MonadLet[S1, S2, B, HKTS1, HKTS2 any](
	mmap func(HKTS1, func(S1) S2) HKTS2,
	first HKTS1,
	key func(B) func(S1) S2,
	f func(S1) B,
) HKTS2 {
	_ = "STUB: not implemented"
	return *new(HKTS2)
}

func Let[S1, S2, B, HKTS1, HKTS2 any](
	mmap func(func(S1) S2) func(HKTS1) HKTS2,
	key func(B) func(S1) S2,
	f func(S1) B,
) func(HKTS1) HKTS2 {
	_ = "STUB: not implemented"
	return nil
}

func LetTo[S1, S2, B, HKTS1, HKTS2 any](
	mmap func(func(S1) S2) func(HKTS1) HKTS2,
	key func(B) func(S1) S2,
	b B,
) func(HKTS1) HKTS2 {
	_ = "STUB: not implemented"
	return nil
}
