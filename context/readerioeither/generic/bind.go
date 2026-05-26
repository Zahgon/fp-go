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
	"context"

	ET "github.com/IBM/fp-go/either"
)

// Bind creates an empty context of type [S] to be used with the [Bind] operation
func Do[GRS ~func(context.Context) GS, GS ~func() ET.Either[error, S], S any](
	empty S,
) GRS {
	_ = "STUB: not implemented"
	return *new(GRS)
}

// Bind attaches the result of a computation to a context [S1] to produce a context [S2]
func Bind[GRS1 ~func(context.Context) GS1, GRS2 ~func(context.Context) GS2, GRT ~func(context.Context) GT, GS1 ~func() ET.Either[error, S1], GS2 ~func() ET.Either[error, S2], GT ~func() ET.Either[error, T], S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) GRT,
) func(GRS1) GRS2 {
	_ = "STUB: not implemented"
	return nil
}

// Let attaches the result of a computation to a context [S1] to produce a context [S2]
func Let[GRS1 ~func(context.Context) GS1, GRS2 ~func(context.Context) GS2, GS1 ~func() ET.Either[error, S1], GS2 ~func() ET.Either[error, S2], S1, S2, T any](
	key func(T) func(S1) S2,
	f func(S1) T,
) func(GRS1) GRS2 {
	_ = "STUB: not implemented"
	return nil
}

// LetTo attaches the a value to a context [S1] to produce a context [S2]
func LetTo[GRS1 ~func(context.Context) GS1, GRS2 ~func(context.Context) GS2, GS1 ~func() ET.Either[error, S1], GS2 ~func() ET.Either[error, S2], S1, S2, B any](
	key func(B) func(S1) S2,
	b B,
) func(GRS1) GRS2 {
	_ = "STUB: not implemented"
	return nil
}

// BindTo initializes a new state [S1] from a value [T]
func BindTo[GRS1 ~func(context.Context) GS1, GRT ~func(context.Context) GT, GS1 ~func() ET.Either[error, S1], GT ~func() ET.Either[error, T], S1, T any](
	setter func(T) S1,
) func(GRT) GRS1 {
	_ = "STUB: not implemented"
	return nil
}

// ApS attaches a value to a context [S1] to produce a context [S2] by considering the context and the value concurrently
func ApS[GRTS1 ~func(context.Context) GTS1, GRS1 ~func(context.Context) GS1, GRS2 ~func(context.Context) GS2, GRT ~func(context.Context) GT, GTS1 ~func() ET.Either[error, func(T) S2], GS1 ~func() ET.Either[error, S1], GS2 ~func() ET.Either[error, S2], GT ~func() ET.Either[error, T], S1, S2, T any](
	setter func(T) func(S1) S2,
	fa GRT,
) func(GRS1) GRS2 {
	_ = "STUB: not implemented"
	return nil
}
