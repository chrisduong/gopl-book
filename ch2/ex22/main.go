package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// XXX: the exercise suggested that the input can say about its Measurement,
// like 30.3ft, 3.3m, 70.23kg, 20lbs
// TODO: how to capture the floating point number prefix

type Feet float64

type Meter float64

// Measurement interface to conditionally print corresponding Unit type
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

// makeMeasurement return the specific measurement
func makeMeasurement(f float64, unit string) (Measurement, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "m":
		return Meter(f), nil
	case "ft":
		return Feet(f), nil
	default:
		return nil, fmt.Errorf("Unexpected unit %v", unit)
	}
}

// analyseInput use Regex to capture Value and Unit kind if it is input correctly
func analyseInput(s string) (float64, string, error) {
	// Can you https://regex101.com to quickly golang regexp
	re := regexp.MustCompile(`(^-?\d+(?:\.\d+)?)([A-Za-z]+)`)
	match := re.FindStringSubmatch(s)
	if match == nil {
		log.Fatalf("Expecting <value>.<unit>, got %q", s)
	}
	v, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		log.Fatalf("%v isn't a number.", match[1])
	}
	if match[2] == "" {
		log.Fatalf("No unit specified.")
	}
	unit := match[2]

	return v, unit, nil
}

// FeetToMeter converts Feet to Meter
func FeetToMeter(ft Feet) Meter {
	return Meter(ft * 0.3)
}

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {

			v, unit, err := analyseInput(arg)
			if err != nil {
				fmt.Printf("An Input error occurred %q", err)
				os.Exit(2)
			}

			m, err := makeMeasurement(v, unit)
			if err != nil {
				fmt.Printf("%q", err)
				os.Exit(2)
			}
			fmt.Printf("You input this measurement %s", m.String())
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			s := scan.Text()
			v, unit, err := analyseInput(s)
			if err != nil {
				fmt.Printf("An Input error occurred %q", err)
				os.Exit(2)
			}

			m, err := makeMeasurement(v, unit)
			if err != nil {
				fmt.Printf("%q", err)
				os.Exit(2)
			}
			fmt.Printf("You input this measurement %s", m.String())
		}
	}
}
