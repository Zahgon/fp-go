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
	"time"

	"github.com/IBM/fp-go/v2/either"
	O "github.com/IBM/fp-go/v2/option"
)

// type IOEither[E, A any] = func() Either[E, A]

// Deprecated:
func MakeIO[GA ~func() either.Either[E, A], E, A any](f GA) GA {
	_ = "STUB: not implemented"

	// Deprecated:
	return *new(GA)
}

func Left[GA ~func() either.Either[E, A], E, A any](l E) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func Right[GA ~func() either.Either[E, A], E, A any](r A) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func Of[GA ~func() either.Either[E, A], E, A any](r A) GA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GA)
}

func MonadOf[GA ~func() either.Either[E, A], E, A any](r A) GA {
	_ = "STUB: not implemented"

	// Deprecated:
	return *new(GA)
}

func LeftIO[GA ~func() either.Either[E, A], GE ~func() E, E, A any](ml GE) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func RightIO[GA ~func() either.Either[E, A], GR ~func() A, E, A any](mr GR) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func FromEither[GA ~func() either.Either[E, A], E, A any](e either.Either[E, A]) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func FromOption[GA ~func() either.Either[E, A], E, A any](onNone func() E) func(o O.Option[A]) GA {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](onNone func() E) func(func(A) O.Option[B]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func FromIO[GA ~func() either.Either[E, A], GR ~func() A, E, A any](mr GR) GA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GA)
}

func MonadMap[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](fa GA, f func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func Map[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](f func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadMapTo[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](fa GA, b B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func MapTo[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](b B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChain[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](fa GA, f func(A) GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func Chain[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](f func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainTo[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](fa GA, fb GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ChainTo[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](fb GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadChainEitherK[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](ma GA, f func(A) either.Either[E, B]) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func MonadChainIOK[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], GR ~func() B, E, A, B any](ma GA, f func(A) GR) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ChainIOK[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], GR ~func() B, E, A, B any](f func(A) GR) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func ChainEitherK[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](f func(A) either.Either[E, B]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadAp[GB ~func() either.Either[E, B], GAB ~func() either.Either[E, func(A) B], GA ~func() either.Either[E, A], E, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func Ap[GB ~func() either.Either[E, B], GAB ~func() either.Either[E, func(A) B], GA ~func() either.Either[E, A], E, A, B any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadApSeq[GB ~func() either.Either[E, B], GAB ~func() either.Either[E, func(A) B], GA ~func() either.Either[E, A], E, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func ApSeq[GB ~func() either.Either[E, B], GAB ~func() either.Either[E, func(A) B], GA ~func() either.Either[E, A], E, A, B any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadApPar[GB ~func() either.Either[E, B], GAB ~func() either.Either[E, func(A) B], GA ~func() either.Either[E, A], E, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func ApPar[GB ~func() either.Either[E, B], GAB ~func() either.Either[E, func(A) B], GA ~func() either.Either[E, A], E, A, B any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func Flatten[GA ~func() either.Either[E, A], GAA ~func() either.Either[E, GA], E, A any](mma GAA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func TryCatch[GA ~func() either.Either[E, A], E, A any](f func() (A, error), onThrow func(error) E) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func TryCatchError[GA ~func() either.Either[error, A], A any](f func() (A, error)) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Memoize computes the value of the provided IO monad lazily but exactly once
//
// Deprecated:
func Memoize[GA ~func() either.Either[E, A], E, A any](ma GA) GA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GA)
}

func MonadMapLeft[GA1 ~func() either.Either[E1, A], GA2 ~func() either.Either[E2, A], E1, E2, A any](fa GA1, f func(E1) E2) GA2 {
	_ = "STUB: not implemented"
	return *new(GA2)
}

// Deprecated:
func MapLeft[GA1 ~func() either.Either[E1, A], GA2 ~func() either.Either[E2, A], E1, E2, A any](f func(E1) E2) func(GA1) GA2 {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some [time.Duration]
//
// Deprecated:
func Delay[GA ~func() either.Either[E, A], E, A any](delay time.Duration) func(GA) GA {
	_ = "STUB: not implemented"
	return nil

	// After creates an operation that passes after the given [time.Time]
	//
	// Deprecated:
}

func After[GA ~func() either.Either[E, A], E, A any](timestamp time.Time) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadBiMap[GA ~func() either.Either[E1, A], GB ~func() either.Either[E2, B], E1, E2, A, B any](fa GA, f func(E1) E2, g func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
//
// Deprecated:
func BiMap[GA ~func() either.Either[E1, A], GB ~func() either.Either[E2, B], E1, E2, A, B any](f func(E1) E2, g func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Fold convers an IOEither into an IO
//
// Deprecated:
func Fold[GA ~func() either.Either[E, A], GB ~func() B, E, A, B any](onLeft func(E) GB, onRight func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadFold[GA ~func() either.Either[E, A], GB ~func() B, E, A, B any](ma GA, onLeft func(E) GB, onRight func(A) GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// GetOrElse extracts the value or maps the error
func GetOrElse[GA ~func() either.Either[E, A], GB ~func() A, E, A any](onLeft func(E) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirst runs the monad returned by the function but returns the result of the original monad
//
// Deprecated:
func MonadChainFirst[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](ma GA, f func(A) GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirst runs the monad returned by the function but returns the result of the original monad
//
// Deprecated:
func ChainFirst[GA ~func() either.Either[E, A], GB ~func() either.Either[E, B], E, A, B any](f func(A) GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK runs the monad returned by the function but returns the result of the original monad
//
// Deprecated:
func MonadChainFirstIOK[GA ~func() either.Either[E, A], GIOB ~func() B, E, A, B any](first GA, f func(A) GIOB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirstIOK runs the monad returned by the function but returns the result of the original monad
//
// Deprecated:
func ChainFirstIOK[GA ~func() either.Either[E, A], GIOB ~func() B, E, A, B any](f func(A) GIOB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstEitherK runs the monad returned by the function but returns the result of the original monad
//
// Deprecated:
func MonadChainFirstEitherK[GA ~func() either.Either[E, A], E, A, B any](first GA, f func(A) either.Either[E, B]) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirstEitherK runs the monad returned by the function but returns the result of the original monad
//
// Deprecated:
func ChainFirstEitherK[GA ~func() either.Either[E, A], E, A, B any](f func(A) either.Either[E, B]) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
//
// Deprecated:
func Swap[GEA ~func() either.Either[E, A], GAE ~func() either.Either[A, E], E, A any](val GEA) GAE {
	_ = "STUB: not implemented"
	return *new(GAE)
}

// FromImpure converts a side effect without a return value into a side effect that returns any
//
// Deprecated:
func FromImpure[GA ~func() either.Either[E, any], IMP ~func(), E any](f IMP) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
//
// Deprecated:
func Defer[GEA ~func() either.Either[E, A], E, A any](gen func() GEA) GEA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GEA)
}

func MonadAlt[LAZY ~func() GIOA, GIOA ~func() either.Either[E, A], E, A any](first GIOA, second LAZY) GIOA {
	_ = "STUB: not implemented"
	return *new(GIOA)
}

// Deprecated:
func Alt[LAZY ~func() GIOA, GIOA ~func() either.Either[E, A], E, A any](second LAZY) func(GIOA) GIOA {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadFlap[GEAB ~func() either.Either[E, func(A) B], GEB ~func() either.Either[E, B], E, B, A any](fab GEAB, a A) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

// Deprecated:
func Flap[GEAB ~func() either.Either[E, func(A) B], GEB ~func() either.Either[E, B], E, B, A any](a A) func(GEAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func ToIOOption[GA ~func() O.Option[A], GEA ~func() either.Either[E, A], E, A any](ioe GEA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func FromIOOption[GEA ~func() either.Either[E, A], GA ~func() O.Option[A], E, A any](onNone func() E) func(ioo GA) GEA {
	_ = "STUB: not implemented"
	return nil
}
