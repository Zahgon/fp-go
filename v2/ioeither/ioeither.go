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

package ioeither

import (
	"time"

	"github.com/IBM/fp-go/v2/either"
	"github.com/IBM/fp-go/v2/io"
	IOO "github.com/IBM/fp-go/v2/iooption"
	"github.com/IBM/fp-go/v2/lazy"
	O "github.com/IBM/fp-go/v2/option"
	R "github.com/IBM/fp-go/v2/reader"
)

type (
	IO[A any]        = io.IO[A]
	Either[E, A any] = either.Either[E, A]

	// IOEither represents a synchronous computation that may fail
	// refer to [https://andywhite.xyz/posts/2021-01-27-rte-foundations/#ioeitherlte-agt] for more details
	IOEither[E, A any] = IO[Either[E, A]]

	Kleisli[E, A, B any]  = R.Reader[A, IOEither[E, B]]
	Operator[E, A, B any] = Kleisli[E, IOEither[E, A], B]
)

// Left constructs an [IOEither] that represents a failure with an error value of type E
func Left[A, E any](l E) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// Right constructs an [IOEither] that represents a successful computation with a value of type A
func Right[E, A any](r A) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// Of constructs an [IOEither] that represents a successful computation with a value of type A.
// This is an alias for [Right] and is the canonical way to lift a pure value into the IOEither context.
func Of[E, A any](r A) IOEither[E, A] {
	_ = "STUB: not implemented"

	// MonadOf is an alias for [Of], provided for consistency with monad naming conventions
	return nil
}

func MonadOf[E, A any](r A) IOEither[E, A] {
	_ = "STUB: not implemented"

	// LeftIO constructs an [IOEither] from an [IO] that produces an error value
	return nil
}

func LeftIO[A, E any](ml IO[E]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// RightIO constructs an [IOEither] from an [IO] that produces a success value
func RightIO[E, A any](mr IO[A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// FromEither lifts an [Either] value into the [IOEither] context
func FromEither[E, A any](e Either[E, A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

func FromOption[A, E any](onNone func() E) func(o O.Option[A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromIOOption[A, E any](onNone func() E) func(o IOO.IOOption[A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[A, B, E any](onNone func() E) func(func(A) O.Option[B]) Operator[E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOK[E, A, B any](ma IOEither[E, A], f io.Kleisli[A, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOK[E, A, B any](f io.Kleisli[A, B]) Operator[E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainLazyK[E, A, B any](f func(A) lazy.Lazy[B]) Operator[E, A, B] {
	_ = "STUB: not implemented"
	return nil

	// FromIO creates an [IOEither] from an [IO] instance, invoking [IO] for each invocation of [IOEither]
}

func FromIO[E, A any](mr IO[A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// FromLazy creates an [IOEither] from a [Lazy] instance, invoking [Lazy] for each invocation of [IOEither]
func FromLazy[E, A any](mr lazy.Lazy[A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// MonadMap applies a function to the value inside a successful [IOEither], leaving errors unchanged
func MonadMap[E, A, B any](fa IOEither[E, A], f func(A) B) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Map returns a function that applies a transformation to the value inside a successful [IOEither]
func Map[E, A, B any](f func(A) B) Operator[E, A, B] { _ = "STUB: not implemented"; return nil }

// MonadMapTo replaces the value inside a successful [IOEither] with a constant value
func MonadMapTo[E, A, B any](fa IOEither[E, A], b B) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MapTo returns a function that replaces the value inside a successful [IOEither] with a constant value
func MapTo[E, A, B any](b B) Operator[E, A, B] { _ = "STUB: not implemented"; return nil }

// MonadChain sequences two [IOEither] computations, where the second depends on the result of the first
func MonadChain[E, A, B any](fa IOEither[E, A], f Kleisli[E, A, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Chain returns a function that sequences two [IOEither] computations
func Chain[E, A, B any](f Kleisli[E, A, B]) Operator[E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[E, A, B any](ma IOEither[E, A], f either.Kleisli[E, A, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[E, A, B any](f either.Kleisli[E, A, B]) Operator[E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAp applies a function wrapped in an [IOEither] to a value wrapped in an [IOEither]
func MonadAp[B, E, A any](mab IOEither[E, func(A) B], ma IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Ap applies a function wrapped in an [IOEither] to a value wrapped in an [IOEither].
// This is an alias of [ApPar] which applies the function and value in parallel.
func Ap[B, E, A any](ma IOEither[E, A]) Operator[E, func(A) B, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApPar applies a function wrapped in an [IOEither] to a value wrapped in an [IOEither] in parallel
func MonadApPar[B, E, A any](mab IOEither[E, func(A) B], ma IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApPar applies a function wrapped in an [IOEither] to a value wrapped in an [IOEither] in parallel
func ApPar[B, E, A any](ma IOEither[E, A]) Operator[E, func(A) B, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApSeq applies a function wrapped in an [IOEither] to a value wrapped in an [IOEither] sequentially
func MonadApSeq[B, E, A any](mab IOEither[E, func(A) B], ma IOEither[E, A]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApSeq applies a function wrapped in an [IOEither] to a value wrapped in an [IOEither] sequentially
func ApSeq[B, E, A any](ma IOEither[E, A]) func(IOEither[E, func(A) B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Flatten removes one level of nesting from a nested [IOEither]
func Flatten[E, A any](mma IOEither[E, IOEither[E, A]]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// TryCatch executes a function that may throw an error and converts it to an [IOEither]
func TryCatch[E, A any](f func() (A, error), onThrow func(error) E) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// TryCatchError executes a function that may throw an error and converts it to an [IOEither] with error type error
func TryCatchError[A any](f func() (A, error)) IOEither[error, A] {
	_ = "STUB: not implemented"
	return nil
}

// Memoize caches the result of an [IOEither] computation so it's only executed once
func Memoize[E, A any](ma IOEither[E, A]) IOEither[E, A] { _ = "STUB: not implemented"; return nil }

// MonadMapLeft applies a function to the error value of a failed [IOEither], leaving successful values unchanged
func MonadMapLeft[A, E1, E2 any](fa IOEither[E1, A], f func(E1) E2) IOEither[E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// MapLeft returns a function that applies a transformation to the error value of a failed [IOEither]
func MapLeft[A, E1, E2 any](f func(E1) E2) func(IOEither[E1, A]) IOEither[E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadBiMap applies one function to the error value and another to the success value of an [IOEither]
func MonadBiMap[E1, E2, A, B any](fa IOEither[E1, A], f func(E1) E2, g func(A) B) IOEither[E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap returns a function that maps a pair of functions over the two type arguments of the bifunctor
func BiMap[E1, E2, A, B any](f func(E1) E2, g func(A) B) func(IOEither[E1, A]) IOEither[E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// Fold converts an [IOEither] into an [IO] by providing handlers for both the error and success cases
func Fold[E, A, B any](onLeft func(E) IO[B], onRight io.Kleisli[A, B]) func(IOEither[E, A]) IO[B] {
	_ = "STUB: not implemented"
	return nil
}

// GetOrElse extracts the value from a successful [IOEither] or computes a default value from the error
func GetOrElse[E, A any](onLeft func(E) IO[A]) func(IOEither[E, A]) IO[A] {
	_ = "STUB: not implemented"
	return nil
}

// GetOrElseOf extracts the value from a successful [IOEither] or computes a default value from the error
func GetOrElseOf[E, A any](onLeft func(E) A) func(IOEither[E, A]) IO[A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainTo sequences two [IOEither] computations, discarding the result of the first
func MonadChainTo[A, E, B any](fa IOEither[E, A], fb IOEither[E, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainTo returns a function that sequences two [IOEither] computations, discarding the result of the first
func ChainTo[A, E, B any](fb IOEither[E, B]) Operator[E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainToIO sequences an [IOEither] with an [IO], discarding the result of the first
func MonadChainToIO[E, A, B any](fa IOEither[E, A], fb IO[B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainToIO returns a function that sequences an [IOEither] with an [IO], discarding the result of the first
func ChainToIO[E, A, B any](fb IO[B]) Operator[E, A, B] { _ = "STUB: not implemented"; return nil }

// MonadChainFirst executes a side-effecting [IOEither] computation but returns the original value
func MonadChainFirst[E, A, B any](ma IOEither[E, A], f Kleisli[E, A, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTap is an alias for [MonadChainFirst], executing a side effect while preserving the original value
func MonadTap[E, A, B any](ma IOEither[E, A], f Kleisli[E, A, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirst returns a function that executes a side-effecting [IOEither] computation but returns the original value
func ChainFirst[E, A, B any](f Kleisli[E, A, B]) Operator[E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// Tap is an alias for [ChainFirst], executing a side effect while preserving the original value
func Tap[E, A, B any](f Kleisli[E, A, B]) Operator[E, A, A] { _ = "STUB: not implemented"; return nil }

// MonadChainFirstEitherK executes a side-effecting [Either] computation but returns the original [IOEither] value
func MonadChainFirstEitherK[A, E, B any](ma IOEither[E, A], f either.Kleisli[E, A, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstEitherK returns a function that executes a side-effecting [Either] computation but returns the original value
func ChainFirstEitherK[A, E, B any](f either.Kleisli[E, A, B]) Operator[E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK executes a side-effecting [IO] computation but returns the original [IOEither] value
func MonadChainFirstIOK[E, A, B any](ma IOEither[E, A], f io.Kleisli[A, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOK returns a function that executes a side-effecting [IO] computation but returns the original value
func ChainFirstIOK[E, A, B any](f io.Kleisli[A, B]) Operator[E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapEitherK is an alias for [MonadChainFirstEitherK], executing an [Either] side effect while preserving the original value
func MonadTapEitherK[A, E, B any](ma IOEither[E, A], f either.Kleisli[E, A, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapEitherK is an alias for [ChainFirstEitherK], executing an [Either] side effect while preserving the original value
func TapEitherK[A, E, B any](f either.Kleisli[E, A, B]) Operator[E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapIOK is an alias for [MonadChainFirstIOK], executing an [IO] side effect while preserving the original value
func MonadTapIOK[E, A, B any](ma IOEither[E, A], f io.Kleisli[A, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapIOK is an alias for [ChainFirstIOK], executing an [IO] side effect while preserving the original value
func TapIOK[E, A, B any](f io.Kleisli[A, B]) Operator[E, A, A] {
	_ = "STUB: not implemented"
	return nil

	// MonadFold eliminates an [IOEither] by providing handlers for both error and success cases, returning an [IO]
}

func MonadFold[E, A, B any](ma IOEither[E, A], onLeft func(E) IO[B], onRight io.Kleisli[A, B]) IO[B] {
	_ = "STUB: not implemented"
	return nil
}

// WithResource constructs a function that safely manages a resource with automatic cleanup.
// It creates a resource, operates on it, and ensures the resource is released even if an error occurs.
func WithResource[A, E, R, ANY any](onCreate IOEither[E, R], onRelease Kleisli[E, R, ANY]) Kleisli[E, Kleisli[E, R, A], A] {
	_ = "STUB: not implemented"
	return nil
}

// Swap exchanges the error and success type parameters of an [IOEither]
func Swap[E, A any](val IOEither[E, A]) IOEither[A, E] { _ = "STUB: not implemented"; return nil }

// FromImpure converts a side effect without a return value into an [IOEither] that returns any
func FromImpure[E any](f func()) IOEither[E, Void] { _ = "STUB: not implemented"; return nil }

// Defer creates an [IOEither] by lazily evaluating a generator function each time the [IOEither] is executed
func Defer[E, A any](gen lazy.Lazy[IOEither[E, A]]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil

	// MonadAlt provides an alternative [IOEither] computation if the first one fails
}

func MonadAlt[E, A any](first IOEither[E, A], second lazy.Lazy[IOEither[E, A]]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Alt returns a function that provides an alternative [IOEither] computation if the first one fails
func Alt[E, A any](second lazy.Lazy[IOEither[E, A]]) Operator[E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadFlap applies a value to a function wrapped in an [IOEither]
func MonadFlap[E, B, A any](fab IOEither[E, func(A) B], a A) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Flap returns a function that applies a value to a function wrapped in an [IOEither]
func Flap[E, B, A any](a A) Operator[E, func(A) B, B] { _ = "STUB: not implemented"; return nil }

// ToIOOption converts an [IOEither] to an [IOO.IOOption], discarding error information
func ToIOOption[E, A any](ioe IOEither[E, A]) IOO.IOOption[A] {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operator that delays the execution of an [IOEither] by the specified duration
func Delay[E, A any](delay time.Duration) Operator[E, A, A] { _ = "STUB: not implemented"; return nil }

// After creates an operator that delays the execution of an [IOEither] until the specified time
func After[E, A any](timestamp time.Time) Operator[E, A, A] { _ = "STUB: not implemented"; return nil }

// MonadChainLeft chains a computation on the left (error) side of an [IOEither].
// If the input is a Left value, it applies the function f to transform the error and potentially
// change the error type from EA to EB. If the input is a Right value, it passes through unchanged.
//
// Note: MonadChainLeft is identical to [OrElse] - both provide the same functionality for error recovery.
//
// This is useful for error recovery or error transformation scenarios where you want to handle
// errors by performing another computation that may also fail.
//
// Parameters:
//   - fa: The input [IOEither] that may contain an error of type EA
//   - f: A function that takes an error of type EA and returns an [IOEither] with error type EB
//
// Returns:
//   - An [IOEither] with the potentially transformed error type EB
//
// Example:
//
//	// Recover from a specific error by trying an alternative computation
//	result := MonadChainLeft(
//	    Left[int]("network error"),
//	    func(err string) IOEither[string, int] {
//	        if err == "network error" {
//	            return Right[string](42) // recover with default value
//	        }
//	        return Left[int]("unrecoverable: " + err)
//	    },
//	)
func MonadChainLeft[EA, EB, A any](fa IOEither[EA, A], f Kleisli[EB, EA, A]) IOEither[EB, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainLeft is the curried version of [MonadChainLeft].
// It returns a function that chains a computation on the left (error) side of an [IOEither].
//
// Note: ChainLeft is identical to [OrElse] - both provide the same functionality for error recovery.
//
// This is particularly useful in functional composition pipelines where you want to handle
// errors by performing another computation that may also fail.
//
// Parameters:
//   - f: A function that takes an error of type EA and returns an [IOEither] with error type EB
//
// Returns:
//   - A function that transforms an [IOEither] with error type EA to one with error type EB
//
// Example:
//
//	// Create a reusable error handler
//	recoverFromNetworkError := ChainLeft(func(err string) IOEither[string, int] {
//	    if strings.Contains(err, "network") {
//	        return Right[string](0) // return default value
//	    }
//	    return Left[int](err) // propagate other errors
//	})
//
//	result := F.Pipe1(
//	    Left[int]("network timeout"),
//	    recoverFromNetworkError,
//	)
func ChainLeft[EA, EB, A any](f Kleisli[EB, EA, A]) func(IOEither[EA, A]) IOEither[EB, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstLeft chains a computation on the left (error) side but always returns the original error.
// If the input is a Left value, it applies the function f to the error and executes the resulting computation,
// but always returns the original Left error regardless of what f returns (Left or Right).
// If the input is a Right value, it passes through unchanged without calling f.
//
// This is useful for side effects on errors (like logging or metrics) where you want to perform an action
// when an error occurs but always propagate the original error, ensuring the error path is preserved.
//
// Parameters:
//   - ma: The input [IOEither] that may contain an error of type EA
//   - f: A function that takes an error of type EA and returns an [IOEither] (typically for side effects)
//
// Returns:
//   - An [IOEither] with the original error preserved if input was Left, or the original Right value
//
// Example:
//
//	// Log errors but always preserve the original error
//	result := MonadChainFirstLeft(
//	    Left[int]("database error"),
//	    func(err string) IOEither[string, int] {
//	        return FromIO[string](func() int {
//	            log.Printf("Error occurred: %s", err)
//	            return 0
//	        })
//	    },
//	)
//	// result will always be Left("database error"), even though f returns Right
func MonadChainFirstLeft[A, EA, EB, B any](ma IOEither[EA, A], f Kleisli[EB, EA, B]) IOEither[EA, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadTapLeft[A, EA, EB, B any](ma IOEither[EA, A], f Kleisli[EB, EA, B]) IOEither[EA, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstLeft is the curried version of [MonadChainFirstLeft].
// It returns a function that chains a computation on the left (error) side while always preserving the original error.
//
// This is particularly useful for adding error handling side effects (like logging, metrics, or notifications)
// in a functional pipeline. The original error is always returned regardless of what f returns (Left or Right),
// ensuring the error path is preserved.
//
// Parameters:
//   - f: A function that takes an error of type EA and returns an [IOEither] (typically for side effects)
//
// Returns:
//   - An [Operator] that performs the side effect but always returns the original error if input was Left
//
// Example:
//
//	// Create a reusable error logger
//	logError := ChainFirstLeft(func(err string) IOEither[any, int] {
//	    return FromIO[any](func() int {
//	        log.Printf("Error: %s", err)
//	        return 0
//	    })
//	})
//
//	result := F.Pipe1(
//	    Left[int]("validation failed"),
//	    logError, // logs the error
//	)
//	// result is always Left("validation failed"), even though f returns Right
func ChainFirstLeft[A, EA, EB, B any](f Kleisli[EB, EA, B]) Operator[EA, A, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func TapLeft[A, EA, EB, B any](f Kleisli[EB, EA, B]) Operator[EA, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// OrElse recovers from a Left (error) by providing an alternative computation.
// If the IOEither is Right, it returns the value unchanged.
// If the IOEither is Left, it applies the provided function to the error value,
// which returns a new IOEither that replaces the original.
//
// Note: OrElse is identical to [ChainLeft] - both provide the same functionality for error recovery.
//
// This is useful for error recovery, fallback logic, or chaining alternative IO computations.
// The error type can be widened from E1 to E2, allowing transformation of error types.
//
// Example:
//
//	// Recover from specific errors with fallback IO operations
//	recover := ioeither.OrElse(func(err error) ioeither.IOEither[error, int] {
//	    if err.Error() == "not found" {
//	        return ioeither.Right[error](0) // default value
//	    }
//	    return ioeither.Left[int](err) // propagate other errors
//	})
//	result := recover(ioeither.Left[int](errors.New("not found"))) // Right(0)
//	result := recover(ioeither.Right[error](42)) // Right(42) - unchanged
//
//go:inline
func OrElse[E1, E2, A any](onLeft Kleisli[E2, E1, A]) Kleisli[E2, IOEither[E1, A], A] {
	_ = "STUB: not implemented"
	return nil
}
