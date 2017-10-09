package main

import (
	"testing"
)

func TestStorage(test *testing.T) {

	var t float64 = 24    /* 24 hours in a day. */
	var p float64 = 5.851 /* Meters a parking space needs to be */
	var r float64 = 1000  /* Meters of roadway distance in calculation */
	var h float64 = 1.8   /* Cars parked per hour in a single parking spot. */
	var e float64 = 8     /* Empty parking spots per hour for the measured distance. */

	if Storage(t, p, r, h, e) != 7037.753272944796 {
		test.Error("Failure to determine correct storage result.")
	}
}

func TestTransport(t *testing.T) {

	var d int = 8000 /* Total cars traveling the roadway for the test time frame. */
	var h int = 1200 /* The number of cars that can travel the roadway per hour. */

	result := Transport(d, h)

	if result != 6.666666666666667 {
		t.Error("Failure to determine correct transportation throughput.")
	}
}

func TestDeadSpace(test *testing.T) {
	var t int = 23 /* Transport or transporting (the result of the "Transport" equation). */
	var e int = 2  /* Empty parking sports per hour. This value is what would be the same \
	as the E value in the Storage function. */

	result := DeadSpace(t, e)
	if result != 25 {
		test.Error("Nope that isn't the correct amount of dead space.")
	}
}

func TestGimmeTheStoragePercent(t *testing.T) {
	result := GimmeTheSolutionAlready(50, 75, 75)

	if result[0] != 25 {
		t.Error("The storage percentage is not correct.")
	}
}

func TestGimmeTheTransportPercent(t *testing.T) {
	result := GimmeTheSolutionAlready(50, 75, 75)

	if result[1] != 37.5 {
		t.Error("The transport percentage is not correct.")
	}
}

func TestGimmeTheDeadspacePercent(t *testing.T) {
	result := GimmeTheSolutionAlready(50, 75, 75)

	if result[2] != 37.5 {
		t.Error("The dead space percentage is not correct.")
	}
}
