package main

import (
	"fmt"

	"github.com/spruceid/didkit-go"
)

func main() {
	fmt.Println(didkit.GenerateEd25519Key())
}
