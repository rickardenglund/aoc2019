package computer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run(t *testing.T) {
	program := []int{1,0,0,0,99}
	assert.Equal(t, []int{2,0,0,0,99}, Run(program))

}

func Test_run2(t *testing.T) {
	program := []int{2,3,0,3,99}
	assert.Equal(t, []int{2,3,0,6,99}, Run(program))
}

func Test_run3(t *testing.T) {
	program := []int{2,4,4,5,99,0}
	assert.Equal(t, []int{2,4,4,5,99,9801}, Run(program))
}
func Test_run4(t *testing.T) {
	program := []int{1,1,1,4,99,5,6,0,99}
	assert.Equal(t, []int{30,1,1,4,2,5,6,0,99}, Run(program))
}