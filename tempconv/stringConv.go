package tempconv
import "fmt"

func (c Celsius) String() string{
	 return fmt.Sprintf("%.4f°C", c) 
}
func (f Fahrenheit) String() string { 
	return fmt.Sprintf("%.4f°F", f) 
}