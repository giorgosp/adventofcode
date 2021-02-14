package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p point) Add(p1 point) point {
	p.x = p.x + p1.x
	p.y = p.y + p1.y
	return p
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type segment struct {
	p1, p2 point
}

// swept keeps segments whose left point has been
// passed by the sweep line.
type swept []*segment

func (a active) Insert() {}
func (a active) Remove() {}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	w1Str, w2Str, err := parseWires(f)
	if err != io.EOF {
		check(err)
	}

	// add points of first wire in a map
	wire1Path := make(map[string]point)
	prev := point{0, 0}
	err = visitWire(w1Str, func(p point) {
		next := prev.Add(p)
		fmt.Println(next)
		wire1Path[next.String()] = next
		prev = next
	})
	check(err)

	fmt.Println("===")

	// for the points of second wire, find overlaps
	var overlaps []point
	prev = point{0, 0}
	err = visitWire(w2Str, func(p point) {
		next := prev.Add(p)
		fmt.Println(next)
		prev = next
		if _, ok := wire1Path[next.String()]; ok {
			overlaps = append(overlaps, next)
		}
	})
	check(err)
	if len(overlaps) == 0 {
		panic("zero overlaps")
	}

	// find the smaller mahnattan distance between 0,0 and the overlaps
	min := math.Inf(1)
	p0 := point{0, 0}
	for _, p := range overlaps {
		m := manhattan(p0, p)
		if m < min {
			min = m
		}
	}
	fmt.Println("manhattan", min)
}

func manhattan(p1, p2 point) float64 {
	return math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y))
}

func visitWire(wire string, handler func(point)) error {
	br := bufio.NewReader(strings.NewReader(wire))
	for {
		token, err := br.ReadString(',')
		if token[len(token)-1] == ',' {
			token = token[:len(token)-1]
		}
		if err != nil && err != io.EOF {
			return err
		}
		p := parsePoint(token)
		handler(p)
		if err == io.EOF {
			return nil
		}
	}
}

func parsePoint(token string) point {
	dir := token[0]
	n, err := strconv.Atoi(token[1:])
	check(err)

	switch {
	case dir == 'U':
		return point{0, n}
	case dir == 'D':
		return point{0, -n}
	case dir == 'R':
		return point{n, 0}
	case dir == 'L':
		return point{-n, 0}
	default:
		panic("unknown direction")
	}
}

func parseWires(r io.Reader) (string, string, error) {
	br := bufio.NewReader(r)
	w1, err := br.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	if w1[len(w1)-1] == '\n' {
		w1 = w1[:len(w1)-1]
	}
	w2, err := br.ReadString('\n')
	if err == io.EOF {
		err = nil
	}
	if w2[len(w2)-1] == '\n' {
		w1 = w2[:len(w1)]
	}
	return w1, w2, err
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
