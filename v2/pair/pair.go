// Copyright (c) 2024 - 2025 IBM Corp.
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

package pair

// Of creates a [Pair] with the same value in both the head and tail positions.
//
// Example:
//
//	p := pair.Of(42)  // Pair[int, int]{42, 42}
func Of[A any](value A) Pair[A, A] { _ = "STUB: not implemented"; return nil }

// FromTuple creates a [Pair] from a [Tuple2].
// The first element of the tuple becomes the head, and the second becomes the tail.
//
// Example:
//
//	t := tuple.MakeTuple2("hello", 42)
//	p := pair.FromTuple(t)  // Pair[string, int]{"hello", 42}
//
//go:inline
func FromTuple[A, B any](t Tuple2[A, B]) Pair[A, B] { _ = "STUB: not implemented"; return nil }

// FromHead creates a function that constructs a [Pair] from a given head value.
// It returns a function that takes a tail value and combines it with the head
// to create a Pair.
//
// This is useful for functional composition where you want to partially apply
// the head value and later provide the tail value.
//
// Example:
//
//	makePair := pair.FromHead[int]("hello")
//	p := makePair(42)  // Pair[string, int]{"hello", 42}
//
//go:inline
func FromHead[B, A any](a A) Kleisli[A, B, B] { _ = "STUB: not implemented"; return nil }

// FromTail creates a function that constructs a [Pair] from a given tail value.
// It returns a function that takes a head value and combines it with the tail
// to create a Pair.
//
// This is useful for functional composition where you want to partially apply
// the tail value and later provide the head value.
//
// Example:
//
//	makePair := pair.FromTail[string](42)
//	p := makePair("hello")  // Pair[string, int]{"hello", 42}
//
//go:inline
func FromTail[A, B any](b B) Kleisli[A, A, B] { _ = "STUB: not implemented"; return nil }

// ToTuple creates a [Tuple2] from a [Pair].
// The head becomes the first element, and the tail becomes the second element.
//
// Example:
//
//	p := pair.MakePair("hello", 42)
//	t := pair.ToTuple(p)  // Tuple2[string, int]{"hello", 42}
//
//go:inline
func ToTuple[A, B any](t Pair[A, B]) Tuple2[A, B] { _ = "STUB: not implemented"; return nil }

// MakePair creates a [Pair] from two values.
// The first value becomes the head, and the second becomes the tail.
//
// Example:
//
//	p := pair.MakePair("hello", 42)  // Pair[string, int]{"hello", 42}
//
//go:inline
func MakePair[A, B any](a A, b B) Pair[A, B] { _ = "STUB: not implemented"; return nil }

// Head returns the head (first) value of the pair.
//
// Example:
//
//	p := pair.MakePair("hello", 42)
//	h := pair.Head(p)  // "hello"
//
//go:inline
func Head[A, B any](fa Pair[A, B]) A {
	_ = "STUB: not implemented"

	// Tail returns the tail (second) value of the pair.
	//
	// Example:
	//
	//	p := pair.MakePair("hello", 42)
	//	t := pair.Tail(p)  // 42
	//
	//go:inline
	return *new(A)
}

func Tail[A, B any](fa Pair[A, B]) B {
	_ = "STUB: not implemented"

	// First returns the first value of the pair (alias for Head).
	//
	// Example:
	//
	//	p := pair.MakePair("hello", 42)
	//	f := pair.First(p)  // "hello"
	//
	//go:inline
	return *new(B)
}

func First[A, B any](fa Pair[A, B]) A {
	_ = "STUB: not implemented"

	// Second returns the second value of the pair (alias for Tail).
	//
	// Example:
	//
	//	p := pair.MakePair("hello", 42)
	//	s := pair.Second(p)  // 42
	//
	//go:inline
	return *new(A)
}

func Second[A, B any](fa Pair[A, B]) B {
	_ = "STUB: not implemented"

	// MonadMapHead maps a function over the head value of the pair, leaving the tail unchanged.
	//
	// Example:
	//
	//	p := pair.MakePair(5, "hello")
	//	p2 := pair.MonadMapHead(p, func(n int) string {
	//	    return fmt.Sprintf("%d", n)
	//	})  // Pair[string, string]{"5", "hello"}
	//
	//go:inline
	return *new(B)
}

func MonadMapHead[B, A, A1 any](fa Pair[A, B], f func(A) A1) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func MonadMap[A, B, B1 any](fa Pair[A, B], f func(B) B1) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil

	// MonadMapTail maps a function over the tail value of the pair, leaving the head unchanged.
	//
	// Example:
	//
	//	p := pair.MakePair(5, "hello")
	//	p2 := pair.MonadMapTail(p, func(s string) int {
	//	    return len(s)
	//	})  // Pair[int, int]{5, 5}
	//
	//go:inline
}

func MonadMapTail[A, B, B1 any](fa Pair[A, B], f func(B) B1) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadBiMap maps functions over both the head and tail values of the pair.
//
// Example:
//
//	p := pair.MakePair(5, "hello")
//	p2 := pair.MonadBiMap(p,
//	    strconv.Itoa,
//	    S.Size,
//	)  // Pair[string, int]{"5", 5}
//
//go:inline
func MonadBiMap[A, B, A1, B1 any](fa Pair[A, B], f func(A) A1, g func(B) B1) Pair[A1, B1] {
	_ = "STUB: not implemented"
	return nil
}

// Map returns a function that maps over the tail value of a pair (alias for MapTail).
// This is the curried version of MonadMapTail.
//
// Example:
//
//	mapper := pair.Map[int](S.Size)
//	p := pair.MakePair(5, "hello")
//	p2 := mapper(p)  // Pair[int, int]{5, 5}
//
//go:inline
func Map[A, B, B1 any](f func(B) B1) Operator[A, B, B1] { _ = "STUB: not implemented"; return nil }

// MapHead returns a function that maps over the head value of a pair.
// This is the curried version of MonadMapHead.
//
// Example:
//
//	mapper := pair.MapHead[string](func(n int) string {
//	    return fmt.Sprintf("%d", n)
//	})
//	p := pair.MakePair(5, "hello")
//	p2 := mapper(p)  // Pair[string, string]{"5", "hello"}
//
//go:inline
func MapHead[B, A, A1 any](f func(A) A1) func(Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MapTail returns a function that maps over the tail value of a pair.
// This is the curried version of MonadMapTail.
//
// Example:
//
//	mapper := pair.MapTail[int](S.Size)
//	p := pair.MakePair(5, "hello")
//	p2 := mapper(p)  // Pair[int, int]{5, 5}
//
//go:inline
func MapTail[A, B, B1 any](f func(B) B1) Operator[A, B, B1] { _ = "STUB: not implemented"; return nil }

// BiMap returns a function that maps over both values of a pair.
// This is the curried version of MonadBiMap.
//
// Example:
//
//	mapper := pair.BiMap(
//	    strconv.Itoa,
//	    S.Size,
//	)
//	p := pair.MakePair(5, "hello")
//	p2 := mapper(p)  // Pair[string, int]{"5", 5}
//
//go:inline
func BiMap[A, B, A1, B1 any](f func(A) A1, g func(B) B1) func(Pair[A, B]) Pair[A1, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainHead chains a function over the head value, combining tail values using a semigroup.
// The function receives the head value and returns a new pair. The tail values from both pairs
// are combined using the provided semigroup.
//
// Example:
//
//	import SG "github.com/IBM/fp-go/v2/semigroup"
//
//	strConcat := SG.MakeSemigroup(func(a, b string) string { return a + b })
//	p := pair.MakePair(5, "hello")
//	p2 := pair.MonadChainHead(strConcat, p, func(n int) pair.Pair[string, string] {
//	    return pair.MakePair(fmt.Sprintf("%d", n), "!")
//	})  // Pair[string, string]{"5", "hello!"}
func MonadChainHead[B, A, A1 any](sg Semigroup[B], fa Pair[A, B], f func(A) Pair[A1, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChainTail chains a function over the tail value, combining head values using a semigroup.
// The function receives the tail value and returns a new pair. The head values from both pairs
// are combined using the provided semigroup.
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	p := pair.MakePair(5, "hello")
//	p2 := pair.MonadChainTail(intSum, p, func(s string) pair.Pair[int, int] {
//	    return pair.MakePair(len(s), len(s) * 2)
//	})  // Pair[int, int]{10, 10}
//
//go:inline
func MonadChainTail[A, B, B1 any](sg Semigroup[A], fb Pair[A, B], f Kleisli[A, B, B1]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadChain chains a function over the tail value (alias for MonadChainTail).
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	p := pair.MakePair(5, "hello")
//	p2 := pair.MonadChain(intSum, p, func(s string) pair.Pair[int, int] {
//	    return pair.MakePair(len(s), len(s) * 2)
//	})  // Pair[int, int]{10, 10}
//
//go:inline
func MonadChain[A, B, B1 any](sg Semigroup[A], fa Pair[A, B], f Kleisli[A, B, B1]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// ChainHead returns a function that chains over the head value.
// This is the curried version of MonadChainHead.
//
// Example:
//
//	import SG "github.com/IBM/fp-go/v2/semigroup"
//
//	strConcat := SG.MakeSemigroup(func(a, b string) string { return a + b })
//	chain := pair.ChainHead(strConcat, func(n int) pair.Pair[string, string] {
//	    return pair.MakePair(fmt.Sprintf("%d", n), "!")
//	})
//	p := pair.MakePair(5, "hello")
//	p2 := chain(p)  // Pair[string, string]{"5", "hello!"}
//
//go:inline
func ChainHead[B, A, A1 any](sg Semigroup[B], f func(A) Pair[A1, B]) func(Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// ChainTail returns a function that chains over the tail value.
// This is the curried version of MonadChainTail.
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	chain := pair.ChainTail(intSum, func(s string) pair.Pair[int, int] {
//	    return pair.MakePair(len(s), len(s) * 2)
//	})
//	p := pair.MakePair(5, "hello")
//	p2 := chain(p)  // Pair[int, int]{10, 10}
//
//go:inline
func ChainTail[A, B, B1 any](sg Semigroup[A], f Kleisli[A, B, B1]) Operator[A, B, B1] {
	_ = "STUB: not implemented"
	return nil
}

// Chain returns a function that chains over the tail value (alias for ChainTail).
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	chain := pair.Chain(intSum, func(s string) pair.Pair[int, int] {
//	    return pair.MakePair(len(s), len(s) * 2)
//	})
//	p := pair.MakePair(5, "hello")
//	p2 := chain(p)  // Pair[int, int]{10, 10}
//
//go:inline
func Chain[A, B, B1 any](sg Semigroup[A], f Kleisli[A, B, B1]) Operator[A, B, B1] {
	_ = "STUB: not implemented"
	return nil

	// MonadApHead applies a function wrapped in a pair to a value wrapped in a pair,
	// operating on the head values and combining tail values using a semigroup.
	//
	// Example:
	//
	//	import SG "github.com/IBM/fp-go/v2/semigroup"
	//
	//	strConcat := SG.MakeSemigroup(func(a, b string) string { return a + b })
	//	pf := pair.MakePair(strconv.Itoa, "!")
	//	pv := pair.MakePair(42, "hello")
	//	result := pair.MonadApHead(strConcat, pf, pv)  // Pair[string, string]{"42", "!hello"}
	//
	//go:inline
}

func MonadApHead[B, A, A1 any](sg Semigroup[B], faa Pair[func(A) A1, B], fa Pair[A, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// MonadApTail applies a function wrapped in a pair to a value wrapped in a pair,
// operating on the tail values and combining head values using a semigroup.
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	pf := pair.MakePair(10, S.Size)
//	pv := pair.MakePair(5, "hello")
//	result := pair.MonadApTail(intSum, pf, pv)  // Pair[int, int]{15, 5}
//
//go:inline
func MonadApTail[A, B, B1 any](sg Semigroup[A], fbb Pair[A, func(B) B1], fb Pair[A, B]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// MonadAp applies a function wrapped in a pair to a value wrapped in a pair,
// operating on the tail values (alias for MonadApTail).
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	pf := pair.MakePair(10, S.Size)
//	pv := pair.MakePair(5, "hello")
//	result := pair.MonadAp(intSum, pf, pv)  // Pair[int, int]{15, 5}
//
//go:inline
func MonadAp[A, B, B1 any](sg Semigroup[A], faa Pair[A, func(B) B1], fa Pair[A, B]) Pair[A, B1] {
	_ = "STUB: not implemented"
	return nil
}

// ApHead returns a function that applies a function in a pair to a value in a pair,
// operating on head values. This is the curried version of MonadApHead.
//
// Example:
//
//	import SG "github.com/IBM/fp-go/v2/semigroup"
//
//	strConcat := SG.MakeSemigroup(func(a, b string) string { return a + b })
//	pv := pair.MakePair(42, "hello")
//	ap := pair.ApHead(strConcat, pv)
//	pf := pair.MakePair(strconv.Itoa, "!")
//	result := ap(pf)  // Pair[string, string]{"42", "!hello"}
func ApHead[B, A, A1 any](sg Semigroup[B], fa Pair[A, B]) func(Pair[func(A) A1, B]) Pair[A1, B] {
	_ = "STUB: not implemented"
	return nil
}

// ApTail returns a function that applies a function in a pair to a value in a pair,
// operating on tail values. This is the curried version of MonadApTail.
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	pv := pair.MakePair(5, "hello")
//	ap := pair.ApTail(intSum, pv)
//	pf := pair.MakePair(10, S.Size)
//	result := ap(pf)  // Pair[int, int]{15, 5}
func ApTail[A, B, B1 any](sg Semigroup[A], fb Pair[A, B]) Operator[A, func(B) B1, B1] {
	_ = "STUB: not implemented"
	return nil
}

// Ap returns a function that applies a function in a pair to a value in a pair,
// operating on tail values (alias for ApTail).
//
// Example:
//
//	import N "github.com/IBM/fp-go/v2/number"
//
//	intSum := N.SemigroupSum[int]()
//	pv := pair.MakePair(5, "hello")
//	ap := pair.Ap(intSum, pv)
//	pf := pair.MakePair(10, S.Size)
//	result := ap(pf)  // Pair[int, int]{15, 5}
//
//go:inline
func Ap[A, B, B1 any](sg Semigroup[A], fa Pair[A, B]) Operator[A, func(B) B1, B1] {
	_ = "STUB: not implemented"
	return nil
}

// Swap swaps the head and tail values of a pair.
//
// Example:
//
//	p := pair.MakePair("hello", 42)
//	swapped := pair.Swap(p)  // Pair[int, string]{42, "hello"}
//
//go:inline
func Swap[A, B any](fa Pair[A, B]) Pair[B, A] { _ = "STUB: not implemented"; return nil }

// Paired converts a function with 2 parameters into a function taking a [Pair].
// The inverse function is [Unpaired].
//
// Example:
//
//	add := func(a, b int) int { return a + b }
//	pairedAdd := pair.Paired(add)
//	result := pairedAdd(pair.MakePair(3, 4))  // 7
func Paired[F ~func(T1, T2) R, T1, T2, R any](f F) func(Pair[T1, T2]) R {
	_ = "STUB: not implemented"
	return nil
}

// Unpaired converts a function with a [Pair] parameter into a function with 2 parameters.
// The inverse function is [Paired].
//
// Example:
//
//	pairedAdd := func(p pair.Pair[int, int]) int {
//	    return pair.Head(p) + pair.Tail(p)
//	}
//	add := pair.Unpaired(pairedAdd)
//	result := add(3, 4)  // 7
func Unpaired[F ~func(Pair[T1, T2]) R, T1, T2, R any](f F) func(T1, T2) R {
	_ = "STUB: not implemented"
	return nil
}

// Merge applies a curried function to a pair by applying the tail value first, then the head value.
//
// Example:
//
//	add := func(b int) func(a int) int {
//	    return func(a int) int { return a + b }
//	}
//	merge := pair.Merge(add)
//	result := merge(pair.MakePair(3, 4))  // 7 (applies 4 then 3)
func Merge[F ~func(B) func(A) R, A, B, R any](f F) func(Pair[A, B]) R {
	_ = "STUB: not implemented"
	return nil
}

// Zero returns the zero value of a [Pair], which is a Pair with zero values for both head and tail.
// This function is useful for creating an empty Pair or as an identity element in monoid operations.
//
// The zero value for a Pair[L, R] has the zero value of type L as the head and the zero value
// of type R as the tail. For reference types (pointers, slices, maps, channels, functions, interfaces),
// the zero value is nil. For value types (numbers, booleans, structs), it's the type's zero value.
//
// Example:
//
//	// Zero pair of integers
//	p1 := pair.Zero[int, int]()  // Pair[int, int]{0, 0}
//
//	// Zero pair of string and int
//	p2 := pair.Zero[string, int]()  // Pair[string, int]{"", 0}
//
//	// Zero pair with pointer types
//	p3 := pair.Zero[*int, *string]()  // Pair[*int, *string]{nil, nil}
func Zero[L, R any]() Pair[L, R] { _ = "STUB: not implemented"; return nil }

// Unpack extracts both values from a [Pair] and returns them as separate values.
// This is the inverse operation of [MakePair], allowing you to destructure a Pair
// back into its constituent head and tail values.
//
// This function is particularly useful when you need to work with both values
// independently or pass them to functions that expect separate parameters rather
// than a Pair.
//
// Example:
//
//	p := pair.MakePair("hello", 42)
//	head, tail := pair.Unpack(p)  // head = "hello", tail = 42
//
//	// Using with function that expects separate parameters
//	result := someFunc(pair.Unpack(p))
//
//	// Destructuring for independent use
//	name, age := pair.Unpack(pair.MakePair("Alice", 30))
//	fmt.Printf("%s is %d years old\n", name, age)
//
//go:inline
func Unpack[L, R any](p Pair[L, R]) (L, R) { _ = "STUB: not implemented"; return *new(L), *new(R) }
