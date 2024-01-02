package anyNoteX

import "fmt"

// InterfaceFunctions struct to hold wails runtime for all interface implementations
type InterfaceFunctions struct {
}

// Greet returns a greeting for the given name
func (i *InterfaceFunctions) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
