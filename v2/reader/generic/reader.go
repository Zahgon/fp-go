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

// Package generic provides generic implementations of Reader operations that work with
// custom reader types. These functions use Go's type constraints to work with any type
// that matches the Reader pattern (func(R) A).
//
// Most functions in this package are deprecated in favor of the non-generic versions
// in the parent reader package, which provide better type inference and simpler usage.
//
// Use this package when you need to work with custom reader types or when you need
// explicit control over type parameters.
package generic

import (
	T "github.com/IBM/fp-go/v2/tuple"
)

// Reader[R, A] = func(R) A

// MakeReader creates a reader, i.e. a method that accepts a context and that returns a value
//
// Deprecated:
func MakeReader[GA ~func(R) A, R, A any](r GA) GA {
	_ = "STUB: not implemented"

	// Ask reads the current context
	//
	// Deprecated:
	return *new(GA)
}

func Ask[GR ~func(R) R, R any]() GR { _ = "STUB: not implemented"; return *new(GR) }

// Asks projects a value from the global context in a Reader
//
// Deprecated:
func Asks[GA ~func(R) A, R, A any](f GA) GA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GA)
}

func AsksReader[GA ~func(R) A, R, A any](f func(R) GA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func MonadMap[GA ~func(E) A, GB ~func(E) B, E, A, B any](fa GA, f func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Map can be used to turn functions `func(A)B` into functions `(fa F[A])F[B]` whose argument and return types
// use the type constructor `F` to represent some computational context.
//
// Deprecated:
func Map[GA ~func(E) A, GB ~func(E) B, E, A, B any](f func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadAp[GA ~func(R) A, GB ~func(R) B, GAB ~func(R) func(A) B, R, A, B any](fab GAB, fa GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Ap applies a function to an argument under a type constructor.
//
// Deprecated:
func Ap[GA ~func(R) A, GB ~func(R) B, GAB ~func(R) func(A) B, R, A, B any](fa GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func Of[GA ~func(R) A, R, A any](a A) GA {
	_ = "STUB: not implemented"
	return *

	// Deprecated:
	new(GA)
}

func MonadChain[GA ~func(R) A, GB ~func(R) B, R, A, B any](ma GA, f func(A) GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Chain composes computations in sequence, using the return value of one computation to determine the next computation.
//
// Deprecated:
func Chain[GA ~func(R) A, GB ~func(R) B, R, A, B any](f func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func Flatten[GA ~func(R) A, GGA ~func(R) GA, R, A any](mma GGA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Deprecated:
func Compose[AB ~func(A) B, BC ~func(B) C, AC ~func(A) C, A, B, C any](ab AB) func(BC) AC {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func First[GAB ~func(A) B, GABC ~func(T.Tuple2[A, C]) T.Tuple2[B, C], A, B, C any](pab GAB) GABC {
	_ = "STUB: not implemented"
	return *new(GABC)
}

// Deprecated:
func Second[GBC ~func(B) C, GABC ~func(T.Tuple2[A, B]) T.Tuple2[A, C], A, B, C any](pbc GBC) GABC {
	_ = "STUB: not implemented"
	return *new(GABC)
}

// Deprecated:
func Promap[GA ~func(E) A, GB ~func(D) B, E, A, D, B any](f func(D) E, g func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Local changes the value of the local context during the execution of the action `ma` (similar to `Contravariant`'s
// `contramap`).
//
// Deprecated:
func Local[GA1 ~func(R1) A, GA2 ~func(R2) A, R2, R1, A any](f func(R2) R1) func(GA1) GA2 {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated:
func MonadFlap[GAB ~func(R) func(A) B, GB ~func(R) B, R, A, B any](fab GAB, a A) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Deprecated:
func Flap[GAB ~func(R) func(A) B, GB ~func(R) B, R, A, B any](a A) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}
