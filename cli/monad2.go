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
	S "github.com/IBM/fp-go/string"
)

var (
	concStrgs     = A.Monoid[string]().Concat
	intercalStrgs = A.Intercalate(S.Monoid)
	concAllStrgs  = A.ConcatAll(A.Monoid[string]())
)

func joinAll(middle string) func(all ...[]string) string { _ = "STUB: not implemented"; return nil }

func generateGenericSequenceT(
	nonGenericType func(string) string,
	genericType func(string) string,
	extra []string,
) func(f, fg *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// tuple

// all types T

// non generic version

// generic version

// map call

// the apply calls

// function parameters

func generateGenericSequenceTuple(
	nonGenericType func(string) string,
	genericType func(string) string,
	extra []string,
) func(f, fg *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// tuple

// all types T

// non generic version

// generic version

// map call

// the apply calls

// function parameters

func generateGenericTraverseTuple(
	nonGenericType func(string) string,
	genericType func(string) string,
	extra []string,
) func(f, fg *os.File, i int) {
	_ = "STUB: not implemented"
	return nil
}

// tuple

// all types T

// all types A

// all function types

// non generic version

// generic version

// map call

// the apply calls

// function parameters

// tuple parameter
