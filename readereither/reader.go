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

package readereither

import (
	ET "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
	R "github.com/IBM/fp-go/reader"
)

type ReaderEither[E, L, A any] R.Reader[E, ET.Either[L, A]]

func MakeReaderEither[L, E, A any](f func(E) ET.Either[L, A]) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromEither[E, L, A any](e ET.Either[L, A]) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func RightReader[L, E, A any](r R.Reader[E, A]) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func LeftReader[A, E, L any](l R.Reader[E, L]) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func Left[E, A, L any](l L) ReaderEither[E, L, A] { _ = "STUB: not implemented"; return nil }

func Right[E, L, A any](r A) ReaderEither[E, L, A] { _ = "STUB: not implemented"; return nil }

func FromReader[E, L, A any](r R.Reader[E, A]) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[E, L, A, B any](fa ReaderEither[E, L, A], f func(A) B) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[E, L, A, B any](f func(A) B) func(ReaderEither[E, L, A]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[E, L, A, B any](ma ReaderEither[E, L, A], f func(A) ReaderEither[E, L, B]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[E, L, A, B any](f func(A) ReaderEither[E, L, B]) func(ReaderEither[E, L, A]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func Of[E, L, A any](a A) ReaderEither[E, L, A] { _ = "STUB: not implemented"; return nil }

func MonadAp[E, L, A, B any](fab ReaderEither[E, L, func(A) B], fa ReaderEither[E, L, A]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, E, L, A any](fa ReaderEither[E, L, A]) func(ReaderEither[E, L, func(A) B]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[E, L, A any](pred func(A) bool, onFalse func(A) L) func(A) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func Fold[E, L, A, B any](onLeft func(L) R.Reader[E, B], onRight func(A) R.Reader[E, B]) func(ReaderEither[E, L, A]) R.Reader[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func GetOrElse[E, L, A any](onLeft func(L) R.Reader[E, A]) func(ReaderEither[E, L, A]) R.Reader[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[E, L1, A, L2 any](onLeft func(L1) ReaderEither[E, L2, A]) func(ReaderEither[E, L1, A]) ReaderEither[E, L2, A] {
	_ = "STUB: not implemented"
	return nil
}

func OrLeft[A, L1, E, L2 any](onLeft func(L1) R.Reader[E, L2]) func(ReaderEither[E, L1, A]) ReaderEither[E, L2, A] {
	_ = "STUB: not implemented"
	return nil
}

func Ask[E, L any]() ReaderEither[E, L, E] { _ = "STUB: not implemented"; return nil }

func Asks[L, E, A any](r R.Reader[E, A]) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[A, B, L, E any](ma ReaderEither[E, L, A], f func(A) ET.Either[L, B]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[A, B, L, E any](f func(A) ET.Either[L, B]) func(ma ReaderEither[E, L, A]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[E, A, B, L any](onNone func() L) func(func(A) O.Option[B]) func(ReaderEither[E, L, A]) ReaderEither[E, L, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[E, L, A any](mma ReaderEither[E, L, ReaderEither[E, L, A]]) ReaderEither[E, L, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadBiMap[E, E1, E2, A, B any](fa ReaderEither[E, E1, A], f func(E1) E2, g func(A) B) ReaderEither[E, E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
func BiMap[E, E1, E2, A, B any](f func(E1) E2, g func(A) B) func(ReaderEither[E, E1, A]) ReaderEither[E, E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// Local changes the value of the local context during the execution of the action `ma` (similar to `Contravariant`'s
// `contramap`).
func Local[E, A, R2, R1 any](f func(R2) R1) func(ReaderEither[R1, E, A]) ReaderEither[R2, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Read applies a context to a reader to obtain its value
func Read[E1, A, E any](e E) func(ReaderEither[E, E1, A]) ET.Either[E1, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[L, E, A, B any](fab ReaderEither[L, E, func(A) B], a A) ReaderEither[L, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[L, E, B, A any](a A) func(ReaderEither[L, E, func(A) B]) ReaderEither[L, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapLeft[C, E1, E2, A any](fa ReaderEither[C, E1, A], f func(E1) E2) ReaderEither[C, E2, A] {
	_ = "STUB: not implemented"
	return nil
}

// MapLeft applies a mapping function to the error channel
func MapLeft[C, E1, E2, A any](f func(E1) E2) func(ReaderEither[C, E1, A]) ReaderEither[C, E2, A] {
	_ = "STUB: not implemented"
	return nil
}
