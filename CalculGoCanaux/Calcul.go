package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// for computing a mathematic function
func doCalc(x float64) float64 {

	return x * x
}

// compute area in the given range
func calcSegment(min float64, max float64, segments int, chResultat chan string) {

	var resultat float64

	var pas = (max - min) / float64(segments)

	var x = min

	for i := 0; i < segments; i++ {
		resultat += doCalc(x) * pas
		x += pas
	}

	// convert number to strings
	strMin := strconv.FormatFloat(min,      'E', -1, 64)
	strMax := strconv.FormatFloat(max,      'E', -1, 64)
	strRes := strconv.FormatFloat(resultat, 'E', -1, 64)
	
	chResultat <- "From " + strMin + " to " + strMax + " = " + strRes
}

func main() {


	// read the number of cores of the machine
	nbCores := runtime.NumCPU()

	var min, max float64
	var nbSegments int64
	
	// verify command line arguments
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
	increment  := (max - min) / float64(nbCores)
	chResultat := make(chan string)
	
	max = min + increment

	for i := 0; i < nbCores; i++ {
		
		// creation of child threads for computing
		go calcSegment(min, max, int(nbSegments), chResultat)

		// shifting the range for the next process
		min += increment
		max += increment
	}

	// we fetch and display data
	for i := 0; i < nbCores; i++ {
		fmt.Println(<-chResultat)
	}
}