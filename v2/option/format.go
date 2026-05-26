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
	"fmt"
	"log/slog"
)

const (
	noneGoTemplate = "option.None[%s]"
	someGoTemplate = "option.Some[%s](%#v)"

	noneFmtTemplate = "None[%s]"
	someFmtTemplate = "Some[%s](%v)"
)

// GoString implements fmt.GoStringer for Option.
// Returns a Go-syntax representation of the Option value.
//
// Example:
//
//	Some(42).GoString() // "option.Some[int](42)"
//	None[int]().GoString() // "option.None[int]()"
//
//go:noinline
func (s Option[A]) GoString() string { _ = "STUB: not implemented"; return "" }

// LogValue implements slog.LogValuer for Option.
// Returns a slog.Value that represents the Option for structured logging.
// Returns a group value with "some" key for Some values and "none" key for None values.
//
// Example:
//
//	logger := slog.Default()
//	result := Some(42)
//	logger.Info("result", "value", result)
//	// Logs: {"msg":"result","value":{"some":42}}
//
//	empty := None[int]()
//	logger.Info("empty", "value", empty)
//	// Logs: {"msg":"empty","value":{"none":{}}}
//
//go:noinline
func (s Option[A]) LogValue() slog.Value { _ = "STUB: not implemented"; return *new(slog.Value) }

// Format implements fmt.Formatter for Option.
// Supports all standard format verbs:
//   - %s, %v, %+v: uses String() representation
//   - %#v: uses GoString() representation
//   - %q: quoted String() representation
//   - other verbs: uses String() representation
//
// Example:
//
//	opt := Some(42)
//	fmt.Printf("%s", opt)   // "Some[int](42)"
//	fmt.Printf("%v", opt)   // "Some[int](42)"
//	fmt.Printf("%#v", opt)  // "option.Some[int](42)"
//
//go:noinline
func (s Option[A]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// optString prints some debug info for the object
//
//go:noinline
func optString(isSome bool, value any) string { _ = "STUB: not implemented"; return "" }

// For None, just show the type without ()
