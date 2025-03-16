package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Crosshell/kpi3-lab2"
)

var (
  inputExpression = flag.String("e", "", "Expression to compute")
  inputFile       = flag.String("f", "", "Input file containing the expression")
  outputFile      = flag.String("o", "", "Output file for the result")
)

func main() {
  flag.Parse()

  if *inputExpression != "" && *inputFile != "" {
    fmt.Fprintln(os.Stderr, "Error: Provide either -e or -f, not both.")
    os.Exit(1)
  }

  var inputSource io.Reader
  var err error

  if *inputFile != "" {
    file, err := os.Open(*inputFile)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error opening input file:", err)
      os.Exit(1)
    }
    defer file.Close()
    inputSource = file
  } else if *inputExpression != "" {
    inputSource = strings.NewReader(*inputExpression)
  } else {
    inputSource = os.Stdin
  }

  var outputDest *os.File = os.Stdout
  if *outputFile != "" {
    outputDest, err = os.Create(*outputFile)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error creating output file:", err)
      os.Exit(1)
    }
    defer outputDest.Close()
  }

  handler := &lab2.ComputeHandler{
    Input:  inputSource,
    Output: outputDest,
  }

  if err := handler.Compute(); err != nil {
    fmt.Fprintln(os.Stderr, "Computation error:", err)
    os.Exit(1)
  }
}
