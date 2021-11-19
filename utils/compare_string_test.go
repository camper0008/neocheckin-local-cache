package utils

import (
	"testing"
)

func negative(a int) bool { return a < 0 }
func positive(a int) bool { return a > 0 }
func neutral(a int) bool  { return a == 0 }

func CompareString_shouldBeNegative(t *testing.T) {
	output := CompareString("aaa", "bbb")
	if !negative(output) {
		t.Errorf("should be negative if first string is bigger")
	}
}

func CompareString_shouldBePositive(t *testing.T) {
	output := CompareString("bbb", "aaa")
	if !positive(output) {
		t.Errorf("should be positive if last string is bigger")
	}
}

func CompareString_shouldPreferUpperCase(t *testing.T) {
	output := CompareString("aaa", "AAA")
	if !positive(output) {
		t.Errorf("should prefer upper case")
	}
}

func CompareString_shouldPreferShortest(t *testing.T) {
	output := CompareString("manem", "maneme")
	if !negative(output) {
		t.Errorf("should prefer shortest")
	}
}

func CompareString_shouldBeNeutral(t *testing.T) {
	output := CompareString("blyat", "blyat")
	if !neutral(output) {
		t.Errorf("should be neutral")
	}
}

func CompareString_shouldPreferEmptyString(t *testing.T) {
	output := CompareString("", "blyat")
	if !negative(output) {
		t.Errorf("should prefer empty string")
	}
}

func TestCompareString(t *testing.T) {
	CompareString_shouldBeNegative(t)
	CompareString_shouldBePositive(t)
	CompareString_shouldPreferUpperCase(t)
	CompareString_shouldPreferShortest(t)
	CompareString_shouldBeNeutral(t)
	CompareString_shouldPreferEmptyString(t)
}
