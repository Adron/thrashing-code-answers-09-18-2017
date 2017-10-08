package main

import (
	"fmt"
	"testing"
)

func TestStorage(test *testing.T) {

	var t float32 = 24    /* 24 hours in a day. */
	var p float32 = 5.851 /* Meters a parking space needs to be */
	var r float32 = 1000  /* Meters of roadway distance in calculation */
	var h float32 = 1.8   /* Cars parked per hour in a single parking spot. */
	var e float32 = 8     /* Empty parking spots per hour for the measured distance. */

	result := Storage(t, p, r, h, e)

	if result != 7037.753 {
		test.Error("Failure to determine correct storage result.")
	}
}

func TestTransport(t *testing.T) {

	var d int = 8000 /* Total cars traveling the roadway for the test time frame. */
	var h int = 1200 /* The number of cars that can travel the roadway per hour. */

	result := Transport(d, h)

	fmt.Println(result)

	if result != 6.6666665 {
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
