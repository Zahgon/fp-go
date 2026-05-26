// Copyright (c) 2024 - 2025 IBM Corp.
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

package statet

import (
	"github.com/IBM/fp-go/v2/pair"
)

func Of[
	HKTSA ~func(S) HKTA,
	HKTA,
	S, A any,
](
	fof func(pair.Pair[S, A]) HKTA,

	a A) HKTSA {
	_ = "STUB: not implemented"
	return *new(HKTSA)
}

func MonadMap[
	HKTSA ~func(S) HKTA,
	HKTSB ~func(S) HKTB,
	HKTA,
	HKTB,
	S, A, B any,
](
	fmap func(HKTA, func(pair.Pair[S, A]) pair.Pair[S, B]) HKTB,

	fa HKTSA,
	f func(A) B,
) HKTSB {
	_ = "STUB: not implemented"
	return *new(HKTSB)
}

func Map[
	HKTSA ~func(S) HKTA,
	HKTSB ~func(S) HKTB,
	HKTA,
	HKTB,
	S, A, B any,
](
	fmap func(func(pair.Pair[S, A]) pair.Pair[S, B]) func(HKTA) HKTB,

	f func(A) B,
) func(HKTSA) HKTSB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[
	HKTSA ~func(S) HKTA,
	HKTSB ~func(S) HKTB,
	HKTA,
	HKTB,
	S, A any,
](
	fchain func(HKTA, func(pair.Pair[S, A]) HKTB) HKTB,

	fa HKTSA,
	f func(A) HKTSB,
) HKTSB {
	_ = "STUB: not implemented"
	return *new(HKTSB)
}

func Chain[
	HKTSA ~func(S) HKTA,
	HKTSB ~func(S) HKTB,
	HKTA,
	HKTB,
	S, A any,
](
	fchain func(func(pair.Pair[S, A]) HKTB) func(HKTA) HKTB,

	f func(A) HKTSB,
) func(HKTSA) HKTSB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[
	HKTSA ~func(S) HKTA,
	HKTSB ~func(S) HKTB,
	HKTSAB ~func(S) HKTAB,
	HKTA,
	HKTB,
	HKTAB,

	S, A, B any,
](
	fmap func(HKTA, func(pair.Pair[S, A]) pair.Pair[S, B]) HKTB,
	fchain func(HKTAB, func(pair.Pair[S, func(A) B]) HKTB) HKTB,

	fab HKTSAB,
	fa HKTSA,
) HKTSB {
	_ = "STUB: not implemented"
	return *new(HKTSB)
}

func Ap[
	HKTSA ~func(S) HKTA,
	HKTSB ~func(S) HKTB,
	HKTSAB ~func(S) HKTAB,
	HKTA,
	HKTB,
	HKTAB,

	S, A, B any,
](
	fmap func(func(pair.Pair[S, A]) pair.Pair[S, B]) func(HKTA) HKTB,
	fchain func(func(pair.Pair[S, func(A) B]) HKTB) func(HKTAB) HKTB,

	fa HKTSA,
) func(HKTSAB) HKTSB {
	_ = "STUB: not implemented"
	return nil
}

func FromF[
	HKTSA ~func(S) HKTA,
	HKTA,

	HKTFA,

	S, A any,
](
	fmap func(HKTFA, func(A) pair.Pair[S, A]) HKTA,
	ma HKTFA) HKTSA {
	_ = "STUB: not implemented"
	return *new(HKTSA)
}

func FromState[
	HKTSA ~func(S) HKTA,
	ST ~func(S) pair.Pair[S, A],
	HKTA,

	S, A any,
](
	fof func(pair.Pair[S, A]) HKTA,
	sa ST,
) HKTSA {
	_ = "STUB: not implemented"
	return *new(HKTSA)
}
