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

// these functions curry a golang function with the context as the firsr parameter into a reader with the context as the last parameter, which
// is a equivalent to a function returning a reader of that context
// this goes back to the advice in https://pkg.go.dev/context to put the context as a first parameter as a convention

func Curry0[GA ~func(R) A, R, A any](f func(R) A) GA { _ = "STUB: not implemented"; return *new(GA) }

func Curry1[GA ~func(R) A, R, T1, A any](f func(R, T1) A) func(T1) GA {
	_ = "STUB: not implemented"
	return nil
}

func Curry2[GA ~func(R) A, R, T1, T2, A any](f func(R, T1, T2) A) func(T1) func(T2) GA {
	_ = "STUB: not implemented"
	return nil
}

func Curry3[GA ~func(R) A, R, T1, T2, T3, A any](f func(R, T1, T2, T3) A) func(T1) func(T2) func(T3) GA {
	_ = "STUB: not implemented"
	return nil
}

func Curry4[GA ~func(R) A, R, T1, T2, T3, T4, A any](f func(R, T1, T2, T3, T4) A) func(T1) func(T2) func(T3) func(T4) GA {
	_ = "STUB: not implemented"
	return nil
}

func Uncurry0[GA ~func(R) A, R, A any](f GA) func(R) A { _ = "STUB: not implemented"; return nil }

func Uncurry1[GA ~func(R) A, R, T1, A any](f func(T1) GA) func(R, T1) A {
	_ = "STUB: not implemented"
	return nil
}

func Uncurry2[GA ~func(R) A, R, T1, T2, A any](f func(T1) func(T2) GA) func(R, T1, T2) A {
	_ = "STUB: not implemented"
	return nil
}

func Uncurry3[GA ~func(R) A, R, T1, T2, T3, A any](f func(T1) func(T2) func(T3) GA) func(R, T1, T2, T3) A {
	_ = "STUB: not implemented"
	return nil
}

func Uncurry4[GA ~func(R) A, R, T1, T2, T3, T4, A any](f func(T1) func(T2) func(T3) func(T4) GA) func(R, T1, T2, T3, T4) A {
	_ = "STUB: not implemented"
	return nil
}
