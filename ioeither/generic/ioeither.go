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
	"time"

	ET "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
)

// type IOEither[E, A any] = func() Either[E, A]

func MakeIO[GA ~func() ET.Either[E, A], E, A any](f GA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func Left[GA ~func() ET.Either[E, A], E, A any](l E) GA { _ = "STUB: not implemented"; return *new(GA) }

func Right[GA ~func() ET.Either[E, A], E, A any](r A) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func Of[GA ~func() ET.Either[E, A], E, A any](r A) GA { _ = "STUB: not implemented"; return *new(GA) }

func MonadOf[GA ~func() ET.Either[E, A], E, A any](r A) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func LeftIO[GA ~func() ET.Either[E, A], GE ~func() E, E, A any](ml GE) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func RightIO[GA ~func() ET.Either[E, A], GR ~func() A, E, A any](mr GR) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func FromEither[GA ~func() ET.Either[E, A], E, A any](e ET.Either[E, A]) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func FromOption[GA ~func() ET.Either[E, A], E, A any](onNone func() E) func(o O.Option[A]) GA {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](onNone func() E) func(func(A) O.Option[B]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func FromIO[GA ~func() ET.Either[E, A], GR ~func() A, E, A any](mr GR) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func MonadMap[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](fa GA, f func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Map[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](f func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](fa GA, b B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func MapTo[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](b B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](fa GA, f func(A) GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Chain[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](f func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainTo[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](fa GA, fb GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ChainTo[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](fb GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](ma GA, f func(A) ET.Either[E, B]) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func MonadChainIOK[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], GR ~func() B, E, A, B any](ma GA, f func(A) GR) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ChainIOK[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], GR ~func() B, E, A, B any](f func(A) GR) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](f func(A) ET.Either[E, B]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[GB ~func() ET.Either[E, B], GAB ~func() ET.Either[E, func(A) B], GA ~func() ET.Either[E, A], E, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Ap[GB ~func() ET.Either[E, B], GAB ~func() ET.Either[E, func(A) B], GA ~func() ET.Either[E, A], E, A, B any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadApSeq[GB ~func() ET.Either[E, B], GAB ~func() ET.Either[E, func(A) B], GA ~func() ET.Either[E, A], E, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ApSeq[GB ~func() ET.Either[E, B], GAB ~func() ET.Either[E, func(A) B], GA ~func() ET.Either[E, A], E, A, B any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadApPar[GB ~func() ET.Either[E, B], GAB ~func() ET.Either[E, func(A) B], GA ~func() ET.Either[E, A], E, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ApPar[GB ~func() ET.Either[E, B], GAB ~func() ET.Either[E, func(A) B], GA ~func() ET.Either[E, A], E, A, B any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GA ~func() ET.Either[E, A], GAA ~func() ET.Either[E, GA], E, A any](mma GAA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func TryCatch[GA ~func() ET.Either[E, A], E, A any](f func() (A, error), onThrow func(error) E) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func TryCatchError[GA ~func() ET.Either[error, A], A any](f func() (A, error)) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Memoize computes the value of the provided IO monad lazily but exactly once
func Memoize[GA ~func() ET.Either[E, A], E, A any](ma GA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func MonadMapLeft[GA1 ~func() ET.Either[E1, A], GA2 ~func() ET.Either[E2, A], E1, E2, A any](fa GA1, f func(E1) E2) GA2 {
	_ = "STUB: not implemented"
	return *new(GA2)
}

func MapLeft[GA1 ~func() ET.Either[E1, A], GA2 ~func() ET.Either[E2, A], E1, E2, A any](f func(E1) E2) func(GA1) GA2 {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some [time.Duration]
func Delay[GA ~func() ET.Either[E, A], E, A any](delay time.Duration) func(GA) GA {
	_ = "STUB: not implemented"
	return nil

	// After creates an operation that passes after the given [time.Time]
}

func After[GA ~func() ET.Either[E, A], E, A any](timestamp time.Time) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

func MonadBiMap[GA ~func() ET.Either[E1, A], GB ~func() ET.Either[E2, B], E1, E2, A, B any](fa GA, f func(E1) E2, g func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
func BiMap[GA ~func() ET.Either[E1, A], GB ~func() ET.Either[E2, B], E1, E2, A, B any](f func(E1) E2, g func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Fold convers an IOEither into an IO
func Fold[GA ~func() ET.Either[E, A], GB ~func() B, E, A, B any](onLeft func(E) GB, onRight func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadFold[GA ~func() ET.Either[E, A], GB ~func() B, E, A, B any](ma GA, onLeft func(E) GB, onRight func(A) GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// GetOrElse extracts the value or maps the error
func GetOrElse[GA ~func() ET.Either[E, A], GB ~func() A, E, A any](onLeft func(E) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirst runs the monad returned by the function but returns the result of the original monad
func MonadChainFirst[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](ma GA, f func(A) GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirst runs the monad returned by the function but returns the result of the original monad
func ChainFirst[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], E, A, B any](f func(A) GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK runs the monad returned by the function but returns the result of the original monad
func MonadChainFirstIOK[GA ~func() ET.Either[E, A], GIOB ~func() B, E, A, B any](first GA, f func(A) GIOB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirstIOK runs the monad returned by the function but returns the result of the original monad
func ChainFirstIOK[GA ~func() ET.Either[E, A], GIOB ~func() B, E, A, B any](f func(A) GIOB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstEitherK runs the monad returned by the function but returns the result of the original monad
func MonadChainFirstEitherK[GA ~func() ET.Either[E, A], E, A, B any](first GA, f func(A) ET.Either[E, B]) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirstEitherK runs the monad returned by the function but returns the result of the original monad
func ChainFirstEitherK[GA ~func() ET.Either[E, A], E, A, B any](f func(A) ET.Either[E, B]) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
func Swap[GEA ~func() ET.Either[E, A], GAE ~func() ET.Either[A, E], E, A any](val GEA) GAE {
	_ = "STUB: not implemented"
	return *new(GAE)
}

// FromImpure converts a side effect without a return value into a side effect that returns any
func FromImpure[GA ~func() ET.Either[E, any], IMP ~func(), E any](f IMP) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[GEA ~func() ET.Either[E, A], E, A any](gen func() GEA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadAlt[LAZY ~func() GIOA, GIOA ~func() ET.Either[E, A], E, A any](first GIOA, second LAZY) GIOA {
	_ = "STUB: not implemented"
	return *new(GIOA)
}

func Alt[LAZY ~func() GIOA, GIOA ~func() ET.Either[E, A], E, A any](second LAZY) func(GIOA) GIOA {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[GEAB ~func() ET.Either[E, func(A) B], GEB ~func() ET.Either[E, B], E, B, A any](fab GEAB, a A) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Flap[GEAB ~func() ET.Either[E, func(A) B], GEB ~func() ET.Either[E, B], E, B, A any](a A) func(GEAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func ToIOOption[GA ~func() O.Option[A], GEA ~func() ET.Either[E, A], E, A any](ioe GEA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// OrElse returns the original IOEither if it is a Right, otherwise it applies the given function to the error and returns the result.
func OrElse[GA ~func() ET.Either[E, A], E, A any](onLeft func(E) GA) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

func FromIOOption[GEA ~func() ET.Either[E, A], GA ~func() O.Option[A], E, A any](onNone func() E) func(ioo GA) GEA {
	_ = "STUB: not implemented"
	return nil
}
