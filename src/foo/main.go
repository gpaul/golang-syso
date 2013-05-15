// main.go
package main

import "fmt"

var Awesome []byte // must match the name in data.S, and must be global, can be unexported
func main() {
	fmt.Printf("len=%d, cap=%d\n", len(Awesome), cap(Awesome))
	fmt.Printf("%v\n", Awesome)
}
