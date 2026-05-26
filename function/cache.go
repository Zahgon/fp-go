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

package function

// Memoize converts a unary function into a unary function that caches the value depending on the parameter
func Memoize[K comparable, T any](f func(K) T) func(K) T { _ = "STUB: not implemented"; return nil }

// ContramapMemoize converts a unary function into a unary function that caches the value depending on the parameter
func ContramapMemoize[T, A any, K comparable](kf func(A) K) func(func(A) T) func(A) T {
	_ = "STUB: not implemented"
	return nil
}

// CacheCallback converts a unary function into a unary function that caches the value depending on the parameter
func CacheCallback[
	T, A any, K comparable](kf func(A) K, getOrCreate func(K, func() func() T) func() T) func(func(A) T) func(A) T {
	_ = "STUB: not implemented"
	return nil
}

// SingleElementCache creates a cache function for use with the [CacheCallback] method that has a maximum capacity of one single item
func SingleElementCache[K comparable, T any]() func(K, func() func() T) func() T {
	_ = "STUB: not implemented"
	return nil
}
