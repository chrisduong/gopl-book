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

// isValidUnit will check if the string name is a valid unit
func isValidUnit(s string) bool {
	switch s {
	case
		"m",
		"ft":
		return true
	default:
		return false
	}
}

// mToFeet convert Meter to Feet
func mToFeet(m Meter) Feet {
	return Feet(float64(m) * 3.28)
}

// makeMeasurement return the specific measurement
func makeMeasurement(f float64, unit string) Measurement {
	unit = strings.ToLower(unit)
	switch unit {
	case "m":
		return Meter(f)
	case "ft":
		return Feet(f)
	default:
		return nil
	}
}

// analyseInput use Regex to capture Value and Unit kind if it is input correctly
func analyseInput(s string) (float64, string) {
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
	if !isValidUnit(unit) {
		log.Fatalf("Not supported unit %q", unit)
	}

	return v, unit
}

// printMeasurement print what was input, and its converted number.
func printMeasurement(m Measurement) {
	meter, ok := m.(Meter)
	if ok {
		v1 := meter.String()
		v2 := mToFeet(m.(Meter))
		fmt.Printf("You input meter: %s", v1)
		fmt.Println()
		fmt.Printf("%s equal to %s", v1, v2)
	}
}

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {

			v, unit := analyseInput(arg)
			m := makeMeasurement(v, unit)
			printMeasurement(m)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			s := scan.Text()
			v, unit := analyseInput(s)
			m := makeMeasurement(v, unit)
			printMeasurement(m)
		}
	}
}
