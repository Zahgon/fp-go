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

package readerio

// MonadApFirst combines two effectful actions, keeping only the result of the first.
func MonadApFirst[A, R, B any](first ReaderIO[R, A], second ReaderIO[R, B]) ReaderIO[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// ApFirst combines two effectful actions, keeping only the result of the first.
func ApFirst[A, R, B any](second ReaderIO[R, B]) func(ReaderIO[R, A]) ReaderIO[R, A] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApSecond combines two effectful actions, keeping only the result of the second.
func MonadApSecond[A, R, B any](first ReaderIO[R, A], second ReaderIO[R, B]) ReaderIO[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApSecond combines two effectful actions, keeping only the result of the second.
func ApSecond[A, R, B any](second ReaderIO[R, B]) func(ReaderIO[R, A]) ReaderIO[R, B] {
	_ = "STUB: not implemented"
	return nil
}
