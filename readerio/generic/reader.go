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

package generic

func FromIO[GEA ~func(E) GIOA, GIOA ~func() A, E, A any](t GIOA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromReader[GA ~func(E) A, GEA ~func(E) GIOA, GIOA ~func() A, E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadMap[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](fa GEA, f func(A) B) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Map[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](f func(A) B) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](ma GEA, f func(A) GEB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Chain[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](f func(A) GEB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func Of[GEA ~func(E) GIOA, GIOA ~func() A, E, A any](a A) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadAp[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Ap[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fa GEA) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadApSeq[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ApSeq[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fa GEA) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadApPar[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ApPar[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GEFAB ~func(E) GIOFAB, GIOA ~func() A, GIOB ~func() B, GIOFAB ~func() func(A) B, E, A, B any](fa GEA) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func Ask[GEE ~func(E) GIOE, GIOE ~func() E, E any]() GEE {
	_ = "STUB: not implemented"
	return *new(GEE)
}

func Asks[GA ~func(E) A, GEA ~func(E) GIOA, GIOA ~func() A, E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadChainIOK[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](ma GEA, f func(A) GIOB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainIOK[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](f func(A) GIOB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstIOK[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](ma GEA, f func(A) GIOB) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func ChainFirstIOK[GEA ~func(E) GIOA, GEB ~func(E) GIOB, GIOA ~func() A, GIOB ~func() B, E, A, B any](f func(A) GIOB) func(GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[GEA ~func(E) GA, GA ~func() A, E, A any](gen func() GEA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Memoize computes the value of the provided reader monad lazily but exactly once
// The context used to compute the value is the context of the first call, so do not use this
// method if the value has a functional dependency on the content of the context
func Memoize[GEA ~func(E) GA, GA ~func() A, E, A any](rdr GEA) GEA {
	_ = "STUB: not implemented"
	// synchronization primitives
	return *new(GEA)
}

// callback

// returns our memoized wrapper

func Flatten[GEA ~func(R) GIOA, GGEA ~func(R) GIOEA, GIOA ~func() A, GIOEA ~func() GEA, R, A any](mma GGEA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadFlap[GEFAB ~func(E) GIOFAB, GEB ~func(E) GIOB, GIOFAB ~func() func(A) B, GIOB ~func() B, E, A, B any](fab GEFAB, a A) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Flap[GEFAB ~func(E) GIOFAB, GEB ~func(E) GIOB, GIOFAB ~func() func(A) B, GIOB ~func() B, E, A, B any](a A) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}
