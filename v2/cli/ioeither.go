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

// [GA ~func() ET.Either[E, A], GB ~func() ET.Either[E, B], GTAB ~func() ET.Either[E, T.Tuple2[A, B]], E, A, B any](a GA, b GB) GTAB {

func nonGenericIOEither(param string) string { _ = "STUB: not implemented"; return "" }

var extrasIOEither = A.From("E")

func generateIOEitherSequenceT(f, fg *os.File, i int) { _ = "STUB: not implemented"; return }

func generateIOEitherSequenceTuple(f, fg *os.File, i int) { _ = "STUB: not implemented"; return }

func generateIOEitherTraverseTuple(f, fg *os.File, i int) { _ = "STUB: not implemented"; return }

func generateIOEitherUneitherize(f, fg *os.File, i int) {
	_ = "STUB: not implemented"
	// non generic version
	return
}

// generic version

func generateIOEitherEitherize(f, fg *os.File, i int) {
	_ = "STUB: not implemented"
	// non generic version
	return
}

// generic version

func generateIOEitherHelpers(filename string, count int) error {
	_ = "STUB: not implemented"
	return nil
}

// construct subdirectory

// log

// some header

// some header

// eitherize

// uneitherize

// eitherize

// uneitherize

// sequenceT

// sequenceTuple

// traverseTuple

func IOEitherCommand() *C.Command { _ = "STUB: not implemented"; return nil }
