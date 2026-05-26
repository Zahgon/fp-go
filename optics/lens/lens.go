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

// Lens is an optic used to zoom inside a product.
package lens

import (
	EM "github.com/IBM/fp-go/endomorphism"
	O "github.com/IBM/fp-go/option"
)

type (
	// Lens is a reference to a subpart of a data type
	Lens[S, A any] struct {
		Get func(s S) A
		Set func(a A) EM.Endomorphism[S]
	}
)

// setCopy wraps a setter for a pointer into a setter that first creates a copy before
// modifying that copy
func setCopy[SET ~func(*S, A) *S, S, A any](setter SET) func(s *S, a A) *S {
	_ = "STUB: not implemented"
	return nil
}

// setCopyCurried wraps a setter for a pointer into a setter that first creates a copy before
// modifying that copy
func setCopyCurried[SET ~func(A) EM.Endomorphism[*S], S, A any](setter SET) func(a A) EM.Endomorphism[*S] {
	_ = "STUB: not implemented"
	return nil
}

// MakeLens creates a [Lens] based on a getter and a setter function. Make sure that the setter creates a (shallow) copy of the
// data. This happens automatically if the data is passed by value. For pointers consider to use `MakeLensRef`
// and for other kinds of data structures that are copied by reference make sure the setter creates the copy.
func MakeLens[GET ~func(S) A, SET ~func(S, A) S, S, A any](get GET, set SET) Lens[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// MakeLensCurried creates a [Lens] based on a getter and a setter function. Make sure that the setter creates a (shallow) copy of the
// data. This happens automatically if the data is passed by value. For pointers consider to use `MakeLensRef`
// and for other kinds of data structures that are copied by reference make sure the setter creates the copy.
func MakeLensCurried[GET ~func(S) A, SET ~func(A) EM.Endomorphism[S], S, A any](get GET, set SET) Lens[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// MakeLensRef creates a [Lens] based on a getter and a setter function. The setter passed in does not have to create a shallow
// copy, the implementation wraps the setter into one that copies the pointer before modifying it
//
// Such a [Lens] assumes that property A of S always exists
func MakeLensRef[GET ~func(*S) A, SET func(*S, A) *S, S, A any](get GET, set SET) Lens[*S, A] {
	_ = "STUB: not implemented"
	return nil
}

// MakeLensRefCurried creates a [Lens] based on a getter and a setter function. The setter passed in does not have to create a shallow
// copy, the implementation wraps the setter into one that copies the pointer before modifying it
//
// Such a [Lens] assumes that property A of S always exists
func MakeLensRefCurried[S, A any](get func(*S) A, set func(A) EM.Endomorphism[*S]) Lens[*S, A] {
	_ = "STUB: not implemented"
	return nil
}

// id returns a [Lens] implementing the identity operation
func id[GET ~func(S) S, SET ~func(S, S) S, S any](creator func(get GET, set SET) Lens[S, S]) Lens[S, S] {
	_ = "STUB: not implemented"
	return nil
}

// Id returns a [Lens] implementing the identity operation
func Id[S any]() Lens[S, S] { _ = "STUB: not implemented"; return nil }

// IdRef returns a [Lens] implementing the identity operation
func IdRef[S any]() Lens[*S, *S] { _ = "STUB: not implemented"; return nil }

// Compose combines two lenses and allows to narrow down the focus to a sub-lens
func compose[GET ~func(S) B, SET ~func(S, B) S, S, A, B any](creator func(get GET, set SET) Lens[S, B], ab Lens[A, B]) func(Lens[S, A]) Lens[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// Compose combines two lenses and allows to narrow down the focus to a sub-lens
func Compose[S, A, B any](ab Lens[A, B]) func(Lens[S, A]) Lens[S, B] {
	_ = "STUB: not implemented"
	return nil
}

// ComposeOption combines a `Lens` that returns an optional value with a `Lens` that returns a definite value
// the getter returns an `Option[B]` because the container `A` could already be an option
// if the setter is invoked with `Some[B]` then the value of `B` will be set, potentially on a default value of `A` if `A` did not exist
// if the setter is invoked with `None[B]` then the container `A` is reset to `None[A]` because this is the only way to remove `B`
func ComposeOption[S, B, A any](defaultA A) func(ab Lens[A, B]) func(Lens[S, O.Option[A]]) Lens[S, O.Option[B]] {
	_ = "STUB: not implemented"
	return nil
}

// set A on S

// remove A from S

// ComposeOptions combines a `Lens` that returns an optional value with a `Lens` that returns another optional value
// the getter returns `None[B]` if either `A` or `B` is `None`
// if the setter is called with `Some[B]` and `A` exists, 'A' is updated with `B`
// if the setter is called with `Some[B]` and `A` does not exist, the default of 'A' is updated with `B`
// if the setter is called with `None[B]` and `A` does not exist this is the identity operation on 'S'
// if the setter is called with `None[B]` and `A` does exist, 'B' is removed from 'A'
func ComposeOptions[S, B, A any](defaultA A) func(ab Lens[A, O.Option[B]]) func(Lens[S, O.Option[A]]) Lens[S, O.Option[B]] {
	_ = "STUB: not implemented"
	return nil
}

// sets an A onto S

// sets a B onto an A

// Compose combines two lenses and allows to narrow down the focus to a sub-lens
func ComposeRef[S, A, B any](ab Lens[A, B]) func(Lens[*S, A]) Lens[*S, B] {
	_ = "STUB: not implemented"
	return nil
}

func modify[FCT ~func(A) A, S, A any](f FCT, sa Lens[S, A], s S) S {
	_ = "STUB: not implemented"
	return *new(S)
}

// Modify changes a property of a [Lens] by invoking a transformation function
// if the transformed property has not changes, the method returns the original state
func Modify[S any, FCT ~func(A) A, A any](f FCT) func(Lens[S, A]) EM.Endomorphism[S] {
	_ = "STUB: not implemented"
	return nil
}

func IMap[E any, AB ~func(A) B, BA ~func(B) A, A, B any](ab AB, ba BA) func(Lens[E, A]) Lens[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// fromPredicate returns a `Lens` for a property accessibly as a getter and setter that can be optional
// if the optional value is set then the nil value will be set instead
func fromPredicate[GET ~func(S) O.Option[A], SET ~func(S, O.Option[A]) S, S, A any](creator func(get GET, set SET) Lens[S, O.Option[A]], pred func(A) bool, nilValue A) func(sa Lens[S, A]) Lens[S, O.Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// FromPredicate returns a `Lens` for a property accessibly as a getter and setter that can be optional
// if the optional value is set then the nil value will be set instead
func FromPredicate[S, A any](pred func(A) bool, nilValue A) func(sa Lens[S, A]) Lens[S, O.Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// FromPredicateRef returns a `Lens` for a property accessibly as a getter and setter that can be optional
// if the optional value is set then the nil value will be set instead
func FromPredicateRef[S, A any](pred func(A) bool, nilValue A) func(sa Lens[*S, A]) Lens[*S, O.Option[A]] {
	_ = "STUB: not implemented"
	return nil
}

// FromPredicate returns a `Lens` for a property accessibly as a getter and setter that can be optional
// if the optional value is set then the `nil` value will be set instead
func FromNillable[S, A any](sa Lens[S, *A]) Lens[S, O.Option[*A]] {
	_ = "STUB: not implemented"
	return nil
}

// FromNillableRef returns a `Lens` for a property accessibly as a getter and setter that can be optional
// if the optional value is set then the `nil` value will be set instead
func FromNillableRef[S, A any](sa Lens[*S, *A]) Lens[*S, O.Option[*A]] {
	_ = "STUB: not implemented"
	return nil
}

// fromNullableProp returns a `Lens` from a property that may be optional. The getter returns a default value for these items
func fromNullableProp[GET ~func(S) A, SET ~func(S, A) S, S, A any](creator func(get GET, set SET) Lens[S, A], isNullable func(A) O.Option[A], defaultValue A) func(sa Lens[S, A]) Lens[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromNullableProp returns a `Lens` from a property that may be optional. The getter returns a default value for these items
func FromNullableProp[S, A any](isNullable func(A) O.Option[A], defaultValue A) func(sa Lens[S, A]) Lens[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromNullablePropRef returns a `Lens` from a property that may be optional. The getter returns a default value for these items
func FromNullablePropRef[S, A any](isNullable func(A) O.Option[A], defaultValue A) func(sa Lens[*S, A]) Lens[*S, A] {
	_ = "STUB: not implemented"
	return nil
}

// fromFromOption returns a `Lens` from an option property. The getter returns a default value the setter will always set the some option
func fromOption[GET ~func(S) A, SET ~func(S, A) S, S, A any](creator func(get GET, set SET) Lens[S, A], defaultValue A) func(sa Lens[S, O.Option[A]]) Lens[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromFromOption returns a `Lens` from an option property. The getter returns a default value the setter will always set the some option
func FromOption[S, A any](defaultValue A) func(sa Lens[S, O.Option[A]]) Lens[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// FromFromOptionRef returns a `Lens` from an option property. The getter returns a default value the setter will always set the some option
func FromOptionRef[S, A any](defaultValue A) func(sa Lens[*S, O.Option[A]]) Lens[*S, A] {
	_ = "STUB: not implemented"
	return nil
}
