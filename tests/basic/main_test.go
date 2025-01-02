package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), output %d, actual = %d ", input, output, actual)
	// }

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
}
