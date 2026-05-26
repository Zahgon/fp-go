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

package readerio

import (
	IO "github.com/IBM/fp-go/io"
	R "github.com/IBM/fp-go/reader"
)

type ReaderIO[E, A any] R.Reader[E, IO.IO[A]]

// FromIO converts an [IO.IO] to a [ReaderIO]
func FromIO[E, A any](t IO.IO[A]) ReaderIO[E, A] { _ = "STUB: not implemented"; return nil }

func MonadMap[E, A, B any](fa ReaderIO[E, A], f func(A) B) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[E, A, B any](f func(A) B) func(ReaderIO[E, A]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[E, A, B any](ma ReaderIO[E, A], f func(A) ReaderIO[E, B]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[E, A, B any](f func(A) ReaderIO[E, B]) func(ReaderIO[E, A]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Of[E, A any](a A) ReaderIO[E, A] { _ = "STUB: not implemented"; return nil }

func MonadAp[B, E, A any](fab ReaderIO[E, func(A) B], fa ReaderIO[E, A]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, E, A any](fa ReaderIO[E, A]) func(ReaderIO[E, func(A) B]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ask[E any]() ReaderIO[E, E] { _ = "STUB: not implemented"; return nil }

func Asks[E, A any](r R.Reader[E, A]) ReaderIO[E, A] { _ = "STUB: not implemented"; return nil }

func MonadChainIOK[E, A, B any](ma ReaderIO[E, A], f func(A) IO.IO[B]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOK[E, A, B any](f func(A) IO.IO[B]) func(ReaderIO[E, A]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[E, A any](gen func() ReaderIO[E, A]) ReaderIO[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Memoize computes the value of the provided [ReaderIO] monad lazily but exactly once
// The context used to compute the value is the context of the first call, so do not use this
// method if the value has a functional dependency on the content of the context
func Memoize[E, A any](rdr ReaderIO[E, A]) ReaderIO[E, A] { _ = "STUB: not implemented"; return nil }

func Flatten[E, A any](mma ReaderIO[E, ReaderIO[E, A]]) ReaderIO[E, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[E, A, B any](fab ReaderIO[E, func(A) B], a A) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[E, A, B any](a A) func(ReaderIO[E, func(A) B]) ReaderIO[E, B] {
	_ = "STUB: not implemented"
	return nil
}
