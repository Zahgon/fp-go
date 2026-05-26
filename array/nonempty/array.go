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

package nonempty

import (
	EM "github.com/IBM/fp-go/endomorphism"
	S "github.com/IBM/fp-go/semigroup"
)

// NonEmptyArray represents an array with at least one element
type NonEmptyArray[A any] []A

// Of constructs a single element array
func Of[A any](first A) NonEmptyArray[A] { _ = "STUB: not implemented"; return nil }

// From constructs a [NonEmptyArray] from a set of variadic arguments
func From[A any](first A, data ...A) NonEmptyArray[A] { _ = "STUB: not implemented"; return nil }

// allocate the requested buffer

func IsEmpty[A any](_ NonEmptyArray[A]) bool { _ = "STUB: not implemented"; return false }

func IsNonEmpty[A any](_ NonEmptyArray[A]) bool { _ = "STUB: not implemented"; return false }

func MonadMap[A, B any](as NonEmptyArray[A], f func(a A) B) NonEmptyArray[B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[A, B any](f func(a A) B) func(NonEmptyArray[A]) NonEmptyArray[B] {
	_ = "STUB: not implemented"
	return nil
}

func Reduce[A, B any](f func(B, A) B, initial B) func(NonEmptyArray[A]) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRight[A, B any](f func(A, B) B, initial B) func(NonEmptyArray[A]) B {
	_ = "STUB: not implemented"
	return nil
}

func Tail[A any](as NonEmptyArray[A]) []A { _ = "STUB: not implemented"; return nil }

func Head[A any](as NonEmptyArray[A]) A { _ = "STUB: not implemented"; return *new(A) }

func First[A any](as NonEmptyArray[A]) A { _ = "STUB: not implemented"; return *new(A) }

func Last[A any](as NonEmptyArray[A]) A { _ = "STUB: not implemented"; return *new(A) }

func Size[A any](as NonEmptyArray[A]) int { _ = "STUB: not implemented"; return 0 }

func Flatten[A any](mma NonEmptyArray[NonEmptyArray[A]]) NonEmptyArray[A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[A, B any](fa NonEmptyArray[A], f func(a A) NonEmptyArray[B]) NonEmptyArray[B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A, B any](f func(A) NonEmptyArray[B]) func(NonEmptyArray[A]) NonEmptyArray[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[B, A any](fab NonEmptyArray[func(A) B], fa NonEmptyArray[A]) NonEmptyArray[B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, A any](fa NonEmptyArray[A]) func(NonEmptyArray[func(A) B]) NonEmptyArray[B] {
	_ = "STUB: not implemented"
	return nil
}

// FoldMap maps and folds a [NonEmptyArray]. Map the [NonEmptyArray] passing each value to the iterating function. Then fold the results using the provided [Semigroup].
func FoldMap[A, B any](s S.Semigroup[B]) func(func(A) B) func(NonEmptyArray[A]) B {
	_ = "STUB: not implemented"
	return nil
}

// Fold folds the [NonEmptyArray] using the provided [Semigroup].
func Fold[A any](s S.Semigroup[A]) func(NonEmptyArray[A]) A { _ = "STUB: not implemented"; return nil }

// Prepend prepends a single value to an array
func Prepend[A any](head A) EM.Endomorphism[NonEmptyArray[A]] {
	_ = "STUB: not implemented"
	return nil
}
