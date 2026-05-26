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

package cli

import (
	"os"
)

// Deprecated:
func tupleType(name string) func(i int) string { _ = "STUB: not implemented"; return nil }

func tupleTypePlain(name string) func(i int) string { _ = "STUB: not implemented"; return nil }

func monadGenerateSequenceTNonGeneric(
	hkt func(string) string,
	fmap func(string, string) string,
	fap func(string, string) string,
) func(f *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// the actual apply callback

// map callback

func monadGenerateSequenceTGeneric(
	hkt func(string) string,
	fmap func(string, string) string,
	fap func(string, string) string,
) func(f *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// the actual apply callback

// map callback

func generateTraverseTuple1(
	hkt func(string) string,
	infix string) func(f *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// functions

// types

// map

// applicatives

func generateSequenceTuple1(
	hkt func(string) string,
	infix string) func(f *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// map

// applicatives

func generateSequenceT1(
	hkt func(string) string,
	infix string) func(f *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// map

// applicatives
