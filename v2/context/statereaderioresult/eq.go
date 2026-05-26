// Copyright (c) 2024 - 2025 IBM Corp.
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

package statereaderioresult

import (
	"context"

	"github.com/IBM/fp-go/v2/eq"
)

// Eq implements the equals predicate for values contained in the [StateReaderIOResult] monad
func Eq[S, A any](eqr eq.Eq[ReaderIOResult[Pair[S, A]]]) func(S) eq.Eq[StateReaderIOResult[S, A]] {
	_ = "STUB: not implemented"
	return nil
}

// FromStrictEquals constructs an [eq.Eq] from the canonical comparison function
func FromStrictEquals[S comparable, A comparable]() func(context.Context) func(S) eq.Eq[StateReaderIOResult[S, A]] {
	_ = "STUB: not implemented"
	return nil
}
