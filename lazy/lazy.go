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

package lazy

import (
	"time"

	G "github.com/IBM/fp-go/io/generic"
)

// Lazy represents a synchronous computation without side effects
type Lazy[A any] func() A

func MakeLazy[A any](f func() A) Lazy[A] { _ = "STUB: not implemented"; return nil }

func Of[A any](a A) Lazy[A] { _ = "STUB: not implemented"; return nil }

func FromLazy[A any](a Lazy[A]) Lazy[A] {
	_ = "STUB: not implemented"

	// FromImpure converts a side effect without a return value into a side effect that returns any
	return nil
}

func FromImpure(f func()) Lazy[any] { _ = "STUB: not implemented"; return nil }

func MonadOf[A any](a A) Lazy[A] { _ = "STUB: not implemented"; return nil }

func MonadMap[A, B any](fa Lazy[A], f func(A) B) Lazy[B] { _ = "STUB: not implemented"; return nil }

func Map[A, B any](f func(A) B) func(fa Lazy[A]) Lazy[B] { _ = "STUB: not implemented"; return nil }

func MonadMapTo[A, B any](fa Lazy[A], b B) Lazy[B] { _ = "STUB: not implemented"; return nil }

func MapTo[A, B any](b B) func(Lazy[A]) Lazy[B] { _ = "STUB: not implemented"; return nil }

// MonadChain composes computations in sequence, using the return value of one computation to determine the next computation.
func MonadChain[A, B any](fa Lazy[A], f func(A) Lazy[B]) Lazy[B] {
	_ = "STUB: not implemented"
	return nil

	// Chain composes computations in sequence, using the return value of one computation to determine the next computation.
}

func Chain[A, B any](f func(A) Lazy[B]) func(Lazy[A]) Lazy[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[B, A any](mab Lazy[func(A) B], ma Lazy[A]) Lazy[B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, A any](ma Lazy[A]) func(Lazy[func(A) B]) Lazy[B] { _ = "STUB: not implemented"; return nil }

func Flatten[A any](mma Lazy[Lazy[A]]) Lazy[A] { _ = "STUB: not implemented"; return nil }

// Memoize computes the value of the provided [Lazy] monad lazily but exactly once
func Memoize[A any](ma Lazy[A]) Lazy[A] { _ = "STUB: not implemented"; return nil }

// MonadChainFirst composes computations in sequence, using the return value of one computation to determine the next computation and
// keeping only the result of the first.
func MonadChainFirst[A, B any](fa Lazy[A], f func(A) Lazy[B]) Lazy[A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirst composes computations in sequence, using the return value of one computation to determine the next computation and
// keeping only the result of the first.
func ChainFirst[A, B any](f func(A) Lazy[B]) func(Lazy[A]) Lazy[A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApFirst combines two effectful actions, keeping only the result of the first.
func MonadApFirst[A, B any](first Lazy[A], second Lazy[B]) Lazy[A] {
	_ = "STUB: not implemented"
	return nil
}

// ApFirst combines two effectful actions, keeping only the result of the first.
func ApFirst[A, B any](second Lazy[B]) func(Lazy[A]) Lazy[A] { _ = "STUB: not implemented"; return nil }

// MonadApSecond combines two effectful actions, keeping only the result of the second.
func MonadApSecond[A, B any](first Lazy[A], second Lazy[B]) Lazy[B] {
	_ = "STUB: not implemented"
	return nil
}

// ApSecond combines two effectful actions, keeping only the result of the second.
func ApSecond[A, B any](second Lazy[B]) func(Lazy[A]) Lazy[B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainTo composes computations in sequence, ignoring the return value of the first computation
func MonadChainTo[A, B any](fa Lazy[A], fb Lazy[B]) Lazy[B] { _ = "STUB: not implemented"; return nil }

// ChainTo composes computations in sequence, ignoring the return value of the first computation
func ChainTo[A, B any](fb Lazy[B]) func(Lazy[A]) Lazy[B] { _ = "STUB: not implemented"; return nil }

// Now returns the current timestamp
var Now = G.Now[Lazy[time.Time]]()

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[A any](gen func() Lazy[A]) Lazy[A] { _ = "STUB: not implemented"; return nil }
