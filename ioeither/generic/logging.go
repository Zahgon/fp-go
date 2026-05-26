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

import (
	ET "github.com/IBM/fp-go/either"
)

// LogJSON converts the argument to JSON and then logs it via the format string
// Can be used with [ChainFirst]
func LogJSON[GA ~func() ET.Either[error, any], A any](prefix string) func(A) GA {
	_ = "STUB: not implemented"
	return nil

	// log this
}

// LogJson converts the argument to JSON and then logs it via the format string
// Can be used with [ChainFirst]
//
// Deprecated: use [LogJSON] instead
func LogJson[GA ~func() ET.Either[error, any], A any](prefix string) func(A) GA {
	_ = "STUB: not implemented"
	return nil
}
