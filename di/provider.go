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
)

func lookupAt[T any](idx int, token Dependency[T]) func(params []any) E.Either[error, T] {
	_ = "STUB: not implemented"
	return nil
}

func eraseTuple[A, R any](f func(A) IOE.IOEither[error, R]) func(E.Either[error, A]) IOE.IOEither[error, any] {
	_ = "STUB: not implemented"
	return nil
}

func eraseProviderFactory0[R any](f IOE.IOEither[error, R]) func(params ...any) IOE.IOEither[error, any] {
	_ = "STUB: not implemented"
	return nil
}

func MakeProviderFactory0[R any](
	fct IOE.IOEither[error, R],
) DIE.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(DIE.ProviderFactory)
}

// MakeTokenWithDefault0 creates a unique [InjectionToken] for a specific type with an attached default [DIE.Provider]
func MakeTokenWithDefault0[R any](name string, fct IOE.IOEither[error, R]) InjectionToken[R] {
	_ = "STUB: not implemented"
	return nil
}

func MakeProvider0[R any](
	token InjectionToken[R],
	fct IOE.IOEither[error, R],
) DIE.Provider {
	_ = "STUB: not implemented"
	return *new(DIE.Provider)
}

// ConstProvider simple implementation for a provider with a constant value
func ConstProvider[R any](token InjectionToken[R], value R) DIE.Provider {
	_ = "STUB: not implemented"
	return *new(DIE.Provider)
}
