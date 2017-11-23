package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
 
func (p Pound) String() string     { return fmt.Sprintf("%glb", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }