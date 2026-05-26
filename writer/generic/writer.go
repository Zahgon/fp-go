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
	M "github.com/IBM/fp-go/monoid"
	P "github.com/IBM/fp-go/pair"
	SG "github.com/IBM/fp-go/semigroup"
)

var (
	undefined any = struct{}{}
)

func Tell[GA ~func() P.Pair[any, W], W any](w W) GA { _ = "STUB: not implemented"; return *new(GA) }

func Of[GA ~func() P.Pair[A, W], W, A any](m M.Monoid[W], a A) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Listen modifies the result to include the changes to the accumulator
func Listen[GA ~func() P.Pair[A, W], GTA ~func() P.Pair[P.Pair[A, W], W], W, A any](fa GA) GTA {
	_ = "STUB: not implemented"
	return *new(GTA)
}

// Pass applies the returned function to the accumulator
func Pass[GFA ~func() P.Pair[P.Pair[A, FCT], W], GA ~func() P.Pair[A, W], FCT ~func(W) W, W, A any](fa GFA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func MonadMap[GB ~func() P.Pair[B, W], GA ~func() P.Pair[A, W], FCT ~func(A) B, W, A, B any](fa GA, f FCT) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Map[GB ~func() P.Pair[B, W], GA ~func() P.Pair[A, W], FCT ~func(A) B, W, A, B any](f FCT) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[GB ~func() P.Pair[B, W], GA ~func() P.Pair[A, W], FCT ~func(A) GB, W, A, B any](s SG.Semigroup[W], fa GA, f FCT) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Chain[GB ~func() P.Pair[B, W], GA ~func() P.Pair[A, W], FCT ~func(A) GB, W, A, B any](s SG.Semigroup[W], f FCT) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[GB ~func() P.Pair[B, W], GAB ~func() P.Pair[func(A) B, W], GA ~func() P.Pair[A, W], W, A, B any](s SG.Semigroup[W], fab GAB, fa GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Ap[GB ~func() P.Pair[B, W], GAB ~func() P.Pair[func(A) B, W], GA ~func() P.Pair[A, W], W, A, B any](s SG.Semigroup[W], ga GA) func(GAB) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[GB ~func() P.Pair[B, W], GA ~func() P.Pair[A, W], FCT ~func(A) GB, W, A, B any](s SG.Semigroup[W], ma GA, f FCT) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func ChainFirst[GB ~func() P.Pair[B, W], GA ~func() P.Pair[A, W], FCT ~func(A) GB, W, A, B any](s SG.Semigroup[W], f FCT) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[GAA ~func() P.Pair[GA, W], GA ~func() P.Pair[A, W], W, A any](s SG.Semigroup[W], mma GAA) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

func Execute[GA ~func() P.Pair[A, W], W, A any](fa GA) W { _ = "STUB: not implemented"; return *new(W) }

func Evaluate[GA ~func() P.Pair[A, W], W, A any](fa GA) A {
	_ = "STUB: not implemented"
	return *

	// MonadCensor modifies the final accumulator value by applying a function
	new(A)
}

func MonadCensor[GA ~func() P.Pair[A, W], FCT ~func(W) W, W, A any](fa GA, f FCT) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// Censor modifies the final accumulator value by applying a function
func Censor[GA ~func() P.Pair[A, W], FCT ~func(W) W, W, A any](f FCT) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadListens projects a value from modifications made to the accumulator during an action
func MonadListens[GA ~func() P.Pair[A, W], GAB ~func() P.Pair[P.Pair[A, B], W], FCT ~func(W) B, W, A, B any](fa GA, f FCT) GAB {
	_ = "STUB: not implemented"
	return *new(GAB)
}

// Listens projects a value from modifications made to the accumulator during an action
func Listens[GA ~func() P.Pair[A, W], GAB ~func() P.Pair[P.Pair[A, B], W], FCT ~func(W) B, W, A, B any](f FCT) func(GA) GAB {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[FAB ~func(A) B, GFAB ~func() P.Pair[FAB, W], GB ~func() P.Pair[B, W], W, A, B any](fab GFAB, a A) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Flap[FAB ~func(A) B, GFAB ~func() P.Pair[FAB, W], GB ~func() P.Pair[B, W], W, A, B any](a A) func(GFAB) GB {
	_ = "STUB: not implemented"
	return nil
}
