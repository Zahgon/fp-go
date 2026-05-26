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

package testing

import (
	"testing"

	EQ "github.com/IBM/fp-go/v2/eq"
	"github.com/IBM/fp-go/v2/internal/functor"
	"github.com/IBM/fp-go/v2/internal/pointed"
)

// SequenceArrayTest tests if the sequence operation works in case the operation cannot error
func SequenceArrayTest[
	HKTA,
	HKTB,
	HKTAA any, // HKT[[]A]
](
	eq EQ.Eq[HKTB],

	pa pointed.Pointed[string, HKTA],
	pb pointed.Pointed[bool, HKTB],
	faa functor.Functor[[]string, bool, HKTAA, HKTB],
	seq func([]HKTA) HKTAA,
) func(count int) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

// SequenceArrayErrorTest tests if the sequence operation works in case the operation can error
func SequenceArrayErrorTest[
	HKTA,
	HKTB,
	HKTAA any, // HKT[[]A]
](
	eq EQ.Eq[HKTB],

	left func(error) HKTA,
	leftB func(error) HKTB,
	pa pointed.Pointed[string, HKTA],
	pb pointed.Pointed[bool, HKTB],
	faa functor.Functor[[]string, bool, HKTAA, HKTB],
	seq func([]HKTA) HKTAA,
) func(count int) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}

// test the good case

// iterate and test the bad cases

// run the test

// the actual

// the expected error

// prepare the values bases on the bit mask

// test the good case

// validate the error

// SequenceRecordTest tests if the sequence operation works in case the operation cannot error
func SequenceRecordTest[
	HKTA,
	HKTB,
	HKTAA any, // HKT[map[string]string]
](
	eq EQ.Eq[HKTB],

	pa pointed.Pointed[string, HKTA],
	pb pointed.Pointed[bool, HKTB],
	faa functor.Functor[map[string]string, bool, HKTAA, HKTB],
	seq func(map[string]HKTA) HKTAA,
) func(count int) func(t *testing.T) {
	_ = "STUB: not implemented"
	return nil
}
