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

package generic

import (
	O "github.com/IBM/fp-go/v2/option"
)

//go:inline
func MakeReaderOption[GEA ~func(E) O.Option[A], E, A any](f func(E) O.Option[A]) GEA {
	_ = "STUB: not implemented"

	//go:inline
	return *new(GEA)
}

func FromOption[GEA ~func(E) O.Option[A], E, A any](e O.Option[A]) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func SomeReader[GA ~func(E) A, GEA ~func(E) O.Option[A], E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Some[GEA ~func(E) O.Option[A], E, A any](r A) GEA { _ = "STUB: not implemented"; return *new(GEA) }

//go:inline
func FromReader[GA ~func(E) A, GEA ~func(E) O.Option[A], E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadMap[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], E, A, B any](fa GEA, f func(A) B) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Map[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], E, A, B any](f func(A) B) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], E, A, B any](ma GEA, f func(A) GEB) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Chain[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], E, A, B any](f func(A) GEB) func(GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func Of[GEA ~func(E) O.Option[A], E, A any](a A) GEA { _ = "STUB: not implemented"; return *new(GEA) }

func MonadAp[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], GEFAB ~func(E) O.Option[func(A) B], E, A, B any](fab GEFAB, fa GEA) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Ap[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], GEFAB ~func(E) O.Option[func(A) B], E, A, B any](fa GEA) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[GEA ~func(E) O.Option[A], E, A any](pred func(A) bool) func(A) GEA {
	_ = "STUB: not implemented"
	return nil
}

func Fold[GEA ~func(E) O.Option[A], GB ~func(E) B, E, A, B any](onNone func() GB, onRight func(A) GB) func(GEA) GB {
	_ = "STUB: not implemented"
	return nil
}

func GetOrElse[GEA ~func(E) O.Option[A], GA ~func(E) A, E, A any](onNone func() GA) func(GEA) GA {
	_ = "STUB: not implemented"
	return nil
}

func Ask[GEE ~func(E) O.Option[E], E any]() GEE { _ = "STUB: not implemented"; return *new(GEE) }

func Asks[GA ~func(E) A, GEA ~func(E) O.Option[A], E, A any](r GA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func MonadChainOptionK[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], E, A, B any](ma GEA, f func(A) O.Option[B]) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func ChainOptionK[GEA ~func(E) O.Option[A], GEB ~func(E) O.Option[B], E, A, B any](f func(A) O.Option[B]) func(ma GEA) GEB {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GEA ~func(E) O.Option[A], GGA ~func(E) O.Option[GEA], E, A any](mma GGA) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

// Local changes the value of the local context during the execution of the action `ma` (similar to `Contravariant`'s
// `contramap`).
func Local[GA1 ~func(R1) O.Option[A], GA2 ~func(R2) O.Option[A], R2, R1, E, A any](f func(R2) R1) func(GA1) GA2 {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[GEFAB ~func(E) O.Option[func(A) B], GEB ~func(E) O.Option[B], E, A, B any](fab GEFAB, a A) GEB {
	_ = "STUB: not implemented"
	return *new(GEB)
}

func Flap[GEFAB ~func(E) O.Option[func(A) B], GEB ~func(E) O.Option[B], E, A, B any](a A) func(GEFAB) GEB {
	_ = "STUB: not implemented"
	return nil
}
