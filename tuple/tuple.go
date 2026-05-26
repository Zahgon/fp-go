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

// Package tuple contains type definitions and functions for data structures for tuples of heterogenous types. For homogeneous types
// consider to use arrays for simplicity
package tuple

func Of[T1 any](t T1) Tuple1[T1] { _ = "STUB: not implemented"; return nil }

// First returns the first element of a [Tuple2]
func First[T1, T2 any](t Tuple2[T1, T2]) T1 {
	_ = "STUB: not implemented"

	// Second returns the second element of a [Tuple2]
	return *new(T1)
}

func Second[T1, T2 any](t Tuple2[T1, T2]) T2 { _ = "STUB: not implemented"; return *new(T2) }

func Swap[T1, T2 any](t Tuple2[T1, T2]) Tuple2[T2, T1] { _ = "STUB: not implemented"; return nil }

func Of2[T1, T2 any](e T2) func(T1) Tuple2[T1, T2] { _ = "STUB: not implemented"; return nil }

func BiMap[E, G, A, B any](mapSnd func(E) G, mapFst func(A) B) func(Tuple2[A, E]) Tuple2[B, G] {
	_ = "STUB: not implemented"
	return nil
}

// marshalJSON marshals the tuple into a JSON array
func tupleMarshalJSON(src ...any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// tupleUnmarshalJSON unmarshals a JSON array into a tuple
		nil
}

func tupleUnmarshalJSON(data []byte, dst ...any) error { _ = "STUB: not implemented"; return nil }

// unmarshal

// successfully decoded the tuple

// tupleString converts a tuple to a string
func tupleString(src ...any) string { _ = "STUB: not implemented"; return "" }
