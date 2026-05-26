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
	ET "github.com/IBM/fp-go/either"
)

// MonadTraverseArray transforms an array
func MonadTraverseArray[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseArray transforms an array
func TraverseArray[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// MonadTraverseArrayWithIndex transforms an array
func MonadTraverseArrayWithIndex[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](tas AAS, f func(int, A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseArrayWithIndex transforms an array
func TraverseArrayWithIndex[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](f func(int, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
func SequenceArray[GA ~func() ET.Either[E, A], GAS ~func() ET.Either[E, AAS], AAS ~[]A, GAAS ~[]GA, E, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseRecord transforms an array
func MonadTraverseRecord[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseRecord transforms an array
func TraverseRecord[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex transforms an array
func TraverseRecordWithIndex[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](f func(K, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecord converts a homogeneous sequence of either into an either of sequence
func SequenceRecord[GA ~func() ET.Either[E, A], GAS ~func() ET.Either[E, AAS], AAS ~map[K]A, GAAS ~map[K]GA, K comparable, E, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseArraySeq transforms an array
func MonadTraverseArraySeq[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseArraySeq transforms an array
func TraverseArraySeq[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// MonadTraverseArrayWithIndexSeq transforms an array
func MonadTraverseArrayWithIndexSeq[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](tas AAS, f func(int, A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseArrayWithIndexSeq transforms an array
func TraverseArrayWithIndexSeq[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](f func(int, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArraySeq converts a homogeneous sequence of either into an either of sequence
func SequenceArraySeq[GA ~func() ET.Either[E, A], GAS ~func() ET.Either[E, AAS], AAS ~[]A, GAAS ~[]GA, E, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseRecordSeq transforms an array
func MonadTraverseRecordSeq[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseRecordSeq transforms an array
func TraverseRecordSeq[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexSeq transforms an array
func TraverseRecordWithIndexSeq[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](f func(K, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordSeq converts a homogeneous sequence of either into an either of sequence
func SequenceRecordSeq[GA ~func() ET.Either[E, A], GAS ~func() ET.Either[E, AAS], AAS ~map[K]A, GAAS ~map[K]GA, K comparable, E, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseArrayPar transforms an array
func MonadTraverseArrayPar[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseArrayPar transforms an array
func TraverseArrayPar[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// MonadTraverseArrayWithIndexPar transforms an array
func MonadTraverseArrayWithIndexPar[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](tas AAS, f func(int, A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseArrayWithIndexPar transforms an array
func TraverseArrayWithIndexPar[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~[]A, BBS ~[]B, E, A, B any](f func(int, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArrayPar converts a homogeneous sequence of either into an either of sequence
func SequenceArrayPar[GA ~func() ET.Either[E, A], GAS ~func() ET.Either[E, AAS], AAS ~[]A, GAAS ~[]GA, E, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseRecordPar transforms an array
func MonadTraverseRecordPar[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseRecordPar transforms an array
func TraverseRecordPar[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexPar transforms an array
func TraverseRecordWithIndexPar[GB ~func() ET.Either[E, B], GBS ~func() ET.Either[E, BBS], AAS ~map[K]A, BBS ~map[K]B, K comparable, E, A, B any](f func(K, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordPar converts a homogeneous sequence of either into an either of sequence
func SequenceRecordPar[GA ~func() ET.Either[E, A], GAS ~func() ET.Either[E, AAS], AAS ~map[K]A, GAAS ~map[K]GA, K comparable, E, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}
