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
	"time"

	R "github.com/IBM/fp-go/retry"
)

// Apply policy and delay by its amount if it results in a R.
// Returns updated status.
// HKTSTATUS = HKT<R.RetryStatus>
func applyAndDelay[HKTSTATUS any](
	monadOf func(R.RetryStatus) HKTSTATUS,
	monadDelay func(time.Duration) func(HKTSTATUS) HKTSTATUS,
) func(policy R.RetryPolicy, status R.RetryStatus) HKTSTATUS {
	_ = "STUB: not implemented"
	return nil
}

// Retry combinator for actions that don't raise exceptions, but
// signal in their type the outcome has failed. Examples are the
// `Option`, `Either` and `EitherT` monads.
//
// policy - refers to the retry policy
// action - converts a status into an operation to be executed
// check  - checks if the result of the action needs to be retried
func Retrying[HKTA, HKTSTATUS, A any](
	monadChain func(func(A) HKTA) func(HKTA) HKTA,
	monadChainStatus func(func(R.RetryStatus) HKTA) func(HKTSTATUS) HKTA,
	monadOf func(A) HKTA,
	monadOfStatus func(R.RetryStatus) HKTSTATUS,
	monadDelay func(time.Duration) func(HKTSTATUS) HKTSTATUS,

	policy R.RetryPolicy,
	action func(R.RetryStatus) HKTA,
	check func(A) bool,
) HKTA {
	_ = "STUB: not implemented"
	// delay callback
	return *new(HKTA)
}

// function to check if we need to retry or not

// need some lazy init because we reference it in the chain

// seed
