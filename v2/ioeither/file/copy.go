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

// CopyFile copies a file from source to destination path with proper resource management.
//
// This is a curried function that follows the "data-last" pattern, where the source path
// is provided first, returning a function that accepts the destination path. This design
// enables partial application and better composition with other functional operations.
//
// The function uses [ioeither.WithResource] to ensure both source and destination files
// are properly closed, even if an error occurs during the copy operation. The copy is
// performed using [io.Copy] which efficiently transfers data between the files.
//
// Parameters:
//   - src: The path to the source file to copy from
//
// Returns:
//   - A function that accepts the destination path and returns an [IOEither] that:
//   - On success: Contains the destination path (Right)
//   - On failure: Contains the error (Left) from opening, copying, or closing files
//
// Example:
//
//	// Create a copy operation for a specific source file
//	copyFromSource := CopyFile("/path/to/source.txt")
//
//	// Execute the copy to a destination
//	result := copyFromSource("/path/to/destination.txt")()
//
//	// Or use it in a pipeline
//	result := F.Pipe1(
//	    CopyFile("/path/to/source.txt"),
//	    ioeither.Map(func(dst string) string {
//	        return "Copied to: " + dst
//	    }),
//	)("/path/to/destination.txt")()
//
//go:inline
func CopyFile(src string) func(dst string) IOEither[error, string] {
	_ = "STUB: not implemented"
	return nil
}
