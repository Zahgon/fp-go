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
	ET "github.com/IBM/fp-go/either"
)

// MonadApFirst combines two effectful actions, keeping only the result of the first.
func MonadApFirst[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], GBA ~func() ET.Either[E, func(B) A], E, A, B any](first GA, second GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ApFirst combines two effectful actions, keeping only the result of the first.
func ApFirst[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], GBA ~func() ET.Either[E, func(B) A], E, A, B any](second GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadApSecond combines two effectful actions, keeping only the result of the second.
func MonadApSecond[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], GBB ~func() ET.Either[E, func(B) B], E, A, B any](first GA, second GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// ApSecond combines two effectful actions, keeping only the result of the second.
func ApSecond[GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], GBB ~func() ET.Either[E, func(B) B], E, A, B any](second GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}
