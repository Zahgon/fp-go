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

package di

import (
	DIE "github.com/IBM/fp-go/v2/di/erasure"
	"github.com/IBM/fp-go/v2/result"
)

var (
	toOptionAny   = toType[Option[any]]()
	toIOEitherAny = toType[IOResult[any]]()
	toIOOptionAny = toType[IOOption[any]]()
	toArrayAny    = toType[[]any]()
)

// asDependency converts a generic type to a [DIE.Dependency]
func asDependency[T DIE.Dependency](t T) DIE.Dependency {
	_ = "STUB: not implemented"

	// toType converts an any to a T
	return *new(DIE.Dependency)
}

func toType[T any]() result.Kleisli[any, T] { _ = "STUB: not implemented"; return nil }

// toOptionType converts an any to an Option[any] and then to an Option[T]
func toOptionType[T any](item result.Kleisli[any, T]) result.Kleisli[any, Option[T]] {
	_ = "STUB: not implemented"
	return nil
}

// toIOEitherType converts an any to an IOEither[error, any] and then to an IOEither[error, T]
func toIOEitherType[T any](item result.Kleisli[any, T]) result.Kleisli[any, IOResult[T]] {
	_ = "STUB: not implemented"
	return nil
}

// toIOOptionType converts an any to an IOOption[any] and then to an IOOption[T]
func toIOOptionType[T any](item result.Kleisli[any, T]) result.Kleisli[any, IOOption[T]] {
	_ = "STUB: not implemented"
	return nil
}

// toArrayType converts an any to a []T
func toArrayType[T any](item result.Kleisli[any, T]) result.Kleisli[any, []T] {
	_ = "STUB: not implemented"
	return nil
}
