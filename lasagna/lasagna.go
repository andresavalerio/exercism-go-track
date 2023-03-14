package lasagna


const OvenTime = 40


func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}


func PreparationTime(numberOfLayers int) int {
	const timePerLayer = 2
	return numberOfLayers * timePerLayer
}


func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	preparationTime := PreparationTime(numberOfLayers)
	return preparationTime + actualMinutesInOven
}
