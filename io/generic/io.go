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

	T "github.com/IBM/fp-go/tuple"
)

var (
	// undefined represents an undefined value
	undefined = struct{}{}
)

// type IO[A any] = func() A

func MakeIO[GA ~func() A, A any](f func() A) GA { _ = "STUB: not implemented"; return *new(GA) }

func Of[GA ~func() A, A any](a A) GA { _ = "STUB: not implemented"; return *new(GA) }

func FromIO[GA ~func() A, A any](a GA) GA {
	_ = "STUB: not implemented"

	// FromImpure converts a side effect without a return value into a side effect that returns any
	return *new(GA)
}

func FromImpure[GA ~func() any, IMP ~func()](f IMP) GA { _ = "STUB: not implemented"; return *new(GA) }

func MonadOf[GA ~func() A, A any](a A) GA { _ = "STUB: not implemented"; return *new(GA) }

func MonadMap[GA ~func() A, GB ~func() B, A, B any](fa GA, f func(A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Map[GA ~func() A, GB ~func() B, A, B any](f func(A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[GA ~func() A, GB ~func() B, A, B any](fa GA, b B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func MapTo[GA ~func() A, GB ~func() B, A, B any](b B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// MonadChain composes computations in sequence, using the return value of one computation to determine the next computation.
func MonadChain[GA ~func() A, GB ~func() B, A, B any](fa GA, f func(A) GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// Chain composes computations in sequence, using the return value of one computation to determine the next computation.
func Chain[GA ~func() A, GB ~func() B, A, B any](f func(A) GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainTo composes computations in sequence, ignoring the return value of the first computation
func MonadChainTo[GA ~func() A, GB ~func() B, A, B any](fa GA, fb GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// ChainTo composes computations in sequence, ignoring the return value of the first computation
func ChainTo[GA ~func() A, GB ~func() B, A, B any](fb GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainFirst composes computations in sequence, using the return value of one computation to determine the next computation and
// keeping only the result of the first.
func MonadChainFirst[GA ~func() A, GB ~func() B, A, B any](fa GA, f func(A) GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ChainFirst composes computations in sequence, using the return value of one computation to determine the next computation and
// keeping only the result of the first.
func ChainFirst[GA ~func() A, GB ~func() B, A, B any](f func(A) GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

func ApSeq[GB ~func() B, GAB ~func() func(A) B, GA ~func() A, B, A any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func ApPar[GB ~func() B, GAB ~func() func(A) B, GA ~func() A, B, A any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func Ap[GB ~func() B, GAB ~func() func(A) B, GA ~func() A, B, A any](ma GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GA ~func() A, GAA ~func() GA, A any](mma GAA) GA {
	_ = "STUB: not implemented"

	// Memoize computes the value of the provided IO monad lazily but exactly once
	return *new(GA)
}

func Memoize[GA ~func() A, A any](ma GA) GA { _ = "STUB: not implemented"; return *new(GA) }

// Delay creates an operation that passes in the value after some delay
func Delay[GA ~func() A, A any](delay time.Duration) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

func after(timestamp time.Time) func() {
	_ = "STUB: not implemented"

	// check if we need to wait
	return nil
}

// After creates an operation that passes after the given timestamp
func After[GA ~func() A, A any](timestamp time.Time) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// wait as long as necessary

// execute after wait

// Now returns the current timestamp
func Now[GA ~func() time.Time]() GA { _ = "STUB: not implemented"; return *new(GA) }

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[GA ~func() A, A any](gen func() GA) GA { _ = "STUB: not implemented"; return *new(GA) }

func MonadFlap[FAB ~func(A) B, GFAB ~func() FAB, GB ~func() B, A, B any](fab GFAB, a A) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Flap[FAB ~func(A) B, GFAB ~func() FAB, GB ~func() B, A, B any](a A) func(GFAB) GB {
	_ = "STUB: not implemented"
	return nil
}

// WithTime returns an operation that measures the start and end timestamp of the operation
func WithTime[GTA ~func() T.Tuple3[A, time.Time, time.Time], GA ~func() A, A any](a GA) GTA {
	_ = "STUB: not implemented"
	return *new(GTA)
}

// WithDuration returns an operation that measures the duration of the operation
func WithDuration[GTA ~func() T.Tuple2[A, time.Duration], GA ~func() A, A any](a GA) GTA {
	_ = "STUB: not implemented"
	return *new(GTA)
}
