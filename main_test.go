package awesomeProject

import (
"testing"
"gotest.tools/assert"
)

// Fixture type data struct.
type multiplyTest struct{
	x, y, a int
}

// t testing.T
// Runs test against the Multiply func
func TestMultiply(t *testing.T) {

	t.Parallel()
	var mapD = []multiplyTest{
		{10, 20, 200},
		{100, -20, -2000},
		{-10, -20, 200},
	}

	var val int
	for i := 0; i < len(mapD); i++ {
		var x = mapD[i].x
		var y = mapD[i].y
		val = Multiply(x, y)
		assert.Equal(t,val, mapD[i].a,"Failed on multiply")
	}
}