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
	ET "github.com/IBM/fp-go/either"
	"github.com/IBM/fp-go/io"
	IOE "github.com/IBM/fp-go/ioeither"
	L "github.com/IBM/fp-go/lazy"
	O "github.com/IBM/fp-go/option"
	RD "github.com/IBM/fp-go/reader"
	RE "github.com/IBM/fp-go/readereither"
	RIO "github.com/IBM/fp-go/readerio"
)

type ReaderIOEither[R, E, A any] RD.Reader[R, IOE.IOEither[E, A]]

// MakeReader constructs an instance of a reader
func MakeReader[R, E, A any](f func(R) IOE.IOEither[E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFromReaderIO[R, E, A any](a A, f func(A) RIO.ReaderIO[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromReaderIO[R, E, A any](f func(A) RIO.ReaderIO[R, A]) func(A) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func RightReaderIO[R, E, A any](ma RIO.ReaderIO[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func LeftReaderIO[A, R, E any](me RIO.ReaderIO[R, E]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[R, E, A, B any](fa ReaderIOEither[R, E, A], f func(A) B) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[R, E, A, B any](f func(A) B) func(fa ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[R, E, A, B any](fa ReaderIOEither[R, E, A], b B) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MapTo[R, E, A, B any](b B) func(ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[R, E, A, B any](fa ReaderIOEither[R, E, A], f func(A) ReaderIOEither[R, E, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[R, E, A, B any](fa ReaderIOEither[R, E, A], f func(A) ReaderIOEither[R, E, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f func(A) ET.Either[E, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[R, E, A, B any](f func(A) ET.Either[E, B]) func(ma ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f func(A) ET.Either[E, B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirstEitherK[R, E, A, B any](f func(A) ET.Either[E, B]) func(ma ReaderIOEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainReaderK[R, E, A, B any](ma ReaderIOEither[R, E, A], f func(A) RD.Reader[R, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainReaderK[E, R, A, B any](f func(A) RD.Reader[R, B]) func(ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOEitherK[R, E, A, B any](ma ReaderIOEither[R, E, A], f func(A) IOE.IOEither[E, B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOEitherK[R, E, A, B any](f func(A) IOE.IOEither[E, B]) func(ma ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOK[R, E, A, B any](ma ReaderIOEither[R, E, A], f func(A) io.IO[B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOK[R, E, A, B any](f func(A) io.IO[B]) func(ma ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstIOK[R, E, A, B any](ma ReaderIOEither[R, E, A], f func(A) io.IO[B]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirstIOK[R, E, A, B any](f func(A) io.IO[B]) func(ma ReaderIOEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[R, A, B, E any](onNone func() E) func(func(A) O.Option[B]) func(ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[R, E, A, B any](fab ReaderIOEither[R, E, func(A) B], fa ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, R, E, A any](fa ReaderIOEither[R, E, A]) func(fab ReaderIOEither[R, E, func(A) B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[R, E, A, B any](f func(A) ReaderIOEither[R, E, B]) func(fa ReaderIOEither[R, E, A]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirst[R, E, A, B any](f func(A) ReaderIOEither[R, E, B]) func(fa ReaderIOEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Right[R, E, A any](a A) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

func Left[R, A, E any](e E) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

func ThrowError[R, A, E any](e E) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

// Of returns a Reader with a fixed value
func Of[R, E, A any](a A) ReaderIOEither[R, E, A] { _ = "STUB: not implemented"; return nil }

func Flatten[R, E, A any](mma ReaderIOEither[R, E, ReaderIOEither[R, E, A]]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromEither[R, E, A any](t ET.Either[E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func RightReader[E, R, A any](ma RD.Reader[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func LeftReader[A, R, E any](ma RD.Reader[R, E]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromReader[E, R, A any](ma RD.Reader[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func RightIO[R, E, A any](ma io.IO[A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func LeftIO[R, A, E any](ma io.IO[E]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromIO[R, E, A any](ma io.IO[A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromIOEither[R, E, A any](ma IOE.IOEither[E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromReaderEither[R, E, A any](ma RE.ReaderEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Ask[R, E any]() ReaderIOEither[R, E, R] { _ = "STUB: not implemented"; return nil }

func Asks[E, R, A any](r RD.Reader[R, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromOption[R, A, E any](onNone func() E) func(O.Option[A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[R, E, A any](pred func(A) bool, onFalse func(A) E) func(A) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Fold[R, E, A, B any](onLeft func(E) RIO.ReaderIO[R, B], onRight func(A) RIO.ReaderIO[R, B]) func(ReaderIOEither[R, E, A]) RIO.ReaderIO[R, B] {
	_ = "STUB: not implemented"
	return nil
}

func GetOrElse[R, E, A any](onLeft func(E) RIO.ReaderIO[R, A]) func(ReaderIOEither[R, E, A]) RIO.ReaderIO[R, A] {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[R, E1, A, E2 any](onLeft func(E1) ReaderIOEither[R, E2, A]) func(ReaderIOEither[R, E1, A]) ReaderIOEither[R, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

func OrLeft[A, E1, R, E2 any](onLeft func(E1) RIO.ReaderIO[R, E2]) func(ReaderIOEither[R, E1, A]) ReaderIOEither[R, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadBiMap[R, E1, E2, A, B any](fa ReaderIOEither[R, E1, A], f func(E1) E2, g func(A) B) ReaderIOEither[R, E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
func BiMap[R, E1, E2, A, B any](f func(E1) E2, g func(A) B) func(ReaderIOEither[R, E1, A]) ReaderIOEither[R, E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
func Swap[R, E, A any](val ReaderIOEither[R, E, A]) ReaderIOEither[R, A, E] {
	_ = "STUB: not implemented"
	return nil
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[R, E, A any](gen L.Lazy[ReaderIOEither[R, E, A]]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// TryCatch wraps a reader returning a tuple as an error into ReaderIOEither
func TryCatch[R, E, A any](f func(R) func() (A, error), onThrow func(error) E) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAlt identifies an associative operation on a type constructor.
func MonadAlt[R, E, A any](first ReaderIOEither[R, E, A], second L.Lazy[ReaderIOEither[R, E, A]]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Alt identifies an associative operation on a type constructor.
func Alt[R, E, A any](second L.Lazy[ReaderIOEither[R, E, A]]) func(ReaderIOEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil

	// Memoize computes the value of the provided [ReaderIOEither] monad lazily but exactly once
	// The context used to compute the value is the context of the first call, so do not use this
	// method if the value has a functional dependency on the content of the context
}

func Memoize[
	R, E, A any](rdr ReaderIOEither[R, E, A]) ReaderIOEither[R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[R, E, B, A any](fab ReaderIOEither[R, E, func(A) B], a A) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[R, E, B, A any](a A) func(ReaderIOEither[R, E, func(A) B]) ReaderIOEither[R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapLeft[R, E1, E2, A any](fa ReaderIOEither[R, E1, A], f func(E1) E2) ReaderIOEither[R, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// MapLeft applies a mapping function to the error channel
func MapLeft[R, A, E1, E2 any](f func(E1) E2) func(ReaderIOEither[R, E1, A]) ReaderIOEither[R, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// Local changes the value of the local context during the execution of the action `ma` (similar to `Contravariant`'s
// `contramap`).
func Local[R1, R2, E, A any](f func(R2) R1) func(ReaderIOEither[R1, E, A]) ReaderIOEither[R2, E, A] {
	_ = "STUB: not implemented"
	return nil
}
