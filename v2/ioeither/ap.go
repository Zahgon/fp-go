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

package ioeither

// MonadApFirst combines two effectful actions, keeping only the result of the first.
func MonadApFirst[A, E, B any](first IOEither[E, A], second IOEither[E, B]) IOEither[E, A] {
	_ = "STUB: not implemented"
	return nil
}

// ApFirst combines two effectful actions, keeping only the result of the first.
func ApFirst[A, E, B any](second IOEither[E, B]) Operator[E, A, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApSecond combines two effectful actions, keeping only the result of the second.
func MonadApSecond[A, E, B any](first IOEither[E, A], second IOEither[E, B]) IOEither[E, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApSecond combines two effectful actions, keeping only the result of the second.
func ApSecond[A, E, B any](second IOEither[E, B]) Operator[E, A, B] {
	_ = "STUB: not implemented"
	return nil
}
