package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {

	n := 200
	theta := make([]float64, n)

	theta[0] = 0.0
	theta[1] = 0.2

	dt := 0.04
	g := 9.8
	l := 1.0

	for i := 2; i < n; i++ {
		a := 2.0 * theta[i-1]
		b := (g / l) * math.Pow(dt, 2) * math.Sin(theta[i-1])
		c := theta[i-2]

		theta[i] = a - b - c
	}

	write(theta)
}

func write(res []float64) {
	f, err := os.Create(fmt.Sprintf("./%v.dat", time.Now().Nanosecond()))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i, el := range res {
		_, err = f.WriteString(fmt.Sprintf("%v %v\n", i, el))
		if err != nil {
			panic(err)
		}
	}

	f.Sync()
}
