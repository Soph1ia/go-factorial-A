package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"math/big"
	"time"
)

type MyEvent struct {
	Message int     `json:"Message"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	benchmark(event.Message)
	return MyResponse{Message: fmt.Sprintf(" factorial took %v ops/ns", event.Message)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}


/**
Method : Benchmark

This method gets the time taken to execute the factorial 40 times.
In total it loops 80 times.
It takes the last 20 execution times.
Gets the average time
Calculates the throughput as time / 40

Prints out the throughput.

returns: none

*/
func benchmark( number int) float64 {
	listofTime := [20]int64{}

	for j := 0; j < 40; j++ {
		start := time.Now().UnixNano()
		// Loop 40 times.
		for i := 0; i <= 40; i++ {
			factorial(number)
		}
		// End time
		end := time.Now().UnixNano()
		// Results
		if j > 20 {
			difference := end - start
			listofTime[j-20] = difference
		}
	}
	// Average Time
	sum := int64(0)
	for i := 0; i < len(listofTime); i++ {
		// adding the values of
		// array to the variable sum
		sum += listofTime[i]
	}
	// avg to find the average
	avg := (float64(sum)) / (float64(len(listofTime)))

	// Throughput Rate
	throughput := avg / 40

	// Response
	return throughput
}

/**
Method: Factorial

Calculates the factorial of the number provided

Returns: pointer to big int
*/
func factorial(n int) *big.Int {
	factVal := big.NewInt(1)
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			//factVal *= uint64(i) // mismatched types int64 and int
			factVal = factVal.Mul(factVal, big.NewInt(int64(i)))
		}
	}
	return factVal
}
