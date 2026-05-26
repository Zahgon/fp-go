// Copyright (c) 2023 - 2025 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache LicensVersion 2.0 (the "License");
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

package ioresult

import (
	"time"

	IOI "github.com/IBM/fp-go/v2/idiomatic/ioresult"
	RI "github.com/IBM/fp-go/v2/idiomatic/result"
	"github.com/IBM/fp-go/v2/io"
	"github.com/IBM/fp-go/v2/ioeither"
	IOO "github.com/IBM/fp-go/v2/iooption"
	O "github.com/IBM/fp-go/v2/option"
	"github.com/IBM/fp-go/v2/result"
)

func fromIOResultKleisliI[A, B any](f IOI.Kleisli[A, B]) Kleisli[A, B] {
	_ = "STUB: not implemented"
	return nil
}

func fromResultKleisliI[A, B any](f RI.Kleisli[A, B]) result.Kleisli[A, B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Left[A any](l error) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func Right[A any](r A) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func Of[A any](r A) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadOf[A any](r A) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func LeftIO[A any](ml IO[error]) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func RightIO[A any](mr IO[A]) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func FromEither[A any](e Result[A]) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func FromResult[A any](e Result[A]) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func FromEitherI[A any](a A, err error) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func FromResultI[A any](a A, err error) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func FromOption[A any](onNone func() error) func(o O.Option[A]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func FromIOOption[A any](onNone func() error) func(o IOO.IOOption[A]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainOptionK[A, B any](onNone func() error) func(O.Kleisli[A, B]) Operator[A, B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadChainIOK[A, B any](ma IOResult[A], f io.Kleisli[A, B]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainIOK[A, B any](f io.Kleisli[A, B]) Operator[A, B] { _ = "STUB: not implemented"; return nil }

//go:inline
func ChainLazyK[A, B any](f func(A) Lazy[B]) Operator[A, B] { _ = "STUB: not implemented"; return nil }

// FromIO creates an [IOResult] from an [IO] instancinvoking [IO] for each invocation of [IOResult]
//
//go:inline
func FromIO[A any](mr IO[A]) IOResult[A] { _ = "STUB: not implemented"; return nil }

// FromLazy creates an [IOResult] from a [Lazy] instancinvoking [Lazy] for each invocation of [IOResult]
//
//go:inline
func FromLazy[A any](mr Lazy[A]) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadMap[A, B any](fa IOResult[A], f func(A) B) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Map[A, B any](f func(A) B) Operator[A, B] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadMapTo[A, B any](fa IOResult[A], b B) IOResult[B] { _ = "STUB: not implemented"; return nil }

//go:inline
func MapTo[A, B any](b B) Operator[A, B] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadChain[A, B any](fa IOResult[A], f Kleisli[A, B]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Chain[A, B any](f Kleisli[A, B]) Operator[A, B] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadChainI[A, B any](fa IOResult[A], f IOI.Kleisli[A, B]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainI[A, B any](f IOI.Kleisli[A, B]) Operator[A, B] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadChainEitherK[A, B any](ma IOResult[A], f result.Kleisli[A, B]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadChainResultK[A, B any](ma IOResult[A], f result.Kleisli[A, B]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainEitherK[A, B any](f result.Kleisli[A, B]) Operator[A, B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainResultK[A, B any](f result.Kleisli[A, B]) Operator[A, B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadAp[B, A any](mab IOResult[func(A) B], ma IOResult[A]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

// Ap is an alias of [ApPar]
//
//go:inline
func Ap[B, A any](ma IOResult[A]) Operator[func(A) B, B] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadApPar[B, A any](mab IOResult[func(A) B], ma IOResult[A]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

// ApPar applies function and value in parallel
//
//go:inline
func ApPar[B, A any](ma IOResult[A]) Operator[func(A) B, B] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadApSeq[B, A any](mab IOResult[func(A) B], ma IOResult[A]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

// ApSeq applies function and value sequentially
//
//go:inline
func ApSeq[B, A any](ma IOResult[A]) func(IOResult[func(A) B]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Flatten[A any](mma IOResult[IOResult[A]]) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func TryCatch[A any](f func() (A, error), onThrow Endomorphism[error]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func TryCatchError[A any](f func() (A, error)) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func Memoize[A any](ma IOResult[A]) IOResult[A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadMapLeft[A, E any](fa IOResult[A], f func(error) E) ioeither.IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MapLeft[A, E any](f func(error) E) func(IOResult[A]) ioeither.IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadBiMap[E, A, B any](fa IOResult[A], f func(error) E, g func(A) B) ioeither.IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
//
//go:inline
func BiMap[E, A, B any](f func(error) E, g func(A) B) func(IOResult[A]) ioeither.IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Fold converts an IOResult into an IO
//
//go:inline
func Fold[A, B any](onLeft func(error) IO[B], onRight io.Kleisli[A, B]) func(IOResult[A]) IO[B] {
	_ = "STUB: not implemented"
	return nil
}

// GetOrElse extracts the value or maps the error
//
//go:inline
func GetOrElse[A any](onLeft func(error) IO[A]) func(IOResult[A]) IO[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func GetOrElseOf[A any](onLeft func(error) A) func(IOResult[A]) IO[A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainTo composes to the second monad ignoring the return value of the first
//
//go:inline
func MonadChainTo[A, B any](fa IOResult[A], fb IOResult[B]) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainTo composes to the second [IOResult] monad ignoring the return value of the first
//
//go:inline
func ChainTo[A, B any](fb IOResult[B]) Operator[A, B] { _ = "STUB: not implemented"; return nil }

// MonadChainFirst runs the [IOResult] monad returned by the function but returns the result of the original monad
//
//go:inline
func MonadChainFirst[A, B any](ma IOResult[A], f Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadTap[A, B any](ma IOResult[A], f Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirst runs the [IOResult] monad returned by the function but returns the result of the original monad
//
//go:inline
func ChainFirst[A, B any](f Kleisli[A, B]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func Tap[A, B any](f Kleisli[A, B]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadChainFirstEitherK[A, B any](ma IOResult[A], f result.Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadTapEitherK[A, B any](ma IOResult[A], f result.Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadChainFirstResultK[A, B any](ma IOResult[A], f result.Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadTapResultK[A, B any](ma IOResult[A], f result.Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainFirstEitherK[A, B any](f result.Kleisli[A, B]) Operator[A, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func TapEitherK[A, B any](f result.Kleisli[A, B]) Operator[A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK runs [IO] the monad returned by the function but returns the result of the original monad
//
//go:inline
func MonadChainFirstIOK[A, B any](ma IOResult[A], f io.Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadTapIOK[A, B any](ma IOResult[A], f io.Kleisli[A, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOK runs the [IO] monad returned by the function but returns the result of the original monad
//
//go:inline
func ChainFirstIOK[A, B any](f io.Kleisli[A, B]) Operator[A, A] {
	_ = "STUB: not implemented"
	return nil
}

func TapIOK[A, B any](f io.Kleisli[A, B]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadFold[A, B any](ma IOResult[A], onLeft func(error) IO[B], onRight io.Kleisli[A, B]) IO[B] {
	_ = "STUB: not implemented"
	return nil
}

// WithResource constructs a function that creates a resourcthen operates on it and then releases the resource
//
//go:inline
func WithResource[A, R, ANY any](onCreate IOResult[R], onRelease Kleisli[R, ANY]) Kleisli[Kleisli[R, A], A] {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
//
//go:inline
func Swap[A any](val IOResult[A]) ioeither.IOEither[A, error] {
	_ = "STUB: not implemented"
	return nil

	// FromImpure converts a side effect without a return value into a side effect that returns any
	//
	//go:inline
}

func FromImpure[E any](f func()) IOResult[Void] { _ = "STUB: not implemented"; return nil }

// Defer creates an IO by creating a brand new IO via a generator function, each time
//
//go:inline
func Defer[A any](gen Lazy[IOResult[A]]) IOResult[A] { _ = "STUB: not implemented"; return nil }

// MonadAlt identifies an associative operation on a type constructor
//
//go:inline
func MonadAlt[A any](first IOResult[A], second Lazy[IOResult[A]]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

// Alt identifies an associative operation on a type constructor
//
//go:inline
func Alt[A any](second Lazy[IOResult[A]]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadFlap[B, A any](fab IOResult[func(A) B], a A) IOResult[B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Flap[B, A any](a A) Operator[func(A) B, B] { _ = "STUB: not implemented"; return nil }

// ToIOOption converts an [IOResult] to an [IOO.IOOption]
//
//go:inline
func ToIOOption[A any](ioe IOResult[A]) IOO.IOOption[A] { _ = "STUB: not implemented"; return nil }

// Delay creates an operation that passes in the value after some delay
//
//go:inline
func Delay[A any](delay time.Duration) Operator[A, A] { _ = "STUB: not implemented"; return nil }

// After creates an operation that passes after the given [time.Time]
//
//go:inline
func After[A any](timestamp time.Time) Operator[A, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadChainLeft[A any](fa IOResult[A], f Kleisli[error, A]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainLeft[A any](f Kleisli[error, A]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadChainFirstLeft[A, B any](fa IOResult[A], f Kleisli[error, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadTapLeft[A, B any](fa IOResult[A], f Kleisli[error, B]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainFirstLeft[A, B any](f Kleisli[error, B]) Operator[A, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func TapLeft[A, B any](f Kleisli[error, B]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

// OrElse recovers from a Left (error) by providing an alternative computation.
// If the IOResult is Right, it returns the value unchanged.
// If the IOResult is Left, it applies the provided function to the error value,
// which returns a new IOResult that replaces the original.
//
// This is useful for error recovery, fallback logic, or chaining alternative computations
// in IO contexts. Since IOResult is specialized for error type, the error type remains error.
//
// Example:
//
//	// Recover from specific errors with fallback IO operations
//	recover := ioresult.OrElse(func(err error) ioresult.IOResult[int] {
//	    if err.Error() == "not found" {
//	        return ioresult.Right[int](0) // default value
//	    }
//	    return ioresult.Left[int](err) // propagate other errors
//	})
//	result := recover(ioresult.Left[int](errors.New("not found"))) // Right(0)
//	result := recover(ioresult.Right[int](42)) // Right(42) - unchanged
//
//go:inline
func OrElse[A any](onLeft Kleisli[error, A]) Operator[A, A] { _ = "STUB: not implemented"; return nil }
