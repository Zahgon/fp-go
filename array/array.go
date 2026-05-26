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

package array

import (
	EM "github.com/IBM/fp-go/endomorphism"
	M "github.com/IBM/fp-go/monoid"
	O "github.com/IBM/fp-go/option"
	"github.com/IBM/fp-go/tuple"
)

// From constructs an array from a set of variadic arguments
func From[A any](data ...A) []A { _ = "STUB: not implemented"; return nil }

// MakeBy returns a `Array` of length `n` with element `i` initialized with `f(i)`.
func MakeBy[F ~func(int) A, A any](n int, f F) []A { _ = "STUB: not implemented"; return nil }

// Replicate creates a `Array` containing a value repeated the specified number of times.
func Replicate[A any](n int, a A) []A { _ = "STUB: not implemented"; return nil }

func MonadMap[A, B any](as []A, f func(a A) B) []B { _ = "STUB: not implemented"; return nil }

func MonadMapRef[A, B any](as []A, f func(a *A) B) []B { _ = "STUB: not implemented"; return nil }

func MapWithIndex[A, B any](f func(int, A) B) func([]A) []B { _ = "STUB: not implemented"; return nil }

func Map[A, B any](f func(a A) B) func([]A) []B { _ = "STUB: not implemented"; return nil }

func MapRef[A, B any](f func(a *A) B) func([]A) []B { _ = "STUB: not implemented"; return nil }

func filterRef[A any](fa []A, pred func(a *A) bool) []A { _ = "STUB: not implemented"; return nil }

func filterMapRef[A, B any](fa []A, pred func(a *A) bool, f func(a *A) B) []B {
	_ = "STUB: not implemented"
	return nil
}

// Filter returns a new array with all elements from the original array that match a predicate
func Filter[A any](pred func(A) bool) EM.Endomorphism[[]A] { _ = "STUB: not implemented"; return nil }

// FilterWithIndex returns a new array with all elements from the original array that match a predicate
func FilterWithIndex[A any](pred func(int, A) bool) EM.Endomorphism[[]A] {
	_ = "STUB: not implemented"
	return nil
}

func FilterRef[A any](pred func(*A) bool) EM.Endomorphism[[]A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFilterMap[A, B any](fa []A, f func(A) O.Option[B]) []B {
	_ = "STUB: not implemented"
	return nil
}

func MonadFilterMapWithIndex[A, B any](fa []A, f func(int, A) O.Option[B]) []B {
	_ = "STUB: not implemented"
	return nil
}

// FilterMap maps an array with an iterating function that returns an [O.Option] and it keeps only the Some values discarding the Nones.
func FilterMap[A, B any](f func(A) O.Option[B]) func([]A) []B {
	_ = "STUB: not implemented"
	return nil
}

// FilterMapWithIndex maps an array with an iterating function that returns an [O.Option] and it keeps only the Some values discarding the Nones.
func FilterMapWithIndex[A, B any](f func(int, A) O.Option[B]) func([]A) []B {
	_ = "STUB: not implemented"
	return nil
}

// FilterChain maps an array with an iterating function that returns an [O.Option] of an array. It keeps only the Some values discarding the Nones and then flattens the result.
func FilterChain[A, B any](f func(A) O.Option[[]B]) func([]A) []B {
	_ = "STUB: not implemented"
	return nil
}

func FilterMapRef[A, B any](pred func(a *A) bool, f func(a *A) B) func([]A) []B {
	_ = "STUB: not implemented"
	return nil
}

func reduceRef[A, B any](fa []A, f func(B, *A) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func MonadReduce[A, B any](fa []A, f func(B, A) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func Reduce[A, B any](f func(B, A) B, initial B) func([]A) B { _ = "STUB: not implemented"; return nil }

func ReduceWithIndex[A, B any](f func(int, B, A) B, initial B) func([]A) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRight[A, B any](f func(A, B) B, initial B) func([]A) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRightWithIndex[A, B any](f func(int, A, B) B, initial B) func([]A) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRef[A, B any](f func(B, *A) B, initial B) func([]A) B {
	_ = "STUB: not implemented"
	return nil
}

func Append[A any](as []A, a A) []A { _ = "STUB: not implemented"; return nil }

func IsEmpty[A any](as []A) bool { _ = "STUB: not implemented"; return false }

func IsNonEmpty[A any](as []A) bool { _ = "STUB: not implemented"; return false }

func Empty[A any]() []A { _ = "STUB: not implemented"; return nil }

func Zero[A any]() []A {
	_ = "STUB: not implemented"

	// Of constructs a single element array
	return nil
}

func Of[A any](a A) []A { _ = "STUB: not implemented"; return nil }

func MonadChain[A, B any](fa []A, f func(a A) []B) []B { _ = "STUB: not implemented"; return nil }

func Chain[A, B any](f func(A) []B) func([]A) []B { _ = "STUB: not implemented"; return nil }

func MonadAp[B, A any](fab []func(A) B, fa []A) []B { _ = "STUB: not implemented"; return nil }

func Ap[B, A any](fa []A) func([]func(A) B) []B { _ = "STUB: not implemented"; return nil }

func Match[A, B any](onEmpty func() B, onNonEmpty func([]A) B) func([]A) B {
	_ = "STUB: not implemented"
	return nil
}

func MatchLeft[A, B any](onEmpty func() B, onNonEmpty func(A, []A) B) func([]A) B {
	_ = "STUB: not implemented"
	return nil
}

func Tail[A any](as []A) O.Option[[]A] { _ = "STUB: not implemented"; return nil }

func Head[A any](as []A) O.Option[A] { _ = "STUB: not implemented"; return nil }

func First[A any](as []A) O.Option[A] { _ = "STUB: not implemented"; return nil }

func Last[A any](as []A) O.Option[A] { _ = "STUB: not implemented"; return nil }

func PrependAll[A any](middle A) EM.Endomorphism[[]A] { _ = "STUB: not implemented"; return nil }

func Intersperse[A any](middle A) EM.Endomorphism[[]A] { _ = "STUB: not implemented"; return nil }

func Intercalate[A any](m M.Monoid[A]) func(A) func([]A) A { _ = "STUB: not implemented"; return nil }

func Flatten[A any](mma [][]A) []A { _ = "STUB: not implemented"; return nil }

func Slice[A any](low, high int) func(as []A) []A { _ = "STUB: not implemented"; return nil }

func Lookup[A any](idx int) func([]A) O.Option[A] { _ = "STUB: not implemented"; return nil }

func UpsertAt[A any](a A) EM.Endomorphism[[]A] { _ = "STUB: not implemented"; return nil }

func Size[A any](as []A) int { _ = "STUB: not implemented"; return 0 }

func MonadPartition[A any](as []A, pred func(A) bool) tuple.Tuple2[[]A, []A] {
	_ = "STUB: not implemented"
	return nil
}

// Partition creates two new arrays out of one, the left result contains the elements
// for which the predicate returns false, the right one those for which the predicate returns true
func Partition[A any](pred func(A) bool) func([]A) tuple.Tuple2[[]A, []A] {
	_ = "STUB: not implemented"
	return nil
}

// IsNil checks if the array is set to nil
func IsNil[A any](as []A) bool { _ = "STUB: not implemented"; return false }

// IsNonNil checks if the array is set to nil
func IsNonNil[A any](as []A) bool { _ = "STUB: not implemented"; return false }

// ConstNil returns a nil array
func ConstNil[A any]() []A { _ = "STUB: not implemented"; return nil }

func SliceRight[A any](start int) EM.Endomorphism[[]A] { _ = "STUB: not implemented"; return nil }

// Copy creates a shallow copy of the array
func Copy[A any](b []A) []A {
	_ = "STUB: not implemented"

	// Clone creates a deep copy of the array using the provided endomorphism to clone the values
	return nil
}

func Clone[A any](f func(A) A) func(as []A) []A { _ = "STUB: not implemented"; return nil }

// FoldMap maps and folds an array. Map the Array passing each value to the iterating function. Then fold the results using the provided Monoid.
func FoldMap[A, B any](m M.Monoid[B]) func(func(A) B) func([]A) B {
	_ = "STUB: not implemented"
	return nil

	// FoldMapWithIndex maps and folds an array. Map the Array passing each value to the iterating function. Then fold the results using the provided Monoid.
}

func FoldMapWithIndex[A, B any](m M.Monoid[B]) func(func(int, A) B) func([]A) B {
	_ = "STUB: not implemented"
	return nil
}

// Fold folds the array using the provided Monoid.
func Fold[A any](m M.Monoid[A]) func([]A) A { _ = "STUB: not implemented"; return nil }

func Push[A any](a A) EM.Endomorphism[[]A] { _ = "STUB: not implemented"; return nil }

func MonadFlap[B, A any](fab []func(A) B, a A) []B { _ = "STUB: not implemented"; return nil }

func Flap[B, A any](a A) func([]func(A) B) []B { _ = "STUB: not implemented"; return nil }

func Prepend[A any](head A) EM.Endomorphism[[]A] { _ = "STUB: not implemented"; return nil }
