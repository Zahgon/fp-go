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

package record

import (
	EM "github.com/IBM/fp-go/endomorphism"
	Mg "github.com/IBM/fp-go/magma"
	Mo "github.com/IBM/fp-go/monoid"
	O "github.com/IBM/fp-go/option"
	"github.com/IBM/fp-go/ord"
	T "github.com/IBM/fp-go/tuple"
)

// IsEmpty tests if a map is empty
func IsEmpty[K comparable, V any](r map[K]V) bool { _ = "STUB: not implemented"; return false }

// IsNonEmpty tests if a map is not empty
func IsNonEmpty[K comparable, V any](r map[K]V) bool { _ = "STUB: not implemented"; return false }

// Keys returns the key in a map
func Keys[K comparable, V any](r map[K]V) []K { _ = "STUB: not implemented"; return nil }

// Values returns the values in a map
func Values[K comparable, V any](r map[K]V) []V { _ = "STUB: not implemented"; return nil }

// Collect applies a collector function to the key value pairs in a map and returns the result as an array
func Collect[K comparable, V, R any](f func(K, V) R) func(map[K]V) []R {
	_ = "STUB: not implemented"
	return nil
}

// CollectOrd applies a collector function to the key value pairs in a map and returns the result as an array
func CollectOrd[V, R any, K comparable](o ord.Ord[K]) func(func(K, V) R) func(map[K]V) []R {
	_ = "STUB: not implemented"
	return nil
}

func Reduce[K comparable, V, R any](f func(R, V) R, initial R) func(map[K]V) R {
	_ = "STUB: not implemented"
	return nil
}

func ReduceWithIndex[K comparable, V, R any](f func(K, R, V) R, initial R) func(map[K]V) R {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRef[K comparable, V, R any](f func(R, *V) R, initial R) func(map[K]V) R {
	_ = "STUB: not implemented"
	return nil
}

func ReduceRefWithIndex[K comparable, V, R any](f func(K, R, *V) R, initial R) func(map[K]V) R {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[K comparable, V, R any](r map[K]V, f func(V) R) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapWithIndex[K comparable, V, R any](r map[K]V, f func(K, V) R) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapRefWithIndex[K comparable, V, R any](r map[K]V, f func(K, *V) R) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapRef[K comparable, V, R any](r map[K]V, f func(*V) R) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

func Map[K comparable, V, R any](f func(V) R) func(map[K]V) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

func MapRef[K comparable, V, R any](f func(*V) R) func(map[K]V) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

func MapWithIndex[K comparable, V, R any](f func(K, V) R) func(map[K]V) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

func MapRefWithIndex[K comparable, V, R any](f func(K, *V) R) func(map[K]V) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

// Lookup returns the entry for a key in a map if it exists
func Lookup[V any, K comparable](k K) func(map[K]V) O.Option[V] {
	_ = "STUB: not implemented"
	return nil
}

// MonadLookup returns the entry for a key in a map if it exists
func MonadLookup[V any, K comparable](m map[K]V, k K) O.Option[V] {
	_ = "STUB: not implemented"
	return nil
}

// Has tests if a key is contained in a map
func Has[K comparable, V any](k K, r map[K]V) bool { _ = "STUB: not implemented"; return false }

func Union[K comparable, V any](m Mg.Magma[V]) func(map[K]V) func(map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil

	// Merge combines two maps giving the values in the right one precedence. Also refer to [MergeMonoid]
}

func Merge[K comparable, V any](right map[K]V) func(map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// Empty creates an empty map
func Empty[K comparable, V any]() map[K]V { _ = "STUB: not implemented"; return nil }

// Size returns the number of elements in a map
func Size[K comparable, V any](r map[K]V) int { _ = "STUB: not implemented"; return 0 }

func ToArray[K comparable, V any](r map[K]V) []T.Tuple2[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func ToEntries[K comparable, V any](r map[K]V) []T.Tuple2[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func FromEntries[K comparable, V any](fa []T.Tuple2[K, V]) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func UpsertAt[K comparable, V any](k K, v V) func(map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func DeleteAt[K comparable, V any](k K) func(map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// Singleton creates a new map with a single entry
func Singleton[K comparable, V any](k K, v V) map[K]V { _ = "STUB: not implemented"; return nil }

// FilterMapWithIndex creates a new map with only the elements for which the transformation function creates a Some
func FilterMapWithIndex[K comparable, V1, V2 any](f func(K, V1) O.Option[V2]) func(map[K]V1) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

// FilterMap creates a new map with only the elements for which the transformation function creates a Some
func FilterMap[K comparable, V1, V2 any](f func(V1) O.Option[V2]) func(map[K]V1) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

// Filter creates a new map with only the elements that match the predicate
func Filter[K comparable, V any](f func(K) bool) func(map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FilterWithIndex creates a new map with only the elements that match the predicate
func FilterWithIndex[K comparable, V any](f func(K, V) bool) func(map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// IsNil checks if the map is set to nil
func IsNil[K comparable, V any](m map[K]V) bool {
	_ = "STUB: not implemented"

	// IsNonNil checks if the map is set to nil
	return false
}

func IsNonNil[K comparable, V any](m map[K]V) bool { _ = "STUB: not implemented"; return false }

// ConstNil return a nil map
func ConstNil[K comparable, V any]() map[K]V { _ = "STUB: not implemented"; return nil }

func MonadChainWithIndex[V1 any, K comparable, V2 any](m Mo.Monoid[map[K]V2], r map[K]V1, f func(K, V1) map[K]V2) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[V1 any, K comparable, V2 any](m Mo.Monoid[map[K]V2], r map[K]V1, f func(V1) map[K]V2) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

func ChainWithIndex[V1 any, K comparable, V2 any](m Mo.Monoid[map[K]V2]) func(func(K, V1) map[K]V2) func(map[K]V1) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

func Chain[V1 any, K comparable, V2 any](m Mo.Monoid[map[K]V2]) func(func(V1) map[K]V2) func(map[K]V1) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

// Flatten converts a nested map into a regular map
func Flatten[K comparable, V any](m Mo.Monoid[map[K]V]) func(map[K]map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FilterChainWithIndex creates a new map with only the elements for which the transformation function creates a Some
func FilterChainWithIndex[V1 any, K comparable, V2 any](m Mo.Monoid[map[K]V2]) func(func(K, V1) O.Option[map[K]V2]) func(map[K]V1) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

// FilterChain creates a new map with only the elements for which the transformation function creates a Some
func FilterChain[V1 any, K comparable, V2 any](m Mo.Monoid[map[K]V2]) func(func(V1) O.Option[map[K]V2]) func(map[K]V1) map[K]V2 {
	_ = "STUB: not implemented"
	return nil
}

// FoldMap maps and folds a record. Map the record passing each value to the iterating function. Then fold the results using the provided Monoid.
func FoldMap[K comparable, A, B any](m Mo.Monoid[B]) func(func(A) B) func(map[K]A) B {
	_ = "STUB: not implemented"
	return nil
}

// FoldMapWithIndex maps and folds a record. Map the record passing each value to the iterating function. Then fold the results using the provided Monoid.
func FoldMapWithIndex[K comparable, A, B any](m Mo.Monoid[B]) func(func(K, A) B) func(map[K]A) B {
	_ = "STUB: not implemented"
	return nil
}

// Fold folds the record using the provided Monoid.
func Fold[K comparable, A any](m Mo.Monoid[A]) func(map[K]A) A {
	_ = "STUB: not implemented"
	return nil

	// ReduceOrdWithIndex reduces a map into a single value via a reducer function making sure that the keys are passed to the reducer in the specified order
}

func ReduceOrdWithIndex[V, R any, K comparable](o ord.Ord[K]) func(func(K, R, V) R, R) func(map[K]V) R {
	_ = "STUB: not implemented"
	return nil
}

// ReduceOrd reduces a map into a single value via a reducer function making sure that the keys are passed to the reducer in the specified order
func ReduceOrd[V, R any, K comparable](o ord.Ord[K]) func(func(R, V) R, R) func(map[K]V) R {
	_ = "STUB: not implemented"
	return nil
}

// FoldMap maps and folds a record. Map the record passing each value to the iterating function. Then fold the results using the provided Monoid and the items in the provided order
func FoldMapOrd[A, B any, K comparable](o ord.Ord[K]) func(m Mo.Monoid[B]) func(func(A) B) func(map[K]A) B {
	_ = "STUB: not implemented"
	return nil
}

// Fold folds the record using the provided Monoid with the items passed in the given order
func FoldOrd[A any, K comparable](o ord.Ord[K]) func(m Mo.Monoid[A]) func(map[K]A) A {
	_ = "STUB: not implemented"
	return nil
}

// FoldMapWithIndex maps and folds a record. Map the record passing each value to the iterating function. Then fold the results using the provided Monoid and the items in the provided order
func FoldMapOrdWithIndex[K comparable, A, B any](o ord.Ord[K]) func(m Mo.Monoid[B]) func(func(K, A) B) func(map[K]A) B {
	_ = "STUB: not implemented"
	return nil
}

// KeysOrd returns the keys in the map in their given order
func KeysOrd[V any, K comparable](o ord.Ord[K]) func(r map[K]V) []K {
	_ = "STUB: not implemented"
	return nil
}

// ValuesOrd returns the values in the map ordered by their keys in the given order
func ValuesOrd[V any, K comparable](o ord.Ord[K]) func(r map[K]V) []V {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[B any, K comparable, A any](fab map[K]func(A) B, a A) map[K]B {
	_ = "STUB: not implemented"
	return nil
}

func Flap[B any, K comparable, A any](a A) func(map[K]func(A) B) map[K]B {
	_ = "STUB: not implemented"
	return nil
}

// Copy creates a shallow copy of the map
func Copy[K comparable, V any](m map[K]V) map[K]V { _ = "STUB: not implemented"; return nil }

// Clone creates a deep copy of the map using the provided endomorphism to clone the values
func Clone[K comparable, V any](f EM.Endomorphism[V]) EM.Endomorphism[map[K]V] {
	_ = "STUB: not implemented"
	return nil

	// FromFoldableMap converts from a reducer to a map
	// Duplicate keys are resolved by the provided [Mg.Magma]
}

func FromFoldableMap[
	FOLDABLE ~func(func(map[K]V, A) map[K]V, map[K]V) func(HKTA) map[K]V, // the reduce function
	A any,
	HKTA any,
	K comparable,
	V any](m Mg.Magma[V], red FOLDABLE) func(f func(A) T.Tuple2[K, V]) func(fa HKTA) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FromArrayMap converts from an array to a map
// Duplicate keys are resolved by the provided [Mg.Magma]
func FromArrayMap[
	A any,
	K comparable,
	V any](m Mg.Magma[V]) func(f func(A) T.Tuple2[K, V]) func(fa []A) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FromFoldable converts from a reducer to a map
// Duplicate keys are resolved by the provided [Mg.Magma]
func FromFoldable[
	HKTA any,
	FOLDABLE ~func(func(map[K]V, T.Tuple2[K, V]) map[K]V, map[K]V) func(HKTA) map[K]V, // the reduce function
	K comparable,
	V any](m Mg.Magma[V], red FOLDABLE) func(fa HKTA) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FromArray converts from an array to a map
// Duplicate keys are resolved by the provided [Mg.Magma]
func FromArray[
	K comparable,
	V any](m Mg.Magma[V]) func(fa []T.Tuple2[K, V]) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[A any, K comparable, B any](m Mo.Monoid[map[K]B], fab map[K]func(A) B, fa map[K]A) map[K]B {
	_ = "STUB: not implemented"
	return nil
}

func Ap[A any, K comparable, B any](m Mo.Monoid[map[K]B]) func(fa map[K]A) func(map[K]func(A) B) map[K]B {
	_ = "STUB: not implemented"
	return nil
}
