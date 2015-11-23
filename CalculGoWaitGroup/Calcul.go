package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"

)

// pour calculer la fonction que l'on veux
func doCalc(x float64) float64 {

	return x * x
}

func calcSegment(min, max float64, segments int, wg *sync.WaitGroup) {

	var pas = (max - min) / float64(segments)
	var resultat float64

	var x = min

	for i := 0; i < int(segments); i++ {
		resultat += doCalc(x) * pas
		x += pas
	}

	fmt.Printf("De %e à %e = %e\n", min, max, resultat)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	// read the number of cores of the machine
	nbCores := runtime.NumCPU()

	// verify command line arguments
	var min, max float64
	var nbSegments int64

	if len(os.Args) == 4 {

		min, _        = strconv.ParseFloat(os.Args[1], 64) 
		max, _        = strconv.ParseFloat(os.Args[2], 64)
		nbSegments, _ = strconv.ParseInt(os.Args[3], 10, 64)

	} else {
		fmt.Printf("Usage : %s <min> <max> <nbSegments>\n", os.Args[0])
		os.Exit(1)
	}

	if min > max {
		fmt.Println("Wrong input, maximum is inferior to minimum...")
		os.Exit(2)
	}

	// compute the increment, to distribute the same part to all the threads
	increment := (max - min) / float64(nbCores)
	
	max = min + increment

	for i := 0; i < nbCores; i++ {
		
		go calcSegment(min, max, int(nbSegments), &wg)
		wg.Add(1)

		min += increment
		max += increment
	}

	// waiting all goroutines to terminate
	wg.Wait()
}
