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

package fromio

func MonadChainFirstIOK[A, B, HKTA, HKTB any, GIOB ~func() B](
	mchain func(HKTA, func(A) HKTA) HKTA,
	mmap func(HKTB, func(B) A) HKTA,
	fromio func(GIOB) HKTB,
	first HKTA, f func(A) GIOB) HKTA {
	_ = "STUB: not implemented"
	// chain
	return *new(HKTA)
}

func ChainFirstIOK[A, B, HKTA, HKTB any, GIOB ~func() B](
	mchain func(func(A) HKTA) func(HKTA) HKTA,
	mmap func(func(B) A) func(HKTB) HKTA,
	fromio func(GIOB) HKTB,
	f func(A) GIOB) func(HKTA) HKTA {
	_ = "STUB: not implemented"
	// chain
	return nil
}

func MonadChainIOK[GR ~func() B, A, B, HKTA, HKTB any](
	mchain func(HKTA, func(A) HKTB) HKTB,
	fromio func(GR) HKTB,
	first HKTA, f func(A) GR) HKTB {
	_ = "STUB: not implemented"
	// chain
	return *new(HKTB)
}

func ChainIOK[GR ~func() B, A, B, HKTA, HKTB any](
	mchain func(func(A) HKTB) func(HKTA) HKTB,
	fromio func(GR) HKTB,
	f func(A) GR) func(HKTA) HKTB {
	_ = "STUB: not implemented"
	// chain
	return nil
}
