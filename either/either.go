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

// package either implements the Either monad
//
// A data type that can be of either of two types but not both. This is
// typically used to carry an error or a return value
package either

import (
	L "github.com/IBM/fp-go/lazy"
	O "github.com/IBM/fp-go/option"
)

// Of is equivalent to [Right]
func Of[E, A any](value A) Either[E, A] { _ = "STUB: not implemented"; return nil }

func FromIO[E any, IO ~func() A, A any](f IO) Either[E, A] { _ = "STUB: not implemented"; return nil }

func MonadAp[B, E, A any](fab Either[E, func(a A) B], fa Either[E, A]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, E, A any](fa Either[E, A]) func(fab Either[E, func(a A) B]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[E, A, B any](fa Either[E, A], f func(a A) B) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadBiMap[E1, E2, A, B any](fa Either[E1, A], f func(E1) E2, g func(a A) B) Either[E2, B] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
func BiMap[E1, E2, A, B any](f func(E1) E2, g func(a A) B) func(Either[E1, A]) Either[E2, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[E, A, B any](fa Either[E, A], b B) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MapTo[E, A, B any](b B) func(Either[E, A]) Either[E, B] { _ = "STUB: not implemented"; return nil }

func MonadMapLeft[E1, A, E2 any](fa Either[E1, A], f func(E1) E2) Either[E2, A] {
	_ = "STUB: not implemented"
	return nil
}

func Map[E, A, B any](f func(a A) B) func(fa Either[E, A]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// MapLeft applies a mapping function to the error channel
func MapLeft[A, E1, E2 any](f func(E1) E2) func(fa Either[E1, A]) Either[E2, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[E, A, B any](fa Either[E, A], f func(a A) Either[E, B]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[E, A, B any](ma Either[E, A], f func(a A) Either[E, B]) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainTo[A, E, B any](_ Either[E, A], mb Either[E, B]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainOptionK[A, B, E any](onNone func() E, ma Either[E, A], f func(A) O.Option[B]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[A, B, E any](onNone func() E) func(func(A) O.Option[B]) func(Either[E, A]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainTo[A, E, B any](mb Either[E, B]) func(Either[E, A]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[E, A, B any](f func(a A) Either[E, B]) func(Either[E, A]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirst[E, A, B any](f func(a A) Either[E, B]) func(Either[E, A]) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[E, A any](mma Either[E, Either[E, A]]) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func TryCatch[FE func(error) E, E, A any](val A, err error, onThrow FE) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func TryCatchError[A any](val A, err error) Either[error, A] { _ = "STUB: not implemented"; return nil }

func Sequence2[E, T1, T2, R any](f func(T1, T2) Either[E, R]) func(Either[E, T1], Either[E, T2]) Either[E, R] {
	_ = "STUB: not implemented"
	return nil
}

func Sequence3[E, T1, T2, T3, R any](f func(T1, T2, T3) Either[E, R]) func(Either[E, T1], Either[E, T2], Either[E, T3]) Either[E, R] {
	_ = "STUB: not implemented"
	return nil
}

func FromOption[A, E any](onNone func() E) func(O.Option[A]) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ToOption[E, A any](ma Either[E, A]) O.Option[A] { _ = "STUB: not implemented"; return nil }

func FromError[A any](f func(a A) error) func(A) Either[error, A] {
	_ = "STUB: not implemented"
	return nil
}

func ToError[A any](e Either[error, A]) error { _ = "STUB: not implemented"; return nil }

func Fold[E, A, B any](onLeft func(E) B, onRight func(A) B) func(Either[E, A]) B {
	_ = "STUB: not implemented"
	return nil
}

// UnwrapError converts an Either into the idiomatic tuple
func UnwrapError[A any](ma Either[error, A]) (A, error) {
	_ = "STUB: not implemented"
	return *new(A), nil
}

func FromPredicate[E, A any](pred func(A) bool, onFalse func(A) E) func(A) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromNillable[A, E any](e E) func(*A) Either[E, *A] { _ = "STUB: not implemented"; return nil }

func GetOrElse[E, A any](onLeft func(E) A) func(Either[E, A]) A {
	_ = "STUB: not implemented"
	return nil
}

func Reduce[E, A, B any](f func(B, A) B, initial B) func(Either[E, A]) B {
	_ = "STUB: not implemented"
	return nil
}

func AltW[E, E1, A any](that L.Lazy[Either[E1, A]]) func(Either[E, A]) Either[E1, A] {
	_ = "STUB: not implemented"
	return nil
}

func Alt[E, A any](that L.Lazy[Either[E, A]]) func(Either[E, A]) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[E, A any](onLeft func(e E) Either[E, A]) func(Either[E, A]) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func ToType[A, E any](onError func(any) E) func(any) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Memoize[E, A any](val Either[E, A]) Either[E, A] { _ = "STUB: not implemented"; return nil }

func MonadSequence2[E, T1, T2, R any](e1 Either[E, T1], e2 Either[E, T2], f func(T1, T2) Either[E, R]) Either[E, R] {
	_ = "STUB: not implemented"
	return nil
}

func MonadSequence3[E, T1, T2, T3, R any](e1 Either[E, T1], e2 Either[E, T2], e3 Either[E, T3], f func(T1, T2, T3) Either[E, R]) Either[E, R] {
	_ = "STUB: not implemented"
	return nil
}

// Swap changes the order of type parameters
func Swap[E, A any](val Either[E, A]) Either[A, E] { _ = "STUB: not implemented"; return nil }

func MonadFlap[E, B, A any](fab Either[E, func(A) B], a A) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[E, B, A any](a A) func(Either[E, func(A) B]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAlt[E, A any](fa Either[E, A], that L.Lazy[Either[E, A]]) Either[E, A] {
	_ = "STUB: not implemented"
	return nil
}
