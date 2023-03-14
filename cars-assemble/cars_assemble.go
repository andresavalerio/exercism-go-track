package cars

const minutesInOneHour = 60.0
const costPerTenCars = 95_000
const costPerIndividualCar = 10_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRateInPercetage := successRate * 0.01
	return float64(productionRate) * successRateInPercetage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsInAnHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	workingCarsInAMinute := workingCarsInAnHour / minutesInOneHour
	return int(workingCarsInAMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	individualCars := carsCount % 10
	totalCost := (groupsOfTen * costPerTenCars) + (individualCars * costPerIndividualCar)
	return uint(totalCost)
}
