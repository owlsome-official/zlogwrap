package zlogwrap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testCaseToString struct {
	name        string
	inputParam1 string
	inputParam2 string
	expected    string
}

func TestToString(t *testing.T) {
	tests := []testCaseToString{
		{
			name:        "Empty Case",
			inputParam1: "",
			inputParam2: "",
			expected:    " ", // note: this join between ("" and "") = " "
		},
		{
			name:        "Only param-1 Case",
			inputParam1: "1",
			inputParam2: "",
			expected:    "1 ",
		},
		{
			name:        "Only param-2 Case",
			inputParam1: "",
			inputParam2: "2",
			expected:    " 2",
		},
		{
			name:        "All params Case",
			inputParam1: "1",
			inputParam2: "2",
			expected:    "1 2",
		},
		{
			name:        "All input Case",
			inputParam1: "1",
			inputParam2: "2",
			expected:    "1 2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputParam1 := test.inputParam1
			inputParam2 := test.inputParam2
			expected := test.expected

			actual := toString(inputParam1, inputParam2)

			require.Equal(t, expected, actual)
		})
	}
}
