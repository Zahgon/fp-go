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

package testing

import (
	"testing"

	E "github.com/IBM/fp-go/v2/eq"
	"github.com/IBM/fp-go/v2/internal/apply"
	"github.com/IBM/fp-go/v2/internal/functor"
	"github.com/IBM/fp-go/v2/internal/pointed"
)

// Apply associative composition law
//
// F.ap(F.ap(F.map(fbc, bc => ab => a => bc(ab(a))), fab), fa) <-> F.ap(fbc, F.ap(fab, fa))
//
// Deprecated: use [ApplyAssertAssociativeComposition] instead
func AssertAssociativeComposition[HKTA, HKTB, HKTC, HKTAB, HKTBC, HKTAC, HKTABAC, A, B, C any](t *testing.T,
	eq E.Eq[HKTC],

	fofab func(func(A) B) HKTAB,
	fofbc func(func(B) C) HKTBC,

	fmap func(HKTBC, func(func(B) C) func(func(A) B) func(A) C) HKTABAC,

	fapab func(HKTAB, HKTA) HKTB,
	fapbc func(HKTBC, HKTB) HKTC,
	fapac func(HKTAC, HKTA) HKTC,

	fapabac func(HKTABAC, HKTAB) HKTAC,

	ab func(A) B,
	bc func(B) C,
) func(fa HKTA) bool {
	_ = "STUB: not implemented"
	return nil
}

// Apply associative composition law
//
// F.ap(F.ap(F.map(fbc, bc => ab => a => bc(ab(a))), fab), fa) <-> F.ap(fbc, F.ap(fab, fa))
func ApplyAssertAssociativeComposition[HKTA, HKTB, HKTC, HKTAB, HKTBC, HKTAC, HKTABAC, A, B, C any](t *testing.T,
	eq E.Eq[HKTC],

	fofab pointed.Pointed[func(A) B, HKTAB],
	fofbc pointed.Pointed[func(B) C, HKTBC],

	fmap functor.Functor[func(B) C, func(func(A) B) func(A) C, HKTBC, HKTABAC],

	fapab apply.Apply[A, B, HKTA, HKTB, HKTAB],
	fapbc apply.Apply[B, C, HKTB, HKTC, HKTBC],
	fapac apply.Apply[A, C, HKTA, HKTC, HKTAC],

	fapabac apply.Apply[func(A) B, func(A) C, HKTAB, HKTAC, HKTABAC],

	ab func(A) B,
	bc func(B) C,
) func(fa HKTA) bool {
	_ = "STUB: not implemented"
	return nil
}

// AssertLaws asserts the apply laws `identity`, `composition` and `associative composition`
//
// Deprecated: use [ApplyAssertLaws] instead
func AssertLaws[HKTA, HKTB, HKTC, HKTAB, HKTBC, HKTAC, HKTABAC, A, B, C any](t *testing.T,
	eqa E.Eq[HKTA],
	eqc E.Eq[HKTC],

	fofab func(func(A) B) HKTAB,
	fofbc func(func(B) C) HKTBC,

	faa func(HKTA, func(A) A) HKTA,
	fab func(HKTA, func(A) B) HKTB,
	fac func(HKTA, func(A) C) HKTC,
	fbc func(HKTB, func(B) C) HKTC,

	fmap func(HKTBC, func(func(B) C) func(func(A) B) func(A) C) HKTABAC,

	fapab func(HKTAB, HKTA) HKTB,
	fapbc func(HKTBC, HKTB) HKTC,
	fapac func(HKTAC, HKTA) HKTC,

	fapabac func(HKTABAC, HKTAB) HKTAC,

	ab func(A) B,
	bc func(B) C,
) func(fa HKTA) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// functor laws

// associative composition laws

// ApplyAssertLaws asserts the apply laws `identity`, `composition` and `associative composition`
func ApplyAssertLaws[HKTA, HKTB, HKTC, HKTAB, HKTBC, HKTAC, HKTABAC, A, B, C any](t *testing.T,
	eqa E.Eq[HKTA],
	eqc E.Eq[HKTC],

	fofab pointed.Pointed[func(A) B, HKTAB],
	fofbc pointed.Pointed[func(B) C, HKTBC],

	faa functor.Functor[A, A, HKTA, HKTA],

	fmap functor.Functor[func(B) C, func(func(A) B) func(A) C, HKTBC, HKTABAC],

	fapab apply.Apply[A, B, HKTA, HKTB, HKTAB],
	fapbc apply.Apply[B, C, HKTB, HKTC, HKTBC],
	fapac apply.Apply[A, C, HKTA, HKTC, HKTAC],

	fapabac apply.Apply[func(A) B, func(A) C, HKTAB, HKTAC, HKTABAC],

	ab func(A) B,
	bc func(B) C,
) func(fa HKTA) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// functor laws

// associative composition laws
