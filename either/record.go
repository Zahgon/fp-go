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

package either

// TraverseRecordG transforms a record of options into an option of a record
func TraverseRecordG[GA ~map[K]A, GB ~map[K]B, K comparable, E, A, B any](f func(A) Either[E, B]) func(GA) Either[E, GB] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecord transforms a record of eithers into an either of a record
func TraverseRecord[K comparable, E, A, B any](f func(A) Either[E, B]) func(map[K]A) Either[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexG transforms a record of options into an option of a record
func TraverseRecordWithIndexG[GA ~map[K]A, GB ~map[K]B, K comparable, E, A, B any](f func(K, A) Either[E, B]) func(GA) Either[E, GB] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex transforms a record of eithers into an either of a record
func TraverseRecordWithIndex[K comparable, E, A, B any](f func(K, A) Either[E, B]) func(map[K]A) Either[E, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

func SequenceRecordG[GA ~map[K]A, GOA ~map[K]Either[E, A], K comparable, E, A any](ma GOA) Either[E, GA] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecord converts a homogeneous sequence of either into an either of sequence
func SequenceRecord[K comparable, E, A any](ma map[K]Either[E, A]) Either[E, map[K]A] {
	_ = "STUB: not implemented"
	return nil
}

func upsertAtReadWrite[M ~map[K]V, K comparable, V any](r M, k K, v V) M {
	_ = "STUB: not implemented"
	return *

	// CompactRecordG discards the noe values and keeps the some values
	new(M)
}

func CompactRecordG[M1 ~map[K]Either[E, A], M2 ~map[K]A, K comparable, E, A any](m M1) M2 {
	_ = "STUB: not implemented"
	return *new(M2)
}

// CompactRecord discards all none values and keeps the somes
func CompactRecord[K comparable, E, A any](m map[K]Either[E, A]) map[K]A {
	_ = "STUB: not implemented"
	return nil
}
