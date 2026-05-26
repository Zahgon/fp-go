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

package testing

import (
	"testing"

	E "github.com/IBM/fp-go/eq"
	"github.com/IBM/fp-go/internal/applicative"
	"github.com/IBM/fp-go/internal/apply"
	"github.com/IBM/fp-go/internal/functor"
	"github.com/IBM/fp-go/internal/pointed"
)

// Applicative identity law
//
// A.ap(A.of(a => a), fa) <-> fa
//
// Deprecated: use [ApplicativeAssertIdentity]
func AssertIdentity[HKTA, HKTAA, A any](t *testing.T,
	eq E.Eq[HKTA],

	fof func(func(A) A) HKTAA,

	fap func(HKTAA, HKTA) HKTA,
) func(fa HKTA) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// Applicative identity law
//
// A.ap(A.of(a => a), fa) <-> fa
func ApplicativeAssertIdentity[HKTA, HKTFAA, A any](t *testing.T,
	eq E.Eq[HKTA],

	ap applicative.Applicative[A, A, HKTA, HKTA, HKTFAA],
	paa pointed.Pointed[func(A) A, HKTFAA],

) func(fa HKTA) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// Applicative homomorphism law
//
// A.ap(A.of(ab), A.of(a)) <-> A.of(ab(a))
//
// Deprecated: use [ApplicativeAssertHomomorphism]
func AssertHomomorphism[HKTA, HKTB, HKTAB, A, B any](t *testing.T,
	eq E.Eq[HKTB],

	fofa func(A) HKTA,
	fofb func(B) HKTB,
	fofab func(func(A) B) HKTAB,

	fap func(HKTAB, HKTA) HKTB,

	ab func(A) B,
) func(a A) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// Applicative homomorphism law
//
// A.ap(A.of(ab), A.of(a)) <-> A.of(ab(a))
func ApplicativeAssertHomomorphism[HKTA, HKTB, HKTFAB, A, B any](t *testing.T,
	eq E.Eq[HKTB],

	apab applicative.Applicative[A, B, HKTA, HKTB, HKTFAB],
	pb pointed.Pointed[B, HKTB],
	pfab pointed.Pointed[func(A) B, HKTFAB],

	ab func(A) B,
) func(a A) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// Applicative interchange law
//
// A.ap(fab, A.of(a)) <-> A.ap(A.of(ab => ab(a)), fab)
//
// Deprecated: use [ApplicativeAssertInterchange]
func AssertInterchange[HKTA, HKTB, HKTAB, HKTABB, A, B any](t *testing.T,
	eq E.Eq[HKTB],

	fofa func(A) HKTA,
	fofab func(func(A) B) HKTAB,
	fofabb func(func(func(A) B) B) HKTABB,

	fapab func(HKTAB, HKTA) HKTB,
	fapabb func(HKTABB, HKTAB) HKTB,

	ab func(A) B,
) func(a A) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// Applicative interchange law
//
// A.ap(fab, A.of(a)) <-> A.ap(A.of(ab => ab(a)), fab)
func ApplicativeAssertInterchange[HKTA, HKTB, HKTFAB, HKTABB, A, B any](t *testing.T,
	eq E.Eq[HKTB],

	apab applicative.Applicative[A, B, HKTA, HKTB, HKTFAB],
	apabb applicative.Applicative[func(A) B, B, HKTFAB, HKTB, HKTABB],
	pabb pointed.Pointed[func(func(A) B) B, HKTABB],

	ab func(A) B,
) func(a A) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// AssertLaws asserts the apply laws `identity`, `composition`, `associative composition`, 'applicative identity', 'homomorphism', 'interchange'
//
// Deprecated: use [ApplicativeAssertLaws] instead
func AssertLaws[HKTA, HKTB, HKTC, HKTAA, HKTAB, HKTBC, HKTAC, HKTABB, HKTABAC, A, B, C any](t *testing.T,
	eqa E.Eq[HKTA],
	eqb E.Eq[HKTB],
	eqc E.Eq[HKTC],

	fofa func(A) HKTA,
	fofb func(B) HKTB,

	fofaa func(func(A) A) HKTAA,
	fofab func(func(A) B) HKTAB,
	fofbc func(func(B) C) HKTBC,
	fofabb func(func(func(A) B) B) HKTABB,

	faa func(HKTA, func(A) A) HKTA,
	fab func(HKTA, func(A) B) HKTB,
	fac func(HKTA, func(A) C) HKTC,
	fbc func(HKTB, func(B) C) HKTC,

	fmap func(HKTBC, func(func(B) C) func(func(A) B) func(A) C) HKTABAC,

	fapaa func(HKTAA, HKTA) HKTA,
	fapab func(HKTAB, HKTA) HKTB,
	fapbc func(HKTBC, HKTB) HKTC,
	fapac func(HKTAC, HKTA) HKTC,

	fapabb func(HKTABB, HKTAB) HKTB,
	fapabac func(HKTABAC, HKTAB) HKTAC,

	ab func(A) B,
	bc func(B) C,
) func(a A) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// apply laws

// applicative laws

// ApplicativeAssertLaws asserts the apply laws `identity`, `composition`, `associative composition`, 'applicative identity', 'homomorphism', 'interchange'
func ApplicativeAssertLaws[HKTA, HKTB, HKTC, HKTAA, HKTAB, HKTBC, HKTAC, HKTABB, HKTABAC, A, B, C any](t *testing.T,
	eqa E.Eq[HKTA],
	eqb E.Eq[HKTB],
	eqc E.Eq[HKTC],

	fofb pointed.Pointed[B, HKTB],

	fofaa pointed.Pointed[func(A) A, HKTAA],
	fofbc pointed.Pointed[func(B) C, HKTBC],

	fofabb pointed.Pointed[func(func(A) B) B, HKTABB],

	faa functor.Functor[A, A, HKTA, HKTA],

	fmap functor.Functor[func(B) C, func(func(A) B) func(A) C, HKTBC, HKTABAC],

	fapaa applicative.Applicative[A, A, HKTA, HKTA, HKTAA],
	fapab applicative.Applicative[A, B, HKTA, HKTB, HKTAB],
	fapbc apply.Apply[B, C, HKTB, HKTC, HKTBC],
	fapac apply.Apply[A, C, HKTA, HKTC, HKTAC],

	fapabb applicative.Applicative[func(A) B, B, HKTAB, HKTB, HKTABB],
	fapabac applicative.Applicative[func(A) B, func(A) C, HKTAB, HKTAC, HKTABAC],

	ab func(A) B,
	bc func(B) C,
) func(a A) bool {
	_ = "STUB: not implemented"
	// mark as test helper
	return nil
}

// apply laws

// applicative laws
