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

package readerioeither

import (
	"context"
	"time"

	ET "github.com/IBM/fp-go/either"
	IO "github.com/IBM/fp-go/io"
	IOE "github.com/IBM/fp-go/ioeither"
	L "github.com/IBM/fp-go/lazy"
	O "github.com/IBM/fp-go/option"
	RIO "github.com/IBM/fp-go/readerio"
)

func FromEither[A any](e ET.Either[error, A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func Left[A any](l error) ReaderIOEither[A] { _ = "STUB: not implemented"; return nil }

func Right[A any](r A) ReaderIOEither[A] { _ = "STUB: not implemented"; return nil }

func MonadMap[A, B any](fa ReaderIOEither[A], f func(A) B) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[A, B any](f func(A) B) func(ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[A, B any](fa ReaderIOEither[A], b B) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MapTo[A, B any](b B) func(ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[A, B any](ma ReaderIOEither[A], f func(A) ReaderIOEither[B]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A, B any](f func(A) ReaderIOEither[B]) func(ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[A, B any](ma ReaderIOEither[A], f func(A) ReaderIOEither[B]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirst[A, B any](f func(A) ReaderIOEither[B]) func(ReaderIOEither[A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func Of[A any](a A) ReaderIOEither[A] { _ = "STUB: not implemented"; return nil }

// MonadAp implements the `Ap` function for a reader with context. It creates a sub-context that will
// be canceled if any of the input operations errors out or
func MonadAp[B, A any](fab ReaderIOEither[func(A) B], fa ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, A any](fa ReaderIOEither[A]) func(ReaderIOEither[func(A) B]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[A any](pred func(A) bool, onFalse func(A) error) func(A) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[A any](onLeft func(error) ReaderIOEither[A]) func(ReaderIOEither[A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func Ask() ReaderIOEither[context.Context] { _ = "STUB: not implemented"; return nil }

func MonadChainEitherK[A, B any](ma ReaderIOEither[A], f func(A) ET.Either[error, B]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[A, B any](f func(A) ET.Either[error, B]) func(ma ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstEitherK[A, B any](ma ReaderIOEither[A], f func(A) ET.Either[error, B]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirstEitherK[A, B any](f func(A) ET.Either[error, B]) func(ma ReaderIOEither[A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[A, B any](onNone func() error) func(func(A) O.Option[B]) func(ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func FromIOEither[A any](t IOE.IOEither[error, A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func FromIO[A any](t IO.IO[A]) ReaderIOEither[A] { _ = "STUB: not implemented"; return nil }

func FromLazy[A any](t L.Lazy[A]) ReaderIOEither[A] { _ = "STUB: not implemented"; return nil }

// Never returns a 'ReaderIOEither' that never returns, except if its context gets canceled
func Never[A any]() ReaderIOEither[A] { _ = "STUB: not implemented"; return nil }

func MonadChainIOK[A, B any](ma ReaderIOEither[A], f func(A) IO.IO[B]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOK[A, B any](f func(A) IO.IO[B]) func(ma ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstIOK[A, B any](ma ReaderIOEither[A], f func(A) IO.IO[B]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirstIOK[A, B any](f func(A) IO.IO[B]) func(ma ReaderIOEither[A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOEitherK[A, B any](f func(A) IOE.IOEither[error, B]) func(ma ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some delay
func Delay[A any](delay time.Duration) func(ma ReaderIOEither[A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

// Timer will return the current time after an initial delay
func Timer(delay time.Duration) ReaderIOEither[time.Time] { _ = "STUB: not implemented"; return nil }

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[A any](gen L.Lazy[ReaderIOEither[A]]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

// TryCatch wraps a reader returning a tuple as an error into ReaderIOEither
func TryCatch[A any](f func(context.Context) func() (A, error)) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAlt identifies an associative operation on a type constructor
func MonadAlt[A any](first ReaderIOEither[A], second L.Lazy[ReaderIOEither[A]]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

// Alt identifies an associative operation on a type constructor
func Alt[A any](second L.Lazy[ReaderIOEither[A]]) func(ReaderIOEither[A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil

	// Memoize computes the value of the provided [ReaderIOEither] monad lazily but exactly once
	// The context used to compute the value is the context of the first call, so do not use this
	// method if the value has a functional dependency on the content of the context
}

func Memoize[A any](rdr ReaderIOEither[A]) ReaderIOEither[A] { _ = "STUB: not implemented"; return nil }

// Flatten converts a nested [ReaderIOEither] into a [ReaderIOEither]
func Flatten[A any](rdr ReaderIOEither[ReaderIOEither[A]]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[B, A any](fab ReaderIOEither[func(A) B], a A) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[B, A any](a A) func(ReaderIOEither[func(A) B]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Fold[A, B any](onLeft func(error) ReaderIOEither[B], onRight func(A) ReaderIOEither[B]) func(ReaderIOEither[A]) ReaderIOEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func GetOrElse[A any](onLeft func(error) RIO.ReaderIO[context.Context, A]) func(ReaderIOEither[A]) RIO.ReaderIO[context.Context, A] {
	_ = "STUB: not implemented"
	return nil
}

func OrLeft[A any](onLeft func(error) RIO.ReaderIO[context.Context, error]) func(ReaderIOEither[A]) ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}
