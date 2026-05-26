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

import (
	M "github.com/IBM/fp-go/monoid"
	S "github.com/IBM/fp-go/semigroup"
)

func Semigroup[A any]() func(S.Semigroup[A]) S.Semigroup[Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// Monoid returning the left-most non-`None` value. If both operands are `Some`s then the inner values are
// concatenated using the provided `Semigroup`
//
// | x       | y       | concat(x, y)       |
// | ------- | ------- | ------------------ |
// | none    | none    | none               |
// | some(a) | none    | some(a)            |
// | none    | some(b) | some(b)            |
// | some(a) | some(b) | some(concat(a, b)) |
func Monoid[A any]() func(S.Semigroup[A]) M.Monoid[Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// AlternativeMonoid is the alternative [Monoid] for an [Option]
func AlternativeMonoid[A any](m M.Monoid[A]) M.Monoid[Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// AltMonoid is the alternative [Monoid] for an [Option]
func AltMonoid[A any]() M.Monoid[Option[A]] { _ = "STUB: not implemented"; return nil }
