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

package stateless

import (
	M "github.com/IBM/fp-go/v2/monoid"
)

// Next returns the [Iterator] for the next element in an iterator [Pair]
func Next[U any](m Pair[Iterator[U], U]) Iterator[U] { _ = "STUB: not implemented"; return nil }

// Current returns the current element in an [Iterator] [Pair]
func Current[U any](m Pair[Iterator[U], U]) U {
	_ = "STUB: not implemented"
	return *

	// Empty returns the empty iterator
	new(U)
}

func Empty[U any]() Iterator[U] { _ = "STUB: not implemented"; return nil }

// Of returns an iterator with one single element
func Of[U any](a U) Iterator[U] { _ = "STUB: not implemented"; return nil }

// FromArray returns an iterator from multiple elements
func FromArray[U any](as []U) Iterator[U] { _ = "STUB: not implemented"; return nil }

// ToArray converts the iterator to an array
func ToArray[U any](u Iterator[U]) []U { _ = "STUB: not implemented"; return nil }

// Reduce applies a function for each value of the iterator with a floating result
func Reduce[U, V any](f func(V, U) V, initial V) func(Iterator[U]) V {
	_ = "STUB: not implemented"
	return nil
}

// MonadMap transforms an [Iterator] of type [U] into an [Iterator] of type [V] via a mapping function
func MonadMap[U, V any](ma Iterator[U], f func(U) V) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Map transforms an [Iterator] of type [U] into an [Iterator] of type [V] via a mapping function
func Map[U, V any](f func(U) V) Operator[U, V] { _ = "STUB: not implemented"; return nil }

func MonadChain[U, V any](ma Iterator[U], f Kleisli[U, V]) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[U, V any](f Kleisli[U, V]) Kleisli[Iterator[U], V] {
	_ = "STUB: not implemented"
	return nil
}

// Flatten converts an [Iterator] of [Iterator] into a simple [Iterator]
func Flatten[U any](ma Iterator[Iterator[U]]) Iterator[U] { _ = "STUB: not implemented"; return nil }

// From constructs an [Iterator] from a set of variadic arguments
func From[U any](data ...U) Iterator[U] { _ = "STUB: not implemented"; return nil }

// MakeBy returns an [Iterator] with an infinite number of elements initialized with `f(i)`
func MakeBy[FCT ~func(int) U, U any](f FCT) Iterator[U] { _ = "STUB: not implemented"; return nil }

// Replicate creates an [Iterator] containing a value repeated an infinite number of times.
func Replicate[U any](a U) Iterator[U] { _ = "STUB: not implemented"; return nil }

// FilterMap filters and transforms the content of an iterator
func FilterMap[U, V any](f func(U) Option[V]) Operator[U, V] { _ = "STUB: not implemented"; return nil }

// Filter filters the content of an iterator
func Filter[U any](f Predicate[U]) Operator[U, U] { _ = "STUB: not implemented"; return nil }

// Ap is the applicative functor for iterators
func Ap[V, U any](ma Iterator[U]) Operator[func(U) V, V] { _ = "STUB: not implemented"; return nil }

// MonadAp is the applicative functor for iterators
func MonadAp[V, U any](fab Iterator[func(U) V], ma Iterator[U]) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Repeat creates an [Iterator] containing a value repeated the specified number of times.
// Alias of [Replicate]
func Repeat[U any](n int, a U) Iterator[U] { _ = "STUB: not implemented"; return nil }

// Count creates an [Iterator] containing a consecutive sequence of integers starting with the provided start value
func Count(start int) Iterator[int] { _ = "STUB: not implemented"; return nil }

// FilterChain filters and transforms the content of an iterator
func FilterChain[U, V any](f func(U) Option[Iterator[V]]) Operator[U, V] {
	_ = "STUB: not implemented"
	return nil
}

// FoldMap maps and folds an iterator. Map the iterator passing each value to the iterating function. Then fold the results using the provided Monoid.
func FoldMap[U, V any](m M.Monoid[V]) func(func(U) V) func(ma Iterator[U]) V {
	_ = "STUB: not implemented"
	return nil
}

// Fold folds the iterator using the provided Monoid.
func Fold[U any](m M.Monoid[U]) func(Iterator[U]) U { _ = "STUB: not implemented"; return nil }

func MonadChainFirst[U, V any](ma Iterator[U], f Kleisli[U, V]) Iterator[U] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirst[U, V any](f Kleisli[U, V]) Operator[U, U] { _ = "STUB: not implemented"; return nil }
