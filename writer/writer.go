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

package writer

import (
	EM "github.com/IBM/fp-go/endomorphism"
	IO "github.com/IBM/fp-go/io"
	M "github.com/IBM/fp-go/monoid"
	P "github.com/IBM/fp-go/pair"
	SG "github.com/IBM/fp-go/semigroup"
)

type Writer[W, A any] IO.IO[P.Pair[A, W]]

// Tell appends a value to the accumulator
func Tell[W any](w W) Writer[W, any] { _ = "STUB: not implemented"; return nil }

func Of[A, W any](m M.Monoid[W], a A) Writer[W, A] { _ = "STUB: not implemented"; return nil }

// Listen modifies the result to include the changes to the accumulator
func Listen[W, A any](fa Writer[W, A]) Writer[W, P.Pair[A, W]] {
	_ = "STUB: not implemented"
	return nil
}

// Pass applies the returned function to the accumulator
func Pass[W, A any](fa Writer[W, P.Pair[A, EM.Endomorphism[W]]]) Writer[W, A] {
	_ = "STUB: not implemented"
	return nil
}

func MonadMap[FCT ~func(A) B, W, A, B any](fa Writer[W, A], f FCT) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[W any, FCT ~func(A) B, A, B any](f FCT) func(Writer[W, A]) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[FCT ~func(A) Writer[W, B], W, A, B any](s SG.Semigroup[W], fa Writer[W, A], fct FCT) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A, B, W any](s SG.Semigroup[W], fa func(A) Writer[W, B]) func(Writer[W, A]) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[B, A, W any](s SG.Semigroup[W], fab Writer[W, func(A) B], fa Writer[W, A]) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[B, A, W any](s SG.Semigroup[W], fa Writer[W, A]) func(Writer[W, func(A) B]) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[FCT ~func(A) Writer[W, B], W, A, B any](s SG.Semigroup[W], fa Writer[W, A], fct FCT) Writer[W, A] {
	_ = "STUB: not implemented"
	return nil
}

func ChainFirst[FCT ~func(A) Writer[W, B], W, A, B any](s SG.Semigroup[W], fct FCT) func(Writer[W, A]) Writer[W, A] {
	_ = "STUB: not implemented"
	return nil
}

func Flatten[W, A any](s SG.Semigroup[W], mma Writer[W, Writer[W, A]]) Writer[W, A] {
	_ = "STUB: not implemented"
	return nil
}

// Execute extracts the accumulator
func Execute[W, A any](fa Writer[W, A]) W {
	_ = "STUB: not implemented"
	return *

	// Evaluate extracts the value
	new(W)
}

func Evaluate[W, A any](fa Writer[W, A]) A {
	_ = "STUB: not implemented"
	return *

	// MonadCensor modifies the final accumulator value by applying a function
	new(A)
}

func MonadCensor[A any, FCT ~func(W) W, W any](fa Writer[W, A], f FCT) Writer[W, A] {
	_ = "STUB: not implemented"
	return nil
}

// Censor modifies the final accumulator value by applying a function
func Censor[A any, FCT ~func(W) W, W any](f FCT) func(Writer[W, A]) Writer[W, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadListens projects a value from modifications made to the accumulator during an action
func MonadListens[A any, FCT ~func(W) B, W, B any](fa Writer[W, A], f FCT) Writer[W, P.Pair[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

// Listens projects a value from modifications made to the accumulator during an action
func Listens[A any, FCT ~func(W) B, W, B any](f FCT) func(Writer[W, A]) Writer[W, P.Pair[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[W, B, A any](fab Writer[W, func(A) B], a A) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[W, B, A any](a A) func(Writer[W, func(A) B]) Writer[W, B] {
	_ = "STUB: not implemented"
	return nil
}
