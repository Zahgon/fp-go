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

package generic

import (
	M "github.com/IBM/fp-go/monoid"
	O "github.com/IBM/fp-go/option"
	"github.com/IBM/fp-go/tuple"
)

// Of constructs a single element array
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

func Empty[GA ~[]A, A any]() GA { _ = "STUB: not implemented"; return *new(GA) }

func UpsertAt[GA ~[]A, A any](a A) func(GA) GA { _ = "STUB: not implemented"; return nil }

func MonadMap[GA ~[]A, GB ~[]B, A, B any](as GA, f func(a A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Map[GA ~[]A, GB ~[]B, A, B any](f func(a A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapWithIndex[GA ~[]A, GB ~[]B, A, B any](as GA, f func(int, A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

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

func FilterChain[GA ~[]A, GB ~[]B, A, B any](f func(a A) O.Option[GB]) func(GA) GB {
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

func MonadPartition[GA ~[]A, A any](as GA, pred func(A) bool) tuple.Tuple2[GA, GA] {
	_ = "STUB: not implemented"
	return nil
}

// returns the partition

func Partition[GA ~[]A, A any](pred func(A) bool) func(GA) tuple.Tuple2[GA, GA] {
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

func Slice[AS ~[]A, A any](start int, end int) func(AS) AS { _ = "STUB: not implemented"; return nil }

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

func Prepend[ENDO ~func(AS) AS, AS []A, A any](head A) ENDO {
	_ = "STUB: not implemented"
	return *new(ENDO)
}
