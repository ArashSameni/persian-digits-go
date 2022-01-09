package persiandigits

import (
	"testing"
)

func TestToEnglish(t *testing.T) {
	testCase := "Hey im testing ۰-۱-۲-۳-۴-۵-۶-۷-۸-۹"
	expected, expectedN := "Hey im testing 0-1-2-3-4-5-6-7-8-9", 10

	output, n := ToEnglish(testCase)
	if expected != output || n != expectedN {
		t.Errorf("Expected: %s n:%d, Output: %s n:%d", expected, expectedN, output, n)
	}
}

func TestToPersian(t *testing.T) {
	testCase := "این یکیو تست میکنیم 0-1-2-3-4-5-6-7-8-9 ببینیم چی میشه"
	expected, expectedN := "این یکیو تست میکنیم ۰-۱-۲-۳-۴-۵-۶-۷-۸-۹ ببینیم چی میشه", 10

	output, n := ToPersian(testCase)
	if expected != output || n != expectedN {
		t.Errorf("Expected: %s n:%d, Output: %s n:%d", expected, expectedN, output, n)
	}
}
