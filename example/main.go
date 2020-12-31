package main

import (
	"fmt"
	"math"
	"time"

	"github.com/epes/etime"
)

func main() {
	ctr := 0

	fn := func() bool {
		ctr++

		f := int(math.Floor(math.Sqrt(float64(ctr))))

		res := f%2 == 0

		fmt.Printf("[%s] %d %d %v\n", time.Now().Format("03:04:05"), ctr, f, res)

		return res
	}

	etime.NewHotColdTicker(5*time.Second, 1*time.Second, 5, fn)
}
