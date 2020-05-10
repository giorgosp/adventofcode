package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func fuel(mass int64) int64 {
	return int64(math.Floor(float64(mass/3))) - 2
}

func fuelfuel(mass int64) int64 {
	newMass := fuel(mass)
	if newMass <= 0 {
		return 0
	}
	return newMass + fuelfuel(newMass)
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	var totalFuel int64
	for s.Scan() {
		mass, err := strconv.Atoi(s.Text())
		checkErr(err)

		totalFuel += fuel(int64(mass)) + fuelfuel(fuel(int64(mass)))
	}

	fmt.Println("total fuel: ", totalFuel)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
