package main

import (
	intmath "aoc2019/mymath"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var gui bool //nolint:unused
const input = "59755896917240436883590128801944128314960209697748772345812613779993681653921392130717892227131006192013685880745266526841332344702777305618883690373009336723473576156891364433286347884341961199051928996407043083548530093856815242033836083385939123450194798886212218010265373470007419214532232070451413688761272161702869979111131739824016812416524959294631126604590525290614379571194343492489744116326306020911208862544356883420805148475867290136336455908593094711599372850605375386612760951870928631855149794159903638892258493374678363533942710253713596745816693277358122032544598918296670821584532099850685820371134731741105889842092969953797293495"

func main() {
	guiPtr := flag.Bool("gui", false, "Add --gui flag to enable graphics")
	flag.Parse()
	gui = *guiPtr

	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %v in %v\n", p1, time.Since(start))

	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

func part1() interface{} {
	return toStr(fft100(input, 0)[0:8])
}

func part2() interface{} {
	res := decode(input)
	fmt.Printf("%v\n", res)
	return res
}

func decode(input string) string {
	input = times10k(input)
	offset, _ := strconv.Atoi(input[0:7])

	res := fft100(input, offset)

	return toStr(res)
}

func times10k(input string) string {
	sb := strings.Builder{}
	for i := 0; i < 10_000; i++ {
		sb.WriteString(input)
	}
	return sb.String()
}

func splitInput(in string) []int {
	numbers := strings.Split(in, "")
	out := make([]int, len(numbers))
	for i, n := range numbers {
		res, _ := strconv.Atoi(n)
		out[i] = res
	}
	return out
}

func fft100(in string, offset int) []int {
	signal := splitInput(in)
	for i := 0; i < 100; i++ {
		start := time.Now()
		signal = fft(signal, offset)
		fmt.Printf("%v: %v\n", i, time.Since(start))
	}

	return signal[offset : offset+8]

}
func toStr(signal []int) string {
	str := ""
	for i := 0; i < len(signal); i++ {
		str += strconv.Itoa(signal[i])
	}
	return str
}

func fft(input []int, offset int) []int {
	var out = make([]int, len(input))

	for i := offset; i < len(input); i++ {
		out[i] = sum(i, input)
	}
	return out
}

func sum(series int, input []int) int {
	sum := 0
	for i := series; i < len(input); i++ {
		k := getPattern(series, i)
		sum += input[i] * k
	}
	return intmath.Abs(sum) % 10
}

var p = [4]int{0, 1, 0, -1}

func getPattern(series, i int) int {
	if true {
		return p[i%4]
	}
	if series == 0 {
		return p[(i+1)%4]
	}
	return p[((i+1)/(series+1))%4]
}
