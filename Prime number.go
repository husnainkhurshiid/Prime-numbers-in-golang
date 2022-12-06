package main

import (
	"fmt"
	"math"
	"time"
)

func CheckPrimeNumber(start, end int) {
	primeSum := 0
	totalNumber := 0
	if start < 2 || end < 2 {
		fmt.Print("Numbers Should Be Greater Than 2")
		return
	}
	for start < end {
		isPrime := true
		sqrt := int(math.Sqrt(float64(start)))
		for i := 2; i <= sqrt; i++ {
			if start%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%v ", start)
			primeSum += start
			totalNumber += 1
		}
		start++
	}
	fmt.Println()
	fmt.Printf("Sum Of Numbers : %d \n", primeSum)
	fmt.Printf("Total Numbers : %d \n", totalNumber)
}

func main() {
	start := time.Now()
	CheckPrimeNumber(2, 1000)
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Printf("Execution Time %s", elapsed)
}
