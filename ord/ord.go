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

package ord

import (
	C "github.com/IBM/fp-go/constraints"
	E "github.com/IBM/fp-go/eq"
)

type Ord[T any] interface {
	E.Eq[T]
	Compare(x, y T) int
}

type ord[T any] struct {
	c func(x, y T) int
	e func(x, y T) bool
}

func (self ord[T]) Equals(x, y T) bool { _ = "STUB: not implemented"; return false }

func (self ord[T]) Compare(x, y T) int { _ = "STUB: not implemented"; return 0 }

// ToEq converts an [Ord] to [E.Eq]
func ToEq[T any](o Ord[T]) E.Eq[T] {
	_ = "STUB: not implemented"

	// MakeOrd creates an instance of an Ord
	return nil
}

func MakeOrd[T any](c func(x, y T) int, e func(x, y T) bool) Ord[T] {
	_ = "STUB: not implemented"
	return nil

	// MakeOrd creates an instance of an Ord from a compare function
}

func FromCompare[T any](compare func(T, T) int) Ord[T] { _ = "STUB: not implemented"; return nil }

// Reverse creates an inverted ordering
func Reverse[T any](o Ord[T]) Ord[T] { _ = "STUB: not implemented"; return nil }

// Contramap creates an ordering under a transformation function
func Contramap[A, B any](f func(B) A) func(Ord[A]) Ord[B] { _ = "STUB: not implemented"; return nil }

// Min takes the minimum of two values. If they are considered equal, the first argument is chosen
func Min[A any](o Ord[A]) func(A, A) A { _ = "STUB: not implemented"; return nil }

// Max takes the maximum of two values. If they are considered equal, the first argument is chosen
func Max[A any](o Ord[A]) func(A, A) A { _ = "STUB: not implemented"; return nil }

// Clamp clamps a value between a minimum and a maximum
func Clamp[A any](o Ord[A]) func(A, A) func(A) A { _ = "STUB: not implemented"; return nil }

func strictCompare[A C.Ordered](a, b A) int { _ = "STUB: not implemented"; return 0 }

func strictEq[A comparable](a, b A) bool {
	_ = "STUB: not implemented"

	// FromStrictCompare implements the ordering based on the built in native order
	return false
}

func FromStrictCompare[A C.Ordered]() Ord[A] { _ = "STUB: not implemented"; return nil }

// Lt tests whether one value is strictly less than another
func Lt[A any](o Ord[A]) func(A) func(A) bool { _ = "STUB: not implemented"; return nil }

// Leq Tests whether one value is less or equal than another
func Leq[A any](O Ord[A]) func(A) func(A) bool { _ = "STUB: not implemented"; return nil }

/**
 * Test whether one value is strictly greater than another
 */
func Gt[A any](O Ord[A]) func(A) func(A) bool { _ = "STUB: not implemented"; return nil }

// Geq tests whether one value is greater or equal than another
func Geq[A any](O Ord[A]) func(A) func(A) bool { _ = "STUB: not implemented"; return nil }

// Between tests whether a value is between a minimum (inclusive) and a maximum (exclusive)
func Between[A any](O Ord[A]) func(A, A) func(A) bool { _ = "STUB: not implemented"; return nil }

// returns the predicate
