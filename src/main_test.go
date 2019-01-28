package main

import (
	"testing"
	"gotest.tools/assert"
	"fmt"
)

type multiplyTest struct{
	x, y, a int
}


func TestMultiply(t *testing.T) {
	println(Multiply(100,10))
	t.Parallel()
	var mapD = []multiplyTest{
		{10, 20, 200},
		{10, 20, 200},
		{10, 20, 200},
	}


	var val int
	for i := 0; i < len(mapD); i++ {
		println(len(mapD))
		val = Multiply(mapD[i].x, mapD[i].y)
		assert.Equal(t,val, mapD[i].a,fmt.Printf("Multiply returned an unexpected value %i", val))
	}

}

