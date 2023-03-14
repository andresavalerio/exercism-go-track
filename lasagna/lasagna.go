package lasagna


const OvenTime = 40
const timePerLayer = 2


func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}


func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * timePerLayer
}


func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	preparationTime := PreparationTime(numberOfLayers)
	return preparationTime + actualMinutesInOven
}
