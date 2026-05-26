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

package readerioresult

import (
	"time"

	"github.com/IBM/fp-go/v2/io"
	"github.com/IBM/fp-go/v2/ioresult"
	"github.com/IBM/fp-go/v2/reader"
	RE "github.com/IBM/fp-go/v2/readereither"
	"github.com/IBM/fp-go/v2/readerio"
	RIOE "github.com/IBM/fp-go/v2/readerioeither"
	"github.com/IBM/fp-go/v2/readeroption"
	"github.com/IBM/fp-go/v2/result"
)

// FromReaderOption converts a ReaderOption to a ReaderIOResult.
// If the ReaderOption is None, the provided function is called to produce the error.
//
//go:inline
func FromReaderOption[R, A any](onNone Lazy[error]) Kleisli[R, ReaderOption[R, A], A] {
	_ = "STUB: not implemented"
	return nil
}

// FromReaderIO creates a function that lifts a ReaderIO-producing function into ReaderIOResult.
// The ReaderIO result is placed in the Right side of the Either.
//
//go:inline
func FromReaderIO[R, A any](ma ReaderIO[R, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// RightReaderIO lifts a ReaderIO into a ReaderIOResult, placing the result in the Right side.
//
//go:inline
func RightReaderIO[R, A any](ma ReaderIO[R, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// LeftReaderIO lifts a ReaderIO into a ReaderIOResult, placing the result in the Left (error) side.
//
//go:inline
func LeftReaderIO[A, R any](me ReaderIO[R, error]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadMap applies a function to the value inside a ReaderIOResult context.
// If the computation is successful (Right), the function is applied to the value.
// If it's an error (Left), the error is propagated unchanged.
//
//go:inline
func MonadMap[R, A, B any](fa ReaderIOResult[R, A], f func(A) B) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// Map returns a function that applies a transformation to the success value of a ReaderIOResult.
// This is the curried version of MonadMap, useful for function composition.
//
//go:inline
func Map[R, A, B any](f func(A) B) Operator[R, A, B] { _ = "STUB: not implemented"; return nil }

// MonadMapTo replaces the success value with a constant value.
// Useful when you want to discard the result but keep the effect.
//
//go:inline
func MonadMapTo[R, A, B any](fa ReaderIOResult[R, A], b B) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// MapTo returns a function that replaces the success value with a constant.
// This is the curried version of MonadMapTo.
//
//go:inline
func MapTo[R, A, B any](b B) Operator[R, A, B] { _ = "STUB: not implemented"; return nil }

// MonadChain sequences two computations where the second depends on the result of the first.
// This is the fundamental operation for composing dependent effectful computations.
// If the first computation fails, the second is not executed.
//
//go:inline
func MonadChain[R, A, B any](fa ReaderIOResult[R, A], f Kleisli[R, A, B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirst sequences two computations but keeps the result of the first.
// Useful for performing side effects while preserving the original value.
//
//go:inline
func MonadChainFirst[R, A, B any](fa ReaderIOResult[R, A], f Kleisli[R, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTap is an alias for MonadChainFirst, executing a side effect while preserving the original value.
// The name "Tap" emphasizes the side-effect nature of the operation.
//
//go:inline
func MonadTap[R, A, B any](fa ReaderIOResult[R, A], f Kleisli[R, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainEitherK chains a computation that returns an Either into a ReaderIOResult.
// The Either is automatically lifted into the ReaderIOResult context.
//
//go:inline
func MonadChainEitherK[R, A, B any](ma ReaderIOResult[R, A], f result.Kleisli[A, B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainEitherK chains a computation that returns an Either into a ReaderIOResult.
// The Either is automatically lifted into the ReaderIOResult context.
//
//go:inline
func MonadChainResultK[R, A, B any](ma ReaderIOResult[R, A], f result.Kleisli[A, B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainEitherK returns a function that chains an Either-returning function into ReaderIOResult.
// This is the curried version of MonadChainEitherK.
//
//go:inline
func ChainEitherK[R, A, B any](f result.Kleisli[A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainResultK returns a function that chains an Either-returning function into ReaderIOResult.
// This is the curried version of MonadChainEitherK.
//
//go:inline
func ChainResultK[R, A, B any](f result.Kleisli[A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstEitherK chains an Either-returning computation but keeps the original value.
// Useful for validation or side effects that return Either.
//
//go:inline
func MonadChainFirstEitherK[R, A, B any](ma ReaderIOResult[R, A], f result.Kleisli[A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapEitherK is an alias for MonadChainFirstEitherK, executing a Result side effect while preserving the original value.
//
//go:inline
func MonadTapEitherK[R, A, B any](ma ReaderIOResult[R, A], f result.Kleisli[A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstEitherK returns a function that chains an Either computation while preserving the original value.
// This is the curried version of MonadChainFirstEitherK.
//
//go:inline
func ChainFirstEitherK[R, A, B any](f result.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapEitherK is an alias for ChainFirstEitherK, executing a Result side effect while preserving the original value.
//
//go:inline
func TapEitherK[R, A, B any](f result.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOEitherK chains an IOResult computation while preserving the original value.
// Useful for performing side effects that may fail while keeping the original value.
//
//go:inline
func ChainFirstIOEitherK[R, A, B any](f ioresult.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapIOEitherK is an alias for ChainFirstIOEitherK, executing an IOResult side effect while preserving the original value.
//
//go:inline
func TapIOEitherK[R, A, B any](f ioresult.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOResultK chains an IOResult computation while preserving the original value.
// This is an alias for ChainFirstIOEitherK with more explicit naming.
//
//go:inline
func ChainFirstIOResultK[R, A, B any](f ioresult.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapIOResultK is an alias for ChainFirstIOResultK, executing an IOResult side effect while preserving the original value.
//
//go:inline
func TapIOResultK[R, A, B any](f ioresult.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstEitherK chains an Either-returning computation but keeps the original value.
// Useful for validation or side effects that return Either.
//
//go:inline
func MonadChainFirstResultK[R, A, B any](ma ReaderIOResult[R, A], f result.Kleisli[A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapResultK is an alias for MonadChainFirstResultK, executing a Result side effect while preserving the original value.
//
//go:inline
func MonadTapResultK[R, A, B any](ma ReaderIOResult[R, A], f result.Kleisli[A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstResultK returns a function that chains a Result computation while preserving the original value.
// This is an alias for ChainFirstEitherK with more explicit naming.
//
//go:inline
func ChainFirstResultK[R, A, B any](f result.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapResultK is an alias for ChainFirstResultK, executing a Result side effect while preserving the original value.
//
//go:inline
func TapResultK[R, A, B any](f result.Kleisli[A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainReaderK chains a Reader-returning computation into a ReaderIOResult.
// The Reader is automatically lifted into the ReaderIOResult context.
//
//go:inline
func MonadChainReaderK[R, A, B any](ma ReaderIOResult[R, A], f reader.Kleisli[R, A, B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderK returns a function that chains a Reader-returning function into ReaderIOResult.
// This is the curried version of MonadChainReaderK.
//
//go:inline
func ChainReaderK[R, A, B any](f reader.Kleisli[R, A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstReaderK chains a Reader computation but keeps the original value.
// Useful for performing Reader-based side effects while preserving the original value.
//
//go:inline
func MonadChainFirstReaderK[R, A, B any](ma ReaderIOResult[R, A], f reader.Kleisli[R, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapReaderK is an alias for MonadChainFirstReaderK, executing a Reader side effect while preserving the original value.
//
//go:inline
func MonadTapReaderK[R, A, B any](ma ReaderIOResult[R, A], f reader.Kleisli[R, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderK returns a function that chains a Reader computation while preserving the original value.
// This is the curried version of MonadChainFirstReaderK.
//
//go:inline
func ChainFirstReaderK[R, A, B any](f reader.Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderK is an alias for ChainFirstReaderK, executing a Reader side effect while preserving the original value.
//
//go:inline
func TapReaderK[R, A, B any](f reader.Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderOptionK returns a function that chains a ReaderOption-returning function into ReaderIOResult.
// If the ReaderOption is None, the provided error function is called.
//
//go:inline
func ChainReaderOptionK[R, A, B any](onNone Lazy[error]) func(readeroption.Kleisli[R, A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderOptionK chains a ReaderOption computation while preserving the original value.
// If the ReaderOption is None, the provided error function is called.
//
//go:inline
func ChainFirstReaderOptionK[R, A, B any](onNone Lazy[error]) func(readeroption.Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderOptionK is an alias for ChainFirstReaderOptionK, executing a ReaderOption side effect while preserving the original value.
//
//go:inline
func TapReaderOptionK[R, A, B any](onNone Lazy[error]) func(readeroption.Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainReaderK chains a Reader-returning computation into a ReaderIOResult.
// The Reader is automatically lifted into the ReaderIOResult context.
//
//go:inline
func MonadChainReaderEitherK[R, A, B any](ma ReaderIOResult[R, A], f RE.Kleisli[R, error, A, B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderK returns a function that chains a Reader-returning function into ReaderIOResult.
// This is the curried version of MonadChainReaderK.
//
//go:inline
func ChainReaderEitherK[R, A, B any](f RE.Kleisli[R, error, A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstReaderEitherK chains a ReaderEither computation but keeps the original value.
// Useful for performing ReaderEither-based side effects while preserving the original value.
//
//go:inline
func MonadChainFirstReaderEitherK[R, A, B any](ma ReaderIOResult[R, A], f RE.Kleisli[R, error, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapReaderEitherK is an alias for MonadChainFirstReaderEitherK, executing a ReaderEither side effect while preserving the original value.
//
//go:inline
func MonadTapReaderEitherK[R, A, B any](ma ReaderIOResult[R, A], f RE.Kleisli[R, error, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderEitherK returns a function that chains a ReaderEither computation while preserving the original value.
// This is the curried version of MonadChainFirstReaderEitherK.
//
//go:inline
func ChainFirstReaderEitherK[R, A, B any](f RE.Kleisli[R, error, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderEitherK is an alias for ChainFirstReaderEitherK, executing a ReaderEither side effect while preserving the original value.
//
//go:inline
func TapReaderEitherK[R, A, B any](f RE.Kleisli[R, error, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainReaderResultK chains a ReaderResult-returning computation into a ReaderIOResult.
// This is an alias for MonadChainReaderEitherK with more explicit naming.
//
//go:inline
func MonadChainReaderResultK[R, A, B any](ma ReaderIOResult[R, A], f RE.Kleisli[R, error, A, B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderResultK returns a function that chains a ReaderResult-returning function into ReaderIOResult.
// This is an alias for ChainReaderEitherK with more explicit naming.
//
//go:inline
func ChainReaderResultK[R, A, B any](f RE.Kleisli[R, error, A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstReaderResultK chains a ReaderResult computation but keeps the original value.
// This is an alias for MonadChainFirstReaderEitherK with more explicit naming.
//
//go:inline
func MonadChainFirstReaderResultK[R, A, B any](ma ReaderIOResult[R, A], f RE.Kleisli[R, error, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapReaderResultK is an alias for MonadChainFirstReaderResultK, executing a ReaderResult side effect while preserving the original value.
//
//go:inline
func MonadTapReaderResultK[R, A, B any](ma ReaderIOResult[R, A], f RE.Kleisli[R, error, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderResultK returns a function that chains a ReaderResult computation while preserving the original value.
// This is an alias for ChainFirstReaderEitherK with more explicit naming.
//
//go:inline
func ChainFirstReaderResultK[R, A, B any](f RE.Kleisli[R, error, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderResultK is an alias for ChainFirstReaderResultK, executing a ReaderResult side effect while preserving the original value.
//
//go:inline
func TapReaderResultK[R, A, B any](f RE.Kleisli[R, error, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainReaderIOK chains a ReaderIO-returning computation into a ReaderIOResult.
// The ReaderIO is automatically lifted into the ReaderIOResult context (always succeeds).
//
//go:inline
func MonadChainReaderIOK[R, A, B any](ma ReaderIOResult[R, A], f readerio.Kleisli[R, A, B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainReaderIOK returns a function that chains a ReaderIO-returning function into ReaderIOResult.
// This is the curried version of MonadChainReaderIOK.
//
//go:inline
func ChainReaderIOK[R, A, B any](f readerio.Kleisli[R, A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstReaderIOK chains a ReaderIO computation but keeps the original value.
// Useful for performing ReaderIO-based side effects while preserving the original value.
//
//go:inline
func MonadChainFirstReaderIOK[R, A, B any](ma ReaderIOResult[R, A], f readerio.Kleisli[R, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapReaderIOK is an alias for MonadChainFirstReaderIOK, executing a ReaderIO side effect while preserving the original value.
//
//go:inline
func MonadTapReaderIOK[R, A, B any](ma ReaderIOResult[R, A], f readerio.Kleisli[R, A, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstReaderIOK returns a function that chains a ReaderIO computation while preserving the original value.
// This is the curried version of MonadChainFirstReaderIOK.
//
//go:inline
func ChainFirstReaderIOK[R, A, B any](f readerio.Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapReaderIOK is an alias for ChainFirstReaderIOK, executing a ReaderIO side effect while preserving the original value.
//
//go:inline
func TapReaderIOK[R, A, B any](f readerio.Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainIOEitherK chains an IOEither-returning computation into a ReaderIOResult.
// The IOEither is automatically lifted into the ReaderIOResult context.
//
//go:inline
func MonadChainIOEitherK[R, A, B any](ma ReaderIOResult[R, A], f func(A) IOResult[B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainIOEitherK chains an IOEither-returning computation into a ReaderIOResult.
// The IOEither is automatically lifted into the ReaderIOResult context.
//
//go:inline
func MonadChainIOResultK[R, A, B any](ma ReaderIOResult[R, A], f func(A) IOResult[B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainIOEitherK returns a function that chains an IOEither-returning function into ReaderIOResult.
// This is the curried version of MonadChainIOEitherK.
//
//go:inline
func ChainIOEitherK[R, A, B any](f func(A) IOResult[B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainIOEitherK returns a function that chains an IOEither-returning function into ReaderIOResult.
// This is the curried version of MonadChainIOEitherK.
//
//go:inline
func ChainIOResultK[R, A, B any](f func(A) IOResult[B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainIOK chains an IO-returning computation into a ReaderIOResult.
// The IO is automatically lifted into the ReaderIOResult context (always succeeds).
//
//go:inline
func MonadChainIOK[R, A, B any](ma ReaderIOResult[R, A], f func(A) IO[B]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainIOK returns a function that chains an IO-returning function into ReaderIOResult.
// This is the curried version of MonadChainIOK.
//
//go:inline
func ChainIOK[R, A, B any](f func(A) IO[B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK chains an IO computation but keeps the original value.
// Useful for performing IO side effects while preserving the original value.
//
//go:inline
func MonadChainFirstIOK[R, A, B any](ma ReaderIOResult[R, A], f func(A) IO[B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapIOK is an alias for MonadChainFirstIOK, executing an IO side effect while preserving the original value.
//
//go:inline
func MonadTapIOK[R, A, B any](ma ReaderIOResult[R, A], f func(A) IO[B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOK returns a function that chains an IO computation while preserving the original value.
// This is the curried version of MonadChainFirstIOK.
//
//go:inline
func ChainFirstIOK[R, A, B any](f func(A) IO[B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapIOK is an alias for ChainFirstIOK, executing an IO side effect while preserving the original value.
//
//go:inline
func TapIOK[R, A, B any](f func(A) IO[B]) Operator[R, A, A] { _ = "STUB: not implemented"; return nil }

// ChainOptionK returns a function that chains an Option-returning function into ReaderIOResult.
// If the Option is None, the provided error function is called to produce the error value.
//
//go:inline
func ChainOptionK[R, A, B any](onNone Lazy[error]) func(func(A) Option[B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAp applies a function wrapped in a context to a value wrapped in a context.
// Both computations are executed (default behavior may be sequential or parallel depending on implementation).
//
//go:inline
func MonadAp[R, A, B any](fab ReaderIOResult[R, func(A) B], fa ReaderIOResult[R, A]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApSeq applies a function in a context to a value in a context, executing them sequentially.
//
//go:inline
func MonadApSeq[R, A, B any](fab ReaderIOResult[R, func(A) B], fa ReaderIOResult[R, A]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApPar applies a function in a context to a value in a context, executing them in parallel.
//
//go:inline
func MonadApPar[R, A, B any](fab ReaderIOResult[R, func(A) B], fa ReaderIOResult[R, A]) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// Ap returns a function that applies a function in a context to a value in a context.
// This is the curried version of MonadAp.
//
//go:inline
func Ap[B, R, A any](fa ReaderIOResult[R, A]) Operator[R, func(A) B, B] {
	_ = "STUB: not implemented"
	return nil

	// Chain returns a function that sequences computations where the second depends on the first.
	// This is the curried version of MonadChain.
	//
	//go:inline
}

func Chain[R, A, B any](f Kleisli[R, A, B]) Operator[R, A, B] {
	_ = "STUB: not implemented"
	return nil

	// ChainFirst returns a function that sequences computations but keeps the first result.
	// This is the curried version of MonadChainFirst.
	//
	//go:inline
}

func ChainFirst[R, A, B any](f Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil

	// Tap is an alias for ChainFirst, executing a side effect while preserving the original value.
	// The name "Tap" emphasizes the side-effect nature of the operation.
	//
	//go:inline
}

func Tap[R, A, B any](f Kleisli[R, A, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"

	// Right creates a successful ReaderIOResult with the given value.
	//
	//go:inline
	return nil
}

func Right[R, A any](a A) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// Left creates a failed ReaderIOResult with the given error.
//
//go:inline
func Left[R, A any](e error) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// ThrowError creates a failed ReaderIOResult with the given error.
// This is an alias for Left, following the naming convention from other functional libraries.
//
//go:inline
func ThrowError[R, A any](e error) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// Of creates a successful ReaderIOResult with the given value.
// This is the pointed functor operation, lifting a pure value into the ReaderIOResult context.
//
//go:inline
func Of[R, A any](a A) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// Flatten removes one level of nesting from a nested ReaderIOResult.
// Converts ReaderIOResult[R, ReaderIOResult[R, A]] to ReaderIOResult[R, A].
//
//go:inline
func Flatten[R, A any](mma ReaderIOResult[R, ReaderIOResult[R, A]]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil

	// FromEither lifts an Either into a ReaderIOResult context.
	// The Either value is independent of any context or IO effects.
	//
	//go:inline
}

func FromEither[R, A any](t Result[A]) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// FromResult lifts an Either into a ReaderIOResult context.
// The Either value is independent of any context or IO effects.
//
//go:inline
func FromResult[R, A any](t Result[A]) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// RightReader lifts a Reader into a ReaderIOResult, placing the result in the Right side.
//
//go:inline
func RightReader[R, A any](ma Reader[R, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// LeftReader lifts a Reader into a ReaderIOResult, placing the result in the Left (error) side.
//
//go:inline
func LeftReader[A, R any](ma Reader[R, error]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromReader lifts a Reader into a ReaderIOResult context.
// The Reader result is placed in the Right side (success).
//
//go:inline
func FromReader[R, A any](ma Reader[R, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// RightIO lifts an IO into a ReaderIOResult, placing the result in the Right side.
//
//go:inline
func RightIO[R, A any](ma IO[A]) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// LeftIO lifts an IO into a ReaderIOResult, placing the result in the Left (error) side.
//
//go:inline
func LeftIO[R, A any](ma IO[error]) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// FromIO lifts an IO into a ReaderIOResult context.
// The IO result is placed in the Right side (success).
//
//go:inline
func FromIO[R, A any](ma IO[A]) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// FromIOEither lifts an IOEither into a ReaderIOResult context.
// The computation becomes independent of any reader context.
//
//go:inline
func FromIOEither[R, A any](ma IOResult[A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromIOEither lifts an IOEither into a ReaderIOResult context.
// The computation becomes independent of any reader context.
//
//go:inline
func FromIOResult[R, A any](ma IOResult[A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromReaderEither lifts a ReaderEither into a ReaderIOResult context.
// The Either result is lifted into an IO effect.
//
//go:inline
func FromReaderEither[R, A any](ma RE.ReaderEither[R, error, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// Ask returns a ReaderIOResult that retrieves the current context.
// Useful for accessing configuration or dependencies.
//
//go:inline
func Ask[R any]() ReaderIOResult[R, R] { _ = "STUB: not implemented"; return nil }

// Asks returns a ReaderIOResult that retrieves a value derived from the context.
// This is useful for extracting specific fields from a configuration object.
//
//go:inline
func Asks[R, A any](r Reader[R, A]) ReaderIOResult[R, A] { _ = "STUB: not implemented"; return nil }

// FromOption converts an Option to a ReaderIOResult.
// If the Option is None, the provided function is called to produce the error.
//
//go:inline
func FromOption[R, A any](onNone Lazy[error]) Kleisli[R, Option[A], A] {
	_ = "STUB: not implemented"
	return nil
}

// FromPredicate creates a ReaderIOResult from a predicate.
// If the predicate returns false, the onFalse function is called to produce the error.
//
//go:inline
func FromPredicate[R, A any](pred func(A) bool, onFalse func(A) error) Kleisli[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// Fold handles both success and error cases, producing a ReaderIO.
// This is useful for converting a ReaderIOResult into a ReaderIO by handling all cases.
//
//go:inline
func Fold[R, A, B any](onLeft readerio.Kleisli[R, error, B], onRight func(A) ReaderIO[R, B]) func(ReaderIOResult[R, A]) ReaderIO[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// GetOrElse provides a default value in case of error.
// The default is computed lazily via a ReaderIO.
//
//go:inline
func GetOrElse[R, A any](onLeft readerio.Kleisli[R, error, A]) func(ReaderIOResult[R, A]) ReaderIO[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// OrElse tries an alternative computation if the first one fails.
//
//go:inline
func OrElse[R, A any](onLeft Kleisli[R, error, A]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil

	// OrLeft transforms the error using a ReaderIO if the computation fails.
	// The success value is preserved unchanged.
	//
	//go:inline
}

func OrLeft[A, R, E any](onLeft readerio.Kleisli[R, error, E]) func(ReaderIOResult[R, A]) RIOE.ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadBiMap applies two functions: one to the error, one to the success value.
// This allows transforming both channels simultaneously.
//
//go:inline
func MonadBiMap[R, E, A, B any](fa ReaderIOResult[R, A], f func(error) E, g func(A) B) RIOE.ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap returns a function that maps over both the error and success channels.
// This is the curried version of MonadBiMap.
//
//go:inline
func BiMap[R, E, A, B any](f func(error) E, g func(A) B) func(ReaderIOResult[R, A]) RIOE.ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil

	// Swap exchanges the error and success types.
	// Left becomes Right and Right becomes Left.
	//
	//go:inline
}

func Swap[R, A any](val ReaderIOResult[R, A]) RIOE.ReaderIOEither[R, A, error] {
	_ = "STUB: not implemented"
	return nil

	// Defer creates a ReaderIOResult lazily via a generator function.
	// The generator is called each time the ReaderIOResult is executed.
	//
	//go:inline
}

func Defer[R, A any](gen Lazy[ReaderIOResult[R, A]]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil

	// TryCatch wraps a function that returns (value, error) into a ReaderIOResult.
	// The onThrow function converts the error into the desired error type.
	//
	//go:inline
}

func TryCatch[R, A any](f func(R) func() (A, error), onThrow Endomorphism[error]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAlt tries the first computation, and if it fails, tries the second.
// This implements the Alternative pattern for error recovery.
//
//go:inline
func MonadAlt[R, A any](first ReaderIOResult[R, A], second Lazy[ReaderIOResult[R, A]]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// Alt returns a function that tries an alternative computation if the first fails.
// This is the curried version of MonadAlt.
//
//go:inline
func Alt[R, A any](second Lazy[ReaderIOResult[R, A]]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil

	// Memoize computes the value of the ReaderIOResult lazily but exactly once.
	// The context used is from the first call. Do not use if the value depends on the context.
	//
	//go:inline
}

func Memoize[R, A any](rdr ReaderIOResult[R, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil

	// MonadFlap applies a value to a function wrapped in a context.
	// This is the reverse of Ap - the value is fixed and the function varies.
	//
	//go:inline
}

func MonadFlap[R, B, A any](fab ReaderIOResult[R, func(A) B], a A) ReaderIOResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// Flap returns a function that applies a fixed value to a function in a context.
// This is the curried version of MonadFlap.
//
//go:inline
func Flap[R, B, A any](a A) Operator[R, func(A) B, B] { _ = "STUB: not implemented"; return nil }

// MonadMapLeft applies a function to the error value, leaving success unchanged.
//
//go:inline
func MonadMapLeft[R, E, A any](fa ReaderIOResult[R, A], f func(error) E) RIOE.ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MapLeft returns a function that transforms the error channel.
// This is the curried version of MonadMapLeft.
//
//go:inline
func MapLeft[R, A, E any](f func(error) E) func(ReaderIOResult[R, A]) RIOE.ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Local runs a computation with a modified context.
// The function f transforms the context before passing it to the computation.
// This is similar to Contravariant's contramap operation.
//
//go:inline
func Local[A, R1, R2 any](f func(R2) R1) func(ReaderIOResult[R1, A]) ReaderIOResult[R2, A] {
	_ = "STUB: not implemented"
	return nil
}

// Read executes a ReaderIOResult by providing a concrete environment value.
// This converts a ReaderIOResult[R, A] into an IOResult[A] by supplying the R value.
//
//go:inline
func Read[A, R any](r R) func(ReaderIOResult[R, A]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainLeft chains a computation on the error channel, allowing error recovery or transformation.
// If the computation is successful (Right), it passes through unchanged.
//
//go:inline
func MonadChainLeft[R, A any](fa ReaderIOResult[R, A], f Kleisli[R, error, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainLeft returns a function that chains a computation on the error channel.
// This is the curried version of MonadChainLeft.
//
//go:inline
func ChainLeft[R, A any](f Kleisli[R, error, A]) func(ReaderIOResult[R, A]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil

	// MonadChainFirstLeft chains an error-handling computation but preserves the original error.
	// Useful for logging or side effects on errors without changing the error value.
	//
	//go:inline
}

func MonadChainFirstLeft[A, R, B any](ma ReaderIOResult[R, A], f Kleisli[R, error, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadTapLeft is an alias for MonadChainFirstLeft, executing a side effect on errors while preserving the original error.
//
//go:inline
func MonadTapLeft[A, R, B any](ma ReaderIOResult[R, A], f Kleisli[R, error, B]) ReaderIOResult[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstLeft returns a function that chains an error-handling computation while preserving the original error.
// This is the curried version of MonadChainFirstLeft.
//
//go:inline
func ChainFirstLeft[A, R, B any](f Kleisli[R, error, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstLeftIOK chains an IO computation on errors while preserving the original error.
// Useful for IO-based error logging or side effects.
//
//go:inline
func ChainFirstLeftIOK[A, R, B any](f io.Kleisli[error, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// TapLeft is an alias for ChainFirstLeft, executing a side effect on errors while preserving the original error.
//
//go:inline
func TapLeft[A, R, B any](f Kleisli[R, error, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil

	// TapLeftIOK is an alias for ChainFirstLeftIOK, executing an IO side effect on errors while preserving the original error.
	//
	//go:inline
}

func TapLeftIOK[A, R, B any](f io.Kleisli[error, B]) Operator[R, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some delay
//
//go:inline
func Delay[R, A any](delay time.Duration) Operator[R, A, A] { _ = "STUB: not implemented"; return nil }

// After creates an operation that passes after the given [time.Time]
//
//go:inline
func After[R, A any](timestamp time.Time) Operator[R, A, A] { _ = "STUB: not implemented"; return nil }

// ReadIOEither executes a ReaderIOResult computation by providing an environment
// obtained from an IOResult. This function bridges the gap between IOResult-based
// environment acquisition and ReaderIOResult-based computations.
//
// The function first executes the IOResult[R] to obtain the environment (or an error),
// then uses that environment to run the ReaderIOResult[R, A] computation.
//
// Type parameters:
//   - A: The success value type of the ReaderIOResult computation
//   - R: The environment/context type required by the ReaderIOResult
//
// Parameters:
//   - r: An IOResult[R] that produces the environment (or an error)
//
// Returns:
//   - A function that takes a ReaderIOResult[R, A] and returns IOResult[A]
//
// Example:
//
//	type Config struct { BaseURL string }
//
//	// Get config from environment with potential error
//	getConfig := func() IOResult[Config] {
//	    return func() Result[Config] {
//	        // Load config, may fail
//	        return result.Of(Config{BaseURL: "https://api.example.com"})
//	    }
//	}
//
//	// A computation that needs config
//	fetchUser := func(id int) ReaderIOResult[Config, User] {
//	    return func(cfg Config) IOResult[User] {
//	        return func() Result[User] {
//	            // Use cfg.BaseURL to fetch user
//	            return result.Of(User{ID: id})
//	        }
//	    }
//	}
//
//	// Execute the computation with the config
//	result := ReadIOEither[User](getConfig())(fetchUser(123))()
//
//go:inline
func ReadIOEither[A, R any](r IOResult[R]) func(ReaderIOResult[R, A]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

// ReadIOResult executes a ReaderIOResult computation by providing an environment
// obtained from an IOResult. This is an alias for ReadIOEither with more explicit naming.
//
// The function first executes the IOResult[R] to obtain the environment (or an error),
// then uses that environment to run the ReaderIOResult[R, A] computation.
//
// Type parameters:
//   - A: The success value type of the ReaderIOResult computation
//   - R: The environment/context type required by the ReaderIOResult
//
// Parameters:
//   - r: An IOResult[R] that produces the environment (or an error)
//
// Returns:
//   - A function that takes a ReaderIOResult[R, A] and returns IOResult[A]
//
// Example:
//
//	type Database struct { ConnectionString string }
//
//	// Get database connection with potential error
//	getDB := func() IOResult[Database] {
//	    return func() Result[Database] {
//	        return result.Of(Database{ConnectionString: "localhost:5432"})
//	    }
//	}
//
//	// Query that needs database
//	queryUsers := ReaderIOResult[Database, []User] {
//	    return func(db Database) IOResult[[]User] {
//	        return func() Result[[]User] {
//	            // Execute query using db
//	            return result.Of([]User{})
//	        }
//	    }
//	}
//
//	// Execute query with database
//	users := ReadIOResult[[]User](getDB())(queryUsers)()
//
//go:inline
func ReadIOResult[A, R any](r IOResult[R]) func(ReaderIOResult[R, A]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}

// ReadIO executes a ReaderIOResult computation by providing an environment
// obtained from an IO computation. Unlike ReadIOEither/ReadIOResult, the environment
// acquisition cannot fail (it's a pure IO, not IOResult).
//
// The function first executes the IO[R] to obtain the environment,
// then uses that environment to run the ReaderIOResult[R, A] computation.
//
// Type parameters:
//   - A: The success value type of the ReaderIOResult computation
//   - R: The environment/context type required by the ReaderIOResult
//
// Parameters:
//   - r: An IO[R] that produces the environment (cannot fail)
//
// Returns:
//   - A function that takes a ReaderIOResult[R, A] and returns IOResult[A]
//
// Example:
//
//	type Logger struct { Level string }
//
//	// Get logger (always succeeds)
//	getLogger := func() IO[Logger] {
//	    return func() Logger {
//	        return Logger{Level: "INFO"}
//	    }
//	}
//
//	// Log operation that may fail
//	logMessage := func(msg string) ReaderIOResult[Logger, string] {
//	    return func(logger Logger) IOResult[string] {
//	        return func() Result[string] {
//	            // Log with logger, may fail
//	            return result.Of(fmt.Sprintf("[%s] %s", logger.Level, msg))
//	        }
//	    }
//	}
//
//	// Execute logging with logger
//	logged := ReadIO[string](getLogger())(logMessage("Hello"))()
//
//go:inline
func ReadIO[A, R any](r IO[R]) func(ReaderIOResult[R, A]) IOResult[A] {
	_ = "STUB: not implemented"
	return nil
}
