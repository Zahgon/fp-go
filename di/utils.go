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

package di

import (
	DIE "github.com/IBM/fp-go/di/erasure"
	E "github.com/IBM/fp-go/either"
	IOE "github.com/IBM/fp-go/ioeither"
	IOO "github.com/IBM/fp-go/iooption"
	O "github.com/IBM/fp-go/option"
)

var (
	toOptionAny   = toType[O.Option[any]]()
	toIOEitherAny = toType[IOE.IOEither[error, any]]()
	toIOOptionAny = toType[IOO.IOOption[any]]()
	toArrayAny    = toType[[]any]()
)

// asDependency converts a generic type to a [DIE.Dependency]
func asDependency[T DIE.Dependency](t T) DIE.Dependency {
	_ = "STUB: not implemented"

	// toType converts an any to a T
	return *new(DIE.Dependency)
}

func toType[T any]() func(t any) E.Either[error, T] { _ = "STUB: not implemented"; return nil }

// toOptionType converts an any to an Option[any] and then to an Option[T]
func toOptionType[T any](item func(any) E.Either[error, T]) func(t any) E.Either[error, O.Option[T]] {
	_ = "STUB: not implemented"
	return nil
}

// toIOEitherType converts an any to an IOEither[error, any] and then to an IOEither[error, T]
func toIOEitherType[T any](item func(any) E.Either[error, T]) func(t any) E.Either[error, IOE.IOEither[error, T]] {
	_ = "STUB: not implemented"
	return nil
}

// toIOOptionType converts an any to an IOOption[any] and then to an IOOption[T]
func toIOOptionType[T any](item func(any) E.Either[error, T]) func(t any) E.Either[error, IOO.IOOption[T]] {
	_ = "STUB: not implemented"
	return nil
}

// toArrayType converts an any to a []T
func toArrayType[T any](item func(any) E.Either[error, T]) func(t any) E.Either[error, []T] {
	_ = "STUB: not implemented"
	return nil
}
