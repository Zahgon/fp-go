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
	"fmt"
)

var (
	// jsonNull is the cached representation of the `null` serialization in JSON
	jsonNull = []byte("null")
)

// Option defines a data structure that logically holds a value or not
type Option[A any] struct {
	isSome bool
	value  A
}

// optString prints some debug info for the object
//
//go:noinline
func optString(isSome bool, value any) string { _ = "STUB: not implemented"; return "" }

// optFormat prints some debug info for the object
//
//go:noinline
func optFormat(isSome bool, value any, f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// String prints some debug info for the object
func (s Option[A]) String() string { _ = "STUB: not implemented"; return "" }

// Format prints some debug info for the object
func (s Option[A]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

func optMarshalJSON(isSome bool, value any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s Option[A]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// optUnmarshalJSON unmarshals the [Option] from a JSON string
//
//go:noinline
func optUnmarshalJSON(isSome *bool, value any, data []byte) error {
	_ = "STUB: not implemented"
	// decode the value
	return nil
}

func (s *Option[A]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func IsNone[T any](val Option[T]) bool { _ = "STUB: not implemented"; return false }

func Some[T any](value T) Option[T] { _ = "STUB: not implemented"; return nil }

func Of[T any](value T) Option[T] { _ = "STUB: not implemented"; return nil }

func None[T any]() Option[T] { _ = "STUB: not implemented"; return nil }

func IsSome[T any](val Option[T]) bool { _ = "STUB: not implemented"; return false }

func MonadFold[A, B any](ma Option[A], onNone func() B, onSome func(A) B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func Unwrap[A any](ma Option[A]) (A, bool) { _ = "STUB: not implemented"; return *new(A), false }
