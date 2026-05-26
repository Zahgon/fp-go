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

func MonadAp[GAB ~func(A) B, B, A any](fab GAB, fa A) B { _ = "STUB: not implemented"; return *new(B) }

func Ap[GAB ~func(A) B, B, A any](fa A) func(GAB) B { _ = "STUB: not implemented"; return nil }

func MonadMap[GAB ~func(A) B, A, B any](fa A, f GAB) B { _ = "STUB: not implemented"; return *new(B) }

func Map[GAB ~func(A) B, A, B any](f GAB) func(A) B { _ = "STUB: not implemented"; return nil }

func MonadChain[GAB ~func(A) B, A, B any](ma A, f GAB) B { _ = "STUB: not implemented"; return *new(B) }

func Chain[GAB ~func(A) B, A, B any](f GAB) func(A) B { _ = "STUB: not implemented"; return nil }

func MonadChainFirst[GAB ~func(A) B, A, B any](fa A, f GAB) A {
	_ = "STUB: not implemented"
	return *new(A)
}

func ChainFirst[GAB ~func(A) B, A, B any](f GAB) func(A) A { _ = "STUB: not implemented"; return nil }

func MonadFlap[GAB ~func(A) B, A, B any](fab GAB, a A) B { _ = "STUB: not implemented"; return *new(B) }

func Flap[GAB ~func(A) B, B, A any](a A) func(GAB) B { _ = "STUB: not implemented"; return nil }
