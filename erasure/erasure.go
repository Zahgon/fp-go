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

package erasure

import (
	E "github.com/IBM/fp-go/either"
)

// Erase converts a variable of type T to an any by returning a pointer to that variable
func Erase[T any](t T) any {
	_ = "STUB: not implemented"

	// Unerase converts an erased variable back to its original value
	return *new(any)
}

func Unerase[T any](t any) T {
	_ = "STUB: not implemented"

	// SafeUnerase converts an erased variable back to its original value
	return *new(T)
}

func SafeUnerase[T any](t any) E.Either[error, T] { _ = "STUB: not implemented"; return nil }

// Erase0 converts a type safe function into an erased function
func Erase0[T1 any](f func() T1) func() any { _ = "STUB: not implemented"; return nil }

// Erase1 converts a type safe function into an erased function
func Erase1[T1, T2 any](f func(T1) T2) func(any) any { _ = "STUB: not implemented"; return nil }

// Erase2 converts a type safe function into an erased function
func Erase2[T1, T2, T3 any](f func(T1, T2) T3) func(any, any) any {
	_ = "STUB: not implemented"
	return nil
}
