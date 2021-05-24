package increment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetNumber(t *testing.T) {
	inc := New(5)
	inc.Increment()

	assert.Equal(t, uint(1), inc.GetNumber(), "Should be initialized with 0")
}

func Test_IncrementValue(t *testing.T) {
	inc := New(10)
	inc.Increment()

	assert.Equal(t, uint(1), inc.i, "Should be incremented from 0 to 1")
}

func Test_SetMaximumValue(t *testing.T) {
	max := uint(10)

	inc := New(0)
	inc.SetMaximumValue(max)

	assert.Equal(t, inc.max, max, "Should be equal")
}

func Test_SetValueToNull(t *testing.T) {
	max := uint(1)

	inc := New(max)
	inc.Increment()

	assert.Equal(t, uint(0), inc.i)
}

func Test_SetMaximumValueToNull(t *testing.T) {
	inc := New(1)
	inc.Increment()

	inc.SetMaximumValue(0)

	assert.Equal(t, uint(0), inc.i)
}
