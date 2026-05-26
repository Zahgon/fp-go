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

package readerioeither

// TraverseArray uses transforms an array [[]A] into [[]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[[]B]]
func TraverseArray[A, B any](f func(A) ReaderIOEither[B]) func([]A) ReaderIOEither[[]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndex uses transforms an array [[]A] into [[]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[[]B]]
func TraverseArrayWithIndex[A, B any](f func(int, A) ReaderIOEither[B]) func([]A) ReaderIOEither[[]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
func SequenceArray[A any](ma []ReaderIOEither[A]) ReaderIOEither[[]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecord uses transforms a record [map[K]A] into [map[K]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[map[K]B]]
func TraverseRecord[K comparable, A, B any](f func(A) ReaderIOEither[B]) func(map[K]A) ReaderIOEither[map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex uses transforms a record [map[K]A] into [map[K]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[map[K]B]]
func TraverseRecordWithIndex[K comparable, A, B any](f func(K, A) ReaderIOEither[B]) func(map[K]A) ReaderIOEither[map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecord converts a homogeneous sequence of either into an either of sequence
func SequenceRecord[K comparable, A any](ma map[K]ReaderIOEither[A]) ReaderIOEither[map[K]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArraySeq uses transforms an array [[]A] into [[]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[[]B]]
func TraverseArraySeq[A, B any](f func(A) ReaderIOEither[B]) func([]A) ReaderIOEither[[]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexSeq uses transforms an array [[]A] into [[]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[[]B]]
func TraverseArrayWithIndexSeq[A, B any](f func(int, A) ReaderIOEither[B]) func([]A) ReaderIOEither[[]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArraySeq converts a homogeneous sequence of either into an either of sequence
func SequenceArraySeq[A any](ma []ReaderIOEither[A]) ReaderIOEither[[]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordSeq uses transforms a record [map[K]A] into [map[K]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[map[K]B]]
func TraverseRecordSeq[K comparable, A, B any](f func(A) ReaderIOEither[B]) func(map[K]A) ReaderIOEither[map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexSeq uses transforms a record [map[K]A] into [map[K]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[map[K]B]]
func TraverseRecordWithIndexSeq[K comparable, A, B any](f func(K, A) ReaderIOEither[B]) func(map[K]A) ReaderIOEither[map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordSeq converts a homogeneous sequence of either into an either of sequence
func SequenceRecordSeq[K comparable, A any](ma map[K]ReaderIOEither[A]) ReaderIOEither[map[K]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayPar uses transforms an array [[]A] into [[]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[[]B]]
func TraverseArrayPar[A, B any](f func(A) ReaderIOEither[B]) func([]A) ReaderIOEither[[]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexPar uses transforms an array [[]A] into [[]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[[]B]]
func TraverseArrayWithIndexPar[A, B any](f func(int, A) ReaderIOEither[B]) func([]A) ReaderIOEither[[]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArrayPar converts a homogeneous sequence of either into an either of sequence
func SequenceArrayPar[A any](ma []ReaderIOEither[A]) ReaderIOEither[[]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordPar uses transforms a record [map[K]A] into [map[K]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[map[K]B]]
func TraverseRecordPar[K comparable, A, B any](f func(A) ReaderIOEither[B]) func(map[K]A) ReaderIOEither[map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexPar uses transforms a record [map[K]A] into [map[K]ReaderIOEither[B]] and then resolves that into a [ReaderIOEither[map[K]B]]
func TraverseRecordWithIndexPar[K comparable, A, B any](f func(K, A) ReaderIOEither[B]) func(map[K]A) ReaderIOEither[map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordPar converts a homogeneous sequence of either into an either of sequence
func SequenceRecordPar[K comparable, A any](ma map[K]ReaderIOEither[A]) ReaderIOEither[map[K]A] {
	_ = "STUB: not implemented"
	return nil
}
