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

package readerioeither

import (
	"time"

	"github.com/IBM/fp-go/v2/either"
	"github.com/IBM/fp-go/v2/io"
	IOE "github.com/IBM/fp-go/v2/ioeither"
	L "github.com/IBM/fp-go/v2/lazy"
	"github.com/IBM/fp-go/v2/reader"
	RE "github.com/IBM/fp-go/v2/readereither"
	"github.com/IBM/fp-go/v2/readerio"
	"github.com/IBM/fp-go/v2/readeroption"
)

// FromReaderOption converts a ReaderOption to a Kleisli arrow that handles None cases.
// When the Option is None, the provided lazy error value is used.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - A: The value type
//   - E: The error type
//
// # Parameters
//
//   - onNone: Lazy function that provides the error value when Option is None
//
// # Returns
//
//   - Kleisli arrow that converts ReaderOption to ReaderIOEither
//
//go:inline
func FromReaderOption[R, A, E any](onNone Lazy[E]) Kleisli[R, E, ReaderOption[R, A], A] {
	_ = "STUB: not implemented"
	return nil
}

// FromReaderIO lifts a ReaderIO into a ReaderIOEither, placing the result in the Right side.
// This is an alias for RightReaderIO, converting a computation that cannot fail into one
// that can fail but never does.
//
// # Type Parameters
//
//   - E: The error type (will never actually contain an error)
//   - R: The context/environment type
//   - A: The value type
//
// # Parameters
//
//   - ma: The ReaderIO to lift
//
// # Returns
//
//   - ReaderIOEither with the ReaderIO result in the Right side
//
//go:inline
func FromReaderIO[E, R, A any](ma ReaderIO[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// RightReaderIO lifts a ReaderIO into a ReaderIOEither, placing the result in the Right side.
//
//go:inline
func RightReaderIO[E, R, A any](ma ReaderIO[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// LeftReaderIO lifts a ReaderIO into a ReaderIOEither, placing the result in the Left (error) side.
//
//go:inline
func LeftReaderIO[A, R, E any](me ReaderIO[R, E]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadMap applies a function to the value inside a ReaderIOEither context.
// If the computation is successful (Right), the function is applied to the value.
// If it's an error (Left), the error is propagated unchanged.
//
//go:inline
func MonadMap[R, E, A, B any](fa ReaderIOEither[R, E, A], f func(A) B) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Map returns a function that applies a transformation to the success value of a ReaderIOEither.
// This is the curried version of MonadMap, useful for function composition.
//
//go:inline
func Map[R, E, A, B any](f func(A) B) Operator[R, E, A, B] { _ = "STUB: not implemented"; return nil }

// MonadMapTo replaces the success value with a constant value.
// Useful when you want to discard the result but keep the effect.
func MonadMapTo[R, E, A, B any](fa ReaderIOEither[R, E, A], b B) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MapTo returns a function that replaces the success value with a constant.
// This is the curried version of MonadMapTo.
func MapTo[R, E, A, B any](b B) Operator[R, E, A, B] { _ = "STUB: not implemented"; return nil }

// MonadChain sequences two computations where the second depends on the result of the first.
// This is the fundamental operation for composing dependent effectful computations.
// If the first computation fails, the second is not executed.
//
//go:inline
func MonadChain[R, E, A, B any](fa ReaderIOEither[R, E, A], f Kleisli[R, E, A, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirst sequences two computations but keeps the result of the first.
// Useful for performing side effects while preserving the original value.
//
//go:inline
func MonadChainFirst[R, E, A, B any](fa ReaderIOEither[R, E, A], f Kleisli[R, E, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTap is an alias for MonadChainFirst.
// It sequences two computations but keeps the result of the first, emphasizing the
// side-effect nature of the operation (like "tapping" into a pipeline).
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - fa: The ReaderIOEither computation
//   - f: The side effect Kleisli arrow
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadTap[R, E, A, B any](fa ReaderIOEither[R, E, A], f Kleisli[R, E, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainEitherK chains a computation that returns an Either into a ReaderIOEither.
// The Either is automatically lifted into the ReaderIOEither context.
//
//go:inline
func MonadChainEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f either.Kleisli[E, A, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainEitherK returns a function that chains an Either-returning function into ReaderIOEither.
// This is the curried version of MonadChainEitherK.
//
//go:inline
func ChainEitherK[R, E, A, B any](f either.Kleisli[E, A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstEitherK chains an Either-returning computation but keeps the original value.
// Useful for validation or side effects that return Either.
//
//go:inline
func MonadChainFirstEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f either.Kleisli[E, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapEitherK is an alias for MonadChainFirstEitherK.
// It chains an Either-returning computation while preserving the original value,
// emphasizing the side-effect nature of the operation.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The Either-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadTapEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f either.Kleisli[E, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstEitherK returns a function that chains an Either computation while preserving the original value.
// This is the curried version of MonadChainFirstEitherK.
//
//go:inline
func ChainFirstEitherK[R, E, A, B any](f either.Kleisli[E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapEitherK is an alias for ChainFirstEitherK.
// It returns a function that chains an Either-returning side effect while preserving
// the original value, emphasizing the "tap" pattern for observing values.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: The Either-returning side effect function
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
//go:inline
func TapEitherK[R, E, A, B any](f either.Kleisli[E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainReaderK chains a Reader-returning computation into a ReaderIOEither.
// The Reader is automatically lifted into the ReaderIOEither context.
//
//go:inline
func MonadChainReaderK[E, R, A, B any](ma ReaderIOEither[R, E, A], f reader.Kleisli[R, A, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderK returns a function that chains a Reader-returning function into ReaderIOEither.
// This is the curried version of MonadChainReaderK.
//
//go:inline
func ChainReaderK[E, R, A, B any](f reader.Kleisli[R, A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstReaderK chains a Reader-returning computation while preserving the original value.
// Useful for performing Reader-based side effects (like logging with context) while keeping
// the original value.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The Reader-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadChainFirstReaderK[E, R, A, B any](ma ReaderIOEither[R, E, A], f reader.Kleisli[R, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapReaderK is an alias for MonadChainFirstReaderK.
// It chains a Reader-returning side effect while preserving the original value.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The Reader-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadTapReaderK[E, R, A, B any](ma ReaderIOEither[R, E, A], f reader.Kleisli[R, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderK returns a function that chains a Reader-returning function into ReaderIOEither
// while preserving the original value. This is the curried version of MonadChainFirstReaderK.
//
//go:inline
func ChainFirstReaderK[E, R, A, B any](f reader.Kleisli[R, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderK is an alias for ChainFirstReaderK.
// It returns a function that chains a Reader-returning side effect while preserving
// the original value.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: The Reader-returning side effect function
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
//go:inline
func TapReaderK[E, R, A, B any](f reader.Kleisli[R, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainReaderIOK chains a ReaderIO-returning computation into a ReaderIOEither.
// The ReaderIO is automatically lifted into the ReaderIOEither context.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type
//   - B: The output value type
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The ReaderIO-returning function
//
// # Returns
//
//   - ReaderIOEither with the result of the ReaderIO computation
//
//go:inline
func MonadChainReaderIOK[E, R, A, B any](ma ReaderIOEither[R, E, A], f readerio.Kleisli[R, A, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderIOK returns a function that chains a ReaderIO-returning function into ReaderIOEither.
// This is the curried version of MonadChainReaderIOK.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type
//   - B: The output value type
//
// # Parameters
//
//   - f: The ReaderIO-returning function
//
// # Returns
//
//   - Operator that chains the ReaderIO computation
//
//go:inline
func ChainReaderIOK[E, R, A, B any](f readerio.Kleisli[R, A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstReaderIOK chains a ReaderIO-returning computation while preserving the original value.
// Useful for performing ReaderIO-based side effects while keeping the original value.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The ReaderIO-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadChainFirstReaderIOK[E, R, A, B any](ma ReaderIOEither[R, E, A], f readerio.Kleisli[R, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapReaderIOK is an alias for MonadChainFirstReaderIOK.
// It chains a ReaderIO-returning side effect while preserving the original value.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The ReaderIO-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadTapReaderIOK[E, R, A, B any](ma ReaderIOEither[R, E, A], f readerio.Kleisli[R, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderIOK returns a function that chains a ReaderIO-returning function while
// preserving the original value. This is the curried version of MonadChainFirstReaderIOK.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: The ReaderIO-returning side effect function
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
//go:inline
func ChainFirstReaderIOK[E, R, A, B any](f readerio.Kleisli[R, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderIOK is an alias for ChainFirstReaderIOK.
// It returns a function that chains a ReaderIO-returning side effect while preserving
// the original value.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: The ReaderIO-returning side effect function
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
//go:inline
func TapReaderIOK[E, R, A, B any](f readerio.Kleisli[R, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainReaderEitherK chains a ReaderEither-returning computation into a ReaderIOEither.
// The ReaderEither is automatically lifted into the ReaderIOEither context.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type
//   - B: The output value type
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The ReaderEither-returning function
//
// # Returns
//
//   - ReaderIOEither with the result of the ReaderEither computation
//
//go:inline
func MonadChainReaderEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f RE.Kleisli[R, E, A, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderEitherK returns a function that chains a ReaderEither-returning function into ReaderIOEither.
// This is the curried version of MonadChainReaderEitherK.
//
//go:inline
func ChainReaderEitherK[E, R, A, B any](f RE.Kleisli[R, E, A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstReaderEitherK chains a ReaderEither-returning computation while preserving the original value.
// Useful for performing ReaderEither-based side effects while keeping the original value.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The ReaderEither-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadChainFirstReaderEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f RE.Kleisli[R, E, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapReaderEitherK is an alias for MonadChainFirstReaderEitherK.
// It chains a ReaderEither-returning side effect while preserving the original value.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The ReaderEither-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadTapReaderEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f RE.Kleisli[R, E, A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderEitherK returns a function that chains a ReaderEither-returning function into ReaderIOEither
// while preserving the original value. This is the curried version of MonadChainFirstReaderEitherK.
//
//go:inline
func ChainFirstReaderEitherK[E, R, A, B any](f RE.Kleisli[R, E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderEitherK is an alias for ChainFirstReaderEitherK.
// It returns a function that chains a ReaderEither-returning side effect while preserving
// the original value.
//
// # Type Parameters
//
//   - E: The error type
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: The ReaderEither-returning side effect function
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
//go:inline
func TapReaderEitherK[E, R, A, B any](f RE.Kleisli[R, E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderOptionK returns a function that chains a ReaderOption-returning function into ReaderIOEither.
// When the Option is None, the provided error value is used.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - A: The input value type
//   - B: The output value type
//   - E: The error type
//
// # Parameters
//
//   - onNone: Lazy function that provides the error value when Option is None
//
// # Returns
//
//   - Function that takes a ReaderOption Kleisli and returns an Operator
//
//go:inline
func ChainReaderOptionK[R, A, B, E any](onNone Lazy[E]) func(readeroption.Kleisli[R, A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderOptionK returns a function that chains a ReaderOption-returning function
// while preserving the original value. When the Option is None, the provided error value is used.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//   - E: The error type
//
// # Parameters
//
//   - onNone: Lazy function that provides the error value when Option is None
//
// # Returns
//
//   - Function that takes a ReaderOption Kleisli and returns an Operator
//
//go:inline
func ChainFirstReaderOptionK[R, A, B, E any](onNone Lazy[E]) func(readeroption.Kleisli[R, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderOptionK is an alias for ChainFirstReaderOptionK.
// It returns a function that chains a ReaderOption-returning side effect while preserving
// the original value.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//   - E: The error type
//
// # Parameters
//
//   - onNone: Lazy function that provides the error value when Option is None
//
// # Returns
//
//   - Function that takes a ReaderOption Kleisli and returns an Operator
//
//go:inline
func TapReaderOptionK[R, A, B, E any](onNone Lazy[E]) func(readeroption.Kleisli[R, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainIOEitherK chains an IOEither-returning computation into a ReaderIOEither.
// The IOEither is automatically lifted into the ReaderIOEither context.
//
//go:inline
func MonadChainIOEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f IOE.Kleisli[E, A, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainIOEitherK returns a function that chains an IOEither-returning function into ReaderIOEither.
// This is the curried version of MonadChainIOEitherK.
//
//go:inline
func ChainIOEitherK[R, E, A, B any](f IOE.Kleisli[E, A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOEitherK chains an IOEither computation while preserving the original value.
// This is useful for performing side effects that may fail (like logging, validation, or
// external API calls) while keeping the original value in the success case.
//
// The function executes the IOEither computation but discards its result, returning the
// original value if both computations succeed. If either computation fails, the error
// is propagated.
//
// This is the curried version that returns an Operator for use in function composition.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: IOEither.Kleisli function that performs the side effect
//
// # Returns
//
//   - Operator that chains the side effect while preserving the original value
//
// # Example Usage
//
//	type Config struct{ LogEnabled bool }
//
//	logValue := func(v int) IOEither[error, string] {
//	    return IOE.Of[error](fmt.Sprintf("Value: %d", v))
//	}
//
//	pipeline := F.Pipe1(
//	    Of[Config, error](42),
//	    ChainFirstIOEitherK[Config](logValue),
//	)
//	result := pipeline(Config{LogEnabled: true})() // Right(42)
//
// # See Also
//
//   - TapIOEitherK: Alias for ChainFirstIOEitherK
//   - ChainIOEitherK: Chains IOEither and uses its result
//   - ChainFirstEitherK: Similar but for Either computations
//
//go:inline
func ChainFirstIOEitherK[R, E, A, B any](f IOE.Kleisli[E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapIOEitherK is an alias for ChainFirstIOEitherK.
// It executes an IOEither side effect while preserving the original value.
//
// The name "Tap" emphasizes the side-effect nature of the operation, similar to
// tapping into a pipeline to observe or log values without modifying the flow.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: IOEither.Kleisli function that performs the side effect
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
// # Example Usage
//
//	type Config struct{}
//
//	validatePositive := func(v int) IOEither[error, bool] {
//	    if v > 0 {
//	        return IOE.Of[error](true)
//	    }
//	    return IOE.Left[bool](errors.New("must be positive"))
//	}
//
//	pipeline := F.Pipe1(
//	    Of[Config, error](42),
//	    TapIOEitherK[Config](validatePositive),
//	)
//	result := pipeline(Config{})() // Right(42) if validation passes
//
// # See Also
//
//   - ChainFirstIOEitherK: The underlying implementation
//   - TapEitherK: Similar but for Either computations
//   - TapIOK: Similar but for IO computations
//
//go:inline
func TapIOEitherK[R, E, A, B any](f IOE.Kleisli[E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainIOK chains an IO-returning computation into a ReaderIOEither.
// The IO is automatically lifted into the ReaderIOEither context (always succeeds).
//
//go:inline
func MonadChainIOK[R, E, A, B any](ma ReaderIOEither[R, E, A], f io.Kleisli[A, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainIOK returns a function that chains an IO-returning function into ReaderIOEither.
// This is the curried version of MonadChainIOK.
//
//go:inline
func ChainIOK[R, E, A, B any](f io.Kleisli[A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK chains an IO computation but keeps the original value.
// Useful for performing IO side effects while preserving the original value.
//
//go:inline
func MonadChainFirstIOK[R, E, A, B any](ma ReaderIOEither[R, E, A], f io.Kleisli[A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapIOK is an alias for MonadChainFirstIOK.
// It chains an IO-returning side effect while preserving the original value.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - ma: The ReaderIOEither computation
//   - f: The IO-returning side effect function
//
// # Returns
//
//   - ReaderIOEither with the original value preserved
//
//go:inline
func MonadTapIOK[R, E, A, B any](ma ReaderIOEither[R, E, A], f io.Kleisli[A, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOK returns a function that chains an IO computation while preserving the original value.
// This is the curried version of MonadChainFirstIOK.
//
//go:inline
func ChainFirstIOK[R, E, A, B any](f io.Kleisli[A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapIOK is an alias for ChainFirstIOK.
// It returns a function that chains an IO-returning side effect while preserving
// the original value.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: The IO-returning side effect function
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
//go:inline
func TapIOK[R, E, A, B any](f io.Kleisli[A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainOptionK returns a function that chains an Option-returning function into ReaderIOEither.
// If the Option is None, the provided error function is called to produce the error value.
//
//go:inline
func ChainOptionK[R, A, B, E any](onNone Lazy[E]) func(func(A) Option[B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAp applies a function wrapped in a context to a value wrapped in a context.
// Both computations are executed (default behavior may be sequential or parallel depending on implementation).
//
//go:inline
func MonadAp[R, E, A, B any](fab ReaderIOEither[R, E, func(A) B], fa ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApSeq applies a function in a context to a value in a context, executing them sequentially.
//
//go:inline
func MonadApSeq[R, E, A, B any](fab ReaderIOEither[R, E, func(A) B], fa ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApPar applies a function in a context to a value in a context, executing them in parallel.
//
//go:inline
func MonadApPar[R, E, A, B any](fab ReaderIOEither[R, E, func(A) B], fa ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Ap returns a function that applies a function in a context to a value in a context.
// This is the curried version of MonadAp.
func Ap[B, R, E, A any](fa ReaderIOEither[R, E, A]) func(fab ReaderIOEither[R, E, func(A) B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Chain returns a function that sequences computations where the second depends on the first.
// This is the curried version of MonadChain.
//
//go:inline
func Chain[R, E, A, B any](f Kleisli[R, E, A, B]) Operator[R, E, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirst returns a function that sequences computations but keeps the first result.
// This is the curried version of MonadChainFirst.
//
//go:inline
func ChainFirst[R, E, A, B any](f Kleisli[R, E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// Tap is an alias for ChainFirst.
// It returns a function that sequences computations but keeps the first result,
// emphasizing the side-effect nature of the operation.
//
// # Type Parameters
//
//   - R: The context/environment type
//   - E: The error type
//   - A: The input value type (preserved in output)
//   - B: The side effect result type (discarded)
//
// # Parameters
//
//   - f: The Kleisli arrow for the side effect
//
// # Returns
//
//   - Operator that executes the side effect while preserving the original value
//
//go:inline
func Tap[R, E, A, B any](f Kleisli[R, E, A, B]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil

	// Right creates a successful ReaderIOEither with the given value.
	//
	//go:inline
}

func Right[R, E, A any](a A) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// Left creates a failed ReaderIOEither with the given error.
//
//go:inline
func Left[R, A, E any](e E) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// ThrowError creates a failed ReaderIOEither with the given error.
// This is an alias for Left, following the naming convention from other functional libraries.
func ThrowError[R, A, E any](e E) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// Of creates a successful ReaderIOEither with the given value.
// This is the pointed functor operation, lifting a pure value into the ReaderIOEither context.
func Of[R, E, A any](a A) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// Flatten removes one level of nesting from a nested ReaderIOEither.
// Converts ReaderIOEither[R, E, ReaderIOEither[R, E, A]] to ReaderIOEither[R, E, A].
func Flatten[R, E, A any](mma ReaderIOEither[R, E, ReaderIOEither[R, E, A]]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromEither lifts an Either into a ReaderIOEither context.
// The Either value is independent of any context or IO effects.
//
//go:inline
func FromEither[R, E, A any](t either.Either[E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil

	// RightReader lifts a Reader into a ReaderIOEither, placing the result in the Right side.
}

func RightReader[E, R, A any](ma Reader[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// LeftReader lifts a Reader into a ReaderIOEither, placing the result in the Left (error) side.
func LeftReader[A, R, E any](ma Reader[R, E]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromReader lifts a Reader into a ReaderIOEither context.
// The Reader result is placed in the Right side (success).
func FromReader[E, R, A any](ma Reader[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil

	// RightIO lifts an IO into a ReaderIOEither, placing the result in the Right side.
}

func RightIO[R, E, A any](ma IO[A]) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// LeftIO lifts an IO into a ReaderIOEither, placing the result in the Left (error) side.
func LeftIO[R, A, E any](ma IO[E]) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// FromIO lifts an IO into a ReaderIOEither context.
// The IO result is placed in the Right side (success).
func FromIO[R, E, A any](ma IO[A]) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// FromIOEither lifts an IOEither into a ReaderIOEither context.
// The computation becomes independent of any reader context.
//
//go:inline
func FromIOEither[R, E, A any](ma IOEither[E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil

	// FromReaderEither lifts a ReaderEither into a ReaderIOEither context.
	// The Either result is lifted into an IO effect.
}

func FromReaderEither[R, E, A any](ma RE.ReaderEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Ask returns a ReaderIOEither that retrieves the current context.
// Useful for accessing configuration or dependencies.
//
//go:inline
func Ask[R, E any]() ReaderIOEither[R, E, R] { _ = "STUB: not implemented"; return nil }

// Asks returns a ReaderIOEither that retrieves a value derived from the context.
// This is useful for extracting specific fields from a configuration object.
//
//go:inline
func Asks[E, R, A any](r Reader[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromOption converts an Option to a ReaderIOEither.
// If the Option is None, the provided function is called to produce the error.
//
//go:inline
func FromOption[R, A, E any](onNone Lazy[E]) func(Option[A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromPredicate creates a ReaderIOEither from a predicate.
// If the predicate returns false, the onFalse function is called to produce the error.
//
//go:inline
func FromPredicate[R, E, A any](pred func(A) bool, onFalse func(A) E) func(A) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Fold handles both success and error cases, producing a ReaderIO.
// This is useful for converting a ReaderIOEither into a ReaderIO by handling all cases.
//
//go:inline
func Fold[R, E, A, B any](onLeft readerio.Kleisli[R, E, B], onRight func(A) ReaderIO[R, B]) func(ReaderIOEither[R, E, A]) ReaderIO[R, B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadFold[R, E, A, B any](ma ReaderIOEither[R, E, A], onLeft readerio.Kleisli[R, E, B], onRight func(A) ReaderIO[R, B]) ReaderIO[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// GetOrElse provides a default value in case of error.
// The default is computed lazily via a ReaderIO.
//
//go:inline
func GetOrElse[R, E, A any](onLeft readerio.Kleisli[R, E, A]) func(ReaderIOEither[R, E, A]) ReaderIO[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// OrLeft transforms the error using a ReaderIO if the computation fails.
// The success value is preserved unchanged.
//
//go:inline
func OrLeft[A, E1, R, E2 any](onLeft func(E1) ReaderIO[R, E2]) func(ReaderIOEither[R, E1, A]) ReaderIOEither[R, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadBiMap applies two functions: one to the error, one to the success value.
// This allows transforming both channels simultaneously.
//
//go:inline
func MonadBiMap[R, E1, E2, A, B any](fa ReaderIOEither[R, E1, A], f func(E1) E2, g func(A) B) ReaderIOEither[R, E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap returns a function that maps over both the error and success channels.
// This is the curried version of MonadBiMap.
//
//go:inline
func BiMap[R, E1, E2, A, B any](f func(E1) E2, g func(A) B) func(ReaderIOEither[R, E1, A]) ReaderIOEither[R, E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// Swap exchanges the error and success types.
// Left becomes Right and Right becomes Left.
//
//go:inline
func Swap[R, E, A any](val ReaderIOEither[R, E, A]) ReaderIOEither[R, A, E] {
	_ = "STUB: not implemented"
	return nil
}

// Defer creates a ReaderIOEither lazily via a generator function.
// The generator is called each time the ReaderIOEither is executed.
//
//go:inline
func Defer[R, E, A any](gen L.Lazy[ReaderIOEither[R, E, A]]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil

	// TryCatch wraps a function that returns (value, error) into a ReaderIOEither.
	// The onThrow function converts the error into the desired error type.
}

func TryCatch[R, E, A any](f func(R) func() (A, error), onThrow func(error) E) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAlt tries the first computation, and if it fails, tries the second.
// This implements the Alternative pattern for error recovery.
//
//go:inline
func MonadAlt[R, E, A any](first ReaderIOEither[R, E, A], second L.Lazy[ReaderIOEither[R, E, A]]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Alt returns a function that tries an alternative computation if the first fails.
// This is the curried version of MonadAlt.
//
//go:inline
func Alt[R, E, A any](second L.Lazy[ReaderIOEither[R, E, A]]) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// Memoize computes the value of the ReaderIOEither lazily but exactly once.
// The context used is from the first call. Do not use if the value depends on the context.
//
//go:inline
func Memoize[
	R, E, A any](rdr ReaderIOEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadFlap applies a value to a function wrapped in a context.
// This is the reverse of Ap - the value is fixed and the function varies.
//
//go:inline
func MonadFlap[R, E, B, A any](fab ReaderIOEither[R, E, func(A) B], a A) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Flap returns a function that applies a fixed value to a function in a context.
// This is the curried version of MonadFlap.
//
//go:inline
func Flap[R, E, B, A any](a A) func(ReaderIOEither[R, E, func(A) B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadMapLeft applies a function to the error value, leaving success unchanged.
//
//go:inline
func MonadMapLeft[R, E1, E2, A any](fa ReaderIOEither[R, E1, A], f func(E1) E2) ReaderIOEither[R, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// MapLeft returns a function that transforms the error channel.
// This is the curried version of MonadMapLeft.
//
//go:inline
func MapLeft[R, A, E1, E2 any](f func(E1) E2) func(ReaderIOEither[R, E1, A]) ReaderIOEither[R, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// Local runs a computation with a modified context.
// The function f transforms the context before passing it to the computation.
// This is similar to Contravariant's contramap operation.
//
//go:inline
func Local[E, A, R1, R2 any](f func(R2) R1) func(ReaderIOEither[R1, E, A]) ReaderIOEither[R2, E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Read[E, A, R any](r R) func(ReaderIOEither[R, E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ReadIOEither executes a ReaderIOEither computation by providing it with an environment
// obtained from an IOEither computation. This is useful when the environment itself needs
// to be computed with side effects and error handling.
//
// The function first executes the IOEither[E, R] to get the environment R (or fail with error E),
// then uses that environment to run the ReaderIOEither[R, E, A] computation.
//
// Type parameters:
//   - A: The success value type of the ReaderIOEither computation
//   - R: The environment/context type required by the ReaderIOEither
//   - E: The error type
//
// Parameters:
//   - r: An IOEither[E, R] that produces the environment (or an error)
//
// Returns:
//   - A function that takes a ReaderIOEither[R, E, A] and returns IOEither[E, A]
//
// Behavior:
//   - If the IOEither[E, R] fails (Left), the error is propagated without running the ReaderIOEither
//   - If the IOEither[E, R] succeeds (Right), the resulting environment is used to execute the ReaderIOEither
//
// Example:
//
//	// Load configuration from a file (may fail)
//	loadConfig := func() IOEither[error, Config] {
//	    return Lazy[E]ither[error, Config] {
//	        // Read config file with error handling
//	        return either.Right(Config{BaseURL: "https://api.example.com"})
//	    }
//	}
//
//	// A computation that needs the config
//	fetchUser := func(id int) ReaderIOEither[Config, error, User] {
//	    return func(cfg Config) IOEither[error, User] {
//	        // Use cfg.BaseURL to fetch user
//	        return ioeither.Right[error](User{ID: id})
//	    }
//	}
//
//	// Execute the computation with dynamically loaded config
//	result := ReadIOEither[User](loadConfig())(fetchUser(123))()
//
//go:inline
func ReadIOEither[A, R, E any](r IOEither[E, R]) func(ReaderIOEither[R, E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ReadIO executes a ReaderIOEither computation by providing it with an environment
// obtained from an IO computation. This is useful when the environment needs to be
// computed with side effects but cannot fail.
//
// The function first executes the IO[R] to get the environment R,
// then uses that environment to run the ReaderIOEither[R, E, A] computation.
//
// Type parameters:
//   - E: The error type of the ReaderIOEither computation
//   - A: The success value type of the ReaderIOEither computation
//   - R: The environment/context type required by the ReaderIOEither
//
// Parameters:
//   - r: An IO[R] that produces the environment
//
// Returns:
//   - A function that takes a ReaderIOEither[R, E, A] and returns IOEither[E, A]
//
// Behavior:
//   - The IO[R] is always executed successfully to obtain the environment
//   - The resulting environment is then used to execute the ReaderIOEither
//   - Only the ReaderIOEither computation can fail with error type E
//
// Example:
//
//	// Get current timestamp (cannot fail)
//	getCurrentTime := func() IO[time.Time] {
//	    return func() time.Time {
//	        return time.Now()
//	    }
//	}
//
//	// A computation that needs the timestamp
//	logWithTimestamp := func(msg string) ReaderIOEither[time.Time, error, string] {
//	    return func(t time.Time) IOEither[error, string] {
//	        logged := fmt.Sprintf("[%s] %s", t.Format(time.RFC3339), msg)
//	        return ioeither.Right[error](logged)
//	    }
//	}
//
//	// Execute the computation with current time
//	result := ReadIO[error, string](getCurrentTime())(logWithTimestamp("Hello"))()
//
//go:inline
func ReadIO[E, A, R any](r IO[R]) func(ReaderIOEither[R, E, A]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainLeft chains a computation on the left (error) side of a ReaderIOEither.
// If the input is a Left value, it applies the function f to transform the error and potentially
// change the error type from EA to EB. If the input is a Right value, it passes through unchanged.
//
// This is useful for error recovery or error transformation scenarios where you want to handle
// errors by performing another computation that may also fail, with access to configuration context.
//
// Note: This is functionally identical to the uncurried form of [OrElse]. Use [ChainLeft] when
// emphasizing the monadic chaining perspective, and [OrElse] for error recovery semantics.
//
// Parameters:
//   - fa: The input ReaderIOEither that may contain an error of type EA
//   - f: A Kleisli function that takes an error of type EA and returns a ReaderIOEither with error type EB
//
// Returns:
//   - A ReaderIOEither with the potentially transformed error type EB
//
// Example:
//
//	type Config struct{ retryCount int }
//	type NetworkError struct{ msg string }
//	type SystemError struct{ code int }
//
//	// Recover from network errors by retrying with config
//	result := MonadChainLeft(
//	    Left[Config, string](NetworkError{"connection failed"}),
//	    func(ne NetworkError) readerioeither.ReaderIOEither[Config, SystemError, string] {
//	        return readerioeither.Asks[SystemError](func(cfg Config) ioeither.IOEither[SystemError, string] {
//	            if cfg.retryCount > 0 {
//	                return ioeither.Right[SystemError]("recovered")
//	            }
//	            return ioeither.Left[string](SystemError{500})
//	        })
//	    },
//	)
//
//go:inline
func MonadChainLeft[R, EA, EB, A any](fa ReaderIOEither[R, EA, A], f Kleisli[R, EB, EA, A]) ReaderIOEither[R, EB, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainLeft is the curried version of [MonadChainLeft].
// It returns a function that chains a computation on the left (error) side of a ReaderIOEither.
//
// This is particularly useful in functional composition pipelines where you want to handle
// errors by performing another computation that may also fail, with access to configuration context.
//
// Note: This is functionally identical to [OrElse]. They are different names for the same operation.
// Use [ChainLeft] when emphasizing the monadic chaining perspective on the error channel,
// and [OrElse] when emphasizing error recovery/fallback semantics.
//
// Parameters:
//   - f: A Kleisli function that takes an error of type EA and returns a ReaderIOEither with error type EB
//
// Returns:
//   - A function that transforms a ReaderIOEither with error type EA to one with error type EB
//
// Example:
//
//	type Config struct{ fallbackService string }
//
//	// Create a reusable error handler with config access
//	recoverFromNetworkError := ChainLeft(func(err string) readerioeither.ReaderIOEither[Config, string, int] {
//	    if strings.Contains(err, "network") {
//	        return readerioeither.Asks[string](func(cfg Config) ioeither.IOEither[string, int] {
//	            return ioeither.TryCatch(
//	                func() (int, error) { return callService(cfg.fallbackService) },
//	                func(e error) string { return e.Error() },
//	            )
//	        })
//	    }
//	    return readerioeither.Left[Config, int](err)
//	})
//
//	result := F.Pipe1(
//	    readerioeither.Left[Config, int]("network timeout"),
//	    recoverFromNetworkError,
//	)(Config{fallbackService: "backup"})()
//
//go:inline
func ChainLeft[R, EA, EB, A any](f Kleisli[R, EB, EA, A]) func(ReaderIOEither[R, EA, A]) ReaderIOEither[R, EB, A] {
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
//   - ma: The input ReaderIOEither that may contain an error of type EA
//   - f: A function that takes an error of type EA and returns a ReaderIOEither (typically for side effects)
//
// Returns:
//   - A ReaderIOEither with the original error preserved if input was Left, or the original Right value
//
//go:inline
func MonadChainFirstLeft[A, R, EA, EB, B any](ma ReaderIOEither[R, EA, A], f Kleisli[R, EB, EA, B]) ReaderIOEither[R, EA, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadTapLeft[A, R, EA, EB, B any](ma ReaderIOEither[R, EA, A], f Kleisli[R, EB, EA, B]) ReaderIOEither[R, EA, A] {
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
//   - f: A function that takes an error of type EA and returns a ReaderIOEither (typically for side effects)
//
// Returns:
//   - An Operator that performs the side effect but always returns the original error if input was Left
//
//go:inline
func ChainFirstLeft[A, R, EA, EB, B any](f Kleisli[R, EB, EA, B]) Operator[R, EA, A, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirstLeftIOK[A, R, EA, B any](f io.Kleisli[EA, B]) Operator[R, EA, A, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func TapLeft[A, R, EA, EB, B any](f Kleisli[R, EB, EA, B]) Operator[R, EA, A, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func TapLeftIOK[A, R, EA, B any](f io.Kleisli[EA, B]) Operator[R, EA, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some delay
//
//go:inline
func Delay[R, E, A any](delay time.Duration) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// After creates an operation that passes after the given [time.Time]
//
//go:inline
func After[R, E, A any](timestamp time.Time) Operator[R, E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// OrElse recovers from a Left (error) by providing an alternative IO computation with access to the reader context.
// If the ReaderIOEither is Right, it returns the value unchanged.
// If the ReaderIOEither is Left, it applies the provided function to the error value,
// which returns a new ReaderIOEither that replaces the original.
//
// Note: OrElse is identical to [ChainLeft] - both provide the same functionality for error recovery.
//
// This is useful for error recovery, fallback logic, or chaining alternative IO computations
// that need access to configuration or dependencies. The error type can be widened from E1 to E2.
//
// Example:
//
//	type Config struct{ retryLimit int }
//
//	// Recover with IO operation using config
//	recover := readerioeither.OrElse(func(err error) readerioeither.ReaderIOEither[Config, error, int] {
//	    if err.Error() == "retryable" {
//	        return readerioeither.Asks[error](func(cfg Config) ioeither.IOEither[error, int] {
//	            if cfg.retryLimit > 0 {
//	                return ioeither.Right[error](42)
//	            }
//	            return ioeither.Left[int](err)
//	        })
//	    }
//	    return readerioeither.Left[Config, int](err)
//	})
//
//go:inline
func OrElse[R, E1, E2, A any](onLeft Kleisli[R, E2, E1, A]) Kleisli[R, E2, ReaderIOEither[R, E1, A], A] {
	_ = "STUB: not implemented"
	return nil
}
