package tailrec

import "fmt"

// Bounce creates a Trampoline that indicates the computation should continue
// with a new intermediate state.
//
// This represents a recursive call in the original algorithm. The computation
// will continue by processing the provided state value in the next iteration.
//
// # Type Parameters
//
//   - B: The intermediate state type (bounce type)
//   - L: The final result type (land type)
//
// # Parameters
//
//   - b: The new intermediate state to process in the next step
//
// # Returns
//
//   - Trampoline[B, L]: A Trampoline in the "bounce" state containing the intermediate value
//
// # Example
//
//	// Countdown that bounces until reaching zero
//	func countdownStep(n int) Trampoline[int, int] {
//	    if n <= 0 {
//	        return Land[int](0)
//	    }
//	    return Bounce[int](n - 1)  // Continue with n-1
//	}
//
//go:inline
func Bounce[L, B any](b B) Trampoline[B, L] { _ = "STUB: not implemented"; return nil }

// Land creates a Trampoline that indicates the computation is complete
// with a final result.
//
// This represents the base case in the original recursive algorithm. When
// a Land trampoline is encountered, the executor should stop iterating and
// return the final result.
//
// # Type Parameters
//
//   - B: The intermediate state type (bounce type)
//   - L: The final result type (land type)
//
// # Parameters
//
//   - l: The final result value
//
// # Returns
//
//   - Trampoline[B, L]: A Trampoline in the "land" state containing the final result
//
// # Example
//
//	// Factorial base case
//	func factorialStep(state State) Trampoline[State, int] {
//	    if state.n <= 1 {
//	        return Land[State](state.acc)  // Computation complete
//	    }
//	    return Bounce[int](State{state.n - 1, state.acc * state.n})
//	}
//
//go:inline
func Land[B, L any](l L) Trampoline[B, L] { _ = "STUB: not implemented"; return nil }

// String implements fmt.Stringer for Trampoline.
//
// Returns a human-readable string representation of the trampoline state.
// For bounce states, returns "Bounce(value)". For land states, returns "Land(value)".
//
// # Returns
//
//   - string: A formatted string representation of the trampoline state
func (t Trampoline[B, L]) String() string { _ = "STUB: not implemented"; return "" }

// Format implements fmt.Formatter for Trampoline.
//
// Supports various formatting verbs for detailed output:
//   - %v: Default format (delegates to String)
//   - %+v: Detailed format with type information
//   - %#v: Go-syntax representation (delegates to GoString)
//   - %s: String format
//   - %q: Quoted string format
//
// # Parameters
//
//   - f: The format state
//   - verb: The formatting verb
func (t Trampoline[B, L]) Format(f fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// %+v: detailed format with type information

// %#v: Go-syntax representation (delegates to GoString)

// %v: default format (delegates to String)

// %s: string format

// %q: quoted string format

// Unknown verb: print with %!verb notation

// GoString implements fmt.GoStringer for Trampoline.
//
// Returns a Go-syntax representation that could be used to recreate the value.
// The output includes the package name, function name, type parameters, and value.
//
// # Returns
//
//   - string: A Go-syntax representation of the trampoline
func (t Trampoline[B, L]) GoString() string { _ = "STUB: not implemented"; return "" }
