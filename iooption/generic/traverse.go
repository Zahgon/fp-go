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
)

func TraverseArray[TB ~func() O.Option[B], TBS ~func() O.Option[GB], GA ~[]A, GB ~[]B, A, B any](f func(A) TB) func(GA) TBS {
	_ = "STUB: not implemented"
	return nil
}

func TraverseArrayWithIndex[TB ~func() O.Option[B], TBS ~func() O.Option[GB], GA ~[]A, GB ~[]B, A, B any](f func(int, A) TB) func(GA) TBS {
	_ = "STUB: not implemented"
	return nil
}

func SequenceArray[TB ~func() O.Option[B], TBS ~func() O.Option[GB], GA ~[]TB, GB ~[]B, A, B any](ma GA) TBS {
	_ = "STUB: not implemented"
	return *new(TBS)
}
