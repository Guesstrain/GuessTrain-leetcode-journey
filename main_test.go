package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNew(t *testing.T) {
	type Teacher struct {
		Name   string
		Age    int
		Gender bool
	}

	teacher := new(Teacher)
	assert.Equal(t, 0, (*teacher).Age)
	assert.Equal(t, "", (*teacher).Name)
	assert.Equal(t, false, (*teacher).Gender)

	intPtr := new(int)
	fmt.Printf("the intPtr pointer address %p", intPtr)
	assert.Equal(t, 0, *intPtr)
}

func TestMake(t *testing.T) {
	numbers := make([]int, 2)
	assert.Equal(t, numbers[0], 0)
	assert.Equal(t, numbers[1], 0)

	numbers[0] = 1
	numbers[1] = 2
	assert.Equal(t, numbers[0], 1)
	assert.Equal(t, numbers[1], 2)
}
