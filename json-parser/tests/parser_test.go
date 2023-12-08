package jsonparsertest

import (
	"fmt"
	"os"
	"testing"

	jsonparser "../../../go-exercises/json-parser"
)

func setupTest(fileName string) (string, func()) {
	data, err := os.ReadFile(fmt.Sprintf("%v", fileName))
	if err != nil {
		panic(err)
	}

	testData := string(data)
	return testData, func() {
		// tear-down code here
	}
}

type TestParam struct {
	fileName string
	expected bool
}

func TestJSONParser(t *testing.T) {
	tests := []TestParam{
		TestParam{"step1/valid.json", true},
		TestParam{"step1/invalid.json", false},
		TestParam{"step2/invalid.json", false},
		TestParam{"step2/invalid2.json", false},
		TestParam{"step2/valid.json", true},
		TestParam{"step2/valid2.json", true},
		TestParam{"step3/invalid.json", false},
		TestParam{"step3/valid.json", true},
		TestParam{"step4/invalid.json", false},
		TestParam{"step4/valid.json", true},
		TestParam{"step4/valid2.json", true},
	}

	for _, testParam := range tests {
		t.Run("Test 1 valid", func(t *testing.T) {
			data, tearDown := setupTest(testParam.fileName)
			defer tearDown()
			result := jsonparser.IsValidValue(data)
			expected := testParam.expected
			t.Log(result, expected)
			if result != expected {
				t.Errorf("Result: Expected JSON parsing to be True but got %v ", result)
			}
		})
	}
}
