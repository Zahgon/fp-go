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

package magma

type Magma[A any] interface {
	Concat(x A, y A) A
}

type magma[A any] struct {
	c func(A, A) A
}

func (m magma[A]) Concat(x A, y A) A { _ = "STUB: not implemented"; return *new(A) }

func MakeMagma[A any](c func(A, A) A) Magma[A] { _ = "STUB: not implemented"; return nil }

func Reverse[A any](m Magma[A]) Magma[A] { _ = "STUB: not implemented"; return nil }

func filterFirst[A any](p func(A) bool, c func(A, A) A, x A, y A) A {
	_ = "STUB: not implemented"
	return *new(A)
}

func filterSecond[A any](p func(A) bool, c func(A, A) A, x A, y A) A {
	_ = "STUB: not implemented"
	return *new(A)
}

func FilterFirst[A any](p func(A) bool) func(Magma[A]) Magma[A] {
	_ = "STUB: not implemented"
	return nil
}

func FilterSecond[A any](p func(A) bool) func(Magma[A]) Magma[A] {
	_ = "STUB: not implemented"
	return nil
}

func first[A any](x, _ A) A { _ = "STUB: not implemented"; return *new(A) }

func second[A any](_, y A) A { _ = "STUB: not implemented"; return *new(A) }

func First[A any]() Magma[A] { _ = "STUB: not implemented"; return nil }

func Second[A any]() Magma[A] { _ = "STUB: not implemented"; return nil }

func endo[A any](f func(A) A, c func(A, A) A, x, y A) A { _ = "STUB: not implemented"; return *new(A) }

func Endo[A any](f func(A) A) func(Magma[A]) Magma[A] { _ = "STUB: not implemented"; return nil }
