package main

import (
	"fmt"

	"github.com/boymeetsblockchain/go-jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}
func main() {
	fmt.Println("Hello world 3")
}
