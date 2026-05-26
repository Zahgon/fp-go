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
	O "github.com/IBM/fp-go/option"
	R "github.com/IBM/fp-go/retry"
)

// Retry combinator for actions that don't raise exceptions, but
// signal in their type the outcome has failed. Examples are the
// `Option`, `Either` and `EitherT` monads.
func Retrying[GA ~func() O.Option[A], A any](
	policy R.RetryPolicy,
	action func(R.RetryStatus) GA,
	check func(A) bool,
) GA {
	_ = "STUB: not implemented"
	// get an implementation for the types
	return *new(GA)
}
