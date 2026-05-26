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

package statereaderioeither

import (
	ET "github.com/IBM/fp-go/either"
	IO "github.com/IBM/fp-go/io"
	IOE "github.com/IBM/fp-go/ioeither"
	RD "github.com/IBM/fp-go/reader"
	RE "github.com/IBM/fp-go/readereither"
	RIOE "github.com/IBM/fp-go/readerioeither"
	ST "github.com/IBM/fp-go/state"
)

func Left[S, R, A, E any](e E) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Right[S, R, E, A any](a A) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Of[S, R, E, A any](a A) StateReaderIOEither[S, R, E, A] { _ = "STUB: not implemented"; return nil }

func MonadMap[S, R, E, A, B any](fa StateReaderIOEither[S, R, E, A], f func(A) B) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[S, R, E, A, B any](f func(A) B) func(StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[S, R, E, A, B any](fa StateReaderIOEither[S, R, E, A], f func(A) StateReaderIOEither[S, R, E, B]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[S, R, E, A, B any](f func(A) StateReaderIOEither[S, R, E, B]) func(StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadAp[S, R, E, A, B any](fab StateReaderIOEither[S, R, E, func(A) B], fa StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[S, R, E, A, B any](fa StateReaderIOEither[S, R, E, A]) func(StateReaderIOEither[S, R, E, func(A) B]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func FromReaderIOEither[S, R, E, A any](fa RIOE.ReaderIOEither[R, E, A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromReaderEither[S, R, E, A any](fa RE.ReaderEither[R, E, A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromIOEither[S, R, E, A any](fa IOE.IOEither[E, A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromState[S, R, E, A any](sa ST.State[S, A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromIO[S, R, E, A any](fa IO.IO[A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromReader[S, R, E, A any](fa RD.Reader[R, A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromEither[S, R, E, A any](ma ET.Either[E, A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

// Combinators

func Local[S, R1, R2, E, A, B any](f func(R2) R1) func(StateReaderIOEither[S, R1, E, A]) StateReaderIOEither[S, R2, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func Asks[
	S, R, E, A any,
](f func(R) StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, A] {
	_ = "STUB: not implemented"
	return nil
}

func FromEitherK[S, R, E, A, B any](f func(A) ET.Either[E, B]) func(A) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func FromIOK[S, R, E, A, B any](f func(A) IO.IO[B]) func(A) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func FromIOEitherK[
	S, R, E, A, B any,
](f func(A) IOE.IOEither[E, B]) func(A) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func FromReaderIOEitherK[S, R, E, A, B any](f func(A) RIOE.ReaderIOEither[R, E, B]) func(A) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainReaderIOEitherK[S, R, E, A, B any](ma StateReaderIOEither[S, R, E, A], f func(A) RIOE.ReaderIOEither[R, E, B]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainReaderIOEitherK[S, R, E, A, B any](f func(A) RIOE.ReaderIOEither[R, E, B]) func(StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainIOEitherK[S, R, E, A, B any](ma StateReaderIOEither[S, R, E, A], f func(A) IOE.IOEither[E, B]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainIOEitherK[S, R, E, A, B any](f func(A) IOE.IOEither[E, B]) func(StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChainEitherK[S, R, E, A, B any](ma StateReaderIOEither[S, R, E, A], f func(A) ET.Either[E, B]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[S, R, E, A, B any](f func(A) ET.Either[E, B]) func(StateReaderIOEither[S, R, E, A]) StateReaderIOEither[S, R, E, B] {
	_ = "STUB: not implemented"
	return nil
}
