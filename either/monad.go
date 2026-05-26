// Copyright (c) 2024 IBM Corp.
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

package either

import (
	"github.com/IBM/fp-go/internal/monad"
)

type eitherMonad[E, A, B any] struct{}

func (o *eitherMonad[E, A, B]) Of(a A) Either[E, A] { _ = "STUB: not implemented"; return nil }

func (o *eitherMonad[E, A, B]) Map(f func(A) B) func(Either[E, A]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func (o *eitherMonad[E, A, B]) Chain(f func(A) Either[E, B]) func(Either[E, A]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func (o *eitherMonad[E, A, B]) Ap(fa Either[E, A]) func(Either[E, func(A) B]) Either[E, B] {
	_ = "STUB: not implemented"
	return nil

	// Monad implements the monadic operations for [Either]
}

func Monad[E, A, B any]() monad.Monad[A, B, Either[E, A], Either[E, B], Either[E, func(A) B]] {
	_ = "STUB: not implemented"
	return nil
}
