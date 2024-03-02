package main

import (
	"bufio"
	"os"
)

// can't use "const" keyword for function definition because its value would vary
var reader = bufio.NewReader(os.Stdin)
