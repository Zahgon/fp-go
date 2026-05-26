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
	OP "github.com/IBM/fp-go/optics/optional"
	O "github.com/IBM/fp-go/option"
)

func setter[M ~map[K]V, K comparable, V any](key K) func(M, V) M {
	_ = "STUB: not implemented"
	return nil
}

func getter[M ~map[K]V, K comparable, V any](key K) func(M) O.Option[V] {
	_ = "STUB: not implemented"
	return nil

	// AtKey returns a Optional that gets and sets properties of a map
}

func AtKey[M ~map[K]V, K comparable, V any](key K) OP.Optional[M, V] {
	_ = "STUB: not implemented"
	return nil
}
