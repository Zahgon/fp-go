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

import (
	O "github.com/IBM/fp-go/option"
	P "github.com/IBM/fp-go/pair"
)

// ZipWith applies a function to pairs of elements at the same index in two iterators, collecting the results in a new iterator. If one
// input iterator is short, excess elements of the longer iterator are discarded.
func ZipWith[AS ~func() O.Option[P.Pair[AS, A]], BS ~func() O.Option[P.Pair[BS, B]], CS ~func() O.Option[P.Pair[CS, C]], FCT ~func(A, B) C, A, B, C any](fa AS, fb BS, f FCT) CS {
	_ = "STUB: not implemented"
	// pre-declare to avoid cyclic reference
	return *new(CS)
}

// combine

// trigger the recursion

// Zip takes two iterators and returns an iterators of corresponding pairs. If one input iterators is short, excess elements of the
// longer iterator are discarded
func Zip[AS ~func() O.Option[P.Pair[AS, A]], BS ~func() O.Option[P.Pair[BS, B]], CS ~func() O.Option[P.Pair[CS, P.Pair[A, B]]], A, B any](fb BS) func(AS) CS {
	_ = "STUB: not implemented"
	return nil
}
