// Copyright (c) 2024 IBM Corp.
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

package pair

import (
	"fmt"

	Sg "github.com/IBM/fp-go/semigroup"
	T "github.com/IBM/fp-go/tuple"
)

type (
	pair struct {
		h, t any
	}

	// Pair defines a data structure that holds two strongly typed values
	Pair[A, B any] pair
)

// String prints some debug info for the object
//
//go:noinline
func pairString(s *pair) string { _ = "STUB: not implemented"; return "" }

// Format prints some debug info for the object
//
//go:noinline
func pairFormat(e *pair, f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// String prints some debug info for the object
func (s Pair[A, B]) String() string { _ = "STUB: not implemented"; return "" }

// Format prints some debug info for the object
func (s Pair[A, B]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// Of creates a [Pair] with the same value to to both fields
func Of[A any](value A) Pair[A, A] { _ = "STUB: not implemented"; return nil }

// FromTuple creates a [Pair] from a [T.Tuple2]
func FromTuple[A, B any](t T.Tuple2[A, B]) Pair[A, B] { _ = "STUB: not implemented"; return nil }

// ToTuple creates a [T.Tuple2] from a [Pair]
func ToTuple[A, B any](t Pair[A, B]) T.Tuple2[A, B] { _ = "STUB: not implemented"; return nil }

// MakePair creates a [Pair] from two values
func MakePair[A, B any](a A, b B) Pair[A, B] { _ = "STUB: not implemented"; return nil }

// Head returns the head value of the pair
func Head[A, B any](fa Pair[A, B]) A {
	_ = "STUB: not implemented"

	// Tail returns the head value of the pair
	return *new(A)
}

func Tail[A, B any](fa Pair[A, B]) B {
	_ = "STUB: not implemented"

	// MonadMapHead maps the head value
	return *new(B)
}

func MonadMapHead[B, A, A1 any](fa Pair[A, B], f func(A) A1) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadMap maps the head value
func MonadMap[B, A, A1 any](fa Pair[A, B], f func(A) A1) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil

	// MonadMapTail maps the Tail value
}

func MonadMapTail[A, B, B1 any](fa Pair[A, B], f func(B) B1) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadBiMap maps both values
func MonadBiMap[A, B, A1, B1 any](fa Pair[A, B], f func(A) A1, g func(B) B1) Pair[A1, B1] {
	_ = "STUB: not implemented"
	return nil
}

// Map maps the head value
func Map[B, A, A1 any](f func(A) A1) func(Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MapHead maps the head value
func MapHead[B, A, A1 any](f func(A) A1) func(Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MapTail maps the Tail value
func MapTail[A, B, B1 any](f func(B) B1) func(Pair[A, B]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// BiMap maps both values
func BiMap[A, B, A1, B1 any](f func(A) A1, g func(B) B1) func(Pair[A, B]) Pair[A1, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainHead chains on the head value
func MonadChainHead[B, A, A1 any](sg Sg.Semigroup[B], fa Pair[A, B], f func(A) Pair[A1, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainTail chains on the Tail value
func MonadChainTail[A, B, B1 any](sg Sg.Semigroup[A], fb Pair[A, B], f func(B) Pair[A, B1]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChain chains on the head value
func MonadChain[B, A, A1 any](sg Sg.Semigroup[B], fa Pair[A, B], f func(A) Pair[A1, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainHead chains on the head value
func ChainHead[B, A, A1 any](sg Sg.Semigroup[B], f func(A) Pair[A1, B]) func(Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainTail chains on the Tail value
func ChainTail[A, B, B1 any](sg Sg.Semigroup[A], f func(B) Pair[A, B1]) func(Pair[A, B]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// Chain chains on the head value
func Chain[B, A, A1 any](sg Sg.Semigroup[B], f func(A) Pair[A1, B]) func(Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApHead applies on the head value
func MonadApHead[B, A, A1 any](sg Sg.Semigroup[B], faa Pair[func(A) A1, B], fa Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApTail applies on the Tail value
func MonadApTail[A, B, B1 any](sg Sg.Semigroup[A], fbb Pair[A, func(B) B1], fb Pair[A, B]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAp applies on the head value
func MonadAp[B, A, A1 any](sg Sg.Semigroup[B], faa Pair[func(A) A1, B], fa Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApHead applies on the head value
func ApHead[B, A, A1 any](sg Sg.Semigroup[B], fa Pair[A, B]) func(Pair[func(A) A1, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApTail applies on the Tail value
func ApTail[A, B, B1 any](sg Sg.Semigroup[A], fb Pair[A, B]) func(Pair[A, func(B) B1]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// Ap applies on the head value
func Ap[B, A, A1 any](sg Sg.Semigroup[B], fa Pair[A, B]) func(Pair[func(A) A1, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// Swap swaps the two channels
func Swap[A, B any](fa Pair[A, B]) Pair[B, A] { _ = "STUB: not implemented"; return nil }

// Paired converts a function with 2 parameters into a function taking a [Pair]
// The inverse function is [Unpaired]
func Paired[F ~func(T1, T2) R, T1, T2, R any](f F) func(Pair[T1, T2]) R {
	_ = "STUB: not implemented"
	return nil
}

// Unpaired converts a function with a [Pair] parameter into a function with 2 parameters
// The inverse function is [Paired]
func Unpaired[F ~func(Pair[T1, T2]) R, T1, T2, R any](f F) func(T1, T2) R {
	_ = "STUB: not implemented"
	return nil
}

func Merge[F ~func(B) func(A) R, A, B, R any](f F) func(Pair[A, B]) R {
	_ = "STUB: not implemented"
	return nil
}
