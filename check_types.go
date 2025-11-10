package main

import (
	"fmt"
	"go-reloaded/pkg/tokenizer"
)

func main() {
	fmt.Println("Word =", tokenizer.Word)
	fmt.Println("Marker =", tokenizer.Marker)
	fmt.Println("Punct =", tokenizer.Punct)
	fmt.Println("Space =", tokenizer.Space)
}
