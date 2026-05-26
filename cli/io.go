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

package cli

import (
	"os"

	A "github.com/IBM/fp-go/array"
	C "github.com/urfave/cli/v2"
)

func nonGenericIO(param string) string { _ = "STUB: not implemented"; return "" }

func genericIO(param string) string { _ = "STUB: not implemented"; return "" }

var extrasIO = A.Empty[string]()

func generateIOSequenceT(f, fg *os.File, i int) { _ = "STUB: not implemented"; return }

func generateIOSequenceTuple(f, fg *os.File, i int) { _ = "STUB: not implemented"; return }

func generateIOTraverseTuple(f, fg *os.File, i int) { _ = "STUB: not implemented"; return }

func generateIOHelpers(filename string, count int) error { _ = "STUB: not implemented"; return nil }

// construct subdirectory

// log

// some header

// some header

// sequenceT

// sequenceTuple

// traverseTuple

func IOCommand() *C.Command { _ = "STUB: not implemented"; return nil }
