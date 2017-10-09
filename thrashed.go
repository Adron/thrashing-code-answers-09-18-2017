package main

func DeadSpace(t int, e int) float64 {
	return float64(t) + float64(e)
}

func Transport(TotalCars int, CarsPerHour int) float64 {
	return float64(TotalCars) / float64(CarsPerHour)
}

func Storage(tbm float64, ps float64, rd float64, tr float64, es float64) float64 {
	return ((rd / ps) - es) * tr * tbm
}

func GimmeTheSolutionAlready(storage float64, transport float64, deadspace float64) []float64 {
	var quotient = (storage + transport + deadspace) / 100
	var storagePercent = storage / quotient
	var transportPercent = transport / quotient
	var deadspacePercent = deadspace / quotient

	var result []float64
	result = append(result, storagePercent)
	result = append(result, transportPercent)
	result = append(result, deadspacePercent)

	return result
}
