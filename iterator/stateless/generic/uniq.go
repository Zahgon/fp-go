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

// addToMap makes a deep copy of a map and adds a value
func addToMap[A comparable](a A, m map[A]bool) map[A]bool { _ = "STUB: not implemented"; return nil }

func Uniq[AS ~func() O.Option[P.Pair[AS, A]], K comparable, A any](f func(A) K) func(as AS) AS {
	_ = "STUB: not implemented"
	return nil
}

func StrictUniq[AS ~func() O.Option[P.Pair[AS, A]], A comparable](as AS) AS {
	_ = "STUB: not implemented"
	return *new(AS)
}
