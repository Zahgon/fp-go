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

// Iso is an optic which converts elements of type `S` into elements of type `A` without loss.
package iso

import (
	EM "github.com/IBM/fp-go/endomorphism"
)

type Iso[S, A any] struct {
	Get        func(s S) A
	ReverseGet func(a A) S
}

func MakeIso[S, A any](get func(S) A, reverse func(A) S) Iso[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// Id returns an iso implementing the identity operation
func Id[S any]() Iso[S, S] { _ = "STUB: not implemented"; return nil }

// Compose combines an ISO with another ISO
func Compose[S, A, B any](ab Iso[A, B]) func(Iso[S, A]) Iso[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// Reverse changes the order of parameters for an iso
func Reverse[S, A any](sa Iso[S, A]) Iso[A, S] { _ = "STUB: not implemented"; return nil }

func modify[FCT ~func(A) A, S, A any](f FCT, sa Iso[S, A], s S) S {
	_ = "STUB: not implemented"
	return *new(S)
}

// Modify applies a transformation
func Modify[S any, FCT ~func(A) A, A any](f FCT) func(Iso[S, A]) EM.Endomorphism[S] {
	_ = "STUB: not implemented"
	return nil
}

// Wrap wraps the value
func Unwrap[A, S any](s S) func(Iso[S, A]) A { _ = "STUB: not implemented"; return nil }

// Unwrap unwraps the value
func Wrap[S, A any](a A) func(Iso[S, A]) S { _ = "STUB: not implemented"; return nil }

// From wraps the value
func To[A, S any](s S) func(Iso[S, A]) A { _ = "STUB: not implemented"; return nil }

// To unwraps the value
func From[S, A any](a A) func(Iso[S, A]) S { _ = "STUB: not implemented"; return nil }

func imap[S, A, B any](sa Iso[S, A], ab func(A) B, ba func(B) A) Iso[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// IMap implements a bidirectional mapping of the transform
func IMap[S, A, B any](ab func(A) B, ba func(B) A) func(Iso[S, A]) Iso[S, B] {
	_ = "STUB: not implemented"
	return nil
}
