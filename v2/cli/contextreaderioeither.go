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

	A "github.com/IBM/fp-go/v2/array"
	C "github.com/urfave/cli/v3"
)

// Deprecated:
func generateNestedCallbacks(i, total int) string { _ = "STUB: not implemented"; return "" }

func generateNestedCallbacksPlain(i, total int) string { _ = "STUB: not implemented"; return "" }

func generateContextReaderIOEitherEitherize(f, fg *os.File, i int) {
	_ = "STUB: not implemented"
	// non generic version
	return
}

// generic version

func generateContextReaderIOEitherUneitherize(f, fg *os.File, i int) {
	_ = "STUB: not implemented"
	// non generic version
	return
}

// generic version

func nonGenericContextReaderIOEither(param string) string { _ = "STUB: not implemented"; return "" }

var extrasContextReaderIOEither = A.Empty[string]()

func generateContextReaderIOEitherSequenceT(f *os.File, i int) { _ = "STUB: not implemented"; return }

func generateContextReaderIOEitherSequenceTuple(f *os.File, i int) {
	_ = "STUB: not implemented"
	return
}

func generateContextReaderIOEitherTraverseTuple(f *os.File, i int) {
	_ = "STUB: not implemented"
	return
}

func generateContextReaderIOEitherHelpers(filename string, count int) error {
	_ = "STUB: not implemented"
	return nil
}

// construct subdirectory

// log

// eitherize

// sequenceT

// sequenceTuple

// traverseTuple

func ContextReaderIOEitherCommand() *C.Command { _ = "STUB: not implemented"; return nil }
