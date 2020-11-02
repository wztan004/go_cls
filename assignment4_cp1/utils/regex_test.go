package utils

import "testing"

func TestVerifyIC(t *testing.T) {
	s1 := []string{"s1234567a", "S1234567A"}
	for _, v := range s1 {
		res, _ := VerifyIC(v)
		if res != true {
			t.Errorf("VerifyIC function. Expecting %v, got %v", !res, res)
		}
	}

	s2 := []string{"1234567A", "1234567"}
	for _, v := range s2 {
		res, _ := VerifyIC(v)
		if res != false {
			t.Errorf("VerifyIC function. Expecting %v, got %v", !res, res)
		}
	}
}