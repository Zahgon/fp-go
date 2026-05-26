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
	IO "github.com/IBM/fp-go/io"
	IOE "github.com/IBM/fp-go/ioeither"
	IOO "github.com/IBM/fp-go/iooption"
	O "github.com/IBM/fp-go/option"
)

// Dependency describes the relationship to a service, that has a type and
// a behaviour such as required, option or lazy
type Dependency[T any] interface {
	DIE.Dependency
	// Unerase converts a value with erased type signature into a strongly typed value
	Unerase(val any) E.Either[error, T]
}

// InjectionToken uniquely identifies a dependency by giving it an Id, Type and name
type InjectionToken[T any] interface {
	Dependency[T]
	// Identity idenifies this dependency as a mandatory, required dependency, it will be resolved eagerly and injected as `T`.
	// If the dependency cannot be resolved, the resolution process fails
	Identity() Dependency[T]
	// Option identifies this dependency as optional, it will be resolved eagerly and injected as [O.Option[T]].
	// If the dependency cannot be resolved, the resolution process continues and the dependency is represented as [O.None[T]]
	Option() Dependency[O.Option[T]]
	// IOEither identifies this dependency as mandatory but it will be resolved lazily as a [IOE.IOEither[error, T]]. This
	// value is memoized to make sure the dependency is a singleton.
	// If the dependency cannot be resolved, the resolution process fails
	IOEither() Dependency[IOE.IOEither[error, T]]
	// IOOption identifies this dependency as optional but it will be resolved lazily as a [IOO.IOOption[T]]. This
	// value is memoized to make sure the dependency is a singleton.
	// If the dependency cannot be resolved, the resolution process continues and the dependency is represented as the none value.
	IOOption() Dependency[IOO.IOOption[T]]
}

// MultiInjectionToken uniquely identifies a dependency by giving it an Id, Type and name that can have multiple implementations.
// Implementations are provided via the [MultiInjectionToken.Item] injection token.
type MultiInjectionToken[T any] interface {
	// Container returns the injection token used to request an array of all provided items
	Container() InjectionToken[[]T]
	// Item returns the injection token used to provide an item
	Item() InjectionToken[T]
}

// makeID creates a generator of unique string IDs
func makeID() IO.IO[string] { _ = "STUB: not implemented"; return nil }

// genID is the common generator of unique string IDs
var genID = makeID()

type tokenBase struct {
	name            string
	id              string
	flag            int
	providerFactory O.Option[DIE.ProviderFactory]
}

type token[T any] struct {
	base   *tokenBase
	toType func(val any) E.Either[error, T]
}

func (t *token[T]) Id() string { _ = "STUB: not implemented"; return "" }

func (t *token[T]) Flag() int { _ = "STUB: not implemented"; return 0 }

func (t *token[T]) String() string { _ = "STUB: not implemented"; return "" }

func (t *token[T]) Unerase(val any) E.Either[error, T] { _ = "STUB: not implemented"; return nil }

func (t *token[T]) ProviderFactory() O.Option[DIE.ProviderFactory] {
	_ = "STUB: not implemented"
	return nil
}

func makeTokenBase(name string, id string, typ int, providerFactory O.Option[DIE.ProviderFactory]) *tokenBase {
	_ = "STUB: not implemented"
	return nil
}

func makeToken[T any](name string, id string, typ int, unerase func(val any) E.Either[error, T], providerFactory O.Option[DIE.ProviderFactory]) Dependency[T] {
	_ = "STUB: not implemented"
	return nil
}

type injectionToken[T any] struct {
	token[T]
	option   Dependency[O.Option[T]]
	ioeither Dependency[IOE.IOEither[error, T]]
	iooption Dependency[IOO.IOOption[T]]
}

type multiInjectionToken[T any] struct {
	container *injectionToken[[]T]
	item      *injectionToken[T]
}

func (i *injectionToken[T]) Identity() Dependency[T] { _ = "STUB: not implemented"; return nil }

func (i *injectionToken[T]) Option() Dependency[O.Option[T]] { _ = "STUB: not implemented"; return nil }

func (i *injectionToken[T]) IOEither() Dependency[IOE.IOEither[error, T]] {
	_ = "STUB: not implemented"
	return nil
}

func (i *injectionToken[T]) IOOption() Dependency[IOO.IOOption[T]] {
	_ = "STUB: not implemented"
	return nil
}

func (i *injectionToken[T]) ProviderFactory() O.Option[DIE.ProviderFactory] {
	_ = "STUB: not implemented"
	return nil
}

func (m *multiInjectionToken[T]) Container() InjectionToken[[]T] {
	_ = "STUB: not implemented"
	return nil
}

func (m *multiInjectionToken[T]) Item() InjectionToken[T] {
	_ = "STUB: not implemented"

	// makeToken create a unique [InjectionToken] for a specific type
	return nil
}

func makeInjectionToken[T any](name string, providerFactory O.Option[DIE.ProviderFactory]) InjectionToken[T] {
	_ = "STUB: not implemented"
	return nil
}

// MakeToken create a unique [InjectionToken] for a specific type
func MakeToken[T any](name string) InjectionToken[T] { _ = "STUB: not implemented"; return nil }

// MakeToken create a unique [InjectionToken] for a specific type
func MakeTokenWithDefault[T any](name string, providerFactory DIE.ProviderFactory) InjectionToken[T] {
	_ = "STUB: not implemented"
	return nil
}

// MakeMultiToken creates a [MultiInjectionToken]
func MakeMultiToken[T any](name string) MultiInjectionToken[T] {
	_ = "STUB: not implemented"
	return nil
}

// empty factory

// container

// item

// returns the token
