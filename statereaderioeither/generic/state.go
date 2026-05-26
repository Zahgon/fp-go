// Copyright (c) 2024 IBM Corp.
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
	P "github.com/IBM/fp-go/pair"
)

func Left[
	SRIOEA ~func(S) RIOEA,
	RIOEA ~func(R) IOEA,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	S, R, E, A any,
](e E) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func Right[
	SRIOEA ~func(S) RIOEA,
	RIOEA ~func(R) IOEA,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	S, R, E, A any,
](a A) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func Of[
	SRIOEA ~func(S) RIOEA,
	RIOEA ~func(R) IOEA,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	S, R, E, A any,
](a A) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func MonadMap[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](fa SRIOEA, f func(A) B) SRIOEB {
	_ = "STUB: not implemented"
	return *new(SRIOEB)
}

func Map[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](f func(A) B) func(SRIOEA) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](fa SRIOEA, f func(A) SRIOEB) SRIOEB {
	_ = "STUB: not implemented"
	return *new(SRIOEB)
}

func Chain[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](f func(A) SRIOEB) func(SRIOEA) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	SRIOEAB ~func(S) RIOEAB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	RIOEAB ~func(R) IOEAB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	IOEAB ~func() ET.Either[E, P.Pair[func(A) B, S]],
	S, R, E, A, B any,
](fab SRIOEAB, fa SRIOEA) SRIOEB {
	_ = "STUB: not implemented"
	return *new(SRIOEB)
}

func Ap[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	SRIOEAB ~func(S) RIOEAB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	RIOEAB ~func(R) IOEAB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	IOEAB ~func() ET.Either[E, P.Pair[func(A) B, S]],
	S, R, E, A, B any,
](fa SRIOEA) func(SRIOEAB) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

// Conversions

func FromReaderIOEither[
	SRIOEA ~func(S) RIOEA,
	RIOEA_IN ~func(R) IOEA_IN,

	RIOEA ~func(R) IOEA,

	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEA_IN ~func() ET.Either[E, A],

	S, R, E, A any,
](fa RIOEA_IN) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func FromReaderEither[
	SRIOEA ~func(S) RIOEA,
	RIOEA_IN ~func(R) IOEA_IN,
	RIOEA ~func(R) IOEA,

	REA_IN ~func(R) ET.Either[E, A],

	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEA_IN ~func() ET.Either[E, A],

	S, R, E, A any,
](fa REA_IN) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func FromIOEither[
	SRIOEA ~func(S) RIOEA,
	RIOEA_IN ~func(R) IOEA_IN,
	IOEA_IN ~func() ET.Either[E, A],
	RIOEA ~func(R) IOEA,

	IOEA ~func() ET.Either[E, P.Pair[A, S]],

	S, R, E, A any,
](fa IOEA_IN) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func FromIO[
	SRIOEA ~func(S) RIOEA,
	RIOEA_IN ~func(R) IOEA_IN,

	IO_IN ~func() A,

	RIOEA ~func(R) IOEA,

	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEA_IN ~func() ET.Either[E, A],

	S, R, E, A any,
](fa IO_IN) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func FromReader[
	SRIOEA ~func(S) RIOEA,
	RIOEA_IN ~func(R) IOEA_IN,

	R_IN ~func(R) A,

	RIOEA ~func(R) IOEA,

	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEA_IN ~func() ET.Either[E, A],

	S, R, E, A any,
](fa R_IN) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func FromEither[
	SRIOEA ~func(S) RIOEA,

	RIOEA ~func(R) IOEA,

	IOEA ~func() ET.Either[E, P.Pair[A, S]],

	S, R, E, A any,
](ma ET.Either[E, A]) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func FromState[
	SRIOEA ~func(S) RIOEA,
	STATE ~func(S) P.Pair[A, S],
	RIOEA ~func(R) IOEA,

	IOEA ~func() ET.Either[E, P.Pair[A, S]],

	S, R, E, A any,
](fa STATE) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

// Combinators

func Local[
	SR1IOEA ~func(S) R1IOEA,
	SR2IOEA ~func(S) R2IOEA,
	R1IOEA ~func(R1) IOEA,
	R2IOEA ~func(R2) IOEA,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	S, R1, R2, E, A any,
](f func(R2) R1) func(SR1IOEA) SR2IOEA {
	_ = "STUB: not implemented"
	return nil
}

func Asks[
	SRIOEA ~func(S) RIOEA,
	RIOEA ~func(R) IOEA,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	S, R, E, A any,
](f func(R) SRIOEA) SRIOEA {
	_ = "STUB: not implemented"
	return *new(SRIOEA)
}

func FromIOEitherK[
	SRIOEB ~func(S) RIOEB,
	RIOEB_IN ~func(R) IOEB_IN,
	IOEB_IN ~func() ET.Either[E, B],
	RIOEB ~func(R) IOEB,
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](f func(A) IOEB_IN) func(A) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func FromEitherK[
	SRIOEB ~func(S) RIOEB,
	RIOEB ~func(R) IOEB,
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](f func(A) ET.Either[E, B]) func(A) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func FromIOK[
	SRIOEB ~func(S) RIOEB,
	RIOEB_IN ~func(R) IOEB_IN,

	IOB_IN ~func() B,

	RIOEB ~func(R) IOEB,

	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	IOEB_IN ~func() ET.Either[E, B],

	S, R, E, A, B any,
](f func(A) IOB_IN) func(A) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func FromReaderIOEitherK[
	SRIOEB ~func(S) RIOEB,
	RIOEB_IN ~func(R) IOEB_IN,
	IOEB_IN ~func() ET.Either[E, B],
	RIOEB ~func(R) IOEB,
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](f func(A) RIOEB_IN) func(A) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainReaderIOEitherK[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEB_IN ~func(R) IOEB_IN,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	IOEB_IN ~func() ET.Either[E, B],
	S, R, E, A, B any,
](ma SRIOEA, f func(A) RIOEB_IN) SRIOEB {
	_ = "STUB: not implemented"
	return *new(SRIOEB)
}

func ChainReaderIOEitherK[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEB_IN ~func(R) IOEB_IN,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	IOEB_IN ~func() ET.Either[E, B],
	S, R, E, A, B any,
](f func(A) RIOEB_IN) func(SRIOEA) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOEitherK[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEB_IN ~func(R) IOEB_IN,
	IOEB_IN ~func() ET.Either[E, B],
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](ma SRIOEA, f func(A) IOEB_IN) SRIOEB {
	_ = "STUB: not implemented"
	return *new(SRIOEB)
}

func ChainIOEitherK[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEB_IN ~func(R) IOEB_IN,
	IOEB_IN ~func() ET.Either[E, B],
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](f func(A) IOEB_IN) func(SRIOEA) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](ma SRIOEA, f func(A) ET.Either[E, B]) SRIOEB {
	_ = "STUB: not implemented"
	return *new(SRIOEB)
}

func ChainEitherK[
	SRIOEA ~func(S) RIOEA,
	SRIOEB ~func(S) RIOEB,
	RIOEA ~func(R) IOEA,
	RIOEB ~func(R) IOEB,
	IOEA ~func() ET.Either[E, P.Pair[A, S]],
	IOEB ~func() ET.Either[E, P.Pair[B, S]],
	S, R, E, A, B any,
](f func(A) ET.Either[E, B]) func(SRIOEA) SRIOEB {
	_ = "STUB: not implemented"
	return nil
}
