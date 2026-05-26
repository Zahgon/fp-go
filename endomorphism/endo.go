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

package endomorphism

func MonadAp[A any](fab Endomorphism[A], fa A) A { _ = "STUB: not implemented"; return *new(A) }

func Ap[A any](fa A) func(Endomorphism[A]) A { _ = "STUB: not implemented"; return nil }

func MonadChain[A any](ma Endomorphism[A], f Endomorphism[A]) Endomorphism[A] {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A any](f Endomorphism[A]) Endomorphism[Endomorphism[A]] {
	_ = "STUB: not implemented"
	return nil
}
