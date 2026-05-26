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

package generic

// These functions curry a Go function with the context as the first parameter into a generic Reader
// with the context as the last parameter, which is equivalent to a function returning a Reader
// of that context.
//
// This follows the Go convention (https://pkg.go.dev/context) of putting the context as the
// first parameter, while Reader monad convention has the context as the last parameter.
//
// The generic versions work with custom reader types that match the pattern ~func(R) A.

// Curry0 converts a function that takes a context and returns a value into a generic Reader.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - A: The result type
func Curry0[GA ~func(R) A, R, A any](f func(R) A) GA {
	_ = "STUB: not implemented"
	return *

	// Curry1 converts a function with context as first parameter into a curried function
	// returning a generic Reader. The context parameter is moved to the end (Reader position).
	//
	// Type Parameters:
	//   - GA: The generic Reader type (~func(R) A)
	//   - R: The environment/context type
	//   - T1: The first parameter type
	//   - A: The result type
	new(GA)
}

func Curry1[GA ~func(R) A, R, T1, A any](f func(R, T1) A) func(T1) GA {
	_ = "STUB: not implemented"
	return nil
}

// Curry2 converts a function with context as first parameter and 2 other parameters
// into a curried function returning a generic Reader.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - T1, T2: The parameter types
//   - A: The result type
func Curry2[GA ~func(R) A, R, T1, T2, A any](f func(R, T1, T2) A) func(T1) func(T2) GA {
	_ = "STUB: not implemented"
	return nil
}

// Curry3 converts a function with context as first parameter and 3 other parameters
// into a curried function returning a generic Reader.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - T1, T2, T3: The parameter types
//   - A: The result type
func Curry3[GA ~func(R) A, R, T1, T2, T3, A any](f func(R, T1, T2, T3) A) func(T1) func(T2) func(T3) GA {
	_ = "STUB: not implemented"
	return nil
}

// Curry4 converts a function with context as first parameter and 4 other parameters
// into a curried function returning a generic Reader.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - T1, T2, T3, T4: The parameter types
//   - A: The result type
func Curry4[GA ~func(R) A, R, T1, T2, T3, T4, A any](f func(R, T1, T2, T3, T4) A) func(T1) func(T2) func(T3) func(T4) GA {
	_ = "STUB: not implemented"
	return nil
}

// Uncurry0 converts a generic Reader back into a regular function with context as first parameter.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - A: The result type
func Uncurry0[GA ~func(R) A, R, A any](f GA) func(R) A {
	_ = "STUB: not implemented"

	// Uncurry1 converts a curried function returning a generic Reader back into a regular function
	// with context as first parameter.
	//
	// Type Parameters:
	//   - GA: The generic Reader type (~func(R) A)
	//   - R: The environment/context type
	//   - T1: The first parameter type
	//   - A: The result type
	return nil
}

func Uncurry1[GA ~func(R) A, R, T1, A any](f func(T1) GA) func(R, T1) A {
	_ = "STUB: not implemented"
	return nil
}

// Uncurry2 converts a curried function with 2 parameters returning a generic Reader back into
// a regular function with context as first parameter.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - T1, T2: The parameter types
//   - A: The result type
func Uncurry2[GA ~func(R) A, R, T1, T2, A any](f func(T1) func(T2) GA) func(R, T1, T2) A {
	_ = "STUB: not implemented"
	return nil
}

// Uncurry3 converts a curried function with 3 parameters returning a generic Reader back into
// a regular function with context as first parameter.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - T1, T2, T3: The parameter types
//   - A: The result type
func Uncurry3[GA ~func(R) A, R, T1, T2, T3, A any](f func(T1) func(T2) func(T3) GA) func(R, T1, T2, T3) A {
	_ = "STUB: not implemented"
	return nil
}

// Uncurry4 converts a curried function with 4 parameters returning a generic Reader back into
// a regular function with context as first parameter.
//
// Type Parameters:
//   - GA: The generic Reader type (~func(R) A)
//   - R: The environment/context type
//   - T1, T2, T3, T4: The parameter types
//   - A: The result type
func Uncurry4[GA ~func(R) A, R, T1, T2, T3, T4, A any](f func(T1) func(T2) func(T3) func(T4) GA) func(R, T1, T2, T3, T4) A {
	_ = "STUB: not implemented"
	return nil
}
