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

import (
	ET "github.com/IBM/fp-go/either"
	T "github.com/IBM/fp-go/tuple"
)

// SequenceT converts n inputs of higher kinded types into a higher kinded types of n strongly typed values, represented as a tuple

func SequenceT1[
	GA ~func(E) GIOA,
	GTA ~func(E) GIOTA,
	GIOA ~func() ET.Either[L, A],
	GIOTA ~func() ET.Either[L, T.Tuple1[A]],
	E, L, A any](a GA) GTA {
	_ = "STUB: not implemented"
	return *new(GTA)
}

func SequenceT2[
	GA ~func(E) GIOA,
	GB ~func(E) GIOB,
	GTAB ~func(E) GIOTAB,
	GIOA ~func() ET.Either[L, A],
	GIOB ~func() ET.Either[L, B],
	GIOTAB ~func() ET.Either[L, T.Tuple2[A, B]],
	E, L, A, B any](a GA, b GB) GTAB {
	_ = "STUB: not implemented"
	return *new(GTAB)
}

func SequenceT3[
	GA ~func(E) GIOA,
	GB ~func(E) GIOB,
	GC ~func(E) GIOC,
	GTABC ~func(E) GIOTABC,
	GIOA ~func() ET.Either[L, A],
	GIOB ~func() ET.Either[L, B],
	GIOC ~func() ET.Either[L, C],
	GIOTABC ~func() ET.Either[L, T.Tuple3[A, B, C]],
	E, L, A, B, C any](a GA, b GB, c GC) GTABC {
	_ = "STUB: not implemented"
	return *new(GTABC)
}

func SequenceT4[
	GA ~func(E) GIOA,
	GB ~func(E) GIOB,
	GC ~func(E) GIOC,
	GD ~func(E) GIOD,
	GTABCD ~func(E) GIOTABCD,
	GIOA ~func() ET.Either[L, A],
	GIOB ~func() ET.Either[L, B],
	GIOC ~func() ET.Either[L, C],
	GIOD ~func() ET.Either[L, D],
	GIOTABCD ~func() ET.Either[L, T.Tuple4[A, B, C, D]],
	E, L, A, B, C, D any](a GA, b GB, c GC, d GD) GTABCD {
	_ = "STUB: not implemented"
	return *new(GTABCD)
}
