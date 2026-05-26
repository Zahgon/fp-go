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

package readerreaderioresult

import (
	"github.com/IBM/fp-go/v2/either"
	"github.com/IBM/fp-go/v2/io"
	"github.com/IBM/fp-go/v2/ioeither"
	"github.com/IBM/fp-go/v2/ioresult"
	"github.com/IBM/fp-go/v2/reader"
	"github.com/IBM/fp-go/v2/readerio"
)

// Do creates an empty context of type [S] to be used with the [Bind] operation.
// This is the starting point for do-notation style composition with two reader contexts.
//
// Example:
//
//	type State struct {
//	    User   User
//	    Posts  []Post
//	}
//	type OuterEnv struct {
//	    Database string
//	}
//	type InnerEnv struct {
//	    UserRepo UserRepository
//	    PostRepo PostRepository
//	}
//	result := readerreaderioeither.Do[OuterEnv, InnerEnv, error](State{})
//
//go:inline
func Do[R, S any](
	empty S,
) ReaderReaderIOResult[R, S] {
	_ = "STUB: not implemented"
	return nil

	// Bind attaches the result of a computation to a context [S1] to produce a context [S2].
	// This enables sequential composition where each step can depend on the results of previous steps
	// and access both the outer (R) and inner (C) reader environments.
	//
	// The setter function takes the result of the computation and returns a function that
	// updates the context from S1 to S2.
	//
	// Example:
	//
	//	type State struct {
	//	    User   User
	//	    Posts  []Post
	//	}
	//	type OuterEnv struct {
	//	    Database string
	//	}
	//	type InnerEnv struct {
	//	    UserRepo UserRepository
	//	    PostRepo PostRepository
	//	}
	//
	//	result := F.Pipe2(
	//	    readerreaderioeither.Do[OuterEnv, InnerEnv, error](State{}),
	//	    readerreaderioeither.Bind(
	//	        func(user User) func(State) State {
	//	            return func(s State) State { s.User = user; return s }
	//	        },
	//	        func(s State) readerreaderioeither.ReaderReaderIOResult[OuterEnv, InnerEnv, error, User] {
	//	            return func(outer OuterEnv) readerioeither.ReaderIOEither[InnerEnv, error, User] {
	//	                return readerioeither.Asks(func(inner InnerEnv) ioeither.IOEither[error, User] {
	//	                    return inner.UserRepo.FindUser(outer.Database)
	//	                })
	//	            }
	//	        },
	//	    ),
	//	    readerreaderioeither.Bind(
	//	        func(posts []Post) func(State) State {
	//	            return func(s State) State { s.Posts = posts; return s }
	//	        },
	//	        func(s State) readerreaderioeither.ReaderReaderIOResult[OuterEnv, InnerEnv, error, []Post] {
	//	            return func(outer OuterEnv) readerioeither.ReaderIOEither[InnerEnv, error, []Post] {
	//	                return readerioeither.Asks(func(inner InnerEnv) ioeither.IOEither[error, []Post] {
	//	                    return inner.PostRepo.FindPostsByUser(outer.Database, s.User.ID)
	//	                })
	//	            }
	//	        },
	//	    ),
	//	)
	//
	//go:inline
}

func Bind[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) ReaderReaderIOResult[R, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// Let attaches a pure computation result to a context [S1] to produce a context [S2].
// Unlike [Bind], the computation function f is pure (doesn't perform effects).
//
// Example:
//
//	readerreaderioeither.Let(
//	    func(fullName string) func(State) State {
//	        return func(s State) State { s.FullName = fullName; return s }
//	    },
//	    func(s State) string {
//	        return s.FirstName + " " + s.LastName
//	    },
//	)
//
//go:inline
func Let[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f func(S1) T,
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// LetTo attaches a constant value to a context [S1] to produce a context [S2].
//
// Example:
//
//	readerreaderioeither.LetTo(
//	    func(status string) func(State) State {
//	        return func(s State) State { s.Status = status; return s }
//	    },
//	    "active",
//	)
//
//go:inline
func LetTo[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	b T,
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindTo wraps a value of type T into a context S1 using the provided setter function.
// This is typically used as the first operation after [Do] to initialize the context.
//
// Example:
//
//	F.Pipe1(
//	    readerreaderioeither.Of[OuterEnv, InnerEnv, error](42),
//	    readerreaderioeither.BindTo(func(n int) State { return State{Count: n} }),
//	)
//
//go:inline
func BindTo[R, S1, T any](
	setter func(T) S1,
) Operator[R, T, S1] {
	_ = "STUB: not implemented"
	return nil
}

// ApS applies a computation in parallel (applicative style) and attaches its result to the context.
// Unlike [Bind], this doesn't allow the computation to depend on the current context state.
//
// Example:
//
//	readerreaderioeither.ApS(
//	    func(count int) func(State) State {
//	        return func(s State) State { s.Count = count; return s }
//	    },
//	    getCount, // ReaderReaderIOResult[OuterEnv, InnerEnv, error, int]
//	)
//
//go:inline
func ApS[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa ReaderReaderIOResult[R, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// ApSL is a lens-based version of [ApS] that uses a lens to focus on a specific field in the context.
//
//go:inline
func ApSL[R, S, T any](
	lens Lens[S, T],
	fa ReaderReaderIOResult[R, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil

	// BindL is a lens-based version of [Bind] that uses a lens to focus on a specific field in the context.
	//
	//go:inline
}

func BindL[R, S, T any](
	lens Lens[S, T],
	f func(T) ReaderReaderIOResult[R, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// LetL is a lens-based version of [Let] that uses a lens to focus on a specific field in the context.
//
//go:inline
func LetL[R, S, T any](
	lens Lens[S, T],
	f func(T) T,
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// LetToL is a lens-based version of [LetTo] that uses a lens to focus on a specific field in the context.
//
//go:inline
func LetToL[R, S, T any](
	lens Lens[S, T],
	b T,
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// BindIOEitherK binds a computation that returns an IOEither to the context.
// The Kleisli function is automatically lifted into ReaderReaderIOResult.
//
//go:inline
func BindIOEitherK[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f ioeither.Kleisli[error, S1, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func BindIOResultK[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f ioresult.Kleisli[S1, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindIOK binds a computation that returns an IO to the context.
// The Kleisli function is automatically lifted into ReaderReaderIOResult.
//
//go:inline
func BindIOK[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f io.Kleisli[S1, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindReaderK binds a computation that returns a Reader to the context.
// The Kleisli function is automatically lifted into ReaderReaderIOResult.
//
//go:inline
func BindReaderK[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f reader.Kleisli[R, S1, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindReaderIOK binds a computation that returns a ReaderIO to the context.
// The Kleisli function is automatically lifted into ReaderReaderIOResult.
//
//go:inline
func BindReaderIOK[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f readerio.Kleisli[R, S1, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindEitherK binds a computation that returns an Either to the context.
// The Kleisli function is automatically lifted into ReaderReaderIOResult.
//
//go:inline
func BindEitherK[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	f either.Kleisli[error, S1, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// BindIOEitherKL is a lens-based version of [BindIOEitherK].
//
//go:inline
func BindIOEitherKL[R, S, T any](
	lens Lens[S, T],
	f ioeither.Kleisli[error, T, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// BindIOKL is a lens-based version of [BindIOK].
//
//go:inline
func BindIOKL[R, S, T any](
	lens Lens[S, T],
	f io.Kleisli[T, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// BindReaderKL is a lens-based version of [BindReaderK].
//
//go:inline
func BindReaderKL[R, S, T any](
	lens Lens[S, T],
	f reader.Kleisli[R, T, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// BindReaderIOKL is a lens-based version of [BindReaderIOK].
//
//go:inline
func BindReaderIOKL[R, S, T any](
	lens Lens[S, T],
	f readerio.Kleisli[R, T, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// ApIOEitherS applies an IOEither computation and attaches its result to the context.
//
//go:inline
func ApIOEitherS[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa IOEither[error, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// ApIOS applies an IO computation and attaches its result to the context.
//
//go:inline
func ApIOS[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa IO[T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// ApReaderS applies a Reader computation and attaches its result to the context.
//
//go:inline
func ApReaderS[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa Reader[R, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// ApReaderIOS applies a ReaderIO computation and attaches its result to the context.
//
//go:inline
func ApReaderIOS[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa ReaderIO[R, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// ApEitherS applies an Either computation and attaches its result to the context.
//
//go:inline
func ApEitherS[R, S1, S2, T any](
	setter func(T) func(S1) S2,
	fa Either[error, T],
) Operator[R, S1, S2] {
	_ = "STUB: not implemented"
	return nil
}

// ApIOEitherSL is a lens-based version of [ApIOEitherS].
//
//go:inline
func ApIOEitherSL[R, S, T any](
	lens Lens[S, T],
	fa IOEither[error, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// ApIOSL is a lens-based version of [ApIOS].
//
//go:inline
func ApIOSL[R, S, T any](
	lens Lens[S, T],
	fa IO[T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// ApReaderSL is a lens-based version of [ApReaderS].
//
//go:inline
func ApReaderSL[R, S, T any](
	lens Lens[S, T],
	fa Reader[R, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// ApReaderIOSL is a lens-based version of [ApReaderIOS].
//
//go:inline
func ApReaderIOSL[R, S, T any](
	lens Lens[S, T],
	fa ReaderIO[R, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}

// ApEitherSL is a lens-based version of [ApEitherS].
//
//go:inline
func ApEitherSL[R, S, T any](
	lens Lens[S, T],
	fa Either[error, T],
) Operator[R, S, S] {
	_ = "STUB: not implemented"
	return nil
}
