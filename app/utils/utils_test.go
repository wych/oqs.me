package utils

import (
	"testing"
)

func TestBase10to62(t *testing.T) {
	var convTest = []struct {
		in       uint
		expected Base62Int
	}{
		{6786990, "STba"},
		{8976600647, "9nUttP"},
		{66580764, "4VMhI"},
		{0, "0"},
		{9448596079, "AJRLSx"},
	}
	for _, ct := range convTest {
		base62 := Base10to62(ct.in)
		if base62 != ct.expected {
			t.Errorf("Base10to62(%d) = %s, excepted %s", ct.in, base62, ct.expected)
		}
	}
}

func TestBase64to10(t *testing.T) {
	var convTest = []struct {
		in       Base62Int
		expected uint
	}{
		{"STba", 6786990},
		{"9nUttP", 8976600647},
		{"4VMhI", 66580764},
		{"0", 0},
		{"AJRLSx", 9448596079},
	}
	for _, ct := range convTest {
		base62, err := Base62to10(ct.in)
		if err != nil {
			t.Errorf("%s", err)
		}
		if base62 != ct.expected {
			t.Errorf("Base10to62(%s) = %d, excepted %d", ct.in, base62, ct.expected)
		}
	}
}

func TestIndexOfbase62Map(t *testing.T) {
	var inArr = []int{1, 3, 6, 54, 15, 0, 61}
	for _, v := range inArr {
		char := base62Map[v]
		index := indexOfbase62Map(char)
		if int(index) != v {
			t.Errorf("excepted: %d, out: %d", v, index)
		}
	}

}
