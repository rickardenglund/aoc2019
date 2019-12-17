package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_goRoutine(t *testing.T) {
	list := splitInput("123")
	assert.Equal(t, []int{1, 2, 3}, list)
}
func Test_getPattern(t *testing.T) {
	assert.Equal(t, 1, getPattern(0, 0))
	assert.Equal(t, 0, getPattern(0, 1))
	assert.Equal(t, -1, getPattern(0, 2))
	assert.Equal(t, 0, getPattern(0, 3))
	assert.Equal(t, 1, getPattern(0, 4))
	//for i := 0; i < 10; i++ {
	//	res := getPattern(1, i)
	//	fmt.Printf("%v, ", res)
	//}

	assert.Equal(t, 0, getPattern(1, 0))
	assert.Equal(t, 1, getPattern(1, 1))
	assert.Equal(t, 1, getPattern(1, 2))
	assert.Equal(t, 0, getPattern(1, 3))

	assert.Equal(t, 1, getPattern(2, 2))
	assert.Equal(t, 1, getPattern(2, 3))
	assert.Equal(t, 1, getPattern(2, 4))
	assert.Equal(t, 0, getPattern(2, 5))

	assert.Equal(t, 1, getPattern(7, 7))
}

func Test_sum(t *testing.T) {
	assert.Equal(t, 4, sum(0, []int{1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 8, sum(1, []int{1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 2, sum(2, []int{1, 2, 3, 4, 5, 6, 7, 8}))

	assert.Equal(t, 2, sum(2, []int{1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 2, sum(2, []int{1, 2, 3, 4, 5, 6, 7, 8}))
}

func Test_fft(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	out := []int{4, 8, 2, 2, 6, 1, 5, 8}
	assert.Equal(t, out, fft(in, 0))

	out2 := []int{3, 4, 0, 4, 0, 4, 3, 8}
	tmp := fft(in, 0)
	assert.Equal(t, out2, fft(tmp, 0))

	in3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	out3 := []int{0, 0, 2, 2, 6, 1, 5, 8}
	assert.Equal(t, out3, fft(in3, 2))

	in4 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	out4 := []int{0, 0, 0, 4, 0, 4, 3, 8}
	offset := 2
	tmp = fft(in4, offset)
	assert.Equal(t, out4, fft(tmp, offset))
}

func Test_fftOffset2(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	out := fft(in, 2)
	assert.Equal(t, "00226158", toStr(out))

	out = fft(out, 2)
	assert.Equal(t, "00040438", toStr(out))

	out = fft(out, 2)
	assert.Equal(t, "00415518", toStr(out))
}

func Test_100fft(t *testing.T) {
	in := "80871224585914546619083218645595"
	assert.Equal(t, "24176176", toStr(fft100(in, 0)))
}

func Test_fftOffset(t *testing.T) {
	in := "80871224585914546619083218645595"
	assert.Equal(t, "17617648", toStr(fft100(in, 2)))
}

func Test_decode(t *testing.T) {
	input := "03036732577212944063491565474664"
	res := decode(input)
	assert.Equal(t, "84462026", res)
}

//func Test_decodep1(t *testing.T) {
//
//	input := "80871224585914546619083218645595"
//	res := decode(input, 1)
//	assert.Equal(t, 24176176, res)
//}
