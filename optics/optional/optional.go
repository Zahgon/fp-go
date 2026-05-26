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

// Optional is an optic used to zoom inside a product. Unlike the `Lens`, the element that the `Optional` focuses
// on may not exist.
package optional

import (
	EM "github.com/IBM/fp-go/endomorphism"
	O "github.com/IBM/fp-go/option"
)

// Optional is an optional reference to a subpart of a data type
type Optional[S, A any] struct {
	GetOption func(s S) O.Option[A]
	Set       func(a A) EM.Endomorphism[S]
}

// setCopy wraps a setter for a pointer into a setter that first creates a copy before
// modifying that copy
func setCopy[SET ~func(*S, A) *S, S, A any](setter SET) func(s *S, a A) *S {
	_ = "STUB: not implemented"
	return nil
}

// MakeOptional creates an Optional based on a getter and a setter function. Make sure that the setter creates a (shallow) copy of the
// data. This happens automatically if the data is passed by value. For pointers consider to use `MakeOptionalRef`
// and for other kinds of data structures that are copied by reference make sure the setter creates the copy.
func MakeOptional[S, A any](get func(S) O.Option[A], set func(S, A) S) Optional[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// MakeOptionalRef creates an Optional based on a getter and a setter function. The setter passed in does not have to create a shallow
// copy, the implementation wraps the setter into one that copies the pointer before modifying it
func MakeOptionalRef[S, A any](get func(*S) O.Option[A], set func(*S, A) *S) Optional[*S, A] {
	_ = "STUB: not implemented"
	return nil
}

// Id returns am optional implementing the identity operation
func id[S any](creator func(get func(S) O.Option[S], set func(S, S) S) Optional[S, S]) Optional[S, S] {
	_ = "STUB: not implemented"
	return nil
}

// Id returns am optional implementing the identity operation
func Id[S any]() Optional[S, S] { _ = "STUB: not implemented"; return nil }

// Id returns am optional implementing the identity operation
func IdRef[S any]() Optional[*S, *S] { _ = "STUB: not implemented"; return nil }

func optionalModifyOption[S, A any](f func(A) A, optional Optional[S, A], s S) O.Option[S] {
	_ = "STUB: not implemented"
	return nil
}

func optionalModify[S, A any](f func(A) A, optional Optional[S, A], s S) S {
	_ = "STUB: not implemented"
	return *new(S)
}

// Compose combines two Optional and allows to narrow down the focus to a sub-Optional
func compose[S, A, B any](creator func(get func(S) O.Option[B], set func(S, B) S) Optional[S, B], ab Optional[A, B]) func(Optional[S, A]) Optional[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// Compose combines two Optional and allows to narrow down the focus to a sub-Optional
func Compose[S, A, B any](ab Optional[A, B]) func(Optional[S, A]) Optional[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// ComposeRef combines two Optional and allows to narrow down the focus to a sub-Optional
func ComposeRef[S, A, B any](ab Optional[A, B]) func(Optional[*S, A]) Optional[*S, B] {
	_ = "STUB: not implemented"
	return nil
}

// fromPredicate implements the function generically for both the ref and the direct case
func fromPredicate[S, A any](creator func(get func(S) O.Option[A], set func(S, A) S) Optional[S, A], pred func(A) bool) func(func(S) A, func(S, A) S) Optional[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromPredicate creates an optional from getter and setter functions. It checks
// for optional values and the correct update procedure
func FromPredicate[S, A any](pred func(A) bool) func(func(S) A, func(S, A) S) Optional[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromPredicate creates an optional from getter and setter functions. It checks
// for optional values and the correct update procedure
func FromPredicateRef[S, A any](pred func(A) bool) func(func(*S) A, func(*S, A) *S) Optional[*S, A] {
	_ = "STUB: not implemented"
	return nil
}

func imap[S, A, B any](sa Optional[S, A], ab func(A) B, ba func(B) A) Optional[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// IMap implements a bidirectional mapping of the transform
func IMap[S, A, B any](ab func(A) B, ba func(B) A) func(Optional[S, A]) Optional[S, B] {
	_ = "STUB: not implemented"
	return nil
}

func ModifyOption[S, A any](f func(A) A) func(Optional[S, A]) func(S) O.Option[S] {
	_ = "STUB: not implemented"
	return nil
}

func SetOption[S, A any](a A) func(Optional[S, A]) func(S) O.Option[S] {
	_ = "STUB: not implemented"
	return nil
}

func ichain[S, A, B any](sa Optional[S, A], ab func(A) O.Option[B], ba func(B) O.Option[A]) Optional[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// IChain implements a bidirectional mapping of the transform if the transform can produce optionals (e.g. in case of type mappings)
func IChain[S, A, B any](ab func(A) O.Option[B], ba func(B) O.Option[A]) func(Optional[S, A]) Optional[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// IChainAny implements a bidirectional mapping to and from any
func IChainAny[S, A any]() func(Optional[S, any]) Optional[S, A] {
	_ = "STUB: not implemented"
	return nil
}
