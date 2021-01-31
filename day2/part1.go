package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	paperNeeded := 0
	for scanner.Scan() {
		dims := parseToDimensions(scanner.Text())
		paperNeeded += dims.getRequiredArea()
	}

	fmt.Println(paperNeeded)
}

type Dimensions struct {
	l int
	w int
	h int
}

func (d Dimensions) getRequiredArea() int {
	basicArea := 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l
	extraArea := d.getSmallestArea()

	return basicArea + extraArea
}

func (d Dimensions) getSmallestArea() int {
	ds := []int{d.l * d.w, d.w * d.h, d.h * d.l}
	sort.Ints(ds)

	return ds[0]
}

func parseToDimensions(dimstr string) Dimensions {
	dimlist := strings.Split(dimstr, "x")

	l, _ := strconv.Atoi(dimlist[0])
	w, _ := strconv.Atoi(dimlist[1])
	h, _ := strconv.Atoi(dimlist[2])

	return Dimensions{l, w, h}
}
