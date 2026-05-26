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

const (
	// useParallel is the feature flag to control if we use the parallel or the sequential implementation of ap
	useParallel = true
)

// MonadApSeq implements the applicative on a single thread by first executing mab and the ma
func MonadApSeq[GA ~func() A, GB ~func() B, GAB ~func() func(A) B, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// MonadApPar implements the applicative on two threads, the main thread executes mab and the actuall
// apply operation and the second thread computes ma. Communication between the threads happens via a channel
func MonadApPar[GA ~func() A, GB ~func() B, GAB ~func() func(A) B, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// MonadAp implements the `ap` operation. Depending on a feature flag this will be sequential or parallel, the preferred implementation
// is parallel
func MonadAp[GA ~func() A, GB ~func() B, GAB ~func() func(A) B, A, B any](mab GAB, ma GA) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// MonadApFirst combines two effectful actions, keeping only the result of the first.
func MonadApFirst[GA ~func() A, GB ~func() B, GBA ~func() func(B) A, A, B any](first GA, second GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// MonadApFirstPar combines two effectful actions, keeping only the result of the first.
func MonadApFirstPar[GA ~func() A, GB ~func() B, GBA ~func() func(B) A, A, B any](first GA, second GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// MonadApFirstSeq combines two effectful actions, keeping only the result of the first.
func MonadApFirstSeq[GA ~func() A, GB ~func() B, GBA ~func() func(B) A, A, B any](first GA, second GB) GA {
	_ = "STUB: not implemented"
	return *new(GA)
}

// ApFirst combines two effectful actions, keeping only the result of the first.
func ApFirst[GA ~func() A, GB ~func() B, GBA ~func() func(B) A, A, B any](second GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// ApFirstPar combines two effectful actions, keeping only the result of the first.
func ApFirstPar[GA ~func() A, GB ~func() B, GBA ~func() func(B) A, A, B any](second GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// ApFirstSeq combines two effectful actions, keeping only the result of the first.
func ApFirstSeq[GA ~func() A, GB ~func() B, GBA ~func() func(B) A, A, B any](second GB) func(GA) GA {
	_ = "STUB: not implemented"
	return nil
}

// MonadApSecond combines two effectful actions, keeping only the result of the second.
func MonadApSecond[GA ~func() A, GB ~func() B, GBB ~func() func(B) B, A, B any](first GA, second GB) GB {
	_ = "STUB: not implemented"
	return *new(GB)
}

// ApSecond combines two effectful actions, keeping only the result of the second.
func ApSecond[GA ~func() A, GB ~func() B, GBB ~func() func(B) B, A, B any](second GB) func(GA) GB {
	_ = "STUB: not implemented"
	return nil
}
