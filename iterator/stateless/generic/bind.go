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

import (
	O "github.com/IBM/fp-go/option"
	P "github.com/IBM/fp-go/pair"
)

// Bind creates an empty context of type [S] to be used with the [Bind] operation
func Do[GS ~func() O.Option[P.Pair[GS, S]], S any](
	empty S,
) GS {
	_ = "STUB: not implemented"
	return *

	// Bind attaches the result of a computation to a context [S1] to produce a context [S2]
	new(GS)
}

func Bind[GS1 ~func() O.Option[P.Pair[GS1, S1]], GS2 ~func() O.Option[P.Pair[GS2, S2]], GA ~func() O.Option[P.Pair[GA, A]], S1, S2, A any](
	setter func(A) func(S1) S2,
	f func(S1) GA,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}

// Let attaches the result of a computation to a context [S1] to produce a context [S2]
func Let[GS1 ~func() O.Option[P.Pair[GS1, S1]], GS2 ~func() O.Option[P.Pair[GS2, S2]], S1, S2, A any](
	key func(A) func(S1) S2,
	f func(S1) A,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}

// LetTo attaches the a value to a context [S1] to produce a context [S2]
func LetTo[GS1 ~func() O.Option[P.Pair[GS1, S1]], GS2 ~func() O.Option[P.Pair[GS2, S2]], S1, S2, B any](
	key func(B) func(S1) S2,
	b B,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}

// BindTo initializes a new state [S1] from a value [T]
func BindTo[GS1 ~func() O.Option[P.Pair[GS1, S1]], GA ~func() O.Option[P.Pair[GA, A]], S1, A any](
	setter func(A) S1,
) func(GA) GS1 {
	_ = "STUB: not implemented"
	return nil
}

// ApS attaches a value to a context [S1] to produce a context [S2] by considering the context and the value concurrently
func ApS[GAS2 ~func() O.Option[P.Pair[GAS2, func(A) S2]], GS1 ~func() O.Option[P.Pair[GS1, S1]], GS2 ~func() O.Option[P.Pair[GS2, S2]], GA ~func() O.Option[P.Pair[GA, A]], S1, S2, A any](
	setter func(A) func(S1) S2,
	fa GA,
) func(GS1) GS2 {
	_ = "STUB: not implemented"
	return nil
}
