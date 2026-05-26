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

package bracket

// Bracket makes sure that a resource is cleaned up in the event of an error. The release action is called regardless of
// whether the body action returns and error or not.
func Bracket[
	GA, // IOEither[E, A]
	GB, // IOEither[E, A]
	GANY, // IOEither[E, ANY]

	EB, // Either[E, B]

	A, B, ANY any](

	ofeb func(EB) GB,

	chainab func(GA, func(A) GB) GB,
	chainebb func(GB, func(EB) GB) GB,
	chainany func(GANY, func(ANY) GB) GB,

	acquire GA,
	use func(A) GB,
	release func(A, EB) GANY,
) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}
