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
	Mg "github.com/IBM/fp-go/v2/magma"
	Mo "github.com/IBM/fp-go/v2/monoid"
	O "github.com/IBM/fp-go/v2/option"
	"github.com/IBM/fp-go/v2/ord"
	"github.com/IBM/fp-go/v2/pair"
)

func IsEmpty[M ~map[K]V, K comparable, V any](r M) bool { _ = "STUB: not implemented"; return false }

func IsNonEmpty[M ~map[K]V, K comparable, V any](r M) bool { _ = "STUB: not implemented"; return false }

func Keys[M ~map[K]V, GK ~[]K, K comparable, V any](r M) GK {
	_ = "STUB: not implemented"
	// fast path
	return *new(GK)
}

// full implementation

func Values[M ~map[K]V, GV ~[]V, K comparable, V any](r M) GV {
	_ = "STUB: not implemented"
	// fast path
	return *new(GV)
}

// full implementation

func KeysOrd[M ~map[K]V, GK ~[]K, K comparable, V any](o ord.Ord[K]) func(r M) GK {
	_ = "STUB: not implemented"
	return nil

	// fast path
}

// full implementation

func ValuesOrd[M ~map[K]V, GV ~[]V, K comparable, V any](o ord.Ord[K]) func(r M) GV {
	_ = "STUB: not implemented"
	return nil

	// fast path
}

// full implementation

func collectOrd[M ~map[K]V, GR ~[]R, K comparable, V, R any](o ord.Ord[K], r M, f func(K, V) R) GR {
	_ = "STUB: not implemented"
	// create the entries
	return *new(GR)
}

// collect this array

// done

func reduceOrd[M ~map[K]V, K comparable, V, R any](o ord.Ord[K], r M, f func(K, R, V) R, initial R) R {
	_ = "STUB: not implemented"
	// create the entries
	return *new(R)
}

// collect this array

// done

func collect[M ~map[K]V, GR ~[]R, K comparable, V, R any](r M, f func(K, V) R) GR {
	_ = "STUB: not implemented"
	return *new(GR)
}

func Collect[M ~map[K]V, GR ~[]R, K comparable, V, R any](f func(K, V) R) func(M) GR {
	_ = "STUB: not implemented"
	// full implementation
	return nil
}

func CollectOrd[M ~map[K]V, GR ~[]R, K comparable, V, R any](o ord.Ord[K]) func(f func(K, V) R) func(M) GR {
	_ = "STUB: not implemented"
	return nil
}

// fast path

// full implementation

func Reduce[M ~map[K]V, K comparable, V, R any](f func(R, V) R, initial R) func(M) R {
	_ = "STUB: not implemented"
	return nil
}

func ReduceWithIndex[M ~map[K]V, K comparable, V, R any](f func(K, R, V) R, initial R) func(M) R {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRef[M ~map[K]V, K comparable, V, R any](f func(R, *V) R, initial R) func(M) R {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRefWithIndex[M ~map[K]V, K comparable, V, R any](f func(K, R, *V) R, initial R) func(M) R {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[BS ~map[K]B, ABS ~map[K]func(A) B, AS ~map[K]A, K comparable, B, A any](m Mo.Monoid[BS], fab ABS, fa AS) BS {
	_ = "STUB: not implemented"
	return *new(BS)
}

func Ap[BS ~map[K]B, ABS ~map[K]func(A) B, AS ~map[K]A, K comparable, B, A any](m Mo.Monoid[BS]) func(fa AS) func(ABS) BS {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[M ~map[K]V, N ~map[K]R, K comparable, V, R any](r M, f func(V) R) N {
	_ = "STUB: not implemented"
	return *new(N)
}

func MonadChainWithIndex[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](m Mo.Monoid[N], r M, f func(K, V1) N) N {
	_ = "STUB: not implemented"
	return *new(N)
}

func MonadChain[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](m Mo.Monoid[N], r M, f func(V1) N) N {
	_ = "STUB: not implemented"
	return *new(N)
}

func ChainWithIndex[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](m Mo.Monoid[N]) func(func(K, V1) N) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

func Chain[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](m Mo.Monoid[N]) func(func(V1) N) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapWithIndex[M ~map[K]V, N ~map[K]R, K comparable, V, R any](r M, f func(K, V) R) N {
	_ = "STUB: not implemented"
	return *new(N)
}

func MonadMapRefWithIndex[M ~map[K]V, N ~map[K]R, K comparable, V, R any](r M, f func(K, *V) R) N {
	_ = "STUB: not implemented"
	return *new(N)
}

func MonadMapRef[M ~map[K]V, N ~map[K]R, K comparable, V, R any](r M, f func(*V) R) N {
	_ = "STUB: not implemented"
	return *new(N)
}

func Map[M ~map[K]V, N ~map[K]R, K comparable, V, R any](f func(V) R) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

func MapRef[M ~map[K]V, N ~map[K]R, K comparable, V, R any](f func(*V) R) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

func MapWithIndex[M ~map[K]V, N ~map[K]R, K comparable, V, R any](f func(K, V) R) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

func MapRefWithIndex[M ~map[K]V, N ~map[K]R, K comparable, V, R any](f func(K, *V) R) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

func MonadLookup[M ~map[K]V, K comparable, V any](m M, k K) O.Option[V] {
	_ = "STUB: not implemented"
	return nil
}

func Lookup[M ~map[K]V, K comparable, V any](k K) func(M) O.Option[V] {
	_ = "STUB: not implemented"
	return nil
}

func Has[M ~map[K]V, K comparable, V any](k K, r M) bool { _ = "STUB: not implemented"; return false }

func union[M ~map[K]V, K comparable, V any](m Mg.Magma[V], left, right M) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func unionLast[M ~map[K]V, K comparable, V any](left, right M) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func Union[M ~map[K]V, K comparable, V any](m Mg.Magma[V]) func(M) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

func UnionLast[M ~map[K]V, K comparable, V any](right M) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

func Merge[M ~map[K]V, K comparable, V any](right M) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

func UnionFirst[M ~map[K]V, K comparable, V any](right M) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

func Empty[M ~map[K]V, K comparable, V any]() M { _ = "STUB: not implemented"; return *new(M) }

func Size[M ~map[K]V, K comparable, V any](r M) int { _ = "STUB: not implemented"; return 0 }

func ToArray[M ~map[K]V, GT ~[]pair.Pair[K, V], K comparable, V any](r M) GT {
	_ = "STUB: not implemented"
	return *new(GT)
}

func toEntriesOrd[M ~map[K]V, GT ~[]pair.Pair[K, V], K comparable, V any](o ord.Ord[K], r M) GT {
	_ = "STUB: not implemented"
	// total number of elements
	return *new(GT)
}

// produce an array that we can sort by key

// final entries

func ToEntriesOrd[M ~map[K]V, GT ~[]pair.Pair[K, V], K comparable, V any](o ord.Ord[K]) func(r M) GT {
	_ = "STUB: not implemented"
	return nil
}

func ToEntries[M ~map[K]V, GT ~[]pair.Pair[K, V], K comparable, V any](r M) GT {
	_ = "STUB: not implemented"
	return *

	// FromFoldableMap uses the reduce method for a higher kinded type to transform
	// its values into a tuple. The key and value are then used to populate the map. Duplicate
	// values are resolved via the provided [Mg.Magma]
	new(GT)
}

func FromFoldableMap[
	FCT ~func(A) pair.Pair[K, V],
	HKTA any,
	FOLDABLE ~func(func(M, A) M, M) func(HKTA) M,
	M ~map[K]V,
	A any,
	K comparable,
	V any](m Mg.Magma[V], fld FOLDABLE) func(f FCT) func(fa HKTA) M {
	_ = "STUB: not implemented"
	return nil
}

func FromFoldable[
	HKTA any,
	FOLDABLE ~func(func(M, pair.Pair[K, V]) M, M) func(HKTA) M,
	M ~map[K]V,
	K comparable,
	V any](m Mg.Magma[V], red FOLDABLE) func(fa HKTA) M {
	_ = "STUB: not implemented"
	return nil
}

func FromArrayMap[
	FCT ~func(A) pair.Pair[K, V],
	GA ~[]A,
	M ~map[K]V,
	A any,
	K comparable,
	V any](m Mg.Magma[V]) func(f FCT) func(fa GA) M {
	_ = "STUB: not implemented"
	return nil
}

func FromArray[
	GA ~[]pair.Pair[K, V],
	M ~map[K]V,
	K comparable,
	V any](m Mg.Magma[V]) func(fa GA) M {
	_ = "STUB: not implemented"
	return nil
}

func FromEntries[M ~map[K]V, GT ~[]pair.Pair[K, V], K comparable, V any](fa GT) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func duplicate[M ~map[K]V, K comparable, V any](r M) M { _ = "STUB: not implemented"; return *new(M) }

func upsertAt[M ~map[K]V, K comparable, V any](r M, k K, v V) M {
	_ = "STUB: not implemented"
	// fast path
	return *new(M)
}

// duplicate and update

func deleteAt[M ~map[K]V, K comparable, V any](r M, k K) M {
	_ = "STUB: not implemented"
	// fast path
	return *new(M)
}

// duplicate and update

func upsertAtReadWrite[M ~map[K]V, K comparable, V any](r M, k K, v V) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func UpsertAt[M ~map[K]V, K comparable, V any](k K, v V) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

func DeleteAt[M ~map[K]V, K comparable, V any](k K) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

func Singleton[M ~map[K]V, K comparable, V any](k K, v V) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func filterMapWithIndex[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](fa M, f func(K, V1) O.Option[V2]) N {
	_ = "STUB: not implemented"
	return *new(N)
}

func filterWithIndex[M ~map[K]V, K comparable, V any](fa M, f func(K, V) bool) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func filter[M ~map[K]V, K comparable, V any](fa M, f func(K) bool) M {
	_ = "STUB: not implemented"
	return *new(M)
}

// Filter creates a new map with only the elements that match the predicate
func Filter[M ~map[K]V, K comparable, V any](f func(K) bool) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

// FilterWithIndex creates a new map with only the elements that match the predicate
func FilterWithIndex[M ~map[K]V, K comparable, V any](f func(K, V) bool) func(M) M {
	_ = "STUB: not implemented"
	return nil
}

// FilterMapWithIndex creates a new map with only the elements for which the transformation function creates a Some
func FilterMapWithIndex[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](f func(K, V1) O.Option[V2]) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

// FilterMap creates a new map with only the elements for which the transformation function creates a Some
func FilterMap[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](f func(V1) O.Option[V2]) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

// Flatten converts a nested map into a regular map
func Flatten[M ~map[K]N, N ~map[K]V, K comparable, V any](m Mo.Monoid[N]) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

// FilterChainWithIndex creates a new map with only the elements for which the transformation function creates a Some
func FilterChainWithIndex[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](m Mo.Monoid[N]) func(func(K, V1) O.Option[N]) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

// FilterChain creates a new map with only the elements for which the transformation function creates a Some
func FilterChain[M ~map[K]V1, N ~map[K]V2, K comparable, V1, V2 any](m Mo.Monoid[N]) func(func(V1) O.Option[N]) func(M) N {
	_ = "STUB: not implemented"
	return nil
}

// IsNil checks if the map is set to nil
func IsNil[M ~map[K]V, K comparable, V any](m M) bool {
	_ = "STUB: not implemented"

	// IsNonNil checks if the map is set to nil
	return false
}

func IsNonNil[M ~map[K]V, K comparable, V any](m M) bool {
	_ = "STUB: not implemented"

	// ConstNil return a nil map
	return false
}

func ConstNil[M ~map[K]V, K comparable, V any]() M { _ = "STUB: not implemented"; return *new(M) }

func FoldMap[AS ~map[K]A, K comparable, A, B any](m Mo.Monoid[B]) func(func(A) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

func Fold[AS ~map[K]A, K comparable, A any](m Mo.Monoid[A]) func(AS) A {
	_ = "STUB: not implemented"
	return nil
}

func FoldMapWithIndex[AS ~map[K]A, K comparable, A, B any](m Mo.Monoid[B]) func(func(K, A) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

func ReduceOrdWithIndex[M ~map[K]V, K comparable, V, R any](o ord.Ord[K]) func(func(K, R, V) R, R) func(M) R {
	_ = "STUB: not implemented"
	return nil
}

func ReduceOrd[M ~map[K]V, K comparable, V, R any](o ord.Ord[K]) func(func(R, V) R, R) func(M) R {
	_ = "STUB: not implemented"
	return nil
}

func FoldMapOrd[AS ~map[K]A, K comparable, A, B any](o ord.Ord[K]) func(m Mo.Monoid[B]) func(func(A) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

func FoldOrd[AS ~map[K]A, K comparable, A any](o ord.Ord[K]) func(m Mo.Monoid[A]) func(AS) A {
	_ = "STUB: not implemented"
	return nil
}

func FoldMapOrdWithIndex[AS ~map[K]A, K comparable, A, B any](o ord.Ord[K]) func(m Mo.Monoid[B]) func(func(K, A) B) func(AS) B {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[GFAB ~map[K]func(A) B, GB ~map[K]B, K comparable, A, B any](fab GFAB, a A) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Flap[GFAB ~map[K]func(A) B, GB ~map[K]B, K comparable, A, B any](a A) func(GFAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func Copy[M ~map[K]V, K comparable, V any](m M) M { _ = "STUB: not implemented"; return *new(M) }

func Clone[M ~map[K]V, K comparable, V any](f func(V) V) func(m M) M {
	_ = "STUB: not implemented"
	// impementation assumes that map does not optimize for the empty map
	return nil
}
