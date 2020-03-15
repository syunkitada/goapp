package str_utils

import "testing"

func TestGenerateRandStr(t *testing.T) {
	if got := GenerateRandStr(3); len(got) != 3 {
		t.Errorf("len(got)=%d, want=%d", len(got), 3)
	}

	if got := GenerateRandStr(5); len(got) != 5 {
		t.Errorf("len(got)=%d, want=%d", len(got), 5)
	}
}
