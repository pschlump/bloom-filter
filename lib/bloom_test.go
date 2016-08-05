//
// Copyright (C) Philip Schlump, 2014-2016
// MIT Licnesed.
//
package bloom

import (
	"fmt"
	"testing"
)

func Test_HashFunctions(t *testing.T) {
	tests := []struct {
		Str         string
		BloomFilter uint32
		Murmur      uint32
	}{
		{Str: "", BloomFilter: 0, Murmur: 1588974330},                 // 0
		{Str: "7", BloomFilter: 1172465224, Murmur: 2694772847},       // 1
		{Str: "75", BloomFilter: 2471797843, Murmur: 2430648749},      // 2
		{Str: "758", BloomFilter: 776002016, Murmur: 3649837500},      // 3
		{Str: "7587", BloomFilter: 36914527, Murmur: 2062256002},      // 4
		{Str: "75872", BloomFilter: 2488709022, Murmur: 2086912925},   // 5
		{Str: "75872a", BloomFilter: 4233351963, Murmur: 2084237965},  // 6
		{Str: "7b5872a", BloomFilter: 2661868145, Murmur: 2049054836}, // 7
	}

	seed := uint32(552211)

	for ii, test := range tests {
		r := SuperFastHash([]byte(test.Str))
		if r != test.BloomFilter {
			t.Errorf("Test %d: Expected %d got %d\n", ii, test.BloomFilter, r)
		}

		s := Murmur([]byte(test.Str), seed)
		if s != test.Murmur {
			t.Errorf("Test %d: Expected %d got %d\n", ii, test.Murmur, s)
		}

		if s == r {
			t.Errorf("Test %d: Expected s not to euqal r,  s=%d r=%d\n", ii, s, r)
		}
	}
}

// func BloomFilter(str string, filterData []byte) (likelyToHaveIt bool, n1, n2 uint32) {
// func AddToBloomFilter(str string, filterData []byte) {

func Test_BloomFilter_01(t *testing.T) {
	tests := []struct {
		Str   string
		Found bool
	}{
		{Str: "", Found: false},
		{Str: "abc", Found: false},
		{Str: "tahiti", Found: true},
		{Str: "Mookie", Found: true},
		{Str: "lala", Found: false},
	}

	bf := NewBloomFilter(5)

	for ii, test := range tests {

		found, _, _ := bf.Found(test.Str)
		if found != test.Found {
			t.Errorf("Test %d: For [%s] Expected %v got %v\n", ii, test.Str, test.Found, found)
		}

		bf.AddTo(test.Str)

	}
}

func Test_BloomFilter_02(t *testing.T) {
	tests := []struct {
		Str   string
		Found bool
	}{
		{Str: "", Found: false},
		{Str: "abc", Found: false},
		{Str: "tahiti", Found: false},
		{Str: "Mookie", Found: false},
		{Str: "lala", Found: false},
		{Str: "lala", Found: true},
		{Str: "abc", Found: true},
	}

	bf := NewBloomFilter(64)

	for ii, test := range tests {

		found, _, _ := bf.Found(test.Str)
		if found != test.Found {
			t.Errorf("Test %d: For [%s] Expected %v got %v\n", ii, test.Str, test.Found, found)
		}

		bf.AddTo(test.Str)

	}
}

func Test_BloomFilter_03(t *testing.T) {
	tests := []struct {
		Str   string
		Found bool
	}{
		{Str: "", Found: false},
		{Str: "abc", Found: false},
		{Str: "tahiti", Found: false},
		{Str: "Mookie", Found: false},
		{Str: "lala", Found: false},
		{Str: "lala", Found: true},
		{Str: "abc", Found: true},
	}

	bf := NewBloomFilter(101)

	for ii, test := range tests {
		found := bf.TestAndSet(test.Str)
		if db1 {
			fmt.Printf("Filter = %s\n", bf)
		}
		if found != test.Found {
			t.Errorf("Test %d: For [%s] Expected %v got %v\n", ii, test.Str, test.Found, found)
		}
	}
	if db1 {
		fmt.Printf("Filter = %s\n", bf)
	}
}

const db1 = false
