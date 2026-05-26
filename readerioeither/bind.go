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

package readerioeither

// Bind creates an empty context of type [S] to be used with the [Bind] operation
func Do[R, E, S any](
	empty S,
) ReaderIOEither[R, E, S] {
	_ = "STUB: not implemented"
	return nil
}

// Bind attaches the result of a computation to a context [S1] to produce a context [S2]
func Bind[R, E, S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) ReaderIOEither[R, E, T],
) func(ReaderIOEither[R, E, S1]) ReaderIOEither[R, E, S2] {
	_ = "STUB: not implemented"
	return nil
}

// Let attaches the result of a computation to a context [S1] to produce a context [S2]
func Let[R, E, S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) T,
) func(ReaderIOEither[R, E, S1]) ReaderIOEither[R, E, S2] {
	_ = "STUB: not implemented"
	return nil
}

// LetTo attaches the a value to a context [S1] to produce a context [S2]
func LetTo[R, E, S1, S2, T any](
	setter func(T) func(S1) S2,
	b T,
) func(ReaderIOEither[R, E, S1]) ReaderIOEither[R, E, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindTo initializes a new state [S1] from a value [T]
func BindTo[R, E, S1, T any](
	setter func(T) S1,
) func(ReaderIOEither[R, E, T]) ReaderIOEither[R, E, S1] {
	_ = "STUB: not implemented"
	return nil
}

// ApS attaches a value to a context [S1] to produce a context [S2] by considering the context and the value concurrently
func ApS[R, E, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa ReaderIOEither[R, E, T],
) func(ReaderIOEither[R, E, S1]) ReaderIOEither[R, E, S2] {
	_ = "STUB: not implemented"
	return nil
}
