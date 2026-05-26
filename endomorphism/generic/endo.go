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

func MonadAp[GA ~func(A) A, A any](fab GA, fa A) A { _ = "STUB: not implemented"; return *new(A) }

func Ap[GA ~func(A) A, A any](fa A) func(GA) A { _ = "STUB: not implemented"; return nil }

func MonadChain[GA ~func(A) A, A any](ma GA, f GA) GA { _ = "STUB: not implemented"; return *new(GA) }

func Chain[ENDO ~func(GA) GA, GA ~func(A) A, A any](f GA) ENDO {
	_ = "STUB: not implemented"
	return *new(ENDO)
}
