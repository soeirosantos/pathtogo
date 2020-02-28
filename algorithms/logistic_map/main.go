package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	linearMap()
	logisticMap()
}

// evolution of the number of bacteria assuming
// there aren't constraints in the growth (not real)
func linearMap() {
	dt := 1.0
	alpha := 0.01
	n := 1000
	N := make([]float64, n)

	N[0] = 0.0001

	for i := 1; i < n; i++ {
		N[i] = N[i-1] * (1 + alpha*dt)
	}

	write(N)
}

// evolution of the density of bacterias x(t) = N(t) / Nm
// assumes that resources in the lamina are limited and
// the number of bacterias will stabilize at some point
func logisticMap() {
	lambda := 1.01
	n := 1000
	x := make([]float64, n)

	x[0] = 0.0001

	for i := 1; i < n; i++ {
		x[i] = lambda * x[i-1] * (1 - x[i-1])
	}

	write(x)

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
