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

package endomorphism

import (
	M "github.com/IBM/fp-go/monoid"
	S "github.com/IBM/fp-go/semigroup"
)

// Endomorphism is a function  that
type Endomorphism[A any] func(A) A

// Of converts any function to an [Endomorphism]
func Of[F ~func(A) A, A any](f F) Endomorphism[A] { _ = "STUB: not implemented"; return nil }

// Wrap converts any function to an [Endomorphism]
func Wrap[F ~func(A) A, A any](f F) Endomorphism[A] { _ = "STUB: not implemented"; return nil }

// Unwrap converts any [Endomorphism] to a function
func Unwrap[F ~func(A) A, A any](f Endomorphism[A]) F {
	_ = "STUB: not implemented"
	return *

	// Identity returns the identity [Endomorphism]
	new(F)
}

func Identity[A any]() Endomorphism[A] { _ = "STUB: not implemented"; return nil }

// Semigroup for the Endomorphism where the `concat` operation is the usual function composition.
func Semigroup[A any]() S.Semigroup[Endomorphism[A]] { _ = "STUB: not implemented"; return nil }

// Monoid for the Endomorphism where the `concat` operation is the usual function composition.
func Monoid[A any]() M.Monoid[Endomorphism[A]] { _ = "STUB: not implemented"; return nil }
