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

package generic

import (
	M "github.com/IBM/fp-go/v2/monoid"
	O "github.com/IBM/fp-go/v2/option"
	"github.com/IBM/fp-go/v2/pair"
)

// Of constructs a single element array
//
//go:inline
func Of[GA ~[]A, A any](value A) GA { _ = "STUB: not implemented"; return *new(GA) }

func Reduce[GA ~[]A, A, B any](f func(B, A) B, initial B) func(GA) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceWithIndex[GA ~[]A, A, B any](f func(int, B, A) B, initial B) func(GA) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRight[GA ~[]A, A, B any](f func(A, B) B, initial B) func(GA) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRightWithIndex[GA ~[]A, A, B any](f func(int, A, B) B, initial B) func(GA) B {
	_ = "STUB: not implemented"
	return nil
}

func MonadReduce[GA ~[]A, A, B any](fa GA, f func(B, A) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func MonadReduceWithIndex[GA ~[]A, A, B any](fa GA, f func(int, B, A) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func MonadReduceRight[GA ~[]A, A, B any](fa GA, f func(A, B) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func MonadReduceRightWithIndex[GA ~[]A, A, B any](fa GA, f func(int, A, B) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

// From constructs an array from a set of variadic arguments
func From[GA ~[]A, A any](data ...A) GA {
	_ = "STUB: not implemented"

	// MakeBy returns a `Array` of length `n` with element `i` initialized with `f(i)`.
	return *new(GA)
}

func MakeBy[AS ~[]A, F ~func(int) A, A any](n int, f F) AS {
	_ = "STUB: not implemented"
	// sanity check
	return *new(AS)
}

// run the generator function across the input

func Replicate[AS ~[]A, A any](n int, a A) AS { _ = "STUB: not implemented"; return *new(AS) }

func Lookup[GA ~[]A, A any](idx int) func(GA) O.Option[A] { _ = "STUB: not implemented"; return nil }

func Tail[GA ~[]A, A any](as GA) O.Option[GA] { _ = "STUB: not implemented"; return nil }

func Head[GA ~[]A, A any](as GA) O.Option[A] { _ = "STUB: not implemented"; return nil }

func First[GA ~[]A, A any](as GA) O.Option[A] { _ = "STUB: not implemented"; return nil }

func Last[GA ~[]A, A any](as GA) O.Option[A] { _ = "STUB: not implemented"; return nil }

func Append[GA ~[]A, A any](as GA, a A) GA { _ = "STUB: not implemented"; return *new(GA) }

func Empty[GA ~[]A, A any]() GA {
	_ = "STUB: not implemented"
	return *

	//go:inline
	new(GA)
}

func UpsertAt[GA ~[]A, A any](a A) func(GA) GA { _ = "STUB: not implemented"; return nil }

//go:inline
func MonadMap[GA ~[]A, GB ~[]B, A, B any](as GA, f func(a A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

//go:inline
func Map[GA ~[]A, GB ~[]B, A, B any](f func(a A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadMapWithIndex[GA ~[]A, GB ~[]B, A, B any](as GA, f func(int, A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

//go:inline
func MapWithIndex[GA ~[]A, GB ~[]B, A, B any](f func(int, A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func Size[GA ~[]A, A any](as GA) int { _ = "STUB: not implemented"; return 0 }

func filterMap[GA ~[]A, GB ~[]B, A, B any](fa GA, f func(A) O.Option[B]) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func filterMapWithIndex[GA ~[]A, GB ~[]B, A, B any](fa GA, f func(int, A) O.Option[B]) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func MonadFilterMap[GA ~[]A, GB ~[]B, A, B any](fa GA, f func(A) O.Option[B]) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func MonadFilterMapWithIndex[GA ~[]A, GB ~[]B, A, B any](fa GA, f func(int, A) O.Option[B]) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func filterWithIndex[AS ~[]A, PRED ~func(int, A) bool, A any](fa AS, pred PRED) AS {
	_ = "STUB: not implemented"
	return *new(AS)
}

func FilterWithIndex[AS ~[]A, PRED ~func(int, A) bool, A any](pred PRED) func(AS) AS {
	_ = "STUB: not implemented"
	return nil
}

func Filter[AS ~[]A, PRED ~func(A) bool, A any](pred PRED) func(AS) AS {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[GA ~[]A, GB ~[]B, A, B any](f func(a A) O.Option[GB]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GAA ~[]GA, GA ~[]A, A any](mma GAA) GA { _ = "STUB: not implemented"; return *new(GA) }

func FilterMap[GA ~[]A, GB ~[]B, A, B any](f func(A) O.Option[B]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func FilterMapWithIndex[GA ~[]A, GB ~[]B, A, B any](f func(int, A) O.Option[B]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadPartition[GA ~[]A, A any](as GA, pred func(A) bool) pair.Pair[GA, GA] {
	_ = "STUB: not implemented"
	return nil
}

// returns the partition

func Partition[GA ~[]A, A any](pred func(A) bool) func(GA) pair.Pair[GA, GA] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[AS ~[]A, BS ~[]B, A, B any](fa AS, f func(a A) BS) BS {
	_ = "STUB: not implemented"
	return *new(BS)
}

func Chain[AS ~[]A, BS ~[]B, A, B any](f func(A) BS) func(AS) BS {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[BS ~[]B, ABS ~[]func(A) B, AS ~[]A, B, A any](fab ABS, fa AS) BS {
	_ = "STUB: not implemented"
	return *new(BS)
}

func Ap[BS ~[]B, ABS ~[]func(A) B, AS ~[]A, B, A any](fa AS) func(ABS) BS {
	_ = "STUB: not implemented"
	return nil
}

func IsEmpty[AS ~[]A, A any](as AS) bool { _ = "STUB: not implemented"; return false }

func IsNil[GA ~[]A, A any](as GA) bool { _ = "STUB: not implemented"; return false }

func IsNonNil[GA ~[]A, A any](as GA) bool { _ = "STUB: not implemented"; return false }

func Match[AS ~[]A, A, B any](onEmpty func() B, onNonEmpty func(AS) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

func MatchLeft[AS ~[]A, A, B any](onEmpty func() B, onNonEmpty func(A, AS) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Slice[AS ~[]A, A any](start, end int) func(AS) AS { _ = "STUB: not implemented"; return nil }

//go:inline
func SliceRight[AS ~[]A, A any](start int) func(AS) AS { _ = "STUB: not implemented"; return nil }

func Copy[AS ~[]A, A any](b AS) AS { _ = "STUB: not implemented"; return *new(AS) }

func Clone[AS ~[]A, A any](f func(A) A) func(as AS) AS {
	_ = "STUB: not implemented"
	// implementation assumes that map does not optimize for the empty array
	return nil
}

func FoldMap[AS ~[]A, A, B any](m M.Monoid[B]) func(func(A) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

func FoldMapWithIndex[AS ~[]A, A, B any](m M.Monoid[B]) func(func(int, A) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

func Fold[AS ~[]A, A any](m M.Monoid[A]) func(AS) A { _ = "STUB: not implemented"; return nil }

func Push[ENDO ~func(GA) GA, GA ~[]A, A any](a A) ENDO {
	_ = "STUB: not implemented"
	return *new(ENDO)
}

func MonadFlap[FAB ~func(A) B, GFAB ~[]FAB, GB ~[]B, A, B any](fab GFAB, a A) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Flap[FAB ~func(A) B, GFAB ~[]FAB, GB ~[]B, A, B any](a A) func(GFAB) GB {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Prepend[ENDO ~func(AS) AS, AS []A, A any](head A) ENDO {
	_ = "STUB: not implemented"
	return *new(ENDO)
}

//go:inline
func Reverse[GT ~[]T, T any](as GT) GT {
	_ = "STUB: not implemented"
	return *

	// Extract returns the first element of an array, or a zero value if empty.
	// This is the comonad extract operation for arrays.
	//
	// Extract is the dual of the monadic return/of operation. While Of wraps a value
	// in a context, Extract unwraps a value from its context.
	//
	// Type Parameters:
	//   - GA: The array type constraint
	//   - A: The type of elements in the array
	//
	// Parameters:
	//   - as: The input array
	//
	// Returns:
	//   - The first element if the array is non-empty, otherwise the zero value of type A
	//
	// Behavior:
	//   - Returns as[0] if the array has at least one element
	//   - Returns the zero value of A if the array is empty
	//   - Does not modify the input array
	//
	// Example:
	//
	//	result := Extract([]int{1, 2, 3})
	//	// result: 1
	//
	// Example with empty array:
	//
	//	result := Extract([]int{})
	//	// result: 0 (zero value for int)
	//
	// Comonad laws:
	//   - Extract ∘ Of == Identity (extracting from a singleton returns the value)
	//   - Extract ∘ Extend(f) == f (extract after extend equals applying f)
	//
	//go:inline
	new(GT)
}

func Extract[GA ~[]A, A any](as GA) A { _ = "STUB: not implemented"; return *new(A) }

// Extend applies a function to every suffix of an array, creating a new array of results.
// This is the comonad extend operation for arrays.
//
// The function f is applied to progressively smaller suffixes of the input array:
//   - f(as[0:]) for the first element
//   - f(as[1:]) for the second element
//   - f(as[2:]) for the third element
//   - and so on...
//
// Type Parameters:
//   - GA: The input array type constraint
//   - GB: The output array type constraint
//   - A: The type of elements in the input array
//   - B: The type of elements in the output array
//
// Parameters:
//   - f: A function that takes an array suffix and returns a value
//
// Returns:
//   - A function that transforms an array of A into an array of B
//
// Behavior:
//   - Creates a new array with the same length as the input
//   - For each position i, applies f to the suffix starting at i
//   - Returns an empty array if the input is empty
//
// Example:
//
//	// Sum all elements from current position to end
//	sumSuffix := Extend[[]int, []int](func(as []int) int {
//	    return MonadReduce(as, func(acc, x int) int { return acc + x }, 0)
//	})
//	result := sumSuffix([]int{1, 2, 3, 4})
//	// result: []int{10, 9, 7, 4}
//	// Explanation: [1+2+3+4, 2+3+4, 3+4, 4]
//
// Example with length:
//
//	// Get remaining length at each position
//	lengths := Extend[[]int, []int](Size[[]int, int])
//	result := lengths([]int{10, 20, 30})
//	// result: []int{3, 2, 1}
//
// Comonad laws:
//   - Left identity: Extend(Extract) == Identity
//   - Right identity: Extract ∘ Extend(f) == f
//   - Associativity: Extend(f) ∘ Extend(g) == Extend(f ∘ Extend(g))
//
//go:inline
func Extend[GA ~[]A, GB ~[]B, A, B any](f func(GA) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func UpdateAt[GT ~[]T, T any](i int, v T) func(GT) O.Option[GT] {
	_ = "STUB: not implemented"
	return nil
}
