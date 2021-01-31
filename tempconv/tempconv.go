package tempconv

type Celsius float64
type Farenheit float64

const (
  BoilingC Celsius = 100
  FreezingC Celsius = 0
  AbsoluteC Celsius = -273.15
)

func CToF(c Celsius) Farenheit {
  return Farenheit(c*9/5 + 32)
}

func FToC(f Farenheit) Celsius {
  return Celsius((f-32) * 5/9)
}

// For every type T, there is a corrsponding conversion operation T(x) that converts the value x of type T.
