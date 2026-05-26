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

package optional

import (
	L "github.com/IBM/fp-go/optics/lens"
	OPT "github.com/IBM/fp-go/optics/optional"
	O "github.com/IBM/fp-go/option"
)

func lensAsOptional[S, A any](creator func(get func(S) O.Option[A], set func(S, A) S) OPT.Optional[S, A], sa L.Lens[S, A]) OPT.Optional[S, A] {
	_ = "STUB: not implemented"
	return nil
}

// LensAsOptional converts a Lens into an Optional
func LensAsOptional[S, A any](sa L.Lens[S, A]) OPT.Optional[S, A] {
	_ = "STUB: not implemented"
	return nil
}
