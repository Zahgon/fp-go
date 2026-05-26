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

func Variadic0[V, R any](f func([]V) (R, error)) func(...V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Variadic1[T1, V, R any](f func(T1, []V) (R, error)) func(T1, ...V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Variadic2[T1, T2, V, R any](f func(T1, T2, []V) (R, error)) func(T1, T2, ...V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Variadic3[T1, T2, T3, V, R any](f func(T1, T2, T3, []V) (R, error)) func(T1, T2, T3, ...V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Variadic4[T1, T2, T3, T4, V, R any](f func(T1, T2, T3, T4, []V) (R, error)) func(T1, T2, T3, T4, ...V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Unvariadic0[V, R any](f func(...V) (R, error)) func([]V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Unvariadic1[T1, V, R any](f func(T1, ...V) (R, error)) func(T1, []V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Unvariadic2[T1, T2, V, R any](f func(T1, T2, ...V) (R, error)) func(T1, T2, []V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Unvariadic3[T1, T2, T3, V, R any](f func(T1, T2, T3, ...V) (R, error)) func(T1, T2, T3, []V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

func Unvariadic4[T1, T2, T3, T4, V, R any](f func(T1, T2, T3, T4, ...V) (R, error)) func(T1, T2, T3, T4, []V) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}
