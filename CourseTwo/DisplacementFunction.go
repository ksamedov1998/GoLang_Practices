package CourseTwo

import (
	"fmt"
	"math"
)

/*
s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity,
and initial displacement. Then the program should prompt the user to enter a value for time and
the program should compute the displacement after the entered time.

GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.

GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values
acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one
float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

*/

type Parameters struct {
	acceleration        float64
	initialVelocity     float64
	initialDisplacement float64
}

func PrintCalculatedAcceleration() {
	var parameters Parameters
	fmt.Print("Enter acceleration :")
	fmt.Scan(&parameters.acceleration)
	fmt.Print("Enter initial velocity :")
	fmt.Scan(&parameters.initialVelocity)
	fmt.Print("Enter initial displacement :")
	fmt.Scan(&parameters.initialDisplacement)
	timeFunction := GenDisplaceFn(parameters.acceleration, parameters.initialVelocity, parameters.initialDisplacement)
	fmt.Printf("Result = %f\n", timeFunction(5))
}

func GenDisplaceFn(acceleration float64, velocity float64, displacement float64) func(float64) float64 {
	timeFunction := func(time float64) float64 {
		return acceleration*math.Pow(time, 2)*0.5 + velocity*time + displacement
	}
	return timeFunction
}
