package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	passwd "github.com/dann1/passwdgen/src/passwd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: passwdgen <length> [--uppercase] [--numbers] [--specials]")
		os.Exit(1)
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid length provided. Length should be an integer.")
		os.Exit(1)
	}

	uppercaseFlag := flag.Bool("uppercase", false, "Include uppercase letters")
	numbersFlag := flag.Bool("numbers", false, "Include numbers")
	specialsFlag := flag.Bool("specials", false, "Include special characters")

	flag.CommandLine.Parse(os.Args[2:])

	password := passwd.Generate(length, *uppercaseFlag, *numbersFlag, *specialsFlag)
	fmt.Println(password)
}
