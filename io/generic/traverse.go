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

func MonadTraverseArray[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

func MonadTraverseArraySeq[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

func MonadTraverseArrayPar[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](tas AAS, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

func TraverseArray[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

func TraverseArraySeq[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

func TraverseArrayPar[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](f func(A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

func TraverseArrayWithIndex[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](f func(int, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

func TraverseArrayWithIndexSeq[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](f func(int, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

func TraverseArrayWithIndexPar[GB ~func() B, GBS ~func() BBS, AAS ~[]A, BBS ~[]B, A, B any](f func(int, A) GB) func(AAS) GBS {
	_ = "STUB: not implemented"
	return nil
}

func SequenceArray[GA ~func() A, GAS ~func() AAS, AAS ~[]A, GAAS ~[]GA, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

func SequenceArraySeq[GA ~func() A, GAS ~func() AAS, AAS ~[]A, GAAS ~[]GA, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

func SequenceArrayPar[GA ~func() A, GAS ~func() AAS, AAS ~[]A, GAAS ~[]GA, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseRecord transforms a record using an IO transform an IO of a record
func MonadTraverseRecord[GBS ~func() MB, MA ~map[K]A, GB ~func() B, MB ~map[K]B, K comparable, A, B any](ma MA, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseRecord transforms a record using an IO transform an IO of a record
func TraverseRecord[GBS ~func() MB, MA ~map[K]A, GB ~func() B, MB ~map[K]B, K comparable, A, B any](f func(A) GB) func(MA) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex transforms a record using an IO transform an IO of a record
func TraverseRecordWithIndex[GB ~func() B, GBS ~func() MB, MA ~map[K]A, MB ~map[K]B, K comparable, A, B any](f func(K, A) GB) func(MA) GBS {
	_ = "STUB: not implemented"
	return nil
}

func SequenceRecord[GA ~func() A, GAS ~func() AAS, AAS ~map[K]A, GAAS ~map[K]GA, K comparable, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseRecordSeq transforms a record using an IO transform an IO of a record
func MonadTraverseRecordSeq[GBS ~func() MB, MA ~map[K]A, GB ~func() B, MB ~map[K]B, K comparable, A, B any](ma MA, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseRecordSeq transforms a record using an IO transform an IO of a record
func TraverseRecordSeq[GBS ~func() MB, MA ~map[K]A, GB ~func() B, MB ~map[K]B, K comparable, A, B any](f func(A) GB) func(MA) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexSeq transforms a record using an IO transform an IO of a record
func TraverseRecordWithIndexSeq[GB ~func() B, GBS ~func() MB, MA ~map[K]A, MB ~map[K]B, K comparable, A, B any](f func(K, A) GB) func(MA) GBS {
	_ = "STUB: not implemented"
	return nil
}

func SequenceRecordSeq[GA ~func() A, GAS ~func() AAS, AAS ~map[K]A, GAAS ~map[K]GA, K comparable, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}

// MonadTraverseRecordPar transforms a record using an IO transform an IO of a record
func MonadTraverseRecordPar[GBS ~func() MB, MA ~map[K]A, GB ~func() B, MB ~map[K]B, K comparable, A, B any](ma MA, f func(A) GB) GBS {
	_ = "STUB: not implemented"
	return *new(GBS)
}

// TraverseRecordPar transforms a record using an IO transform an IO of a record
func TraverseRecordPar[GBS ~func() MB, MA ~map[K]A, GB ~func() B, MB ~map[K]B, K comparable, A, B any](f func(A) GB) func(MA) GBS {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexPar transforms a record using an IO transform an IO of a record
func TraverseRecordWithIndexPar[GB ~func() B, GBS ~func() MB, MA ~map[K]A, MB ~map[K]B, K comparable, A, B any](f func(K, A) GB) func(MA) GBS {
	_ = "STUB: not implemented"
	return nil
}

func SequenceRecordPar[GA ~func() A, GAS ~func() AAS, AAS ~map[K]A, GAAS ~map[K]GA, K comparable, A any](tas GAAS) GAS {
	_ = "STUB: not implemented"
	return *new(GAS)
}
