// Copyright (c) 2024 IBM Corp.
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

package state

import (
	P "github.com/IBM/fp-go/pair"
	R "github.com/IBM/fp-go/reader"
)

// State represents an operation on top of a current [State] that produces a value and a new [State]
type State[S, A any] R.Reader[S, P.Pair[A, S]]

func Get[S any]() State[S, S] { _ = "STUB: not implemented"; return nil }

func Gets[FCT ~func(S) A, A, S any](f FCT) State[S, A] { _ = "STUB: not implemented"; return nil }

func Put[S any]() State[S, any] { _ = "STUB: not implemented"; return nil }

func Modify[FCT ~func(S) S, S any](f FCT) State[S, any] { _ = "STUB: not implemented"; return nil }

func Of[S, A any](a A) State[S, A] { _ = "STUB: not implemented"; return nil }

func MonadMap[S any, FCT ~func(A) B, A, B any](fa State[S, A], f FCT) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[S any, FCT ~func(A) B, A, B any](f FCT) func(State[S, A]) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[S any, FCT ~func(A) State[S, B], A, B any](fa State[S, A], f FCT) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[S any, FCT ~func(A) State[S, B], A, B any](f FCT) func(State[S, A]) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[S, A, B any](fab State[S, func(A) B], fa State[S, A]) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[S, A, B any](ga State[S, A]) func(State[S, func(A) B]) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[S any, FCT ~func(A) State[S, B], A, B any](ma State[S, A], f FCT) State[S, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirst[S any, FCT ~func(A) State[S, B], A, B any](f FCT) func(State[S, A]) State[S, A] {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[S, A any](mma State[S, State[S, A]]) State[S, A] {
	_ = "STUB: not implemented"
	return nil
}

func Execute[A, S any](s S) func(State[S, A]) S { _ = "STUB: not implemented"; return nil }

func Evaluate[A, S any](s S) func(State[S, A]) A { _ = "STUB: not implemented"; return nil }

func MonadFlap[FAB ~func(A) B, S, A, B any](fab State[S, FAB], a A) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[S, A, B any](a A) func(State[S, func(A) B]) State[S, B] {
	_ = "STUB: not implemented"
	return nil
}
