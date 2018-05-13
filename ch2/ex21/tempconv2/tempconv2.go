// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package tempconv2

import "fmt"

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

const (
	// AbsoluteZeroC of Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC of Celsius
	FreezingC Celsius = 0
	// BoilingC of Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
