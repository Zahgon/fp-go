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

package either

import (
	"fmt"
	"log/slog"
)

const (
	leftGoTemplate  = "either.Left[%s](%#v)"
	rightGoTemplate = "either.Right[%s](%#v)"

	leftFmtTemplate  = "Left[%T](%v)"
	rightFmtTemplate = "Right[%T](%v)"
)

func goString(template string, other, v any) string { _ = "STUB: not implemented"; return "" }

// String prints some debug info for the object
//
//go:noinline
func (s Either[E, A]) String() string { _ = "STUB: not implemented"; return "" }

// Format implements fmt.Formatter for Either.
// Supports all standard format verbs:
//   - %s, %v, %+v: uses String() representation
//   - %#v: uses GoString() representation
//   - %q: quoted String() representation
//   - other verbs: uses String() representation
//
// Example:
//
//	e := either.Right[error](42)
//	fmt.Printf("%s", e)   // "Right[int](42)"
//	fmt.Printf("%v", e)   // "Right[int](42)"
//	fmt.Printf("%#v", e)  // "either.Right[error](42)"
//
//go:noinline
func (s Either[E, A]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// GoString implements fmt.GoStringer for Either.
// Returns a Go-syntax representation of the Either value.
//
// Example:
//
//	either.Right[error](42).GoString() // "either.Right[error](42)"
//	either.Left[int](errors.New("fail")).GoString() // "either.Left[int](error)"
//
//go:noinline
func (s Either[E, A]) GoString() string { _ = "STUB: not implemented"; return "" }

// LogValue implements slog.LogValuer for Either.
// Returns a slog.Value that represents the Either for structured logging.
// Returns a group value with "right" key for Right values and "left" key for Left values.
//
// Example:
//
//	logger := slog.Default()
//	result := either.Right[error](42)
//	logger.Info("result", "value", result)
//	// Logs: {"msg":"result","value":{"right":42}}
//
//	err := either.Left[int](errors.New("failed"))
//	logger.Error("error", "value", err)
//	// Logs: {"msg":"error","value":{"left":"failed"}}
//
//go:noinline
func (s Either[E, A]) LogValue() slog.Value { _ = "STUB: not implemented"; return *new(slog.Value) }
