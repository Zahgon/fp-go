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

package array

import (
	M "github.com/IBM/fp-go/monoid"
	S "github.com/IBM/fp-go/semigroup"
)

func concat[T any](left, right []T) []T {
	_ = "STUB: not implemented"
	// some performance checks
	return nil
}

// need to copy

func Monoid[T any]() M.Monoid[[]T] { _ = "STUB: not implemented"; return nil }

func Semigroup[T any]() S.Semigroup[[]T] { _ = "STUB: not implemented"; return nil }

func addLen[A any](count int, data []A) int { _ = "STUB: not implemented"; return 0 }

// ConcatAll efficiently concatenates the input arrays into a final array
func ArrayConcatAll[A any](data ...[]A) []A {
	_ = "STUB: not implemented"
	// get the full size
	return nil
}

// copy

// returns the final array
