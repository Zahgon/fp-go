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
	"context"
	"time"

	E "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
)

const (
	// useParallel is the feature flag to control if we use the parallel or the sequential implementation of ap
	useParallel = true
)

func FromEither[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],
	A any](e E.Either[error, A]) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func RightReader[
	GRA ~func(context.Context) GIOA,
	GR ~func(context.Context) A,
	GIOA ~func() E.Either[error, A],
	A any](r GR) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func LeftReader[
	GRA ~func(context.Context) GIOA,
	GR ~func(context.Context) error,
	GIOA ~func() E.Either[error, A],
	A any](l GR) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func Left[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],
	A any](l error) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func Right[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],
	A any](r A) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func FromReader[
	GRA ~func(context.Context) GIOA,
	GR ~func(context.Context) A,
	GIOA ~func() E.Either[error, A],
	A any](r GR) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func MonadMap[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](fa GRA, f func(A) B) GRB {
	_ = "STUB: not implemented"
	return *new(GRB)
}

func Map[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](f func(A) B) func(GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapTo[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](fa GRA, b B) GRB {
	_ = "STUB: not implemented"
	return *new(GRB)
}

func MapTo[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](b B) func(GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](ma GRA, f func(A) GRB) GRB {
	_ = "STUB: not implemented"
	return *new(GRB)
}

func Chain[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](f func(A) GRB) func(GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirst[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](ma GRA, f func(A) GRB) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func ChainFirst[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	A, B any](f func(A) GRB) func(GRA) GRA {
	_ = "STUB: not implemented"
	return nil
}

func Of[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any](a A) GRA {
	_ = "STUB: not implemented"
	return *

	// withCancelCauseFunc wraps an IOEither such that in case of an error the cancel function is invoked
	new(GRA)
}

func withCancelCauseFunc[
	GIOA ~func() E.Either[error, A],
	A any](cancel context.CancelCauseFunc, ma GIOA) GIOA {
	_ = "STUB: not implemented"
	return *new(GIOA)
}

// MonadApSeq implements the `Ap` function for a reader with context. It creates a sub-context that will
// be canceled if any of the input operations errors out or
func MonadApSeq[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GRAB ~func(context.Context) GIOAB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],
	GIOAB ~func() E.Either[error, func(A) B],

	A, B any](fab GRAB, fa GRA) GRB {
	_ = "STUB: not implemented"
	return *new(GRB)
}

// MonadAp implements the `Ap` function for a reader with context. It creates a sub-context that will
// be canceled if any of the input operations errors out or
func MonadApPar[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GRAB ~func(context.Context) GIOAB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],
	GIOAB ~func() E.Either[error, func(A) B],

	A, B any](fab GRAB, fa GRA) GRB {
	_ = "STUB: not implemented"
	// context sensitive input
	return *new(GRB)
}

// quick check for cancellation

// quick check for cancellation

// create sub-contexts for fa and fab, so they can cancel one other

// cancel has to be called in all paths

// MonadAp implements the `Ap` function for a reader with context. It creates a sub-context that will
// be canceled if any of the input operations errors out or
func MonadAp[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GRAB ~func(context.Context) GIOAB,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],
	GIOAB ~func() E.Either[error, func(A) B],

	A, B any](fab GRAB, fa GRA) GRB {
	_ = "STUB: not implemented"
	// dispatch to the configured version
	return *new(GRB)
}

func Ap[
	GRB ~func(context.Context) GIOB,
	GRAB ~func(context.Context) GIOAB,
	GRA ~func(context.Context) GIOA,

	GIOB ~func() E.Either[error, B],
	GIOAB ~func() E.Either[error, func(A) B],
	GIOA ~func() E.Either[error, A],

	A, B any](fa GRA) func(GRAB) GRB {
	_ = "STUB: not implemented"
	return nil
}

func ApSeq[
	GRB ~func(context.Context) GIOB,
	GRAB ~func(context.Context) GIOAB,
	GRA ~func(context.Context) GIOA,

	GIOB ~func() E.Either[error, B],
	GIOAB ~func() E.Either[error, func(A) B],
	GIOA ~func() E.Either[error, A],

	A, B any](fa GRA) func(GRAB) GRB {
	_ = "STUB: not implemented"
	return nil
}

func ApPar[
	GRB ~func(context.Context) GIOB,
	GRAB ~func(context.Context) GIOAB,
	GRA ~func(context.Context) GIOA,

	GIOB ~func() E.Either[error, B],
	GIOAB ~func() E.Either[error, func(A) B],
	GIOA ~func() E.Either[error, A],

	A, B any](fa GRA) func(GRAB) GRB {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any](pred func(A) bool, onFalse func(A) error) func(A) GRA {
	_ = "STUB: not implemented"
	return nil
}

func Fold[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() B,

	A, B any](onLeft func(error) GRB, onRight func(A) GRB) func(GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func GetOrElse[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	GIOB ~func() A,

	A any](onLeft func(error) GRB) func(GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any](onLeft func(error) GRA) func(GRA) GRA {
	_ = "STUB: not implemented"
	return nil
}

func OrLeft[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,
	GIOA ~func() E.Either[error, A],

	GIOB ~func() error,

	A any](onLeft func(error) GRB) func(GRA) GRA {
	_ = "STUB: not implemented"
	return nil
}

func Ask[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, context.Context],

]() GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func Asks[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) A,
	GIOA ~func() E.Either[error, A],

	A any](r GRB) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func MonadChainEitherK[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,
	GIOA ~func() E.Either[error, A],

	GIOB ~func() E.Either[error, B],

	A, B any](ma GRA, f func(A) E.Either[error, B]) GRB {
	_ = "STUB: not implemented"
	return *new(GRB)
}

func ChainEitherK[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,
	GIOA ~func() E.Either[error, A],

	GIOB ~func() E.Either[error, B],

	A, B any](f func(A) E.Either[error, B]) func(ma GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstEitherK[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A, B any](ma GRA, f func(A) E.Either[error, B]) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func ChainFirstEitherK[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A, B any](f func(A) E.Either[error, B]) func(ma GRA) GRA {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,

	GIOB ~func() E.Either[error, B],

	GIOA ~func() E.Either[error, A],

	A, B any](onNone func() error) func(func(A) O.Option[B]) func(GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func FromIOEither[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any](t GIOA) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func FromIO[
	GRA ~func(context.Context) GIOA,
	GIOB ~func() A,

	GIOA ~func() E.Either[error, A],

	A any](t GIOB) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

// Never returns a 'ReaderIOEither' that never returns, except if its context gets canceled
func Never[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any]() GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func MonadChainIOK[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	GIOB ~func() E.Either[error, B],

	GIO ~func() B,

	A, B any](ma GRA, f func(A) GIO) GRB {
	_ = "STUB: not implemented"
	return *new(GRB)
}

func ChainIOK[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	GIOB ~func() E.Either[error, B],

	GIO ~func() B,

	A, B any](f func(A) GIO) func(ma GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainReaderIOK[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GRIO ~func(context.Context) GIO,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	GIO ~func() B,

	A, B any](ma GRA, f func(A) GRIO) GRB {
	_ = "STUB: not implemented"
	return *new(GRB)
}

func ChainReaderIOK[
	GRB ~func(context.Context) GIOB,
	GRA ~func(context.Context) GIOA,
	GRIO ~func(context.Context) GIO,

	GIOA ~func() E.Either[error, A],
	GIOB ~func() E.Either[error, B],

	GIO ~func() B,

	A, B any](f func(A) GRIO) func(ma GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainFirstIOK[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	GIO ~func() B,

	A, B any](ma GRA, f func(A) GIO) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func ChainFirstIOK[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	GIO ~func() B,

	A, B any](f func(A) GIO) func(ma GRA) GRA {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOEitherK[
	GRA ~func(context.Context) GIOA,
	GRB ~func(context.Context) GIOB,
	GIOA ~func() E.Either[error, A],

	GIOB ~func() E.Either[error, B],

	A, B any](f func(A) GIOB) func(ma GRA) GRB {
	_ = "STUB: not implemented"
	return nil
}

// Delay creates an operation that passes in the value after some delay
func Delay[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any](delay time.Duration) func(ma GRA) GRA {
	_ = "STUB: not implemented"
	return nil
}

// manage the timeout

// whatever comes first

// Timer will return the current time after an initial delay
func Timer[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, time.Time],

](delay time.Duration) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

// Defer creates an IO by creating a brand new IO via a generator function, each time
func Defer[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any](gen func() GRA) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

// TryCatch wraps a reader returning a tuple as an error into ReaderIOEither
func TryCatch[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],

	A any](f func(context.Context) func() (A, error)) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func MonadAlt[LAZY ~func() GEA, GEA ~func(context.Context) GIOA, GIOA ~func() E.Either[error, A], A any](first GEA, second LAZY) GEA {
	_ = "STUB: not implemented"
	return *new(GEA)
}

func Alt[LAZY ~func() GEA, GEA ~func(context.Context) GIOA, GIOA ~func() E.Either[error, A], A any](second LAZY) func(GEA) GEA {
	_ = "STUB: not implemented"
	return nil

	// Memoize computes the value of the provided monad lazily but exactly once
	// The context used to compute the value is the context of the first call, so do not use this
	// method if the value has a functional dependency on the content of the context
}

func Memoize[
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],
	A any](rdr GRA) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func Flatten[
	GGRA ~func(context.Context) GGIOA,
	GGIOA ~func() E.Either[error, GRA],
	GRA ~func(context.Context) GIOA,
	GIOA ~func() E.Either[error, A],
	A any](rdr GGRA) GRA {
	_ = "STUB: not implemented"
	return *new(GRA)
}

func MonadFromReaderIO[
	GRIOEA ~func(context.Context) GIOEA,
	GIOEA ~func() E.Either[error, A],

	GRIOA ~func(context.Context) GIOA,
	GIOA ~func() A,

	A any](a A, f func(A) GRIOA) GRIOEA {
	_ = "STUB: not implemented"
	return *new(GRIOEA)
}

func FromReaderIO[
	GRIOEA ~func(context.Context) GIOEA,
	GIOEA ~func() E.Either[error, A],

	GRIOA ~func(context.Context) GIOA,
	GIOA ~func() A,

	A any](f func(A) GRIOA) func(A) GRIOEA {
	_ = "STUB: not implemented"
	return nil
}

func RightReaderIO[
	GRIOEA ~func(context.Context) GIOEA,
	GIOEA ~func() E.Either[error, A],

	GRIOA ~func(context.Context) GIOA,
	GIOA ~func() A,

	A any](ma GRIOA) GRIOEA {
	_ = "STUB: not implemented"
	return *new(GRIOEA)
}

func LeftReaderIO[
	GRIOEA ~func(context.Context) GIOEA,
	GIOEA ~func() E.Either[error, A],

	GRIOE ~func(context.Context) GIOE,
	GIOE ~func() error,

	A any](ma GRIOE) GRIOEA {
	_ = "STUB: not implemented"
	return *new(GRIOEA)
}

func MonadFlap[GREAB ~func(context.Context) GEAB, GREB ~func(context.Context) GEB, GEAB ~func() E.Either[error, func(A) B], GEB ~func() E.Either[error, B], B, A any](fab GREAB, a A) GREB {
	_ = "STUB: not implemented"
	return *new(GREB)
}

func Flap[GREAB ~func(context.Context) GEAB, GREB ~func(context.Context) GEB, GEAB ~func() E.Either[error, func(A) B], GEB ~func() E.Either[error, B], B, A any](a A) func(GREAB) GREB {
	_ = "STUB: not implemented"
	return nil
}
