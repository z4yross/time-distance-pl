package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"math"
)

func main() {
	jvC := exec.Command("javac", "scripts/jv/Fibonacci.java")
	csC := exec.Command("mcs", "scripts/cs/fibonacci.cs")
	cppC := exec.Command("g++", "-o", "scripts/cpp/fibonacci", "scripts/cpp/fibonacci.cpp")

	jvE, errJ := jvC.Output()
	csE, errS := csC.Output()
	cppE, errC := cppC.Output()

	if errJ != nil {
		log.Fatal(errJ)
	} else 	if errS != nil {
		log.Fatal(errS)
	} else 	if errC != nil {
		log.Fatal(errC)
	} 

	ttJ := 0.0
	ttS := 0.0
	ttC := 0.0

	cicles := 200

	for i := 0; i < cicles; i++{
		jvC = exec.Command("java", "scripts.jv.Main")
		csC = exec.Command("./scripts/cs/fibonacci.exe")
		cppC = exec.Command("./scripts/cpp/fibonacci")

		jvE, errJ = jvC.Output()
		csE, errS = csC.Output()
		cppE, errC = cppC.Output()

		if errJ != nil {
			log.Fatal(errJ)
		} else 	if errS != nil {
			log.Fatal(errS)
		} else 	if errC != nil {
			log.Fatal(errC)
		} 
		
		tJ, errTJ := strconv.ParseFloat(string(jvE), 64)
		tS, errTS := strconv.ParseFloat(string(csE), 64)
		tC, errTC := strconv.ParseFloat(string(cppE), 64)

		if errTJ != nil {
			log.Fatal(errTJ)
		} else 	if errTS != nil {
			log.Fatal(errTS)
		} else 	if errTC != nil {
			log.Fatal(errTC)
		} 	

		ttJ += tJ
		ttS += tS
		ttC += tC

		fmt.Print("\r")
		fmt.Print("Loading -> ",math.Ceil(float64(i) / float64(cicles) * 100),"%")
	}

	ttJ /= float64(cicles)
	ttS /= float64(cicles)
	ttC /= float64(cicles)	

	fmt.Println("\ndistance(java, c#) =", distance(ttJ, ttS))
	fmt.Println("distance(java, cpp) =", distance(ttJ, ttC))
	fmt.Println("distance(c#, cpp) =", distance(ttS, ttC))	
}

func distance(a float64, b float64) float64{
	return math.Round(math.Abs(a - b) * 10000) / 10000
}