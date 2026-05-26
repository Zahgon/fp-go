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
//go:build either_pointers

package either

import "fmt"

type Either[E, A any] struct {
	left  *E
	right *A
}

// String prints some debug info for the object
//
//go:noinline
func eitherString[E, A any](s *Either[E, A]) string { _ = "STUB: not implemented"; return "" }

// Format prints some debug info for the object
//
//go:noinline
func eitherFormat[E, A any](e *Either[E, A], f fmt.State, c rune) {
	_ = "STUB: not implemented"
	return
}

// String prints some debug info for the object
func (s Either[E, A]) String() string { _ = "STUB: not implemented"; return "" }

// Format prints some debug info for the object
func (s Either[E, A]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

//go:inline
func Left[A, E any](value E) Either[E, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func Right[E, A any](value A) Either[E, A] { _ = "STUB: not implemented"; return nil }

//go:inline
func IsLeft[E, A any](e Either[E, A]) bool { _ = "STUB: not implemented"; return false }

//go:inline
func IsRight[E, A any](e Either[E, A]) bool { _ = "STUB: not implemented"; return false }

//go:inline
func MonadFold[E, A, B any](ma Either[E, A], onLeft func(E) B, onRight func(A) B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

//go:inline
func Unwrap[E, A any](ma Either[E, A]) (A, E) { _ = "STUB: not implemented"; return *new(A), *new(E) }
