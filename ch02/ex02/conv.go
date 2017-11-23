package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FtToM(ft Feet) Meter { return Meter(ft / 3.28084) }
func MToFt(m Meter) Feet  { return Feet(m * 3.28084) }
 
func PToKg(p Pound) Kilogram  { return Kilogram(p / 2.20462) }
func KgToP(kg Kilogram) Pound { return Pound(kg * 2.20462) }