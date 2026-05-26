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

package function

// Identity returns the value 'a'
func Identity[A any](a A) A {
	_ = "STUB: not implemented"

	// Constant creates a nullary function that returns the constant value 'a'
	return *new(A)
}

func Constant[A any](a A) func() A { _ = "STUB: not implemented"; return nil }

// Constant1 creates a unary function that returns the constant value 'a' and ignores its input
func Constant1[B, A any](a A) func(B) A { _ = "STUB: not implemented"; return nil }

// Constant2 creates a binary function that returns the constant value 'a' and ignores its inputs
func Constant2[B, C, A any](a A) func(B, C) A { _ = "STUB: not implemented"; return nil }

func IsNil[A any](a *A) bool { _ = "STUB: not implemented"; return false }

func IsNonNil[A any](a *A) bool {
	_ = "STUB: not implemented"

	// Swap returns a new binary function that changes the order of input parameters
	return false
}

func Swap[T1, T2, R any](f func(T1, T2) R) func(T2, T1) R { _ = "STUB: not implemented"; return nil }

// First returns the first out of two input values
func First[T1, T2 any](t1 T1, _ T2) T1 {
	_ = "STUB: not implemented"

	// Second returns the second out of two input values
	// Identical to [SK]
	return *new(T1)
}

func Second[T1, T2 any](_ T1, t2 T2) T2 { _ = "STUB: not implemented"; return *new(T2) }
