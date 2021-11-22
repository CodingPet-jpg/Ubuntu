package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
} // 面向对象思想,函数绑定对象

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°C", f)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%g°K", k)
}
