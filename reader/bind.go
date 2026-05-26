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

package reader

// Bind creates an empty context of type [S] to be used with the [Bind] operation
func Do[R, S any](
	empty S,
) Reader[R, S] {
	_ = "STUB: not implemented"
	return nil
}

// Bind attaches the result of a computation to a context [S1] to produce a context [S2]
func Bind[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) Reader[R, T],
) func(Reader[R, S1]) Reader[R, S2] {
	_ = "STUB: not implemented"
	return nil
}

// Let attaches the result of a computation to a context [S1] to produce a context [S2]
func Let[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) T,
) func(Reader[R, S1]) Reader[R, S2] {
	_ = "STUB: not implemented"
	return nil
}

// LetTo attaches the a value to a context [S1] to produce a context [S2]
func LetTo[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	b T,
) func(Reader[R, S1]) Reader[R, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindTo initializes a new state [S1] from a value [T]
func BindTo[R, S1, T any](
	setter func(T) S1,
) func(Reader[R, T]) Reader[R, S1] {
	_ = "STUB: not implemented"
	return nil
}

// ApS attaches a value to a context [S1] to produce a context [S2] by considering the context and the value concurrently
func ApS[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa Reader[R, T],
) func(Reader[R, S1]) Reader[R, S2] {
	_ = "STUB: not implemented"
	return nil
}
