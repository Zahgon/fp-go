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

package constant

import (
	M "github.com/IBM/fp-go/monoid"
	S "github.com/IBM/fp-go/semigroup"
)

type Const[E, A any] struct {
	value E
}

func Make[E, A any](e E) Const[E, A] { _ = "STUB: not implemented"; return nil }

func Unwrap[E, A any](c Const[E, A]) E { _ = "STUB: not implemented"; return *new(E) }

func Of[E, A any](m M.Monoid[E]) func(A) Const[E, A] { _ = "STUB: not implemented"; return nil }

func MonadMap[E, A, B any](fa Const[E, A], _ func(A) B) Const[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[E, A, B any](s S.Semigroup[E]) func(fab Const[E, func(A) B], fa Const[E, A]) Const[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[E, A, B any](f func(A) B) func(fa Const[E, A]) Const[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[E, A, B any](s S.Semigroup[E]) func(fa Const[E, A]) func(fab Const[E, func(A) B]) Const[E, B] {
	_ = "STUB: not implemented"
	return nil
}
