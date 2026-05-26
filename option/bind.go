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

package option

// Bind creates an empty context of type [S] to be used with the [Bind] operation
func Do[S any](
	empty S,
) Option[S] {
	_ = "STUB: not implemented"

	// Bind attaches the result of a computation to a context [S1] to produce a context [S2]
	return nil
}

func Bind[S1, S2, A any](
	setter func(A) func(S1) S2,
	f func(S1) Option[A],
) func(Option[S1]) Option[S2] {
	_ = "STUB: not implemented"
	return nil
}

// Let attaches the result of a computation to a context [S1] to produce a context [S2]
func Let[S1, S2, B any](
	key func(B) func(S1) S2,
	f func(S1) B,
) func(Option[S1]) Option[S2] {
	_ = "STUB: not implemented"
	return nil
}

// LetTo attaches the a value to a context [S1] to produce a context [S2]
func LetTo[S1, S2, B any](
	key func(B) func(S1) S2,
	b B,
) func(Option[S1]) Option[S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindTo initializes a new state [S1] from a value [T]
func BindTo[S1, T any](
	setter func(T) S1,
) func(Option[T]) Option[S1] {
	_ = "STUB: not implemented"
	return nil
}

// ApS attaches a value to a context [S1] to produce a context [S2] by considering the context and the value concurrently
func ApS[S1, S2, T any](
	setter func(T) func(S1) S2,
	fa Option[T],
) func(Option[S1]) Option[S2] {
	_ = "STUB: not implemented"
	return nil
}
