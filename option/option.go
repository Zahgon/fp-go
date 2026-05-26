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

// package option implements the Option monad, a data type that can have a defined value or none
package option

func fromPredicate[A any](a A, pred func(A) bool) Option[A] { _ = "STUB: not implemented"; return nil }

func FromPredicate[A any](pred func(A) bool) func(A) Option[A] {
	_ = "STUB: not implemented"
	return nil
}

func FromNillable[A any](a *A) Option[*A] { _ = "STUB: not implemented"; return nil }

func FromValidation[A, B any](f func(A) (B, bool)) func(A) Option[B] {
	_ = "STUB: not implemented"
	return nil

	// MonadAp is the applicative functor of Option
}

func MonadAp[B, A any](fab Option[func(A) B], fa Option[A]) Option[B] {
	_ = "STUB: not implemented"
	return nil
}

// Ap is the applicative functor of Option
func Ap[B, A any](fa Option[A]) func(Option[func(A) B]) Option[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[A, B any](fa Option[A], f func(A) B) Option[B] { _ = "STUB: not implemented"; return nil }

func Map[A, B any](f func(a A) B) func(Option[A]) Option[B] { _ = "STUB: not implemented"; return nil }

func MonadMapTo[A, B any](fa Option[A], b B) Option[B] { _ = "STUB: not implemented"; return nil }

func MapTo[A, B any](b B) func(Option[A]) Option[B] { _ = "STUB: not implemented"; return nil }

func TryCatch[A any](f func() (A, error)) Option[A] { _ = "STUB: not implemented"; return nil }

func Fold[A, B any](onNone func() B, onSome func(a A) B) func(ma Option[A]) B {
	_ = "STUB: not implemented"
	return nil
}

func MonadGetOrElse[A any](fa Option[A], onNone func() A) A {
	_ = "STUB: not implemented"
	return *new(A)
}

func GetOrElse[A any](onNone func() A) func(Option[A]) A { _ = "STUB: not implemented"; return nil }

func MonadChain[A, B any](fa Option[A], f func(A) Option[B]) Option[B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A, B any](f func(A) Option[B]) func(Option[A]) Option[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainTo[A, B any](_ Option[A], mb Option[B]) Option[B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainTo[A, B any](mb Option[B]) func(Option[A]) Option[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[A, B any](ma Option[A], f func(A) Option[B]) Option[A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirst[A, B any](f func(A) Option[B]) func(Option[A]) Option[A] {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[A any](mma Option[Option[A]]) Option[A] { _ = "STUB: not implemented"; return nil }

func MonadAlt[A any](fa Option[A], that func() Option[A]) Option[A] {
	_ = "STUB: not implemented"
	return nil
}

func Alt[A any](that func() Option[A]) func(Option[A]) Option[A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadSequence2[T1, T2, R any](o1 Option[T1], o2 Option[T2], f func(T1, T2) Option[R]) Option[R] {
	_ = "STUB: not implemented"
	return nil
}

func Sequence2[T1, T2, R any](f func(T1, T2) Option[R]) func(Option[T1], Option[T2]) Option[R] {
	_ = "STUB: not implemented"
	return nil
}

func Reduce[A, B any](f func(B, A) B, initial B) func(Option[A]) B {
	_ = "STUB: not implemented"
	return nil
}

// Filter converts an optional onto itself if it is some and the predicate is true
func Filter[A any](pred func(A) bool) func(Option[A]) Option[A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[B, A any](fab Option[func(A) B], a A) Option[B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[B, A any](a A) func(Option[func(A) B]) Option[B] { _ = "STUB: not implemented"; return nil }
