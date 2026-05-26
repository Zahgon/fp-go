// Copyright (c) 2024 - 2025 IBM Corp.
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

package iterresult

import (
	"github.com/IBM/fp-go/v2/internal/functor"
	"github.com/IBM/fp-go/v2/internal/monad"
	"github.com/IBM/fp-go/v2/internal/pointed"
)

// Pointed implements the pointed operations for [SeqResult]
func Pointed[A any]() pointed.Pointed[A, SeqResult[A]] { _ = "STUB: not implemented"; return nil }

// Functor implements the monadic operations for [SeqResult]
func Functor[A, B any]() functor.Functor[A, B, SeqResult[A], SeqResult[B]] {
	_ = "STUB: not implemented"
	return nil
}

// Monad implements the monadic operations for [SeqResult]
func Monad[A, B any]() monad.Monad[A, B, SeqResult[A], SeqResult[B], SeqResult[func(A) B]] {
	_ = "STUB: not implemented"
	return nil
}
