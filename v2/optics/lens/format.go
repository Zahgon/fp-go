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

package lens

import (
	"fmt"
	"log/slog"
)

// String returns the name of the lens for debugging and display purposes.
//
// Example:
//
//	nameLens := lens.MakeLensWithName(..., "Person.Name")
//	fmt.Println(nameLens)  // Prints: "Person.Name"
func (l Lens[S, T]) String() string {
	_ = "STUB: not implemented"

	// Format implements fmt.Formatter for Lens.
	// Supports all standard format verbs:
	//   - %s, %v, %+v: uses String() representation (lens name)
	//   - %#v: uses GoString() representation
	//   - %q: quoted String() representation
	//   - other verbs: uses String() representation
	//
	// Example:
	//
	//	nameLens := lens.MakeLensWithName(..., "Person.Name")
	//	fmt.Printf("%s", nameLens)   // "Person.Name"
	//	fmt.Printf("%v", nameLens)   // "Person.Name"
	//	fmt.Printf("%#v", nameLens)  // "lens.Lens[Person, string]{name: \"Person.Name\"}"
	//
	//go:noinline
	return ""
}

func (l Lens[S, T]) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

// GoString implements fmt.GoStringer for Lens.
// Returns a Go-syntax representation of the Lens value.
//
// Example:
//
//	nameLens := lens.MakeLensWithName(..., "Person.Name")
//	nameLens.GoString() // "lens.Lens[Person, string]{name: \"Person.Name\"}"
//
//go:noinline
func (l Lens[S, T]) GoString() string { _ = "STUB: not implemented"; return "" }

// LogValue implements slog.LogValuer for Lens.
// Returns a slog.Value that represents the Lens for structured logging.
// Logs the lens name as a string value.
//
// Example:
//
//	logger := slog.Default()
//	nameLens := lens.MakeLensWithName(..., "Person.Name")
//	logger.Info("using lens", "lens", nameLens)
//	// Logs: {"msg":"using lens","lens":"Person.Name"}
//
//go:noinline
func (l Lens[S, T]) LogValue() slog.Value { _ = "STUB: not implemented"; return *new(slog.Value) }
