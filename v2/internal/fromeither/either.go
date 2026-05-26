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

package fromeither

import (
	ET "github.com/IBM/fp-go/v2/either"
	O "github.com/IBM/fp-go/v2/option"
)

//go:inline
func FromOption[A, HKTEA, E any](fromEither func(ET.Either[E, A]) HKTEA, onNone func() E) func(ma O.Option[A]) HKTEA {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func FromPredicate[E, A, HKTEA any](fromEither func(ET.Either[E, A]) HKTEA, pred func(A) bool, onFalse func(A) E) func(A) HKTEA {
	_ = "STUB: not implemented"
	return nil
}

func MonadFromOption[E, A, HKTEA any](
	fromEither func(ET.Either[E, A]) HKTEA,
	onNone func() E,
	ma O.Option[A],
) HKTEA {
	_ = "STUB: not implemented"
	return *new(HKTEA)
}

//go:inline
func FromOptionK[A, E, B, HKTEB any](
	fromEither func(ET.Either[E, B]) HKTEB,
	onNone func() E) func(f func(A) O.Option[B]) func(A) HKTEB {
	_ = "STUB: not implemented"
	// helper
	return nil
}

//go:inline
func MonadChainEitherK[A, E, B, HKTEA, HKTEB any](
	mchain func(HKTEA, func(A) HKTEB) HKTEB,
	fromEither func(ET.Either[E, B]) HKTEB,
	ma HKTEA,
	f func(A) ET.Either[E, B]) HKTEB {
	_ = "STUB: not implemented"
	return *new(HKTEB)
}

//go:inline
func ChainEitherK[A, E, B, HKTEA, HKTEB any](
	mchain func(func(A) HKTEB) func(HKTEA) HKTEB,
	fromEither func(ET.Either[E, B]) HKTEB,
	f func(A) ET.Either[E, B]) func(HKTEA) HKTEB {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func ChainOptionK[A, E, B, HKTEA, HKTEB any](
	mchain func(HKTEA, func(A) HKTEB) HKTEB,
	fromEither func(ET.Either[E, B]) HKTEB,
	onNone func() E,
) func(f func(A) O.Option[B]) func(ma HKTEA) HKTEB {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadChainFirstEitherK[A, E, B, HKTEA, HKTEB any](
	mchain func(HKTEA, func(A) HKTEA) HKTEA,
	mmap func(HKTEB, func(B) A) HKTEA,
	fromEither func(ET.Either[E, B]) HKTEB,
	ma HKTEA,
	f func(A) ET.Either[E, B]) HKTEA {
	_ = "STUB: not implemented"
	return *new(HKTEA)
}

//go:inline
func ChainFirstEitherK[A, E, B, HKTEA, HKTEB any](
	mchain func(func(A) HKTEA) func(HKTEA) HKTEA,
	mmap func(func(B) A) func(HKTEB) HKTEA,
	fromEither func(ET.Either[E, B]) HKTEB,
	f func(A) ET.Either[E, B]) func(HKTEA) HKTEA {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func BindEitherK[
	E, S1, S2, T,
	HKTET,
	HKTES1,
	HKTES2 any](
	mchain func(func(S1) HKTES2) func(HKTES1) HKTES2,
	mmap func(func(T) S2) func(HKTET) HKTES2,
	fromEither func(ET.Either[E, T]) HKTET,
	setter func(T) func(S1) S2,
	f func(S1) ET.Either[E, T],
) func(HKTES1) HKTES2 {
	_ = "STUB: not implemented"
	return nil
}
