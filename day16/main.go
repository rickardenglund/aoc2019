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
	fmt.Printf("part1: %-10v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

func part1() int {
	return fft100(input)
}

func part2() int {
	return decode(input)
}
func decode(input string) int {
	sb := strings.Builder{}
	for i := 0; i < 10_000; i++ {
		sb.WriteString(input)
	}

	ints := splitInput(sb.String())

	for i := 0; i < 100; i++ {
		//fmt.Printf("%v", i)
		ints = fft(ints)

	}
	fmt.Printf("\n%+v\n", ints)

	return len(ints)

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
func genPattern(n int) []int {
	p := []int{0, 1, 0, -1}
	out := []int{}
	for _, val := range p {
		for i := 0; i < n; i++ {
			out = append(out, val)
		}
	}
	return out
}
func fft100(in string) int {
	signal := splitInput(in)
	for i := 0; i < 100; i++ {
		signal = fft(signal)
	}

	res := toInt(signal, 8)
	return res

}

func toInt(signal []int, n int) int {
	str := ""
	for i := 0; i < len(signal); i++ {
		str += strconv.Itoa(signal[i])
	}
	res, _ := strconv.Atoi(str[:n])
	return res
}

func fftOffset(input []int, offset int) int {
	return do(genPattern(offset+1), input)
}

func fft(input []int) []int {
	var out = make([]int, len(input))

	for i := 0; i < len(input); i++ {
		out[i] = do(genPattern(i+1), input)
	}
	return out
}

func do(pattern []int, input []int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i] * pattern[(i+1)%len(pattern)]
	}
	return intmath.Abs(sum) % 10
}
