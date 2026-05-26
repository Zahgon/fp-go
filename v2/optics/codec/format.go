package codec

import (
	"fmt"
	"log/slog"
)

// String implements the fmt.Stringer interface for typeImpl.
// It returns the name of the type, which is used for simple string representation.
//
// Example:
//
//	stringType := codec.String()
//	fmt.Println(stringType) // Output: "string"
func (t *typeImpl[A, O, I]) String() string {
	_ = "STUB: not implemented"

	// Format implements the fmt.Formatter interface for typeImpl.
	// It provides custom formatting based on the format verb:
	//   - %s, %v: Returns the type name
	//   - %q: Returns the type name in quotes
	//   - %#v: Returns a detailed Go-syntax representation
	//
	// Example:
	//
	//	intType := codec.Int()
	//	fmt.Printf("%s\n", intType)   // Output: int
	//	fmt.Printf("%q\n", intType)   // Output: "int"
	//	fmt.Printf("%#v\n", intType)  // Output: codec.Type[int, int, any]{name: "int"}
	return ""
}

func (t *typeImpl[A, O, I]) Format(f fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// GoString implements the fmt.GoStringer interface for typeImpl.
// It returns a Go-syntax representation of the type that could be used
// to recreate the type (though not executable due to function values).
//
// This is called when using the %#v format verb with fmt.Printf.
//
// Example:
//
//	stringType := codec.String()
//	fmt.Printf("%#v\n", stringType)
//	// Output: codec.Type[string, string, any]{name: "string"}
func (t *typeImpl[A, O, I]) GoString() string { _ = "STUB: not implemented"; return "" }

// LogValue implements the slog.LogValuer interface for typeImpl.
// It provides structured logging representation of the codec type.
// Returns a slog.Value containing the type information as a group with
// the codec name and type parameters.
//
// This method is called automatically when logging a codec with slog.
//
// Example:
//
//	stringType := codec.String()
//	slog.Info("codec created", "codec", stringType)
//	// Logs: codec={name=string type_a=string type_o=string type_i=interface {}}
func (t *typeImpl[A, O, I]) LogValue() slog.Value {
	_ = "STUB: not implemented"
	return *new(slog.Value)
}

// typeNameOf returns a string representation of the type T.
// It handles the special case where T is 'any' (interface{}).
func typeNameOf[T any]() string { _ = "STUB: not implemented"; return "" }

// Handle the case where %T prints "<nil>" for interface{} types
