package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
  BoilingC Celsius = 100
  FreezingC Celsius = 0
  AbsoluteC Celsius = -273.15
)

func CToF(c Celsius) Fahrenheit {
  return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
  return Celsius((f-32) * 5/9)
}

func (c Celsius) String() string {
  return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
  return fmt.Sprintf("%g F", f)
}
// For every type T, there is a corrsponding conversion operation T(x) that converts the value x of type T.
