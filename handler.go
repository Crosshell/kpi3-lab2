package lab2

import (
  "fmt"
  "io"
  "strings"
)

type ComputeHandler struct {
  Input  io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  buf := new(strings.Builder)
  _, err := io.Copy(buf, ch.Input)
  if err != nil {
    return err
  }

  result, err := EvaluatePrefixExpression(buf.String())
  if err != nil {
    return err
  }

  _, err = fmt.Fprintf(ch.Output, "%f\n", result)
  return err
}
