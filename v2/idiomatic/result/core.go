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

package result

// String prints some debug info for the object
func ToString[A any](a A, err error) string { _ = "STUB: not implemented"; return "" }

// IsLeft tests if the Either is a Left value.
// Rather use [Fold] or [MonadFold] if you need to access the values.
// Inverse is [IsRight].
//
// Example:
//
//	either.IsLeft(either.Left[int](errors.New("err"))) // true
//	either.IsLeft(either.Right[error](42)) // false
//
//go:inline
func IsLeft[A any](_ A, err error) bool {
	_ = "STUB: not implemented"

	// IsRight tests if the Either is a Right value.
	// Rather use [Fold] or [MonadFold] if you need to access the values.
	// Inverse is [IsLeft].
	//
	// Example:
	//
	//	either.IsRight(either.Right[error](42)) // true
	//	either.IsRight(either.Left[int](errors.New("err"))) // false
	//
	//go:inline
	return false
}

func IsRight[A any](_ A, err error) bool {
	_ = "STUB: not implemented"

	// Left creates a new Either representing a Left (error/failure) value.
	// By convention, Left represents the error case.
	//
	// Example:
	//
	//	result := either.Left[int](errors.New("something went wrong"))
	//
	//go:inline
	return false
}

func Left[A any](err error) (A, error) {
	_ = "STUB: not implemented"
	return *

	// Right creates a new Either representing a Right (success) value.
	// By convention, Right represents the success case.
	//
	// Example:
	//
	//	result := either.Right[error](42)
	//
	//go:inline
	new(A), nil
}

func Right[A any](a A) (A, error) { _ = "STUB: not implemented"; return *new(A), nil }
