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

package generic

import (
	P "github.com/IBM/fp-go/pair"
)

var (
	undefined any = struct{}{}
)

func Get[GA ~func(S) P.Pair[S, S], S any]() GA { _ = "STUB: not implemented"; return *new(GA) }

func Gets[GA ~func(S) P.Pair[A, S], FCT ~func(S) A, A, S any](f FCT) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func Put[GA ~func(S) P.Pair[any, S], S any]() GA { _ = "STUB: not implemented"; return *new(GA) }

func Modify[GA ~func(S) P.Pair[any, S], FCT ~func(S) S, S any](f FCT) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func Of[GA ~func(S) P.Pair[A, S], S, A any](a A) GA { _ = "STUB: not implemented"; return *new(GA) }

func MonadMap[GB ~func(S) P.Pair[B, S], GA ~func(S) P.Pair[A, S], FCT ~func(A) B, S, A, B any](fa GA, f FCT) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Map[GB ~func(S) P.Pair[B, S], GA ~func(S) P.Pair[A, S], FCT ~func(A) B, S, A, B any](f FCT) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GB ~func(S) P.Pair[B, S], GA ~func(S) P.Pair[A, S], FCT ~func(A) GB, S, A, B any](fa GA, f FCT) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Chain[GB ~func(S) P.Pair[B, S], GA ~func(S) P.Pair[A, S], FCT ~func(A) GB, S, A, B any](f FCT) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[GB ~func(S) P.Pair[B, S], GAB ~func(S) P.Pair[func(A) B, S], GA ~func(S) P.Pair[A, S], S, A, B any](fab GAB, fa GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Ap[GB ~func(S) P.Pair[B, S], GAB ~func(S) P.Pair[func(A) B, S], GA ~func(S) P.Pair[A, S], S, A, B any](ga GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[GB ~func(S) P.Pair[B, S], GA ~func(S) P.Pair[A, S], FCT ~func(A) GB, S, A, B any](ma GA, f FCT) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func ChainFirst[GB ~func(S) P.Pair[B, S], GA ~func(S) P.Pair[A, S], FCT ~func(A) GB, S, A, B any](f FCT) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GAA ~func(S) P.Pair[GA, S], GA ~func(S) P.Pair[A, S], S, A any](mma GAA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func Execute[GA ~func(S) P.Pair[A, S], S, A any](s S) func(GA) S {
	_ = "STUB: not implemented"
	return nil
}

func Evaluate[GA ~func(S) P.Pair[A, S], S, A any](s S) func(GA) A {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[FAB ~func(A) B, GFAB ~func(S) P.Pair[FAB, S], GB ~func(S) P.Pair[B, S], S, A, B any](fab GFAB, a A) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Flap[FAB ~func(A) B, GFAB ~func(S) P.Pair[FAB, S], GB ~func(S) P.Pair[B, S], S, A, B any](a A) func(GFAB) GB {
	_ = "STUB: not implemented"
	return nil
}
