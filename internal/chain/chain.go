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

package chain

// HKTA=HKT[A]
// HKTB=HKT[B]
func MonadChainFirst[A, B, HKTA, HKTB any](
	mchain func(HKTA, func(A) HKTA) HKTA,
	mmap func(HKTB, func(B) A) HKTA,
	first HKTA,
	f func(A) HKTB,
) HKTA {
	_ = "STUB: not implemented"
	return *new(HKTA)
}

func MonadChain[A, B, HKTA, HKTB any](
	mchain func(HKTA, func(A) HKTB) HKTB,
	first HKTA,
	f func(A) HKTB,
) HKTB {
	_ = "STUB: not implemented"
	return *

	// HKTA=HKT[A]
	// HKTB=HKT[B]
	new(HKTB)
}

func ChainFirst[A, B, HKTA, HKTB any](
	mchain func(func(A) HKTA) func(HKTA) HKTA,
	mmap func(func(B) A) func(HKTB) HKTA,
	f func(A) HKTB) func(HKTA) HKTA {
	_ = "STUB: not implemented"
	return nil
}

func Chain[A, B, HKTA, HKTB any](
	mchain func(func(A) HKTB) func(HKTA) HKTB,
	f func(A) HKTB,
) func(HKTA) HKTB {
	_ = "STUB: not implemented"
	return nil
}

func MonadBind[S1, S2, B, HKTS1, HKTS2, HKTB any](
	mchain func(HKTS1, func(S1) HKTS2) HKTS2,
	mmap func(HKTB, func(B) S2) HKTS2,
	first HKTS1,
	key func(B) func(S1) S2,
	f func(S1) HKTB,
) HKTS2 {
	_ = "STUB: not implemented"
	return *new(HKTS2)
}

func Bind[S1, S2, B, HKTS1, HKTS2, HKTB any](
	mchain func(func(S1) HKTS2) func(HKTS1) HKTS2,
	mmap func(func(B) S2) func(HKTB) HKTS2,
	key func(B) func(S1) S2,
	f func(S1) HKTB,
) func(HKTS1) HKTS2 {
	_ = "STUB: not implemented"
	return nil
}

func BindTo[S1, B, HKTS1, HKTB any](
	mmap func(func(B) S1) func(HKTB) HKTS1,
	key func(B) S1,
) func(fa HKTB) HKTS1 {
	_ = "STUB: not implemented"
	return nil
}

func MonadBindTo[S1, B, HKTS1, HKTB any](
	mmap func(HKTB, func(B) S1) HKTS1,
	first HKTB,
	key func(B) S1,
) HKTS1 {
	_ = "STUB: not implemented"
	return *new(HKTS1)
}
