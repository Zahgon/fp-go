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

package generic

// Bind creates an empty context of type [S] to be used with the [Bind] operation
func Do[GS ~func() S, S any](
	empty S,
) GS {
	_ = "STUB: not implemented"
	return *

	// Bind attaches the result of a computation to a context [S1] to produce a context [S2]
	new(GS)
}

func Bind[GS1 ~func() S1, GS2 ~func() S2, GT ~func() T, S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) GT,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}

// Let attaches the result of a computation to a context [S1] to produce a context [S2]
func Let[GS1 ~func() S1, GS2 ~func() S2, S1, S2, T any](
	key func(T) func(S1) S2,
	f func(S1) T,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}

// LetTo attaches the a value to a context [S1] to produce a context [S2]
func LetTo[GS1 ~func() S1, GS2 ~func() S2, S1, S2, B any](
	key func(B) func(S1) S2,
	b B,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}

// BindTo initializes a new state [S1] from a value [T]
func BindTo[GS1 ~func() S1, GT ~func() T, S1, T any](
	setter func(T) S1,
) func(GT) GS1 {
	_ = "STUB: not implemented"
	return nil
}

// ApS attaches a value to a context [S1] to produce a context [S2] by considering the context and the value concurrently
func ApS[GS1 ~func() S1, GS2 ~func() S2, GT ~func() T, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa GT,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}
