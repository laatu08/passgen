package main

import (
	"flag"
	"fmt"
	"os"

	"passgen/internal/generator"
)

func main() {
	length := flag.Int("length", 12, "Password length")
	upper := flag.Bool("upper", false, "Include uppercase letters")
	lower := flag.Bool("lower", false, "Include lowercase letters")
	digits := flag.Bool("digits", false, "Include digits")
	symbols := flag.Bool("symbols", false, "Include symbols")
	flag.Parse()

	password, err := generator.Generate(*length, *upper, *lower, *digits, *symbols)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Generated Password:")
	fmt.Println(password)
}
