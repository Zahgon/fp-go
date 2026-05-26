// Copyright (c) 2023 - 2025 IBM Corp.
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

import (
	O "github.com/IBM/fp-go/v2/option"
	T "github.com/IBM/fp-go/v2/tuple"
)

// SequenceT converts n inputs of higher kinded types into a higher kinded types of n strongly typed values, represented as a tuple

func SequenceT1[
	GA ~func(E) O.Option[A],
	GTA ~func(E) O.Option[T.Tuple1[A]],
	E, A any](a GA) GTA {
	_ = "STUB: not implemented"
	return *new(GTA)
}

func SequenceT2[
	GA ~func(E) O.Option[A],
	GB ~func(E) O.Option[B],
	GTAB ~func(E) O.Option[T.Tuple2[A, B]],
	E, A, B any](a GA, b GB) GTAB {
	_ = "STUB: not implemented"
	return *new(GTAB)
}

func SequenceT3[
	GA ~func(E) O.Option[A],
	GB ~func(E) O.Option[B],
	GC ~func(E) O.Option[C],
	GTABC ~func(E) O.Option[T.Tuple3[A, B, C]],
	E, A, B, C any](a GA, b GB, c GC) GTABC {
	_ = "STUB: not implemented"
	return *new(GTABC)
}

func SequenceT4[
	GA ~func(E) O.Option[A],
	GB ~func(E) O.Option[B],
	GC ~func(E) O.Option[C],
	GD ~func(E) O.Option[D],
	GTABCD ~func(E) O.Option[T.Tuple4[A, B, C, D]],
	E, A, B, C, D any](a GA, b GB, c GC, d GD) GTABCD {
	_ = "STUB: not implemented"
	return *new(GTABCD)
}
