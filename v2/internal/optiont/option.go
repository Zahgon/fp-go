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

package optiont

import (
	O "github.com/IBM/fp-go/v2/option"
)

func Of[A, HKTA any](fof func(O.Option[A]) HKTA, a A) HKTA {
	_ = "STUB: not implemented"
	return *new(HKTA)
}

func None[A, HKTA any](fof func(O.Option[A]) HKTA) HKTA {
	_ = "STUB: not implemented"
	return *new(HKTA)
}

func OfF[A, HKTA, HKTEA any](fmap func(HKTA, func(A) O.Option[A]) HKTEA, fa HKTA) HKTEA {
	_ = "STUB: not implemented"
	return *new(HKTEA)
}

func MonadMap[A, B, HKTFA, HKTFB any](fmap func(HKTFA, func(O.Option[A]) O.Option[B]) HKTFB, fa HKTFA, f func(A) B) HKTFB {
	_ = "STUB: not implemented"
	// HKTGA = Either[E, A]
	// HKTGB = Either[E, B]
	return *new(HKTFB)
}

func Map[A, B, HKTFA, HKTFB any](fmap func(func(O.Option[A]) O.Option[B]) func(HKTFA) HKTFB, f func(A) B) func(HKTFA) HKTFB {
	_ = "STUB: not implemented"
	// HKTGA = Either[E, A]
	// HKTGB = Either[E, B]
	return nil
}

func MonadChain[A, B, HKTFA, HKTFB any](
	fchain func(HKTFA, func(O.Option[A]) HKTFB) HKTFB,
	fof func(O.Option[B]) HKTFB,
	ma HKTFA,
	f func(A) HKTFB) HKTFB {
	_ = "STUB: not implemented"
	// dispatch to the even more generic implementation
	return *new(HKTFB)
}

func Chain[A, B, HKTFA, HKTFB any](
	fchain func(func(O.Option[A]) HKTFB) func(HKTFA) HKTFB,
	fof func(O.Option[B]) HKTFB,
	f func(A) HKTFB) func(ma HKTFA) HKTFB {
	_ = "STUB: not implemented"
	// dispatch to the even more generic implementation
	return nil
}

func MonadAp[A, B, HKTFAB, HKTFGAB, HKTFA, HKTFB any](
	fap func(HKTFGAB, HKTFA) HKTFB,
	fmap func(HKTFAB, func(O.Option[func(A) B]) func(O.Option[A]) O.Option[B]) HKTFGAB,
	fab HKTFAB,
	fa HKTFA) HKTFB {
	_ = "STUB: not implemented"
	return *new(HKTFB)
}

func Ap[A, B, HKTFAB, HKTFGAB, HKTFA, HKTFB any](
	fap func(HKTFA) func(HKTFGAB) HKTFB,
	fmap func(func(O.Option[func(A) B]) func(O.Option[A]) O.Option[B]) func(HKTFAB) HKTFGAB,
	fa HKTFA) func(HKTFAB) HKTFB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMatchE[A, HKTEA, HKTB any](
	fa HKTEA,
	mchain func(HKTEA, func(O.Option[A]) HKTB) HKTB,
	onNone func() HKTB,
	onSome func(A) HKTB) HKTB {
	_ = "STUB: not implemented"
	return *new(HKTB)
}

func MatchE[A, HKTEA, HKTB any](
	mchain func(func(O.Option[A]) HKTB) func(HKTEA) HKTB,
	onNone func() HKTB,
	onSome func(A) HKTB) func(HKTEA) HKTB {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func GetOrElse[A, HKTEA, HKTB any](
	mchain func(func(O.Option[A]) HKTB) func(HKTEA) HKTB,
	onNone func() HKTB,
	onSome func(A) HKTB) func(HKTEA) HKTB {
	_ = "STUB: not implemented"
	return nil
}

func FromOptionK[A, B, HKTB any](
	fof func(O.Option[B]) HKTB,
	f func(A) O.Option[B]) func(A) HKTB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainOptionK[A, B, HKTA, HKTB any](
	fchain func(HKTA, func(O.Option[A]) HKTB) HKTB,
	fof func(O.Option[B]) HKTB,
	ma HKTA,
	f func(A) O.Option[B],
) HKTB {
	_ = "STUB: not implemented"
	return *new(HKTB)
}

func ChainOptionK[A, B, HKTA, HKTB any](
	fchain func(func(O.Option[A]) HKTB) func(HKTA) HKTB,
	fof func(O.Option[B]) HKTB,
	f func(A) O.Option[B],
) func(HKTA) HKTB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAlt[LAZY ~func() HKTFA, A, HKTFA any](
	fof func(O.Option[A]) HKTFA,
	fchain func(HKTFA, func(O.Option[A]) HKTFA) HKTFA,

	first HKTFA,
	second LAZY) HKTFA {
	_ = "STUB: not implemented"
	return *new(HKTFA)
}

func Alt[LAZY ~func() HKTFA, A, HKTFA any](
	fof func(O.Option[A]) HKTFA,
	fchain func(func(O.Option[A]) HKTFA) func(HKTFA) HKTFA,

	second LAZY) func(HKTFA) HKTFA {
	_ = "STUB: not implemented"
	return nil
}

func SomeF[A, HKTA, HKTEA any](fmap func(HKTA, func(A) O.Option[A]) HKTEA, fa HKTA) HKTEA {
	_ = "STUB: not implemented"
	return *new(HKTEA)
}
