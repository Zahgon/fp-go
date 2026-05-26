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

import (
	T "github.com/IBM/fp-go/tuple"
)

// SequenceT converts n inputs of higher kinded types into a higher kinded types of n strongly typed values, represented as a tuple

func SequenceT1[R, E, A any](a ReaderIOEither[R, E, A]) ReaderIOEither[R, E, T.Tuple1[A]] {
	_ = "STUB: not implemented"
	return nil
}

func SequenceT2[R, E, A, B any](a ReaderIOEither[R, E, A], b ReaderIOEither[R, E, B]) ReaderIOEither[R, E, T.Tuple2[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

func SequenceT3[R, E, A, B, C any](a ReaderIOEither[R, E, A], b ReaderIOEither[R, E, B], c ReaderIOEither[R, E, C]) ReaderIOEither[R, E, T.Tuple3[A, B, C]] {
	_ = "STUB: not implemented"
	return nil
}

func SequenceT4[R, E, A, B, C, D any](a ReaderIOEither[R, E, A], b ReaderIOEither[R, E, B], c ReaderIOEither[R, E, C], d ReaderIOEither[R, E, D]) ReaderIOEither[R, E, T.Tuple4[A, B, C, D]] {
	_ = "STUB: not implemented"
	return nil
}
