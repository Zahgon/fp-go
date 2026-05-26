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

package ioeither

// TraverseArray transforms an array
func TraverseArray[E, A, B any](f func(A) IOEither[E, B]) func([]A) IOEither[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndex transforms an array
func TraverseArrayWithIndex[E, A, B any](f func(int, A) IOEither[E, B]) func([]A) IOEither[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
func SequenceArray[E, A any](ma []IOEither[E, A]) IOEither[E, []A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecord transforms a record
func TraverseRecord[K comparable, E, A, B any](f func(A) IOEither[E, B]) func(map[K]A) IOEither[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex transforms a record
func TraverseRecordWithIndex[K comparable, E, A, B any](f func(K, A) IOEither[E, B]) func(map[K]A) IOEither[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecord converts a homogeneous sequence of either into an either of sequence
func SequenceRecord[K comparable, E, A any](ma map[K]IOEither[E, A]) IOEither[E, map[K]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArraySeq transforms an array
func TraverseArraySeq[E, A, B any](f func(A) IOEither[E, B]) func([]A) IOEither[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexSeq transforms an array
func TraverseArrayWithIndexSeq[E, A, B any](f func(int, A) IOEither[E, B]) func([]A) IOEither[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArraySeq converts a homogeneous sequence of either into an either of sequence
func SequenceArraySeq[E, A any](ma []IOEither[E, A]) IOEither[E, []A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordSeq transforms a record
func TraverseRecordSeq[K comparable, E, A, B any](f func(A) IOEither[E, B]) func(map[K]A) IOEither[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexSeq transforms a record
func TraverseRecordWithIndexSeq[K comparable, E, A, B any](f func(K, A) IOEither[E, B]) func(map[K]A) IOEither[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordSeq converts a homogeneous sequence of either into an either of sequence
func SequenceRecordSeq[K comparable, E, A any](ma map[K]IOEither[E, A]) IOEither[E, map[K]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayPar transforms an array
func TraverseArrayPar[E, A, B any](f func(A) IOEither[E, B]) func([]A) IOEither[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexPar transforms an array
func TraverseArrayWithIndexPar[E, A, B any](f func(int, A) IOEither[E, B]) func([]A) IOEither[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArrayPar converts a homogeneous Paruence of either into an either of Paruence
func SequenceArrayPar[E, A any](ma []IOEither[E, A]) IOEither[E, []A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordPar transforms a record
func TraverseRecordPar[K comparable, E, A, B any](f func(A) IOEither[E, B]) func(map[K]A) IOEither[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexPar transforms a record
func TraverseRecordWithIndexPar[K comparable, E, A, B any](f func(K, A) IOEither[E, B]) func(map[K]A) IOEither[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordPar converts a homogeneous Paruence of either into an either of Paruence
func SequenceRecordPar[K comparable, E, A any](ma map[K]IOEither[E, A]) IOEither[E, map[K]A] {
	_ = "STUB: not implemented"
	return nil
}
