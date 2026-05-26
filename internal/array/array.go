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

package array

func Slice[GA ~[]A, A any](low, high int) func(as GA) GA { _ = "STUB: not implemented"; return nil }

// Handle negative indices - count backward from the end

// Start index > array length: return empty array

// End index > array length: slice to the end

// Start >= end: return empty array

func IsEmpty[GA ~[]A, A any](as GA) bool { _ = "STUB: not implemented"; return false }

func IsNil[GA ~[]A, A any](as GA) bool { _ = "STUB: not implemented"; return false }

func IsNonNil[GA ~[]A, A any](as GA) bool { _ = "STUB: not implemented"; return false }

func Reduce[GA ~[]A, A, B any](fa GA, f func(B, A) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func ReduceWithIndex[GA ~[]A, A, B any](fa GA, f func(int, B, A) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func ReduceRight[GA ~[]A, A, B any](fa GA, f func(A, B) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func ReduceRightWithIndex[GA ~[]A, A, B any](fa GA, f func(int, A, B) B, initial B) B {
	_ = "STUB: not implemented"
	return *new(B)
}

func Append[GA ~[]A, A any](as GA, a A) GA { _ = "STUB: not implemented"; return *new(GA) }

func Push[GA ~[]A, A any](as GA, a A) GA { _ = "STUB: not implemented"; return *new(GA) }

func Empty[GA ~[]A, A any]() GA { _ = "STUB: not implemented"; return *new(GA) }

func upsertAt[GA ~[]A, A any](fa GA, a A) GA { _ = "STUB: not implemented"; return *new(GA) }

func UpsertAt[GA ~[]A, A any](a A) func(GA) GA { _ = "STUB: not implemented"; return nil }

func MonadMap[GA ~[]A, GB ~[]B, A, B any](as GA, f func(a A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func Map[GA ~[]A, GB ~[]B, A, B any](f func(a A) B) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}

func MonadMapWithIndex[GA ~[]A, GB ~[]B, A, B any](as GA, f func(idx int, a A) B) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

func ConstNil[GA ~[]A, A any]() GA { _ = "STUB: not implemented"; return *new(GA) }
