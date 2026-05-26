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

var (
	// jsonNull is the cached representation of the `null` serialization in JSON
	jsonNull = []byte("null")
)

// Option defines a data structure that logically holds a value or not.
// It represents an optional value: every Option is either Some and contains a value,
// or None, and does not contain a value.
//
// Option is commonly used to represent the result of operations that may fail,
// as an alternative to returning nil pointers or using error values.
//
// Example:
//
//	var opt Option[int] = Some(42)  // Contains a value
//	var opt Option[int] = None[int]() // Contains no value
type Option[A any] struct {
	value  A
	isSome bool
}

type (
	Kleisli[A, B any]  = func(A) Option[B]
	Operator[A, B any] = Kleisli[Option[A], B]
)

// String implements fmt.Stringer for Option.
// Returns a human-readable string representation.
//
// Example:
//
//	Some(42).String() // "Some[int](42)"
//	None[int]().String() // "None[int]"
func (s Option[A]) String() string { _ = "STUB: not implemented"; return "" }

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

// IsNone checks if an Option is None (contains no value).
//
// Example:
//
//	opt := None[int]()
//	IsNone(opt) // true
//	opt := Some(42)
//	IsNone(opt) // false
//
//go:inline
func IsNone[T any](val Option[T]) bool {
	_ = "STUB: not implemented"

	// Some creates an Option that contains a value.
	//
	// Example:
	//
	//	opt := Some(42) // Option containing 42
	//	opt := Some("hello") // Option containing "hello"
	//
	//go:inline
	return false
}

func Some[T any](value T) Option[T] { _ = "STUB: not implemented"; return nil }

// Of creates an Option that contains a value.
// This is an alias for Some and is used in monadic contexts.
//
// Example:
//
//	opt := Of(42) // Option containing 42
//
//go:inline
func Of[T any](value T) Option[T] {
	_ = "STUB: not implemented"

	// None creates an Option that contains no value.
	//
	// Example:
	//
	//	opt := None[int]() // Empty Option of type int
	//	opt := None[string]() // Empty Option of type string
	//
	//go:inline
	return nil
}

func None[T any]() Option[T] { _ = "STUB: not implemented"; return nil }

// IsSome checks if an Option contains a value.
//
// Example:
//
//	opt := Some(42)
//	IsSome(opt) // true
//	opt := None[int]()
//	IsSome(opt) // false
//
//go:inline
func IsSome[T any](val Option[T]) bool {
	_ = "STUB: not implemented"

	// MonadFold performs a fold operation on an Option.
	// If the Option is Some, applies onSome to the value.
	// If the Option is None, calls onNone.
	//
	// Example:
	//
	//	opt := Some(42)
	//	result := MonadFold(opt,
	//	    func() string { return "no value" },
	//	    func(x int) string { return fmt.Sprintf("value: %d", x) },
	//	) // "value: 42"
	return false
}

func MonadFold[A, B any](ma Option[A], onNone func() B, onSome func(A) B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

// Unwrap extracts the value and presence flag from an Option.
// Returns the value and true if Some, or zero value and false if None.
//
// Example:
//
//	opt := Some(42)
//	val, ok := Unwrap(opt) // val = 42, ok = true
//	opt := None[int]()
//	val, ok := Unwrap(opt) // val = 0, ok = false
//
//go:inline
func Unwrap[A any](ma Option[A]) (A, bool) { _ = "STUB: not implemented"; return *new(A), false }
