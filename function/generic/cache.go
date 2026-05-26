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

// Memoize converts a unary function into a unary function that caches the value depending on the parameter
func Memoize[F ~func(K) T, K comparable, T any](f F) F { _ = "STUB: not implemented"; return *new(F) }

// ContramapMemoize converts a unary function into a unary function that caches the value depending on the parameter
func ContramapMemoize[F ~func(A) T, KF func(A) K, A any, K comparable, T any](kf KF) func(F) F {
	_ = "STUB: not implemented"
	return nil
}

// getOrCreate is a naive implementation of a cache, without bounds
func getOrCreate[K comparable, T any]() func(K, func() func() T) func() T {
	_ = "STUB: not implemented"
	return nil
}

// only lock to access a lazy accessor to the value

// compute the value outside of the lock

// SingleElementCache is a cache with a capacity of a single element
func SingleElementCache[
	LLT ~func() LT, // generator of the generator
	K comparable, // key into the cache
	LT ~func() T, // generator of a value
	T any, // the cached data type
]() func(K, LLT) LT {
	_ = "STUB: not implemented"
	return nil
}

// update state

// CacheCallback converts a unary function into a unary function that caches the value depending on the parameter
func CacheCallback[
	EM ~func(F) F, // endomorphism of the function
	LLT ~func() LT, // generator of the generator
	LT ~func() T, // generator of a value
	F ~func(A) T, // function to actually cache
	KF func(A) K, // extracts the cache key from the input
	C ~func(K, LLT) LT, // the cache callback function
	A any, K comparable, T any](kf KF, getOrCreate C) EM {
	_ = "STUB: not implemented"
	return *new(EM)
}

// cache entry
