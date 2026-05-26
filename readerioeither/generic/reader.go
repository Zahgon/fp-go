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

import (
	ET "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
)

// MakeReader constructs an instance of a reader
func MakeReader[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](f func(R) GIOA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadAlt[LAZY ~func() GEA, GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](first GEA, second LAZY) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Alt[LAZY ~func() GEA, GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](second LAZY) func(GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](fa GEA, f func(A) B) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Map[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](f func(A) B) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](fa GEA, b B) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func MapTo[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](b B) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](fa GEA, f func(A) GEB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Chain[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](f func(A) GEB) func(fa GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](fa GEA, f func(A) GEB) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func ChainFirst[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](f func(A) GEB) func(fa GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](ma GEA, f func(A) ET.Either[E, B]) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](f func(A) ET.Either[E, B]) func(ma GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstEitherK[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A, B any](ma GEA, f func(A) ET.Either[E, B]) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func ChainFirstEitherK[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A, B any](f func(A) ET.Either[E, B]) func(ma GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstIOK[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GIO ~func() B, R, E, A, B any](ma GEA, f func(A) GIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func ChainFirstIOK[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GIO ~func() B, R, E, A, B any](f func(A) GIO) func(GEA) GEA {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainReaderK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], GB ~func(R) B, R, E, A, B any](ma GEA, f func(A) GB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainReaderK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], GB ~func(R) B, R, E, A, B any](f func(A) GB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainReaderIOK[GEA ~func(R) GIOEA, GEB ~func(R) GIOEB, GIOEA ~func() ET.Either[E, A], GIOEB ~func() ET.Either[E, B], GIOB ~func() B, GB ~func(R) GIOB, R, E, A, B any](ma GEA, f func(A) GB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainReaderIOK[GEA ~func(R) GIOEA, GEB ~func(R) GIOEB, GIOEA ~func() ET.Either[E, A], GIOEB ~func() ET.Either[E, B], GIOB ~func() B, GB ~func(R) GIOB, R, E, A, B any](f func(A) GB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](ma GEA, f func(A) GIOB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainIOEitherK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](f func(A) GIOB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], GIO ~func() B, R, E, A, B any](ma GEA, f func(A) GIO) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainIOK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], GIO ~func() B, R, E, A, B any](f func(A) GIO) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[GEA ~func(R) GIOA, GEB ~func(R) GIOB, GIOA ~func() ET.Either[E, A], GIOB ~func() ET.Either[E, B], R, E, A, B any](onNone func() E) func(func(A) O.Option[B]) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() ET.Either[E, A],
	GIOB ~func() ET.Either[E, B],
	GIOFAB ~func() ET.Either[E, func(A) B],
	R, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Ap[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() ET.Either[E, A],
	GIOB ~func() ET.Either[E, B],
	GIOFAB ~func() ET.Either[E, func(A) B],
	R, E, A, B any](fa GEA) func(fab GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadApSeq[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() ET.Either[E, A],
	GIOB ~func() ET.Either[E, B],
	GIOFAB ~func() ET.Either[E, func(A) B],
	R, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ApSeq[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() ET.Either[E, A],
	GIOB ~func() ET.Either[E, B],
	GIOFAB ~func() ET.Either[E, func(A) B],
	R, E, A, B any](fa GEA) func(fab GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadApPar[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() ET.Either[E, A],
	GIOB ~func() ET.Either[E, B],
	GIOFAB ~func() ET.Either[E, func(A) B],
	R, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ApPar[
	GEA ~func(R) GIOA,
	GEB ~func(R) GIOB,
	GEFAB ~func(R) GIOFAB,
	GIOA ~func() ET.Either[E, A],
	GIOB ~func() ET.Either[E, B],
	GIOFAB ~func() ET.Either[E, func(A) B],
	R, E, A, B any](fa GEA) func(fab GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func Right[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](a A) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Left[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](e E) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func ThrowError[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](e E) GEA {
	_ = "STUB: not implemented"
	return *

	// Of returns a Reader with a fixed value
	new(GEA)
}

func Of[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](a A) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Flatten[GEA ~func(R) GIOA, GGEA ~func(R) GIOEA, GIOA ~func() ET.Either[E, A], GIOEA ~func() ET.Either[E, GEA], R, E, A any](mma GGEA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromIOEither[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](t GIOA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromEither[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](t ET.Either[E, A]) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func RightReader[GA ~func(R) A, GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](ma GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func LeftReader[GE ~func(R) E, GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](ma GE) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromReader[GA ~func(R) A, GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](ma GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadFromReaderIO[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GRIO ~func(R) GIO, GIO ~func() A, R, E, A any](a A, f func(A) GRIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromReaderIO[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GRIO ~func(R) GIO, GIO ~func() A, R, E, A any](f func(A) GRIO) func(A) GEA {
	_ = "STUB: not implemented"
	return nil
}

func RightReaderIO[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GRIO ~func(R) GIO, GIO ~func() A, R, E, A any](ma GRIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func LeftReaderIO[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GRIO ~func(R) GIO, GIO ~func() E, R, E, A any](me GRIO) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func RightIO[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GR ~func() A, R, E, A any](ma GR) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func LeftIO[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GR ~func() E, R, E, A any](ma GR) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromIO[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], GR ~func() A, R, E, A any](ma GR) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromReaderEither[GA ~func(R) ET.Either[E, A], GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](ma GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Ask[GER ~func(R) GIOR, GIOR ~func() ET.Either[E, R], R, E any]() GER {
	_ = "STUB: not implemented"
	return *new(GER)
}

func Asks[GA ~func(R) A, GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromOption[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](onNone func() E) func(O.Option[A]) GEA {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](pred func(A) bool, onFalse func(A) E) func(A) GEA {
	_ = "STUB: not implemented"
	return nil
}

func Fold[GB ~func(R) GIOB, GEA ~func(R) GIOA, GIOB ~func() B, GIOA ~func() ET.Either[E, A], R, E, A, B any](onLeft func(E) GB, onRight func(A) GB) func(GEA) GB {
	_ = "STUB: not implemented"
	return nil
}

func GetOrElse[GA ~func(R) GIOB, GEA ~func(R) GIOA, GIOB ~func() A, GIOA ~func() ET.Either[E, A], R, E, A any](onLeft func(E) GA) func(GEA) GA {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[GEA1 ~func(R) GIOA1, GEA2 ~func(R) GIOA2, GIOA1 ~func() ET.Either[E1, A], GIOA2 ~func() ET.Either[E2, A], R, E1, A, E2 any](onLeft func(E1) GEA2) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}

func OrLeft[GEA1 ~func(R) GIOA1, GE2 ~func(R) GIOE2, GEA2 ~func(R) GIOA2, GIOA1 ~func() ET.Either[E1, A], GIOE2 ~func() E2, GIOA2 ~func() ET.Either[E2, A], E1, R, E2, A any](onLeft func(E1) GE2) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}

func MonadBiMap[GA ~func(R) GE1A, GB ~func(R) GE2B, GE1A ~func() ET.Either[E1, A], GE2B ~func() ET.Either[E2, B], R, E1, E2, A, B any](fa GA, f func(E1) E2, g func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
func BiMap[GA ~func(R) GE1A, GB ~func(R) GE2B, GE1A ~func() ET.Either[E1, A], GE2B ~func() ET.Either[E2, B], R, E1, E2, A, B any](f func(E1) E2, g func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
func Swap[GREA ~func(R) GEA, GRAE ~func(R) GAE, GEA ~func() ET.Either[E, A], GAE ~func() ET.Either[A, E], R, E, A any](val GREA) GRAE {
	_ = "STUB: not implemented"
	return *new(GRAE)
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[GEA ~func(R) GA, GA ~func() ET.Either[E, A], R, E, A any](gen func() GEA) GEA {
	_ = "STUB: not implemented"
	return *

	// TryCatch wraps a reader returning a tuple as an error into ReaderIOEither
	new(GEA)
}

func TryCatch[GEA ~func(R) GA, GA ~func() ET.Either[E, A], R, E, A any](f func(R) func() (A, error), onThrow func(error) E) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Memoize computes the value of the provided monad lazily but exactly once
// The context used to compute the value is the context of the first call, so do not use this
// method if the value has a functional dependency on the content of the context
func Memoize[
	GEA ~func(R) GIOA, GIOA ~func() ET.Either[E, A], R, E, A any](rdr GEA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadFlap[GREAB ~func(R) GEAB, GREB ~func(R) GEB, GEAB ~func() ET.Either[E, func(A) B], GEB ~func() ET.Either[E, B], R, E, B, A any](fab GREAB, a A) GREB {
	_ = "STUB: not implemented"
	return *new(GREB)
}

func Flap[GREAB ~func(R) GEAB, GREB ~func(R) GEB, GEAB ~func() ET.Either[E, func(A) B], GEB ~func() ET.Either[E, B], R, E, B, A any](a A) func(GREAB) GREB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapLeft[GREA1 ~func(R) GEA1, GREA2 ~func(R) GEA2, GEA1 ~func() ET.Either[E1, A], GEA2 ~func() ET.Either[E2, A], R, E1, E2, A any](fa GREA1, f func(E1) E2) GREA2 {
	_ = "STUB: not implemented"
	return *new(GREA2)
}

// MapLeft applies a mapping function to the error channel
func MapLeft[GREA1 ~func(R) GEA1, GREA2 ~func(R) GEA2, GEA1 ~func() ET.Either[E1, A], GEA2 ~func() ET.Either[E2, A], R, E1, E2, A any](f func(E1) E2) func(GREA1) GREA2 {
	_ = "STUB: not implemented"
	return nil
}

// Local changes the value of the local context during the execution of the action `ma` (similar to `Contravariant`'s
// `contramap`).
func Local[
	GEA1 ~func(R1) GIOA,
	GEA2 ~func(R2) GIOA,

	GIOA ~func() ET.Either[E, A],
	R1, R2, E, A any,
](f func(R2) R1) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}
