// Package weather can be used to forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string
// CurrentLocation stores the city in Goblinocus that the forecast referes to.
var CurrentLocation string

// Forecast returns the city and the current weather conditions of this city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
