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
)

// Next returns the iterator for the next element in an iterator `Pair`
func Next[GU ~func() Option[Pair[GU, U]], U any](m Pair[GU, U]) GU {
	_ = "STUB: not implemented"

	// Current returns the current element in an iterator `Pair`
	return *new(GU)
}

func Current[GU ~func() Option[Pair[GU, U]], U any](m Pair[GU, U]) U {
	_ = "STUB: not implemented"

	// From constructs an array from a set of variadic arguments
	return *new(U)
}

func From[GU ~func() Option[Pair[GU, U]], U any](data ...U) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

// Empty returns the empty iterator
func Empty[GU ~func() Option[Pair[GU, U]], U any]() GU {
	_ = "STUB: not implemented"
	return *

	// Of returns an iterator with one single element
	new(GU)
}

func Of[GU ~func() Option[Pair[GU, U]], U any](a U) GU { _ = "STUB: not implemented"; return *new(GU) }

// FromArray returns an iterator from multiple elements
func FromArray[GU ~func() Option[Pair[GU, U]], US ~[]U, U any](as US) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

// reduce applies a function for each value of the iterator with a floating result
func reduce[GU ~func() Option[Pair[GU, U]], U, V any](as GU, f func(V, U) V, initial V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

// next (with bad side effect)

// Reduce applies a function for each value of the iterator with a floating result
func Reduce[GU ~func() Option[Pair[GU, U]], U, V any](f func(V, U) V, initial V) func(GU) V {
	_ = "STUB: not implemented"
	return nil
}

// ToArray converts the iterator to an array
func ToArray[GU ~func() Option[Pair[GU, U]], US ~[]U, U any](u GU) US {
	_ = "STUB: not implemented"
	return *new(US)
}

func Map[GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], FCT ~func(U) V, U, V any](f FCT) func(ma GU) GV {
	_ = "STUB: not implemented"
	// pre-declare to avoid cyclic reference
	return nil
}

func MonadMap[GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], U, V any](ma GU, f func(U) V) GV {
	_ = "STUB: not implemented"
	return *new(GV)
}

func concat[GU ~func() Option[Pair[GU, U]], U any](right, left GU) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

func Chain[GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], U, V any](f func(U) GV) func(GU) GV {
	_ = "STUB: not implemented"
	// pre-declare to avoid cyclic reference
	return nil
}

func MonadChain[GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], U, V any](ma GU, f func(U) GV) GV {
	_ = "STUB: not implemented"
	return *new(GV)
}

func MonadChainFirst[GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], U, V any](ma GU, f func(U) GV) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

func ChainFirst[GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], U, V any](f func(U) GV) func(GU) GU {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GV ~func() Option[Pair[GV, GU]], GU ~func() Option[Pair[GU, U]], U any](ma GV) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

// MakeBy returns an [Iterator] with an infinite number of elements initialized with `f(i)`
func MakeBy[GU ~func() Option[Pair[GU, U]], FCT ~func(int) U, U any](f FCT) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

// bootstrap

// Replicate creates an infinite [Iterator] containing a value.
func Replicate[GU ~func() Option[Pair[GU, U]], U any](a U) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

// Repeat creates an [Iterator] containing a value repeated the specified number of times.
// Alias of [Replicate] combined with [Take]
func Repeat[GU ~func() Option[Pair[GU, U]], U any](n int, a U) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

// Count creates an [Iterator] containing a consecutive sequence of integers starting with the provided start value
func Count[GU ~func() Option[Pair[GU, int]]](start int) GU {
	_ = "STUB: not implemented"
	return *new(GU)
}

func FilterMap[GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], FCT ~func(U) Option[V], U, V any](f FCT) func(ma GU) GV {
	_ = "STUB: not implemented"
	// pre-declare to avoid cyclic reference
	return nil
}

func Filter[GU ~func() Option[Pair[GU, U]], FCT ~Predicate[U], U any](f FCT) func(ma GU) GU {
	_ = "STUB: not implemented"
	return nil
}

func Ap[GUV ~func() Option[Pair[GUV, func(U) V]], GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], U, V any](ma GU) func(fab GUV) GV {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[GUV ~func() Option[Pair[GUV, func(U) V]], GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], U, V any](fab GUV, ma GU) GV {
	_ = "STUB: not implemented"
	return *new(GV)
}

func FilterChain[GVV ~func() Option[Pair[GVV, GV]], GV ~func() Option[Pair[GV, V]], GU ~func() Option[Pair[GU, U]], FCT ~func(U) Option[GV], U, V any](f FCT) func(ma GU) GV {
	_ = "STUB: not implemented"
	return nil
}

func FoldMap[GU ~func() Option[Pair[GU, U]], FCT ~func(U) V, U, V any](m M.Monoid[V]) func(FCT) func(ma GU) V {
	_ = "STUB: not implemented"
	return nil
}

func Fold[GU ~func() Option[Pair[GU, U]], U any](m M.Monoid[U]) func(ma GU) U {
	_ = "STUB: not implemented"
	return nil
}
