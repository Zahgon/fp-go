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

package ioeither

import (
	"time"

	ET "github.com/IBM/fp-go/either"
	I "github.com/IBM/fp-go/io"
	IOO "github.com/IBM/fp-go/iooption"
	L "github.com/IBM/fp-go/lazy"
	O "github.com/IBM/fp-go/option"
)

// IOEither represents a synchronous computation that may fail
// refer to [https://andywhite.xyz/posts/2021-01-27-rte-foundations/#ioeitherlte-agt] for more details
type IOEither[E, A any] I.IO[ET.Either[E, A]]

func MakeIO[E, A any](f IOEither[E, A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func Left[A, E any](l E) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func Right[E, A any](r A) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func Of[E, A any](r A) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func MonadOf[E, A any](r A) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func LeftIO[A, E any](ml I.IO[E]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func RightIO[E, A any](mr I.IO[A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func FromEither[E, A any](e ET.Either[E, A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func FromOption[A, E any](onNone func() E) func(o O.Option[A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromIOOption[A, E any](onNone func() E) func(o IOO.IOOption[A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[A, B, E any](onNone func() E) func(func(A) O.Option[B]) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOK[E, A, B any](ma IOEither[E, A], f func(A) I.IO[B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOK[E, A, B any](f func(A) I.IO[B]) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainLazyK[E, A, B any](f func(A) L.Lazy[B]) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// FromIO creates an [IOEither] from an [IO] instance, invoking [IO] for each invocation of [IOEither]
func FromIO[E, A any](mr I.IO[A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// FromLazy creates an [IOEither] from a [Lazy] instance, invoking [Lazy] for each invocation of [IOEither]
func FromLazy[E, A any](mr L.Lazy[A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func MonadMap[E, A, B any](fa IOEither[E, A], f func(A) B) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[E, A, B any](f func(A) B) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[E, A, B any](fa IOEither[E, A], b B) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MapTo[E, A, B any](b B) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[E, A, B any](fa IOEither[E, A], f func(A) IOEither[E, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[E, A, B any](f func(A) IOEither[E, B]) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[E, A, B any](ma IOEither[E, A], f func(A) ET.Either[E, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[E, A, B any](f func(A) ET.Either[E, B]) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[B, E, A any](mab IOEither[E, func(A) B], ma IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Ap is an alias of [ApPar]
func Ap[B, E, A any](ma IOEither[E, A]) func(IOEither[E, func(A) B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadApPar[B, E, A any](mab IOEither[E, func(A) B], ma IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApPar applies function and value in parallel
func ApPar[B, E, A any](ma IOEither[E, A]) func(IOEither[E, func(A) B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadApSeq[B, E, A any](mab IOEither[E, func(A) B], ma IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApSeq applies function and value sequentially
func ApSeq[B, E, A any](ma IOEither[E, A]) func(IOEither[E, func(A) B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[E, A any](mma IOEither[E, IOEither[E, A]]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func TryCatch[E, A any](f func() (A, error), onThrow func(error) E) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func TryCatchError[A any](f func() (A, error)) IOEither[error, A] {
	_ = "STUB: not implemented"
	return nil
}

func Memoize[E, A any](ma IOEither[E, A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func MonadMapLeft[E1, E2, A any](fa IOEither[E1, A], f func(E1) E2) IOEither[E2, A] {
	_ = "STUB: not implemented"
	return nil
}

func MapLeft[A, E1, E2 any](f func(E1) E2) func(IOEither[E1, A]) IOEither[E2, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadBiMap[E1, E2, A, B any](fa IOEither[E1, A], f func(E1) E2, g func(A) B) IOEither[E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
func BiMap[E1, E2, A, B any](f func(E1) E2, g func(A) B) func(IOEither[E1, A]) IOEither[E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// Fold converts an IOEither into an IO
func Fold[E, A, B any](onLeft func(E) I.IO[B], onRight func(A) I.IO[B]) func(IOEither[E, A]) I.IO[B] {
	_ = "STUB: not implemented"
	return nil
}

// GetOrElse extracts the value or maps the error
func GetOrElse[E, A any](onLeft func(E) I.IO[A]) func(IOEither[E, A]) I.IO[A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainTo composes to the second monad ignoring the return value of the first
func MonadChainTo[A, E, B any](fa IOEither[E, A], fb IOEither[E, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainTo composes to the second [IOEither] monad ignoring the return value of the first
func ChainTo[A, E, B any](fb IOEither[E, B]) func(IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirst runs the [IOEither] monad returned by the function but returns the result of the original monad
func MonadChainFirst[E, A, B any](ma IOEither[E, A], f func(A) IOEither[E, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirst runs the [IOEither] monad returned by the function but returns the result of the original monad
func ChainFirst[E, A, B any](f func(A) IOEither[E, B]) func(IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstEitherK[A, E, B any](ma IOEither[E, A], f func(A) ET.Either[E, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirstEitherK[A, E, B any](f func(A) ET.Either[E, B]) func(ma IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK runs [IO] the monad returned by the function but returns the result of the original monad
func MonadChainFirstIOK[E, A, B any](ma IOEither[E, A], f func(A) I.IO[B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOK runs the [IO] monad returned by the function but returns the result of the original monad
func ChainFirstIOK[E, A, B any](f func(A) I.IO[B]) func(IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// WithResource constructs a function that creates a resource, then operates on it and then releases the resource
func WithResource[A, E, R, ANY any](onCreate IOEither[E, R], onRelease func(R) IOEither[E, ANY]) func(func(R) IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
func Swap[E, A any](val IOEither[E, A]) IOEither[A, E] { _ = "STUB: not implemented"; return nil }

// FromImpure converts a side effect without a return value into a side effect that returns any
func FromImpure[E any](f func()) IOEither[E, any] { _ = "STUB: not implemented"; return nil }

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[E, A any](gen L.Lazy[IOEither[E, A]]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAlt identifies an associative operation on a type constructor
func MonadAlt[E, A any](first IOEither[E, A], second L.Lazy[IOEither[E, A]]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Alt identifies an associative operation on a type constructor
func Alt[E, A any](second L.Lazy[IOEither[E, A]]) func(IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil

	// OrElse returns the original IOEither if it is a Right, otherwise it applies the given function to the error and returns the result.
}

func OrElse[E, A any](onLeft func(E) IOEither[E, A]) func(IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[E, B, A any](fab IOEither[E, func(A) B], a A) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[E, B, A any](a A) func(IOEither[E, func(A) B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ToIOOption converts an [IOEither] to an [IOO.IOOption]
func ToIOOption[E, A any](ioe IOEither[E, A]) IOO.IOOption[A] {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some delay
func Delay[E, A any](delay time.Duration) func(IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// After creates an operation that passes after the given [time.Time]
func After[E, A any](timestamp time.Time) func(IOEither[E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}
