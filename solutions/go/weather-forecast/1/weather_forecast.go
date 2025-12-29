// Package weather provides forecast based on current weather condition of various cities.
package weather

var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string
    // CurrentLocation represents the current location for which the current weather condition has to be checked.
	CurrentLocation  string
)

// Forecast returns statement with current location and its respective current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
