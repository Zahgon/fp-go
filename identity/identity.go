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

package identity

func MonadAp[B, A any](fab func(A) B, fa A) B { _ = "STUB: not implemented"; return *new(B) }

func Ap[B, A any](fa A) func(func(A) B) B { _ = "STUB: not implemented"; return nil }

func MonadMap[A, B any](fa A, f func(A) B) B { _ = "STUB: not implemented"; return *new(B) }

func Map[A, B any](f func(A) B) func(A) B { _ = "STUB: not implemented"; return nil }

func MonadMapTo[A, B any](_ A, b B) B { _ = "STUB: not implemented"; return *new(B) }

func MapTo[A, B any](b B) func(A) B { _ = "STUB: not implemented"; return nil }

func Of[A any](a A) A { _ = "STUB: not implemented"; return *new(A) }

func MonadChain[A, B any](ma A, f func(A) B) B { _ = "STUB: not implemented"; return *new(B) }

func Chain[A, B any](f func(A) B) func(A) B { _ = "STUB: not implemented"; return nil }

func MonadChainFirst[A, B any](fa A, f func(A) B) A { _ = "STUB: not implemented"; return *new(A) }

func ChainFirst[A, B any](f func(A) B) func(A) A { _ = "STUB: not implemented"; return nil }

func MonadFlap[B, A any](fab func(A) B, a A) B { _ = "STUB: not implemented"; return *new(B) }

func Flap[B, A any](a A) func(func(A) B) B { _ = "STUB: not implemented"; return nil }
