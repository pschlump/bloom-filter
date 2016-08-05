//
// Copyright (C) Philip Schlump, 2014-2016
// MIT Licnesed.
//
package bloom

import "testing"

func Test_SuperFashHash(t *testing.T) {
	tests := []struct {
		Str         string
		BloomFilter uint32
	}{
		{Str: "", BloomFilter: 0},                 // 0
		{Str: "7", BloomFilter: 1172465224},       // 1
		{Str: "75", BloomFilter: 2471797843},      // 2
		{Str: "758", BloomFilter: 776002016},      // 3
		{Str: "7587", BloomFilter: 36914527},      // 4
		{Str: "75872", BloomFilter: 2488709022},   // 5
		{Str: "75872a", BloomFilter: 4233351963},  // 6
		{Str: "7b5872a", BloomFilter: 2661868145}, // 7
	}

	for ii, test := range tests {
		r := SuperFastHash([]byte(test.Str))
		if r != test.BloomFilter {
			t.Errorf("Test %d: Expected %d got %d\n", ii, test.BloomFilter, r)
		}
	}
}
