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

// TraverseArrayG transforms an array
func TraverseArrayG[GA ~[]A, GB ~[]B, E, A, B any](f func(A) Either[E, B]) func(GA) Either[E, GB] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArray transforms an array
func TraverseArray[E, A, B any](f func(A) Either[E, B]) func([]A) Either[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndexG transforms an array
func TraverseArrayWithIndexG[GA ~[]A, GB ~[]B, E, A, B any](f func(int, A) Either[E, B]) func(GA) Either[E, GB] {
	_ = "STUB: not implemented"
	return nil
}

// TraverseArrayWithIndex transforms an array
func TraverseArrayWithIndex[E, A, B any](f func(int, A) Either[E, B]) func([]A) Either[E, []B] {
	_ = "STUB: not implemented"
	return nil
}

func SequenceArrayG[GA ~[]A, GOA ~[]Either[E, A], E, A any](ma GOA) Either[E, GA] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArray converts a homogeneous sequence of either into an either of sequence
func SequenceArray[E, A any](ma []Either[E, A]) Either[E, []A] {
	_ = "STUB: not implemented"
	return nil
}

// CompactArrayG discards the none values and keeps the right values
func CompactArrayG[A1 ~[]Either[E, A], A2 ~[]A, E, A any](fa A1) A2 {
	_ = "STUB: not implemented"
	return *new(A2)
}

// CompactArray discards the none values and keeps the right values
func CompactArray[E, A any](fa []Either[E, A]) []A { _ = "STUB: not implemented"; return nil }
