package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	intcode, err := loadIntcode("input.txt")
	checkErr(err)

	target := 19690720
	noun, verb := findInputs(intcode, target)
	fmt.Println(100*noun+verb, noun, verb)
}

func findInputs(intcode []int, target int) (int, int) {
	program := make([]int, len(intcode))
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(program, intcode)
			result := runIntcode(program, noun, verb)
			if result == target {
				return noun, verb
			}
		}
	}
	panic("didn't find inputs")
}

func runIntcode(intcode []int, in1, in2 int) int {
	intcode[1] = in1
	intcode[2] = in2

	for i := 0; i < len(intcode); i += 4 {
		op := intcode[i]

		if op == 99 {
			break
		}

		inpos1 := intcode[i+1]
		inpos2 := intcode[i+2]
		outpos := intcode[i+3]

		if op == 1 {
			intcode[outpos] = intcode[inpos1] + intcode[inpos2]
		} else if op == 2 {
			intcode[outpos] = intcode[inpos1] * intcode[inpos2]
		} else {
			panic("unknown opcode " + string(op))
		}
	}

	return intcode[0]
}

func loadIntcode(filename string) ([]int, error) {
	var intcode []int

	f, err := os.Open(filename)
	if err != nil {
		return intcode, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		op, readErr := r.ReadString(',')
		if readErr != nil && readErr != io.EOF {
			return intcode, readErr
		}

		// Check op read. we may get a read even if EOF is returned, so EOF is checked later.
		// Also remove the included delimiter.
		if op[len(op)-1] == ',' {
			op = op[:len(op)-1]
		}

		opInt, err := strconv.Atoi(op)
		if err != nil {
			return intcode, readErr
		}
		intcode = append(intcode, opInt)

		if readErr == io.EOF {
			return intcode, nil
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
