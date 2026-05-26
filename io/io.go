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

package io

import (
	"time"

	G "github.com/IBM/fp-go/io/generic"
	T "github.com/IBM/fp-go/tuple"
)

// IO represents a synchronous computation that cannot fail
// refer to [https://andywhite.xyz/posts/2021-01-27-rte-foundations/#ioltagt] for more details
type IO[A any] func() A

func MakeIO[A any](f func() A) IO[A] { _ = "STUB: not implemented"; return nil }

func Of[A any](a A) IO[A] { _ = "STUB: not implemented"; return nil }

func FromIO[A any](a IO[A]) IO[A] {
	_ = "STUB: not implemented"

	// FromImpure converts a side effect without a return value into a side effect that returns any
	return nil
}

func FromImpure(f func()) IO[any] { _ = "STUB: not implemented"; return nil }

func MonadOf[A any](a A) IO[A] { _ = "STUB: not implemented"; return nil }

func MonadMap[A, B any](fa IO[A], f func(A) B) IO[B] { _ = "STUB: not implemented"; return nil }

func Map[A, B any](f func(A) B) func(fa IO[A]) IO[B] { _ = "STUB: not implemented"; return nil }

func MonadMapTo[A, B any](fa IO[A], b B) IO[B] { _ = "STUB: not implemented"; return nil }

func MapTo[A, B any](b B) func(IO[A]) IO[B] { _ = "STUB: not implemented"; return nil }

// MonadChain composes computations in sequence, using the return value of one computation to determine the next computation.
func MonadChain[A, B any](fa IO[A], f func(A) IO[B]) IO[B] { _ = "STUB: not implemented"; return nil }

// Chain composes computations in sequence, using the return value of one computation to determine the next computation.
func Chain[A, B any](f func(A) IO[B]) func(IO[A]) IO[B] { _ = "STUB: not implemented"; return nil }

func MonadAp[B, A any](mab IO[func(A) B], ma IO[A]) IO[B] { _ = "STUB: not implemented"; return nil }

func Ap[B, A any](ma IO[A]) func(IO[func(A) B]) IO[B] { _ = "STUB: not implemented"; return nil }

func Flatten[A any](mma IO[IO[A]]) IO[A] { _ = "STUB: not implemented"; return nil }

// Memoize computes the value of the provided [IO] monad lazily but exactly once
func Memoize[A any](ma IO[A]) IO[A] { _ = "STUB: not implemented"; return nil }

// MonadChainFirst composes computations in sequence, using the return value of one computation to determine the next computation and
// keeping only the result of the first.
func MonadChainFirst[A, B any](fa IO[A], f func(A) IO[B]) IO[A] {
	_ = "STUB: not implemented"
	return nil
}

// ChainFirst composes computations in sequence, using the return value of one computation to determine the next computation and
// keeping only the result of the first.
func ChainFirst[A, B any](f func(A) IO[B]) func(IO[A]) IO[A] { _ = "STUB: not implemented"; return nil }

// MonadApFirst combines two effectful actions, keeping only the result of the first.
func MonadApFirst[A, B any](first IO[A], second IO[B]) IO[A] { _ = "STUB: not implemented"; return nil }

// ApFirst combines two effectful actions, keeping only the result of the first.
func ApFirst[A, B any](second IO[B]) func(IO[A]) IO[A] { _ = "STUB: not implemented"; return nil }

// MonadApSecond combines two effectful actions, keeping only the result of the second.
func MonadApSecond[A, B any](first IO[A], second IO[B]) IO[B] {
	_ = "STUB: not implemented"
	return nil
}

// ApSecond combines two effectful actions, keeping only the result of the second.
func ApSecond[A, B any](second IO[B]) func(IO[A]) IO[B] { _ = "STUB: not implemented"; return nil }

// MonadChainTo composes computations in sequence, ignoring the return value of the first computation
func MonadChainTo[A, B any](fa IO[A], fb IO[B]) IO[B] { _ = "STUB: not implemented"; return nil }

// ChainTo composes computations in sequence, ignoring the return value of the first computation
func ChainTo[A, B any](fb IO[B]) func(IO[A]) IO[B] { _ = "STUB: not implemented"; return nil }

// Now returns the current timestamp
var Now = G.Now[IO[time.Time]]()

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[A any](gen func() IO[A]) IO[A] { _ = "STUB: not implemented"; return nil }

func MonadFlap[B, A any](fab IO[func(A) B], a A) IO[B] { _ = "STUB: not implemented"; return nil }

func Flap[B, A any](a A) func(IO[func(A) B]) IO[B] { _ = "STUB: not implemented"; return nil }

// Delay creates an operation that passes in the value after some delay
func Delay[A any](delay time.Duration) func(IO[A]) IO[A] { _ = "STUB: not implemented"; return nil }

// After creates an operation that passes after the given timestamp
func After[A any](timestamp time.Time) func(IO[A]) IO[A] { _ = "STUB: not implemented"; return nil }

// WithTime returns an operation that measures the start and end [time.Time] of the operation
func WithTime[A any](a IO[A]) IO[T.Tuple3[A, time.Time, time.Time]] {
	_ = "STUB: not implemented"
	return nil
}

// WithDuration returns an operation that measures the [time.Duration]
func WithDuration[A any](a IO[A]) IO[T.Tuple2[A, time.Duration]] {
	_ = "STUB: not implemented"
	return nil
}
