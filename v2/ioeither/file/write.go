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

package file

import (
	"io"
)

func onWriteAll[W io.Writer](data []byte) Kleisli[error, W, []byte] {
	_ = "STUB: not implemented"
	return nil
}

// WriteAll writes data to a WriteCloser and ensures it is properly closed.
// It takes the data to write and returns an Operator that accepts an IOEither
// that creates the WriteCloser. The WriteCloser is automatically closed after
// the write operation, even if an error occurs.
//
// Example:
//
//	writeOp := F.Pipe2(
//		Open("output.txt"),
//		WriteAll([]byte("Hello, World!")),
//	)
//	result := writeOp() // Either[error, []byte]
func WriteAll[W io.WriteCloser](data []byte) Operator[error, W, []byte] {
	_ = "STUB: not implemented"
	return nil
}

// Write creates a resource-safe writer that automatically manages the lifecycle of a WriteCloser.
// It takes an IOEither that acquires the WriteCloser and returns a Kleisli arrow that accepts
// a write operation. The WriteCloser is automatically closed after the operation completes,
// even if an error occurs.
//
// This is useful for composing multiple write operations with proper resource management.
//
// Example:
//
//	writeOp := Write[int](Open("output.txt"))
//	result := writeOp(func(f *os.File) IOEither[error, int] {
//		return ioeither.TryCatchError(func() (int, error) {
//			return f.Write([]byte("Hello"))
//		})
//	})
func Write[R any, W io.WriteCloser](acquire IOEither[error, W]) Kleisli[error, Kleisli[error, W, R], R] {
	_ = "STUB: not implemented"
	return nil
}
