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
	T "github.com/IBM/fp-go/tuple"
)

// ZipWith applies a function to pairs of elements at the same index in two arrays, collecting the results in a new array. If one
// input array is short, excess elements of the longer array are discarded.
func ZipWith[AS ~[]A, BS ~[]B, CS ~[]C, FCT ~func(A, B) C, A, B, C any](fa AS, fb BS, f FCT) CS {
	_ = "STUB: not implemented"
	return *new(CS)
}

// Zip takes two arrays and returns an array of corresponding pairs. If one input array is short, excess elements of the
// longer array are discarded
func Zip[AS ~[]A, BS ~[]B, CS ~[]T.Tuple2[A, B], A, B any](fb BS) func(AS) CS {
	_ = "STUB: not implemented"
	return nil
}

// Unzip is the function is reverse of [Zip]. Takes an array of pairs and return two corresponding arrays
func Unzip[AS ~[]A, BS ~[]B, CS ~[]T.Tuple2[A, B], A, B any](cs CS) T.Tuple2[AS, BS] {
	_ = "STUB: not implemented"
	return nil
}
