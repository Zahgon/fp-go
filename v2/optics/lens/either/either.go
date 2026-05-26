// Copyright (c) 2023 - 2025 IBM Corp.
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

package either

import (
	L "github.com/IBM/fp-go/v2/optics/lens"
	T "github.com/IBM/fp-go/v2/optics/traversal/either"
)

func AsTraversal[E, S, A any]() func(L.Lens[S, A]) T.Traversal[E, S, A] {
	_ = "STUB: not implemented"
	return nil
}
