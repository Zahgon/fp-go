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

package iso

import (
	"fmt"
	"log/slog"
)

// String returns a string representation of the isomorphism.
//
// Example:
//
//	tempIso := iso.MakeIso(...)
//	fmt.Println(tempIso)  // Prints: "Iso"
func (i Iso[S, T]) String() string {
	_ = "STUB: not implemented"

	// Format implements fmt.Formatter for Iso.
	// Supports all standard format verbs:
	//   - %s, %v, %+v: uses String() representation
	//   - %#v: uses GoString() representation
	//   - %q: quoted String() representation
	//   - other verbs: uses String() representation
	//
	// Example:
	//
	//	tempIso := iso.MakeIso(...)
	//	fmt.Printf("%s", tempIso)   // "Iso"
	//	fmt.Printf("%v", tempIso)   // "Iso"
	//	fmt.Printf("%#v", tempIso)  // "iso.Iso[Celsius, Fahrenheit]"
	//
	//go:noinline
	return ""
}

func (i Iso[S, T]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// GoString implements fmt.GoStringer for Iso.
// Returns a Go-syntax representation of the Iso value.
//
// Example:
//
//	tempIso := iso.MakeIso(...)
//	tempIso.GoString() // "iso.Iso[Celsius, Fahrenheit]"
//
//go:noinline
func (i Iso[S, T]) GoString() string { _ = "STUB: not implemented"; return "" }

// LogValue implements slog.LogValuer for Iso.
// Returns a slog.Value that represents the Iso for structured logging.
// Logs the type information as a string value.
//
// Example:
//
//	logger := slog.Default()
//	tempIso := iso.MakeIso(...)
//	logger.Info("using iso", "iso", tempIso)
//	// Logs: {"msg":"using iso","iso":"Iso"}
//
//go:noinline
func (i Iso[S, T]) LogValue() slog.Value { _ = "STUB: not implemented"; return *new(slog.Value) }
