// Copyright (c) 2023 - 2025 IBM Corp.
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
	"github.com/IBM/fp-go/v2/either"
)

// MonadTraverseArray transforms an array
func MonadTraverseArray[GB ~func(E) GIOB, GBS ~func(E) GIOBS, GIOB ~func() either.Either[L, B], GIOBS ~func() either.Either[L, BBS], AAS ~[]A, BBS ~[]B, E, L, A, B any](ma AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseArray transforms an array
func TraverseArray[GB ~func(E) GIOB, GBS ~func(E) GIOBS, GIOB ~func() either.Either[L, B], GIOBS ~func() either.Either[L, BBS], AAS ~[]A, BBS ~[]B, E, L, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndex transforms an array
func TraverseArrayWithIndex[GB ~func(E) GIOB, GBS ~func(E) GIOBS, GIOB ~func() either.Either[L, B], GIOBS ~func() either.Either[L, BBS], AAS ~[]A, BBS ~[]B, E, L, A, B any](f func(int, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
func SequenceArray[GA ~func(E) GIOA, GAS ~func(E) GIOAS, GIOA ~func() either.Either[L, A], GIOAS ~func() either.Either[L, AAS], AAS ~[]A, GAAS ~[]GA, E, L, A any](ma GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseRecord transforms an array
func MonadTraverseRecord[GB ~func(C) GIOB, GBS ~func(C) GIOBS, GIOB ~func() either.Either[E, B], GIOBS ~func() either.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, C, E, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseRecord transforms a record
func TraverseRecord[GB ~func(C) GIOB, GBS ~func(C) GIOBS, GIOB ~func() either.Either[E, B], GIOBS ~func() either.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, C, E, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex transforms a record
func TraverseRecordWithIndex[GB ~func(C) GIOB, GBS ~func(C) GIOBS, GIOB ~func() either.Either[E, B], GIOBS ~func() either.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, C, E, A, B any](f func(K, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecord converts a homogeneous sequence of either into an either of sequence
func SequenceRecord[GA ~func(C) GIOA, GAS ~func(C) GIOAS, GIOA ~func() either.Either[E, A], GIOAS ~func() either.Either[E, AAS], AAS ~map[K]A, GAAS ~map[K]GA, K comparable, C, E, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}
