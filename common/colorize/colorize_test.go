package colorize

import (
	"fmt"
	"testing"

	"github.com/MochamadAkbar/ordent-test/common/constants"
	"github.com/stretchr/testify/assert"
)

type typeMessageColorized struct {
	name, request, color string
}

func TestMessageColorize(t *testing.T) {
	tests := []typeMessageColorized{
		{
			name:    "MessageColorized('Green', Green)",
			request: "Green",
			color:   constants.Green,
		},
		{
			name:    "MessageColorized('Orange', Orange)",
			request: "Orange",
			color:   constants.Orange,
		},
		{
			name:    "MessageColorized('Blue', Blue)",
			request: "Blue",
			color:   constants.Blue,
		},
		{
			name:    "MessageColorized('Red', Red)",
			request: "Red",
			color:   constants.Red,
		},
		{
			name:    "MessageColorized('DefaultColor', DefaultColor)",
			request: "DefaultColor",
			color:   constants.DefaultColor,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := fmt.Sprintf("%v%v%v", test.color, test.request, constants.DefaultColor)
			result := MessageColorized(test.color, test.request)

			assert.Equal(t, expected, result)
		})
	}
}
