package main

func DeadSpace(t int, e int) int {
	return t + e
}

func Transport(TotalCars int, CarsPerHour int) float32 {
	return float32(TotalCars) / float32(CarsPerHour)
}

func Storage(tbm float32, ps float32, rd float32, tr float32, es float32) float32 {
	return ((rd / ps) - es) * tr * tbm
}
