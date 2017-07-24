package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the StringParse CLI")
	// const CLR_0 = "\x1b[30;1m"
	// const CLR_R = "\x1b[31;1m"
	// const CLR_G = "\x1b[32;1m"
	// const CLR_Y = "\x1b[33;1m"
	// const CLR_B = "\x1b[34;1m"
	// const CLR_M = "\x1b[35;1m"
	// const CLR_C = "\x1b[36;1m"
	// const CLR_W = "\x1b[37;1m"
	// const CLR_N = "\x1b[0m"
	// `arguments` are strings that follow a CLI command
	// for Go, strings without dashes that are not user input are arguments
	// `options` are strings with dashes, follow user input, modify command
	// `flags` are boolean options, do not take user input
	// for Go, strings beginning with dashes (single/double) are considered flags

	// examples of string flag syntax
	// -example
	// -example=text
	// -example text (Non-boolean flags only)
	textPtr := flag.String("text", "", "Text to parse. (**required**)")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};. (**required**)")
	// boolean flag will parse to `true` if flag set and no user input provided
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	// 1. Go automatically checks for user input for all flags that require it
	// 2. Go automatically checks that all flags provided by the user are known
	// if either 1. or 2. are not met, the help text is automatically displayed

	// after parsing, arguments available as slice `flag.Args()` (flag.Arg(i))
	flag.Parse()

	// check if the required text flag is provided, exit if missing
	if *textPtr == "" {
		fmt.Println("\x1b[31;1m* One or more required flags were not provided\x1b[0m")
		fmt.Println("\x1b[33;1m* Please review the instructions below\x1b[0m")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}
