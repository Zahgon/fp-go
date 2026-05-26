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

	E "github.com/IBM/fp-go/either"
)

// MonadTraverseArray transforms an array
func MonadTraverseArray[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](as AS, f func(A) GRB) GRBS {
	_ = "STUB: not implemented"
	return *new(GRBS)
}

// TraverseArray transforms an array
func TraverseArray[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](f func(A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndex transforms an array
func TraverseArrayWithIndex[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](f func(int, A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
func SequenceArray[
	AS ~[]A,
	GAS ~[]GRA,
	GRAS ~func(context.Context) GIOAS,
	GRA ~func(context.Context) GIOA,
	GIOAS ~func() E.Either[error, AS],
	GIOA ~func() E.Either[error, A],
	A any](ma GAS) GRAS {
	_ = "STUB: not implemented"
	return *new(GRAS)
}

// MonadTraverseRecord transforms a record
func MonadTraverseRecord[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](ma AS, f func(A) GRB) GRBS {
	_ = "STUB: not implemented"
	return *new(GRBS)
}

// TraverseRecord transforms a record
func TraverseRecord[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](f func(A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex transforms a record
func TraverseRecordWithIndex[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](f func(K, A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecord converts a homogeneous sequence of either into an either of sequence
func SequenceRecord[K comparable,
	AS ~map[K]A,
	GAS ~map[K]GRA,
	GRAS ~func(context.Context) GIOAS,
	GRA ~func(context.Context) GIOA,
	GIOAS ~func() E.Either[error, AS],
	GIOA ~func() E.Either[error, A],
	A any](ma GAS) GRAS {
	_ = "STUB: not implemented"
	return *new(GRAS)
}

// MonadTraverseArraySeq transforms an array
func MonadTraverseArraySeq[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](as AS, f func(A) GRB) GRBS {
	_ = "STUB: not implemented"
	return *new(GRBS)
}

// TraverseArraySeq transforms an array
func TraverseArraySeq[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](f func(A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexSeq transforms an array
func TraverseArrayWithIndexSeq[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](f func(int, A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArraySeq converts a homogeneous sequence of either into an either of sequence
func SequenceArraySeq[
	AS ~[]A,
	GAS ~[]GRA,
	GRAS ~func(context.Context) GIOAS,
	GRA ~func(context.Context) GIOA,
	GIOAS ~func() E.Either[error, AS],
	GIOA ~func() E.Either[error, A],
	A any](ma GAS) GRAS {
	_ = "STUB: not implemented"
	return *new(GRAS)
}

// MonadTraverseRecordSeq transforms a record
func MonadTraverseRecordSeq[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](ma AS, f func(A) GRB) GRBS {
	_ = "STUB: not implemented"
	return *new(GRBS)
}

// TraverseRecordSeq transforms a record
func TraverseRecordSeq[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](f func(A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexSeq transforms a record
func TraverseRecordWithIndexSeq[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](f func(K, A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordSeq converts a homogeneous sequence of either into an either of sequence
func SequenceRecordSeq[K comparable,
	AS ~map[K]A,
	GAS ~map[K]GRA,
	GRAS ~func(context.Context) GIOAS,
	GRA ~func(context.Context) GIOA,
	GIOAS ~func() E.Either[error, AS],
	GIOA ~func() E.Either[error, A],
	A any](ma GAS) GRAS {
	_ = "STUB: not implemented"
	return *new(GRAS)
}

// MonadTraverseArrayPar transforms an array
func MonadTraverseArrayPar[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](as AS, f func(A) GRB) GRBS {
	_ = "STUB: not implemented"
	return *new(GRBS)
}

// TraverseArrayPar transforms an array
func TraverseArrayPar[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](f func(A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexPar transforms an array
func TraverseArrayWithIndexPar[
	AS ~[]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~[]B,
	A, B any](f func(int, A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArrayPar converts a homogeneous sequence of either into an either of sequence
func SequenceArrayPar[
	AS ~[]A,
	GAS ~[]GRA,
	GRAS ~func(context.Context) GIOAS,
	GRA ~func(context.Context) GIOA,
	GIOAS ~func() E.Either[error, AS],
	GIOA ~func() E.Either[error, A],
	A any](ma GAS) GRAS {
	_ = "STUB: not implemented"
	return *new(GRAS)
}

// MonadTraverseRecordPar transforms a record
func MonadTraverseRecordPar[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](ma AS, f func(A) GRB) GRBS {
	_ = "STUB: not implemented"
	return *new(GRBS)
}

// TraverseRecordPar transforms a record
func TraverseRecordPar[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](f func(A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexPar transforms a record
func TraverseRecordWithIndexPar[K comparable,
	AS ~map[K]A,
	GRBS ~func(context.Context) GIOBS,
	GRB ~func(context.Context) GIOB,
	GIOBS ~func() E.Either[error, BS],
	GIOB ~func() E.Either[error, B],
	BS ~map[K]B,

	A, B any](f func(K, A) GRB) func(AS) GRBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordPar converts a homogeneous sequence of either into an either of sequence
func SequenceRecordPar[K comparable,
	AS ~map[K]A,
	GAS ~map[K]GRA,
	GRAS ~func(context.Context) GIOAS,
	GRA ~func(context.Context) GIOA,
	GIOAS ~func() E.Either[error, AS],
	GIOA ~func() E.Either[error, A],
	A any](ma GAS) GRAS {
	_ = "STUB: not implemented"
	return *new(GRAS)
}
