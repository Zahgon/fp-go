// Copyright (c) 2025 IBM Corp.
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
	M "github.com/IBM/fp-go/v2/monoid"
	S "github.com/IBM/fp-go/v2/semigroup"
)

// Semigroup returns a function that lifts a Semigroup over type A to a Semigroup over Option[A].
// The resulting semigroup combines two Options according to these rules:
//   - If both are Some, concatenates their values using the provided Semigroup
//   - If one is None, returns the other
//   - If both are None, returns None
//
// Example:
//
//	intSemigroup := semigroup.MakeSemigroup(func(a, b int) int { return a + b })
//	optSemigroup := Semigroup[int]()(intSemigroup)
//	optSemigroup.Concat(Some(2), Some(3)) // Some(5)
//	optSemigroup.Concat(Some(2), None[int]()) // Some(2)
//	optSemigroup.Concat(None[int](), Some(3)) // Some(3)
func Semigroup[A any]() func(S.Semigroup[A]) S.Semigroup[Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// Monoid returns a function that lifts a Semigroup over type A to a Monoid over Option[A].
// The monoid returns the left-most non-None value. If both operands are Some, their inner
// values are concatenated using the provided Semigroup. The empty value is None.
//
// Truth table:
//
//	| x       | y       | concat(x, y)       |
//	| ------- | ------- | ------------------ |
//	| none    | none    | none               |
//	| some(a) | none    | some(a)            |
//	| none    | some(b) | some(b)            |
//	| some(a) | some(b) | some(concat(a, b)) |
//
// Example:
//
//	intSemigroup := semigroup.MakeSemigroup(func(a, b int) int { return a + b })
//	optMonoid := Monoid[int]()(intSemigroup)
//	optMonoid.Concat(Some(2), Some(3)) // Some(5)
//	optMonoid.Empty() // None
func Monoid[A any]() func(S.Semigroup[A]) M.Monoid[Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// AlternativeMonoid creates a Monoid for Option[A] using the alternative semantics.
// This combines the applicative functor structure with the alternative (Alt) operation.
//
// Example:
//
//	intMonoid := monoid.MakeMonoid(func(a, b int) int { return a + b }, 0)
//	optMonoid := AlternativeMonoid(intMonoid)
//	result := optMonoid.Concat(Some(2), Some(3)) // Some(5)
//
//go:inline
func AlternativeMonoid[A any](m M.Monoid[A]) M.Monoid[Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// AltMonoid creates a Monoid for Option[A] using the Alt operation.
// This monoid returns the first Some value, or None if both are None.
// The empty value is None.
//
// Example:
//
//	optMonoid := AltMonoid[int]()
//	optMonoid.Concat(Some(2), Some(3)) // Some(2) - returns first Some
//	optMonoid.Concat(None[int](), Some(3)) // Some(3)
//	optMonoid.Empty() // None
//
//go:inline
func AltMonoid[A any]() M.Monoid[Option[A]] { _ = "STUB: not implemented"; return nil }

// takeFirst is a helper function that returns the first Some value, or the second if the first is None.
func takeFirst[A any](l, r Option[A]) Option[A] { _ = "STUB: not implemented"; return nil }

// FirstMonoid creates a Monoid for Option[A] that returns the first Some value.
// This monoid prefers the left operand when it is Some, otherwise returns the right operand.
// The empty value is None.
//
// This is equivalent to AltMonoid but implemented more directly.
//
// Truth table:
//
//	| x       | y       | concat(x, y) |
//	| ------- | ------- | ------------ |
//	| none    | none    | none         |
//	| some(a) | none    | some(a)      |
//	| none    | some(b) | some(b)      |
//	| some(a) | some(b) | some(a)      |
//
// Example:
//
//	optMonoid := FirstMonoid[int]()
//	optMonoid.Concat(Some(2), Some(3)) // Some(2) - returns first Some
//	optMonoid.Concat(None[int](), Some(3)) // Some(3)
//	optMonoid.Concat(Some(2), None[int]()) // Some(2)
//	optMonoid.Empty() // None
//
//go:inline
func FirstMonoid[A any]() M.Monoid[Option[A]] { _ = "STUB: not implemented"; return nil }

// takeLast is a helper function that returns the last Some value, or the first if the last is None.
func takeLast[A any](l, r Option[A]) Option[A] { _ = "STUB: not implemented"; return nil }

// LastMonoid creates a Monoid for Option[A] that returns the last Some value.
// This monoid prefers the right operand when it is Some, otherwise returns the left operand.
// The empty value is None.
//
// Truth table:
//
//	| x       | y       | concat(x, y) |
//	| ------- | ------- | ------------ |
//	| none    | none    | none         |
//	| some(a) | none    | some(a)      |
//	| none    | some(b) | some(b)      |
//	| some(a) | some(b) | some(b)      |
//
// Example:
//
//	optMonoid := LastMonoid[int]()
//	optMonoid.Concat(Some(2), Some(3)) // Some(3) - returns last Some
//	optMonoid.Concat(None[int](), Some(3)) // Some(3)
//	optMonoid.Concat(Some(2), None[int]()) // Some(2)
//	optMonoid.Empty() // None
//
//go:inline
func LastMonoid[A any]() M.Monoid[Option[A]] { _ = "STUB: not implemented"; return nil }
