package main

import (
	"fmt"
	"log"
	"time"

	"github.com/barbucatalinn/standard-deviation/internal"
)

func main() {
	app, err := internal.New("./data/sample.csv")
	if err != nil {
		log.Fatal(err)
	}

	start1 := time.Now()
	stdDev1 := app.CalculateStandardDeviationV1()
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	stdDev2 := app.CalculateStandardDeviationV2()
	elapsed2 := time.Since(start2)

	fmt.Println("Standard deviation v1:", fmt.Sprintf("%.2f", stdDev1))
	fmt.Println("Standard deviation v2:", fmt.Sprintf("%.2f", stdDev2))

	fmt.Printf("Standard deviation v1 took %s\n", elapsed1)
	fmt.Printf("Standard deviation v2 took %s\n", elapsed2)
}
