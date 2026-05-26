// Copyright (c) 2024 IBM Corp.
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

package monad

import (
	"github.com/IBM/fp-go/internal/applicative"
	"github.com/IBM/fp-go/internal/apply"
	"github.com/IBM/fp-go/internal/chain"
	"github.com/IBM/fp-go/internal/functor"
	"github.com/IBM/fp-go/internal/pointed"
)

type Monad[A, B, HKTA, HKTB, HKTFAB any] interface {
	applicative.Applicative[A, B, HKTA, HKTB, HKTFAB]
	chain.Chainable[A, B, HKTA, HKTB, HKTFAB]
}

// ToFunctor converts from [Monad] to [functor.Functor]
func ToFunctor[A, B, HKTA, HKTB, HKTFAB any](ap Monad[A, B, HKTA, HKTB, HKTFAB]) functor.Functor[A, B, HKTA, HKTB] {
	_ = "STUB: not implemented"

	// ToApply converts from [Monad] to [apply.Apply]
	return nil
}

func ToApply[A, B, HKTA, HKTB, HKTFAB any](ap Monad[A, B, HKTA, HKTB, HKTFAB]) apply.Apply[A, B, HKTA, HKTB, HKTFAB] {
	_ = "STUB: not implemented"

	// ToPointed converts from [Monad] to [pointed.Pointed]
	return nil
}

func ToPointed[A, B, HKTA, HKTB, HKTFAB any](ap Monad[A, B, HKTA, HKTB, HKTFAB]) pointed.Pointed[A, HKTA] {
	_ = "STUB: not implemented"

	// ToApplicative converts from [Monad] to [applicative.Applicative]
	return nil
}

func ToApplicative[A, B, HKTA, HKTB, HKTFAB any](ap Monad[A, B, HKTA, HKTB, HKTFAB]) applicative.Applicative[A, B, HKTA, HKTB, HKTFAB] {
	_ = "STUB: not implemented"

	// ToChainable converts from [Monad] to [chain.Chainable]
	return nil
}

func ToChainable[A, B, HKTA, HKTB, HKTFAB any](ap Monad[A, B, HKTA, HKTB, HKTFAB]) chain.Chainable[A, B, HKTA, HKTB, HKTFAB] {
	_ = "STUB: not implemented"
	return nil
}
