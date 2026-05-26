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
	"time"

	ET "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
)

// type IOOption[A any] = func() Option[A]

func MakeIO[GA ~func() O.Option[A], A any](f GA) GA { _ = "STUB: not implemented"; return *new(GA) }

func Of[GA ~func() O.Option[A], A any](r A) GA { _ = "STUB: not implemented"; return *new(GA) }

func Some[GA ~func() O.Option[A], A any](r A) GA { _ = "STUB: not implemented"; return *new(GA) }

func None[GA ~func() O.Option[A], A any]() GA { _ = "STUB: not implemented"; return *new(GA) }

func MonadOf[GA ~func() O.Option[A], A any](r A) GA { _ = "STUB: not implemented"; return *new(GA) }

func FromIO[GA ~func() O.Option[A], GR ~func() A, A any](mr GR) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func FromOption[GA ~func() O.Option[A], A any](o O.Option[A]) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func FromEither[GA ~func() O.Option[A], E, A any](e ET.Either[E, A]) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func MonadMap[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](fa GA, f func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Map[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](f func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](fa GA, f func(A) GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// MonadChainFirst runs the monad returned by the function but returns the result of the original monad
func MonadChainFirst[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](ma GA, f func(A) GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirst runs the monad returned by the function but returns the result of the original monad
func ChainFirst[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](f func(A) GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirstIOK runs the monad returned by the function but returns the result of the original monad
func MonadChainFirstIOK[GA ~func() O.Option[A], GIOB ~func() B, A, B any](first GA, f func(A) GIOB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirstIOK runs the monad returned by the function but returns the result of the original monad
func ChainFirstIOK[GA ~func() O.Option[A], GIOB ~func() B, A, B any](f func(A) GIOB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

func Chain[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](f func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainOptionK[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](ma GA, f func(A) O.Option[B]) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ChainOptionK[GA ~func() O.Option[A], GB ~func() O.Option[B], A, B any](f func(A) O.Option[B]) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOK[GA ~func() O.Option[A], GB ~func() O.Option[B], GR ~func() B, A, B any](ma GA, f func(A) GR) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ChainIOK[GA ~func() O.Option[A], GB ~func() O.Option[B], GR ~func() B, A, B any](f func(A) GR) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[GB ~func() O.Option[B], GAB ~func() O.Option[func(A) B], GA ~func() O.Option[A], A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Ap[GB ~func() O.Option[B], GAB ~func() O.Option[func(A) B], GA ~func() O.Option[A], A, B any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GA ~func() O.Option[A], GAA ~func() O.Option[GA], A any](mma GAA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func Optionize0[GA ~func() O.Option[A], A any](f func() (A, bool)) func() GA {
	_ = "STUB: not implemented"
	return nil
}

func Optionize1[GA ~func() O.Option[A], T1, A any](f func(t1 T1) (A, bool)) func(T1) GA {
	_ = "STUB: not implemented"
	return nil
}

func Optionize2[GA ~func() O.Option[A], T1, T2, A any](f func(t1 T1, t2 T2) (A, bool)) func(T1, T2) GA {
	_ = "STUB: not implemented"
	return nil
}

func Optionize3[GA ~func() O.Option[A], T1, T2, T3, A any](f func(t1 T1, t2 T2, t3 T3) (A, bool)) func(T1, T2, T3) GA {
	_ = "STUB: not implemented"
	return nil
}

func Optionize4[GA ~func() O.Option[A], T1, T2, T3, T4, A any](f func(t1 T1, t2 T2, t3 T3, t4 T4) (A, bool)) func(T1, T2, T3, T4) GA {
	_ = "STUB: not implemented"
	return nil
}

// Memoize computes the value of the provided IO monad lazily but exactly once
func Memoize[GA ~func() O.Option[A], A any](ma GA) GA {
	_ = "STUB: not implemented"
	return *

	// Delay creates an operation that passes in the value after some delay
	new(GA)
}

func Delay[GA ~func() O.Option[A], A any](delay time.Duration) func(GA) GA {
	_ = "STUB: not implemented"
	return nil

	// After creates an operation that passes after the given [time.Time]
}

func After[GA ~func() O.Option[A], A any](timestamp time.Time) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// Fold convers an IOOption into an IO
func Fold[GA ~func() O.Option[A], GB ~func() B, A, B any](onNone func() GB, onSome func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[GA ~func() O.Option[A], A any](gen func() GA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func MonadAlt[LAZY ~func() GIOA, GIOA ~func() O.Option[A], A any](first GIOA, second LAZY) GIOA {
	_ = "STUB: not implemented"
	return *new(GIOA)
}

func Alt[LAZY ~func() GIOA, GIOA ~func() O.Option[A], A any](second LAZY) func(GIOA) GIOA {
	_ = "STUB: not implemented"
	return nil
}
