package zlogwrap

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

type testCaseDefault struct {
	name     string
	input    interface{}
	expected interface{}
}

func TestConfigDefault(t *testing.T) {
	tests := []testCaseDefault{
		{
			name:     "Empty Input",
			input:    nil,
			expected: &ConfigDefault,
		},
		{
			name:     "With an Input",
			input:    Config{Hidden: true, ServiceName: "test"},
			expected: &Config{Hidden: true, ServiceName: "test", Logger: log.Logger},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != nil {
				actual := configDefault(test.input.(Config))
				require.Equal(t, test.expected.(*Config), actual)
				return
			}
			actual := configDefault()
			require.Equal(t, test.expected.(*Config), actual)
		})
	}
}

func TestSetField(t *testing.T) {
	t.Run("TestSetField", func(t *testing.T) {
		var buf bytes.Buffer
		logger := configDefault(Config{Logger: log.Output(&buf)})
		logger.
			SetField("bool", true).
			SetField("int", 12345).
			SetField("int64", int64(12345)).
			SetField("float64", float64(12345.01)).
			SetField("[]byte", []byte(`{"json_key": "json_value"}`)).
			SetField("[]int", []int{1, 2, 3, 4, 5}).
			SetField("[]int64", []int64{1, 2, 3, 4, 5}).
			SetField("[]float64", []float64{1.1, 2.2, 3.3, 4.4, 5.5}).
			SetField("[]string", []string{"a", "b", "c"}).
			SetField("", ""). // << expected skip this line
			SetField("str", "string").
			Debug()

		require.Contains(t, buf.String(), "bool")
		require.Contains(t, buf.String(), "true")
		require.Contains(t, buf.String(), "int")
		require.Contains(t, buf.String(), "12345")
		require.Contains(t, buf.String(), "int64")
		// require.Contains(t, buf.String(), "12345") // the same as `int`
		require.Contains(t, buf.String(), "float64")
		require.Contains(t, buf.String(), "12345.01")
		require.Contains(t, buf.String(), "[]byte")
		require.Contains(t, buf.String(), "{\"json_key\": \"json_value\"}")
		require.Contains(t, buf.String(), "[]int")
		require.Contains(t, buf.String(), "[1,2,3,4,5]")
		require.Contains(t, buf.String(), "[]int64")
		// require.Contains(t, buf.String(), "[1,2,3,4,5]") // the same as `int`
		require.Contains(t, buf.String(), "[]float64")
		require.Contains(t, buf.String(), "[1.1,2.2,3.3,4.4,5.5]")
		require.Contains(t, buf.String(), "[]string")
		require.Contains(t, buf.String(), "[\"a\",\"b\",\"c\"]")
		require.Contains(t, buf.String(), "str")
		require.Contains(t, buf.String(), "string")
	})
}

func TestZLogWrapDefault(t *testing.T) {
	tests := []testCaseDefault{
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "With text",
			input:    "testing logger",
			expected: "testing logger",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := configDefault(Config{Logger: log.Output(&buf)})

			logger.Debug(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			logger.Info(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			logger.Warn(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			logger.Error(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			// logger.Fatal(test.input.(string))                         // TODO: I'll do it later XD
			// require.Contains(t, buf.String(), test.expected.(string)) // TODO: I'll do it later XD

			panicMaker := func() {
				logger.Panic(test.input.(string))
				require.Contains(t, buf.String(), test.expected.(string))
			}
			require.NotPanics(t, panicMaker)
		})
	}

}
func TestZLogWrapTerminateOnCriticalLevel(t *testing.T) {
	t.Run("Critical Log with terminate enabled", func(t *testing.T) {
		var foo string = "bar"
		var buf bytes.Buffer
		logger := configDefault(Config{Logger: log.Output(&buf), TerminateOnCriticalLevel: true})
		panicMaker := func() {
			logger.Panic(foo)
			require.Contains(t, buf.String(), foo)
		}
		require.Panics(t, panicMaker)
	})
}

func TestZLogWrapHiddenTrue(t *testing.T) {
	tests := []testCaseDefault{
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "With text",
			input:    "testing logger",
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// actualDebugOut := captureOutput(func() {
			// 	logger := Config{Hidden: true}
			// 	logger.Debug(test.input.(string))
			// })
			// require.Contains(t, actualDebugOut, test.expected.(string))

			// actualInfoOut := captureOutput(func() {
			// 	logger := Config{Hidden: true}
			// 	logger.Info(test.input.(string))
			// })
			// require.Contains(t, actualInfoOut, test.expected.(string))

			// actualErrorOut := captureOutput(func() {
			// 	logger := Config{Hidden: true}
			// 	logger.Error(test.input.(string))
			// })
			// require.Contains(t, actualErrorOut, test.expected.(string))
		})
	}
}
