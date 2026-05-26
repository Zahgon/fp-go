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

package testing

import (
	"testing"

	"github.com/IBM/fp-go/v2/eq"

	M "github.com/IBM/fp-go/v2/monoid"
)

// AssertLaws asserts the apply monad laws for the [P.Pair] monad
func assertLawsHead[E, A, B, C any](t *testing.T,
	m M.Monoid[E],

	eqe eq.Eq[E],
	eqa eq.Eq[A],
	eqb eq.Eq[B],
	eqc eq.Eq[C],

	ab func(A) B,
	bc func(B) C,
) func(a A) bool {
	_ = "STUB: not implemented"
	return nil
}

// AssertLaws asserts the apply monad laws for the [P.Pair] monad
func assertLawsTail[E, A, B, C any](t *testing.T,
	m M.Monoid[E],

	eqe eq.Eq[E],
	eqa eq.Eq[A],
	eqb eq.Eq[B],
	eqc eq.Eq[C],

	ab func(A) B,
	bc func(B) C,
) func(a A) bool {
	_ = "STUB: not implemented"
	return nil
}

// AssertLaws asserts the apply monad laws for the [P.Pair] monad
func AssertLaws[E, A, B, C any](t *testing.T,
	m M.Monoid[E],

	eqe eq.Eq[E],
	eqa eq.Eq[A],
	eqb eq.Eq[B],
	eqc eq.Eq[C],

	ab func(A) B,
	bc func(B) C,
) func(A) bool {
	_ = "STUB: not implemented"
	return nil
}
