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

package lens

import (
	L "github.com/IBM/fp-go/v2/optics/lens"
	OPT "github.com/IBM/fp-go/v2/optics/optional"
)

// Compose composes a lens with an optional
func Compose[S, A, B any](ab L.Lens[A, B]) func(sa OPT.Optional[S, A]) OPT.Optional[S, B] {
	_ = "STUB: not implemented"
	return nil
}
