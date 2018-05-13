package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// XXX: the exercise suggested that the input can say about its Measurement,
// like 30.3ft, 3.3m, 70.23kg, 20lbs
// TODO: how to capture the floating point number prefix

type Feet float64

type Meter float64

type Measurement interface {
	String() string
}

type Pound float64

type Kilogram float64

func (ft Feet) String() string {
	return fmt.Sprintf("%.3gft", ft)
}
func (m Meter) String() string {
	return fmt.Sprintf("%.3gm", m)
}

func FeetToMeter(ft Feet) Meter {
	return Meter(ft * 0.3)
}

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("An error occurred %q", err)
				os.Exit(2)
			}
			fmt.Printf("The Feet is %s;", Feet(v).String())
			fmt.Printf("then converted to Meter is %s", FeetToMeter(Feet(v)).String())
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			v, err := strconv.ParseFloat(scan.Text(), 64)
			if err != nil {
				fmt.Printf("An error occurred %q", err)
				os.Exit(2)
			}
			fmt.Printf("The Feet is %s;", Feet(v).String())
		}
	}
}
