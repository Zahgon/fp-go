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

package iooption

import (
	"time"

	"github.com/IBM/fp-go/v2/io"
)

func Of[A any](r A) IOOption[A] { _ = "STUB: not implemented"; return nil }

func Some[A any](r A) IOOption[A] { _ = "STUB: not implemented"; return nil }

func None[A any]() IOOption[A] { _ = "STUB: not implemented"; return nil }

func MonadOf[A any](r A) IOOption[A] { _ = "STUB: not implemented"; return nil }

func FromOption[A any](o Option[A]) IOOption[A] { _ = "STUB: not implemented"; return nil }

func ChainOptionK[A, B any](f func(A) Option[B]) Operator[A, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOK[A, B any](ma IOOption[A], f io.Kleisli[A, B]) IOOption[B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOK[A, B any](f io.Kleisli[A, B]) Operator[A, B] { _ = "STUB: not implemented"; return nil }

func FromIO[A any](mr IO[A]) IOOption[A] { _ = "STUB: not implemented"; return nil }

func MonadMap[A, B any](fa IOOption[A], f func(A) B) IOOption[B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[A, B any](f func(A) B) Operator[A, B] { _ = "STUB: not implemented"; return nil }

func MonadChain[A, B any](fa IOOption[A], f Kleisli[A, B]) IOOption[B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A, B any](f Kleisli[A, B]) Operator[A, B] { _ = "STUB: not implemented"; return nil }

func MonadAp[B, A any](mab IOOption[func(A) B], ma IOOption[A]) IOOption[B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, A any](ma IOOption[A]) Operator[func(A) B, B] { _ = "STUB: not implemented"; return nil }

func ApSeq[B, A any](ma IOOption[A]) Operator[func(A) B, B] { _ = "STUB: not implemented"; return nil }

func ApPar[B, A any](ma IOOption[A]) Operator[func(A) B, B] { _ = "STUB: not implemented"; return nil }

func Flatten[A any](mma IOOption[IOOption[A]]) IOOption[A] { _ = "STUB: not implemented"; return nil }

func Optionize0[A any](f func() (A, bool)) Lazy[IOOption[A]] { _ = "STUB: not implemented"; return nil }

func Optionize1[T1, A any](f func(t1 T1) (A, bool)) Kleisli[T1, A] {
	_ = "STUB: not implemented"
	return nil
}

func Optionize2[T1, T2, A any](f func(t1 T1, t2 T2) (A, bool)) func(T1, T2) IOOption[A] {
	_ = "STUB: not implemented"
	return nil
}

func Optionize3[T1, T2, T3, A any](f func(t1 T1, t2 T2, t3 T3) (A, bool)) func(T1, T2, T3) IOOption[A] {
	_ = "STUB: not implemented"
	return nil
}

func Optionize4[T1, T2, T3, T4, A any](f func(t1 T1, t2 T2, t3 T3, t4 T4) (A, bool)) func(T1, T2, T3, T4) IOOption[A] {
	_ = "STUB: not implemented"
	return nil
}

func Memoize[A any](ma IOOption[A]) IOOption[A] { _ = "STUB: not implemented"; return nil }

// Fold convers an [IOOption] into an [IO]
func Fold[A, B any](onNone IO[B], onSome io.Kleisli[A, B]) func(IOOption[A]) IO[B] {
	_ = "STUB: not implemented"
	return nil
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[A any](gen func() IOOption[A]) IOOption[A] { _ = "STUB: not implemented"; return nil }

// FromEither converts an [Either] into an [IOOption]
func FromEither[E, A any](e Either[E, A]) IOOption[A] { _ = "STUB: not implemented"; return nil }

// MonadAlt identifies an associative operation on a type constructor
func MonadAlt[A any](first, second IOOption[A]) IOOption[A] { _ = "STUB: not implemented"; return nil }

// Alt identifies an associative operation on a type constructor
func Alt[A any](second IOOption[A]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

// MonadChainFirst runs the monad returned by the function but returns the result of the original monad
func MonadChainFirst[A, B any](ma IOOption[A], f Kleisli[A, B]) IOOption[A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirst runs the monad returned by the function but returns the result of the original monad
func ChainFirst[A, B any](f Kleisli[A, B]) Operator[A, A] { _ = "STUB: not implemented"; return nil }

// MonadChainFirstIOK runs the monad returned by the function but returns the result of the original monad
func MonadChainFirstIOK[A, B any](first IOOption[A], f io.Kleisli[A, B]) IOOption[A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirstIOK runs the monad returned by the function but returns the result of the original monad
func ChainFirstIOK[A, B any](f io.Kleisli[A, B]) Operator[A, A] {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some delay
func Delay[A any](delay time.Duration) Operator[A, A] { _ = "STUB: not implemented"; return nil }

// After creates an operation that passes after the given [time.Time]
func After[A any](timestamp time.Time) Operator[A, A] { _ = "STUB: not implemented"; return nil }
