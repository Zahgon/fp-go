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

package exec

import (
	RIOE "github.com/IBM/fp-go/context/readerioeither"
	"github.com/IBM/fp-go/exec"
	F "github.com/IBM/fp-go/function"
)

var (
	// Command executes a cancelable command
	Command = F.Curry3(command)
)

func command(name string, args []string, in []byte) RIOE.ReaderIOEither[exec.CommandOutput] {
	_ = "STUB: not implemented"
	return nil
}
