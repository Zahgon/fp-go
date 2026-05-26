// Copyright (c) 2023 - 2025 IBM Corp.
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

// Package string provides functional programming utilities for working with strings.
// It includes functions for string manipulation, comparison, conversion, and formatting,
// following functional programming principles with curried functions and composable operations.
package string

import (
	"strings"

	F "github.com/IBM/fp-go/v2/function"
	"github.com/IBM/fp-go/v2/ord"
)

var (
	// ToUpperCase converts the string to uppercase
	ToUpperCase = strings.ToUpper

	// ToLowerCase converts the string to lowercase
	ToLowerCase = strings.ToLower

	// Ord implements the default ordering for strings
	Ord = ord.FromStrictCompare[string]()

	// Join joins strings
	Join = F.Curry2(F.Bind2nd[[]string, string, string])(strings.Join)

	// Equals returns a predicate that tests if a string is equal
	Equals = F.Curry2(Eq)

	// Includes returns a predicate that tests for the existence of the search string
	Includes = F.Bind2of2(strings.Contains)

	// HasPrefix returns a predicate that checks if the prefix is included in the string
	HasPrefix = F.Bind2of2(strings.HasPrefix)

	// HasSuffix
	HasSuffix = F.Bind2of2(strings.HasSuffix)
)

// Eq tests if two strings are equal
func Eq(left, right string) bool { _ = "STUB: not implemented"; return false }

// ToBytes converts a string to a byte slice
func ToBytes(s string) []byte {
	_ = "STUB: not implemented"

	// ToRunes converts a string to a rune slice
	return nil
}

func ToRunes(s string) []rune {
	_ = "STUB: not implemented"

	// IsEmpty returns true if the string is empty
	//
	//go:inline
	return nil
}

func IsEmpty(s string) bool {
	_ = "STUB: not implemented"

	// IsNonEmpty returns true if the string is not empty
	//
	//go:inline
	return false
}

func IsNonEmpty(s string) bool {
	_ = "STUB: not implemented"

	// Size returns the length of the string in bytes
	//
	//go:inline
	return false
}

func Size(s string) int {
	_ = "STUB: not implemented"

	// Format applies a format string to an arbitrary value and returns a function
	// that formats values of type T using the provided format string
	return 0
}

func Format[T any](format string) func(T) string { _ = "STUB: not implemented"; return nil }

// Intersperse returns a function that concatenates two strings with a middle string in between.
// If either string is empty, the middle string is not added (to satisfy monoid identity laws).
func Intersperse(middle string) func(string, string) string { _ = "STUB: not implemented"; return nil }

// Prepend returns a function that prepends a prefix to a string.
// This is a curried function that takes a prefix and returns a function
// that prepends that prefix to any string passed to it.
//
// Example:
//
//	addHello := Prepend("Hello, ")
//	result := addHello("World") // "Hello, World"
func Prepend(prefix string) func(string) string { _ = "STUB: not implemented"; return nil }

// Append returns a function that appends a suffix to a string.
// This is a curried function that takes a suffix and returns a function
// that appends that suffix to any string passed to it.
//
// Example:
//
//	addExclamation := Append("!")
//	result := addExclamation("Hello") // "Hello!"
func Append(suffix string) func(string) string { _ = "STUB: not implemented"; return nil }
