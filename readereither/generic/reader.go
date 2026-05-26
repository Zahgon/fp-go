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

package generic

import (
	ET "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
)

func MakeReaderEither[GEA ~func(E) ET.Either[L, A], L, E, A any](f func(E) ET.Either[L, A]) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromEither[GEA ~func(E) ET.Either[L, A], L, E, A any](e ET.Either[L, A]) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func RightReader[GA ~func(E) A, GEA ~func(E) ET.Either[L, A], L, E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func LeftReader[GL ~func(E) L, GEA ~func(E) ET.Either[L, A], L, E, A any](l GL) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Left[GEA ~func(E) ET.Either[L, A], L, E, A any](l L) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Right[GEA ~func(E) ET.Either[L, A], L, E, A any](r A) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func FromReader[GA ~func(E) A, GEA ~func(E) ET.Either[L, A], L, E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadMap[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], L, E, A, B any](fa GEA, f func(A) B) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Map[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], L, E, A, B any](f func(A) B) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], L, E, A, B any](ma GEA, f func(A) GEB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Chain[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], L, E, A, B any](f func(A) GEB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func Of[GEA ~func(E) ET.Either[L, A], L, E, A any](a A) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadAp[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], GEFAB ~func(E) ET.Either[L, func(A) B], L, E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Ap[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], GEFAB ~func(E) ET.Either[L, func(A) B], L, E, A, B any](fa GEA) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[GEA ~func(E) ET.Either[L, A], L, E, A any](pred func(A) bool, onFalse func(A) L) func(A) GEA {
	_ = "STUB: not implemented"
	return nil
}

func Fold[GEA ~func(E) ET.Either[L, A], GB ~func(E) B, E, L, A, B any](onLeft func(L) GB, onRight func(A) GB) func(GEA) GB {
	_ = "STUB: not implemented"
	return nil
}

func GetOrElse[GEA ~func(E) ET.Either[L, A], GA ~func(E) A, E, L, A any](onLeft func(L) GA) func(GEA) GA {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[GEA1 ~func(E) ET.Either[L1, A], GEA2 ~func(E) ET.Either[L2, A], E, L1, A, L2 any](onLeft func(L1) GEA2) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}

func OrLeft[GEA1 ~func(E) ET.Either[L1, A], GEA2 ~func(E) ET.Either[L2, A], GE2 ~func(E) L2, L1, E, L2, A any](onLeft func(L1) GE2) func(GEA1) GEA2 {
	_ = "STUB: not implemented"
	return nil
}

func Ask[GEE ~func(E) ET.Either[L, E], E, L any]() GEE { _ = "STUB: not implemented"; return *new(GEE) }

func Asks[GA ~func(E) A, GEA ~func(E) ET.Either[L, A], E, L, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadChainEitherK[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], L, E, A, B any](ma GEA, f func(A) ET.Either[L, B]) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainEitherK[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], L, E, A, B any](f func(A) ET.Either[L, B]) func(ma GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[GEA ~func(E) ET.Either[L, A], GEB ~func(E) ET.Either[L, B], L, E, A, B any](onNone func() L) func(func(A) O.Option[B]) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GEA ~func(E) ET.Either[L, A], GGA ~func(E) ET.Either[L, GEA], L, E, A any](mma GGA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadBiMap[GA ~func(E) ET.Either[E1, A], GB ~func(E) ET.Either[E2, B], E, E1, E2, A, B any](fa GA, f func(E1) E2, g func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// BiMap maps a pair of functions over the two type arguments of the bifunctor.
func BiMap[GA ~func(E) ET.Either[E1, A], GB ~func(E) ET.Either[E2, B], E, E1, E2, A, B any](f func(E1) E2, g func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Local changes the value of the local context during the execution of the action `ma` (similar to `Contravariant`'s
// `contramap`).
func Local[GA1 ~func(R1) ET.Either[E, A], GA2 ~func(R2) ET.Either[E, A], R2, R1, E, A any](f func(R2) R1) func(GA1) GA2 {
	_ = "STUB: not implemented"
	return nil
}

// Read applies a context to a reader to obtain its value
func Read[GA ~func(E) ET.Either[E1, A], E, E1, A any](e E) func(GA) ET.Either[E1, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[GEFAB ~func(E) ET.Either[L, func(A) B], GEB ~func(E) ET.Either[L, B], L, E, A, B any](fab GEFAB, a A) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Flap[GEFAB ~func(E) ET.Either[L, func(A) B], GEB ~func(E) ET.Either[L, B], L, E, A, B any](a A) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapLeft[GA1 ~func(C) ET.Either[E1, A], GA2 ~func(C) ET.Either[E2, A], C, E1, E2, A any](fa GA1, f func(E1) E2) GA2 {
	_ = "STUB: not implemented"
	return *new(GA2)
}

// MapLeft applies a mapping function to the error channel
func MapLeft[GA1 ~func(C) ET.Either[E1, A], GA2 ~func(C) ET.Either[E2, A], C, E1, E2, A any](f func(E1) E2) func(GA1) GA2 {
	_ = "STUB: not implemented"
	return nil
}
