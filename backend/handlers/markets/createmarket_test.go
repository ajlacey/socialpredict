package marketshandlers

import (
	"strings"
	"testing"
)

// TestCheckQuestionTitleLength_invalid tests the question titles that should generate an error
func TestCheckQuestionTitleLength_invalid(t *testing.T) {
	tests := []struct {
		testname      string
		questionTitle string
	}{
		{
			testname:      "TitleExceedsLength",
			questionTitle: strings.Repeat("a", maxQuestionTitleLength+1),
		},
		{
			testname:      "EmptyTitle",
			questionTitle: "",
		},
	}
	for _, test := range tests {
		err := checkQuestionTitleLength(test.questionTitle)
		if err == nil {
			t.Errorf("Expected error in test %s", test.testname)
		}
	}
}

// TODO: Add test description
func TestCheckQuestionTitleLength_valid(t *testing.T) {
	tests := []struct {
		testname      string
		questionTitle string
	}{
		{
			testname:      "Single character title",
			questionTitle: "a",
		},
		{
			testname:      "Max length title",
			questionTitle: strings.Repeat("a", maxQuestionTitleLength),
		},
	}
	for _, test := range tests {
		err := checkQuestionTitleLength(test.questionTitle)
		if err != nil {
			t.Errorf("Unexpected error in test %s", test.testname)
		}
	}
}

// TestCheckQuestionDescriptionLength_invalid ensures an error is returned if
// the question exceeds the maxDescriptionLength
func TestCheckQuestionDescriptionLength_invalid(t *testing.T) {
	test := struct {
		testname    string
		description string
	}{
		testname:    "One character too many",
		description: strings.Repeat("a", maxDescriptionLength+1),
	}

	if err := checkQuestionDescriptionLength(test.description); err == nil {
		t.Errorf("Expected error in test %s, got nil", test.testname)
	}
}

// TestCheckQuestionDescriptionLength_valid checks the upper and lower bounds
// of the permissible description length, ensuring no errors are returned.
func TestCheckQuestionDescriptionLength_valid(t *testing.T) {
	tests := []struct {
		testname    string
		description string
	}{
		{
			testname:    "Maximum length",
			description: strings.Repeat("a", maxDescriptionLength),
		},
		{
			testname:    "Empty description",
			description: "",
		},
	}
	for _, test := range tests {
		if err := checkQuestionDescriptionLength(test.description); err != nil {
			t.Errorf("Unexpected error in test %s, got nil", test.testname)
		}
	}

}
