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

package file

import (
	"io"

	IO "github.com/IBM/fp-go/io"
)

// Close closes a closeable resource and ignores a potential error
func Close[R io.Closer](r R) IO.IO[R] { _ = "STUB: not implemented"; return nil }

// #nosec: G104

// Remove removes a resource and ignores a potential error
func Remove(name string) IO.IO[string] { _ = "STUB: not implemented"; return nil }

// #nosec: G104
