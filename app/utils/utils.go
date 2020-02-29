package utils

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

var base62Map []byte = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

const base62MapLastIndex = 61

const base62 = 62

type Base62Int string

func Base10to62(b10n uint) Base62Int {
	var i uint8 = 11
	buffer := make([]byte, i)
	for {
		i--
		buffer[i] = base62Map[b10n%base62]
		b10n = b10n / base62
		if b10n == 0 {
			break
		}
	}
	return Base62Int(buffer[i:])
}

func Base62to10(b62n Base62Int) (uint, error) {
	var rt uint = 0
	length := len(b62n)
	b62nBytes := []byte(b62n)
	var value int8
	for i, v := range b62nBytes {
		value = indexOfbase62Map(v)
		if value == -1 {
			err := errors.New("not a base 62 number")
			return 0, err
		}
		pos := length - (i + 1)
		rt = rt + uint(value)*uint(math.Pow(base62, float64(pos)))
	}
	return rt, nil
}

func indexOfbase62Map(e byte) int8 {
	var start int8 = 0
	var last int8 = base62MapLastIndex
	var middle int8
	for start < last {
		middle = start + ((last - start) >> 1)
		if base62Map[middle] < e {
			start = middle + 1
		} else {
			last = middle
		}
	}
	if base62Map[last] == e {
		return last
	}
	return -1
}

func GenChallenge() Base62Int {
	rand.Seed(time.Now().UnixNano())
	cInt := rand.Intn(62 * 62)
	return Base10to62(uint(cInt))
}

func (b62 Base62Int) BiggerThan(b62n Base62Int) bool {
	lLength := len(b62)
	rLength := len(b62n)
	if lLength < rLength {
		return false
	}
	if lLength > rLength {
		return true
	}
	// In situation both of base62 number are in same length
	minLength := int(math.Min(float64(lLength), float64(rLength)))
	for i := 0; i < minLength; i++ {
		if uint8(b62[i]) < uint8(b62n[i]) {
			return false
		}
	}
	return true
}
