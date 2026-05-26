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

package generic

import (
	"github.com/IBM/fp-go/v2/either"
	O "github.com/IBM/fp-go/v2/option"
)

// MakeReader constructs an instance of a reader
// Deprecated:
func MakeReader[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](f func(R) GIOA) GEA {
	_ = "STUB: not implemented"

	// Deprecated:
	return *new(GEA)
}

func MonadAlt[LAZY ~func() GEA, GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](first GEA, second LAZY) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func Alt[LAZY ~func() GEA, GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](second LAZY) func(GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadMap[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](fa GEA, f func(A) B) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func Map[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](f func(A) B) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadMapTo[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](fa GEA, b B) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func MapTo[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](b B) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChain[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](fa GEA, f func(A) GEB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func Chain[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](f func(A) GEB) func(fa GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainFirst[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](fa GEA, f func(A) GEB) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func ChainFirst[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](f func(A) GEB) func(fa GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](ma GEA, f func(A) either.Either[E, B]) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func ChainEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](f func(A) either.Either[E, B]) func(ma GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainFirstEitherK[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A, B any](ma GEA, f func(A) either.Either[E, B]) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func ChainFirstEitherK[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A, B any](f func(A) either.Either[E, B]) func(ma GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainFirstIOK[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GIO ~func() B, R, E, A, B any](ma GEA, f func(A) GIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func ChainFirstIOK[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GIO ~func() B, R, E, A, B any](f func(A) GIO) func(GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainReaderK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], GB ~func(R) B, R, E, A, B any](ma GEA, f func(A) GB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func ChainReaderK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], GB ~func(R) B, R, E, A, B any](f func(A) GB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainReaderIOK[GEA ~func(R) GIOEA, GEB ~func(R) GIOEB, GIOEA ~func() either.Either[E, A], GIOEB ~func() either.Either[E, B], GIOB ~func() B, GB ~func(R) GIOB, R, E, A, B any](ma GEA, f func(A) GB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func ChainReaderIOK[GEA ~func(R) GIOEA, GEB ~func(R) GIOEB, GIOEA ~func() either.Either[E, A], GIOEB ~func() either.Either[E, B], GIOB ~func() B, GB ~func(R) GIOB, R, E, A, B any](f func(A) GB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainIOEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](ma GEA, f func(A) GIOB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func ChainIOEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](f func(A) GIOB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainIOK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], GIO ~func() B, R, E, A, B any](ma GEA, f func(A) GIO) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func ChainIOK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], GIO ~func() B, R, E, A, B any](f func(A) GIO) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func ChainOptionK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() either.Either[E, A], GIOB ~func() either.Either[E, B], R, E, A, B any](onNone func() E) func(func(A) O.Option[B]) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadAp[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() either.Either[E, A],
	GIOB ~func() either.Either[E, B],
	GIOFAB ~func() either.Either[E, func(A) B],
	R, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func Ap[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() either.Either[E, A],
	GIOB ~func() either.Either[E, B],
	GIOFAB ~func() either.Either[E, func(A) B],
	R, E, A, B any](fa GEA) func(fab GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadApSeq[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() either.Either[E, A],
	GIOB ~func() either.Either[E, B],
	GIOFAB ~func() either.Either[E, func(A) B],
	R, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func ApSeq[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() either.Either[E, A],
	GIOB ~func() either.Either[E, B],
	GIOFAB ~func() either.Either[E, func(A) B],
	R, E, A, B any](fa GEA) func(fab GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadApPar[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() either.Either[E, A],
	GIOB ~func() either.Either[E, B],
	GIOFAB ~func() either.Either[E, func(A) B],
	R, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func ApPar[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() either.Either[E, A],
	GIOB ~func() either.Either[E, B],
	GIOFAB ~func() either.Either[E, func(A) B],
	R, E, A, B any](fa GEA) func(fab GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func Right[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](a A) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func Left[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](e E) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func ThrowError[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](e E) GEA {
	_ = "STUB: not implemented"
	return *

	// Of returns a Reader with a fixed value
	// Deprecated:
	new(GEA)
}

func Of[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](a A) GEA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GEA)
}

func Flatten[GEA ~func(R) GIOA, GGEA ~func(R) GIOEA, GIOA ~func() either.Either[E, A], GIOEA ~func() either.Either[E, GEA], R, E, A any](mma GGEA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func FromIOEither[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](t GIOA) GEA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GEA)
}

func FromEither[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](t either.Either[E, A]) GEA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GEA)
}

func RightReader[GA ~func(R) A, GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](ma GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func LeftReader[GE ~func(R) E, GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](ma GE) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func FromReader[GA ~func(R) A, GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](ma GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func MonadFromReaderIO[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GRIO ~func(R) GIO, GIO ~func() A, R, E, A any](a A, f func(A) GRIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func FromReaderIO[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GRIO ~func(R) GIO, GIO ~func() A, R, E, A any](f func(A) GRIO) func(A) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func RightReaderIO[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GRIO ~func(R) GIO, GIO ~func() A, R, E, A any](ma GRIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func LeftReaderIO[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GRIO ~func(R) GIO, GIO ~func() E, R, E, A any](me GRIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func RightIO[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GR ~func() A, R, E, A any](ma GR) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func LeftIO[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GR ~func() E, R, E, A any](ma GR) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func FromIO[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], GR ~func() A, R, E, A any](ma GR) GEA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GEA)
}

func FromReaderEither[GA ~func(R) either.Either[E, A], GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](ma GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func Ask[GER ~func(R) GIOR, GIOR ~func() either.Either[E, R], R, E any]() GER {
	_ = "STUB: not implemented"
	return *new(GER)
}

// Deprecated:
func Asks[GA ~func(R) A, GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Deprecated:
func FromOption[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](onNone func() E) func(O.Option[A]) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func FromPredicate[GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](pred func(A) bool, onFalse func(A) E) func(A) GEA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func Fold[GB ~func(R) GIOB, GEA ~func(R) GIOA, GIOB ~func() B, GIOA ~func() either.Either[E, A], R, E, A, B any](onLeft func(E) GB, onRight func(A) GB) func(GEA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func GetOrElse[GA ~func(R) GIOB, GEA ~func(R) GIOA, GIOB ~func() A, GIOA ~func() either.Either[E, A], R, E, A any](onLeft func(E) GA) func(GEA) GA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func OrElse[GEA1 ~func(R) GIOA1, GEA2 ~func(R) GIOA2, GIOA1 ~func() either.Either[E1, A], GIOA2 ~func() either.Either[E2, A], R, E1, A, E2 any](onLeft func(E1) GEA2) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func OrLeft[GEA1 ~func(R) GIOA1, GE2 ~func(R) GIOE2, GEA2 ~func(R) GIOA2, GIOA1 ~func() either.Either[E1, A], GIOE2 ~func() E2, GIOA2 ~func() either.Either[E2, A], E1, R, E2, A any](onLeft func(E1) GE2) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadBiMap[GA ~func(R) GE1A, GB ~func(R) GE2B, GE1A ~func() either.Either[E1, A], GE2B ~func() either.Either[E2, B], R, E1, E2, A, B any](fa GA, f func(E1) E2, g func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
// Deprecated:
func BiMap[GA ~func(R) GE1A, GB ~func(R) GE2B, GE1A ~func() either.Either[E1, A], GE2B ~func() either.Either[E2, B], R, E1, E2, A, B any](f func(E1) E2, g func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
// Deprecated:
func Swap[GREA ~func(R) GEA, GRAE ~func(R) GAE, GEA ~func() either.Either[E, A], GAE ~func() either.Either[A, E], R, E, A any](val GREA) GRAE {
	_ = "STUB: not implemented"
	return *new(GRAE)
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
// Deprecated:
func Defer[GEA ~func(R) GA, GA ~func() either.Either[E, A], R, E, A any](gen func() GEA) GEA {
	_ = "STUB: not implemented"
	return *

	// TryCatch wraps a reader returning a tuple as an error into ReaderIOEither
	// Deprecated:
	new(GEA)
}

func TryCatch[GEA ~func(R) GA, GA ~func() either.Either[E, A], R, E, A any](f func(R) func() (A, error), onThrow func(error) E) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Memoize computes the value of the provided monad lazily but exactly once
// The context used to compute the value is the context of the first call, so do not use this
// method if the value has a functional dependency on the content of the context
// Deprecated:
func Memoize[
	GEA ~func(R) GIOA, GIOA ~func() either.Either[E, A], R, E, A any](rdr GEA) GEA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GEA)
}

func MonadFlap[GREAB ~func(R) GEAB, GREB ~func(R) GEB, GEAB ~func() either.Either[E, func(A) B], GEB ~func() either.Either[E, B], R, E, B, A any](fab GREAB, a A) GREB {
	_ = "STUB: not implemented"
	return *new(GREB)
}

// Deprecated:
func Flap[GREAB ~func(R) GEAB, GREB ~func(R) GEB, GEAB ~func() either.Either[E, func(A) B], GEB ~func() either.Either[E, B], R, E, B, A any](a A) func(GREAB) GREB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadMapLeft[GREA1 ~func(R) GEA1, GREA2 ~func(R) GEA2, GEA1 ~func() either.Either[E1, A], GEA2 ~func() either.Either[E2, A], R, E1, E2, A any](fa GREA1, f func(E1) E2) GREA2 {
	_ = "STUB: not implemented"
	return *new(GREA2)
}

// MapLeft applies a mapping function to the error channel
// Deprecated:
func MapLeft[GREA1 ~func(R) GEA1, GREA2 ~func(R) GEA2, GEA1 ~func() either.Either[E1, A], GEA2 ~func() either.Either[E2, A], R, E1, E2, A any](f func(E1) E2) func(GREA1) GREA2 {
	_ = "STUB: not implemented"
	return nil
}

// Local changes the value of the local context during the execution of the action `ma` (similar to `Contravariant`'s
// `contramap`).
// Deprecated:
func Local[
	GEA1 ~func(R1) GIOA,
	GEA2 ~func(R2) GIOA,

	GIOA ~func() either.Either[E, A],
	R1, R2, E, A any,
](f func(R2) R1) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}
