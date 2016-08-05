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
		{Str: "75872", BloomFilter: 2488709022}, //
	}

	for ii, test := range tests {
		r := SuperFastHash([]byte(test.Str))
		if r != test.BloomFilter {
			t.Errorf("Test %d: Expected %d got %d\n", ii, test.BloomFilter, r)
		}
	}
}
