// Copyright (c) 2023 - 2025 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache LicensVersion 2.0 (the "License");
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

package ioresult

// TraverseArray transforms an array
//
//go:inline
func TraverseArray[A, B any](f Kleisli[A, B]) Kleisli[[]A, []B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndex transforms an array
//
//go:inline
func TraverseArrayWithIndex[A, B any](f func(int, A) IOResult[B]) Kleisli[[]A, []B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
//
//go:inline
func SequenceArray[A any](ma []IOResult[A]) IOResult[[]A] { _ = "STUB: not implemented"; return nil }

// TraverseRecord transforms a record
//
//go:inline
func TraverseRecord[K comparable, A, B any](f Kleisli[A, B]) Kleisli[map[K]A, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndex transforms a record
//
//go:inline
func TraverseRecordWithIndex[K comparable, A, B any](f func(K, A) IOResult[B]) Kleisli[map[K]A, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecord converts a homogeneous sequence of either into an either of sequence
//
//go:inline
func SequenceRecord[K comparable, A any](ma map[K]IOResult[A]) IOResult[map[K]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArraySeq transforms an array
//
//go:inline
func TraverseArraySeq[A, B any](f Kleisli[A, B]) Kleisli[[]A, []B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexSeq transforms an array
//
//go:inline
func TraverseArrayWithIndexSeq[A, B any](f func(int, A) IOResult[B]) Kleisli[[]A, []B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArraySeq converts a homogeneous sequence of either into an either of sequence
//
//go:inline
func SequenceArraySeq[A any](ma []IOResult[A]) IOResult[[]A] { _ = "STUB: not implemented"; return nil }

// TraverseRecordSeq transforms a record
//
//go:inline
func TraverseRecordSeq[K comparable, A, B any](f Kleisli[A, B]) Kleisli[map[K]A, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexSeq transforms a record
//
//go:inline
func TraverseRecordWithIndexSeq[K comparable, A, B any](f func(K, A) IOResult[B]) Kleisli[map[K]A, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordSeq converts a homogeneous sequence of either into an either of sequence
//
//go:inline
func SequenceRecordSeq[K comparable, A any](ma map[K]IOResult[A]) IOResult[map[K]A] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayPar transforms an array
//
//go:inline
func TraverseArrayPar[A, B any](f Kleisli[A, B]) Kleisli[[]A, []B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexPar transforms an array
//
//go:inline
func TraverseArrayWithIndexPar[A, B any](f func(int, A) IOResult[B]) Kleisli[[]A, []B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArrayPar converts a homogeneous Paruence of either into an either of Paruence
//
//go:inline
func SequenceArrayPar[A any](ma []IOResult[A]) IOResult[[]A] { _ = "STUB: not implemented"; return nil }

// TraverseRecordPar transforms a record
//
//go:inline
func TraverseRecordPar[K comparable, A, B any](f Kleisli[A, B]) Kleisli[map[K]A, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseRecordWithIndexPar transforms a record
//
//go:inline
func TraverseRecordWithIndexPar[K comparable, A, B any](f func(K, A) IOResult[B]) Kleisli[map[K]A, map[K]B] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceRecordPar converts a homogeneous Paruence of either into an either of Paruence
//
//go:inline
func SequenceRecordPar[K comparable, A any](ma map[K]IOResult[A]) IOResult[map[K]A] {
	_ = "STUB: not implemented"
	return nil
}
