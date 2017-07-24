package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the StringParse CLI")
	// `arguments` are strings that follow a CLI command
	// for Go, strings without dashes that are not user input are arguments
	// `options` are strings with dashes, follow user input, modify command
	// `flags` are boolean options, do not take user input
	// for Go, strings beginning with dashes (single/double) are considered flags

	// examples of string flag syntax
	// -example
	// -example=text
	// -example text (Non-boolean flags only)
	textPtr := flag.String("text", "", "Text to parse.")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	// boolean flag will parse to `true` if flag set and no user input provided
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")

	// after parsing, arguments available as slice `flag.Args()` (flag.Arg(i))
	flag.Parse()

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}
