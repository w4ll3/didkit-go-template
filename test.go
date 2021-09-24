package main

import (
	"fmt"

	"github.com/spruceid/didkit/tree/feature/go-package/lib/didkit-go-test"
)

func main() {
	fmt.Println(didkit.GenerateEd25519Key())
}
