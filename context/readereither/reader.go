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

package readereither

import (
	"context"

	ET "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
)

func MakeReaderEither[A any](f func(context.Context) ET.Either[error, A]) ReaderEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func FromEither[A any](e ET.Either[error, A]) ReaderEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func Left[A any](l error) ReaderEither[A] { _ = "STUB: not implemented"; return nil }

func Right[A any](r A) ReaderEither[A] { _ = "STUB: not implemented"; return nil }

func MonadMap[A, B any](fa ReaderEither[A], f func(A) B) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Map[A, B any](f func(A) B) func(ReaderEither[A]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadChain[A, B any](ma ReaderEither[A], f func(A) ReaderEither[B]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A, B any](f func(A) ReaderEither[B]) func(ReaderEither[A]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Of[A any](a A) ReaderEither[A] { _ = "STUB: not implemented"; return nil }

func MonadAp[A, B any](fab ReaderEither[func(A) B], fa ReaderEither[A]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Ap[A, B any](fa ReaderEither[A]) func(ReaderEither[func(A) B]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func FromPredicate[A any](pred func(A) bool, onFalse func(A) error) func(A) ReaderEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func OrElse[A any](onLeft func(error) ReaderEither[A]) func(ReaderEither[A]) ReaderEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func Ask() ReaderEither[context.Context] { _ = "STUB: not implemented"; return nil }

func MonadChainEitherK[A, B any](ma ReaderEither[A], f func(A) ET.Either[error, B]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainEitherK[A, B any](f func(A) ET.Either[error, B]) func(ma ReaderEither[A]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func ChainOptionK[A, B any](onNone func() error) func(func(A) O.Option[B]) func(ReaderEither[A]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func MonadFlap[B, A any](fab ReaderEither[func(A) B], a A) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}

func Flap[B, A any](a A) func(ReaderEither[func(A) B]) ReaderEither[B] {
	_ = "STUB: not implemented"
	return nil
}
