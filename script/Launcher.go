package main 

import (
    "os/exec"
    "fmt"
    "math"
    "strconv"
)



func main() {
    
    tabCores := []int{1, 2, 4, 8, 16}
	
	cmdName := "/usr/bin/time"
	goArgs := []string{"--format", 
						"%e", 
						"./CalculGoCanaux", 
						"-1", 
						"-100", 
						"100", 
						"0"}
	
	cppArgs := []string{"--format", 
						"%e", 
						"./Calcul.cxx", 
						"-1", 
						"-100", 
						"100", 
						"0"}
	
	// testing GO
	for _, nbCores := range tabCores {
		
		fmt.Printf("Using %d cores\n", nbCores)
		goArgs[3] = strconv.Itoa(nbCores)
		cppArgs[3] = goArgs[3]

		var j float64
		for j = 2. ; j < math.Pow(2., 4096.) ; j = math.Pow(j, 2.) {
			
			goArgs[6] = strconv.FormatInt(int64(j), 10) 
			cppArgs[6] = goArgs[6]

			fmt.Printf("\tTrying with %e rectangles...\n", j)
			if cmdOut, err := exec.Command(cmdName, goArgs...).CombinedOutput(); err == nil {	
				fmt.Printf("GO,%d,%d,%s", nbCores, nbRectangles, string(cmdOut))
			} else {
				fmt.Println(err)
			}
			
			if cmdOut, err := exec.Command(cmdName, cppArgs...).CombinedOutput(); err == nil {	
				fmt.Printf("C++,%d,%d,%s", nbCores, nbRectangles, string(cmdOut))
			}
		}
	}
	
}
