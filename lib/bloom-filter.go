//
// Copyright (C) Philip Schlump, 2014-2016
// MIT Licnesed.
//
package bloom

/*
uint32_t SuperFastHash (const char * data, int len) {
  uint32_t hash = len, tmp;
  int rem;

  if (len <= 0 || data == NULL) return 0;
  rem = len & 3;
  len >>= 2;

  // Main loop
  for (;len > 0; len--) {
    hash += get16bits (data);
    tmp = (get16bits (data + 2) << 11) ^ hash;
    hash = (hash << 16) ^ tmp;
    data += 2 * sizeof (uint16_t);
    hash += hash >> 11;
  }

  // Handle end cases
  switch (rem) {
    case 3: hash += get16bits (data);
      hash ^= hash << 16;
      hash ^= ((signed char)data[sizeof (uint16_t)]) << 18;
      hash += hash >> 11;
      break;
    case 2: hash += get16bits (data);
      hash ^= hash << 11;
      hash += hash >> 17;
      break;
    case 1: hash += (signed char)*data;
      hash ^= hash << 10;
      hash += hash >> 1;
  }

  // Force "avalanching" of final 127 bits
  hash ^= hash << 3;
  hash += hash >> 5;
  hash ^= hash << 4;
  hash += hash >> 17;
  hash ^= hash << 25;
  hash += hash >> 6;

  return hash;
}
*/

func SuperFastHash(data []byte) (hash uint32) {
	llen := len(data)
	if llen <= 0 {
		return 0
	}
	hash = uint32(llen)
	rem := llen & 3
	llen >>= 2

	// Main loop
	for ; llen > 0; llen-- {
		d16 := uint32(data[0]) | uint32(data[1]<<8)
		data = data[2:]
		hash += d16
		d16 = uint32(data[0]) | uint32(data[1]<<8)
		data = data[2:]
		tmp := (d16 << 11) ^ hash
		hash = (hash << 16) ^ tmp
		hash += hash >> 11
	}

	// Handle end cases
	switch rem {
	case 3:
		d16 := uint32(data[0]) | uint32(data[1]<<8)
		data = data[2:]
		hash += d16
		hash ^= hash << 16
		d16 = uint32(data[0])
		hash ^= d16 << 18
		hash += hash >> 11
	case 2:
		d16 := uint32(data[0]) | uint32(data[1]<<8)
		hash += d16
		hash ^= hash << 11
		hash += hash >> 17
	case 1:
		d16 := uint32(data[0])
		hash += d16
		hash ^= hash << 10
		hash += hash >> 1
	}

	// Force "avalanching" of final 127 bits
	hash ^= hash << 3
	hash += hash >> 5
	hash ^= hash << 4
	hash += hash >> 17
	hash ^= hash << 25
	hash += hash >> 6

	return
}

/*
//via https://gist.github.com/588423
//thanks github.com/raycmorgan!
function murmur(str, seed) {
  var m = 0x5bd1e995;
  var r = 24;
  var h = seed ^ str.length;
  var length = str.length;
  var currentIndex = 0;

  while (length >= 4) {
    var k = UInt32(str, currentIndex);

    k = Umul32(k, m);
    k ^= k >>> r;
    k = Umul32(k, m);

    h = Umul32(h, m);
    h ^= k;

    currentIndex += 4;
    length -= 4;
  }

  switch (length) {
  case 3:
    h ^= UInt16(str, currentIndex);
    h ^= str.charCodeAt(currentIndex + 2) << 16;
    h = Umul32(h, m);
    break;

  case 2:
    h ^= UInt16(str, currentIndex);
    h = Umul32(h, m);
    break;

  case 1:
    h ^= str.charCodeAt(currentIndex);
    h = Umul32(h, m);
    break;
  }

  h ^= h >>> 13;
  h = Umul32(h, m);
  h ^= h >>> 15;

  return h >>> 0;
}
*/

func Murmur(data []byte, seed uint32) (hash uint32) {
	mm := uint32(0x5bd1e995)
	length := len(data)
	hash = seed ^ uint32(length)

	rem := length & 3
	length >>= 2

	// Main loop
	for ; length > 0; length-- {
		kk := uint32(data[0]) | uint32(data[1]<<8) | uint32(data[2]<<16) | uint32(data[3]<<24)
		data = data[4:]

		kk = kk * mm
		kk ^= kk >> 24
		kk = kk * mm

		hash = hash * mm
		hash ^= kk
	}

	switch rem {
	case 3:
		d16 := uint32(data[0]) | uint32(data[1]<<8)
		data = data[2:]
		hash ^= d16
		d16 = uint32(data[0])
		hash ^= d16 << 16
		hash = hash * mm

	case 2:
		d16 := uint32(data[0]) | uint32(data[1]<<8)
		hash ^= d16
		hash = hash * mm

	case 1:
		d16 := uint32(data[0])
		hash ^= d16
		hash = hash * mm
	}

	hash ^= hash >> 13
	hash = hash * mm
	hash ^= hash >> 15

	return hash
}

func BloomFilter(str string, filterData []byte) (likelyToHaveIt bool, n1, n2 uint32) {
	modSize := uint32(len(filterData))
	likelyToHaveIt = true

	n1 = Murmur([]byte(str), uint32(552211)) % modSize
	n2 = SuperFastHash([]byte(str)) % modSize

	if filterData[n1] == 'n' || filterData[n2] == 'n' {
		likelyToHaveIt = false
	}

	return
}

func AddToBloomFilter(str string, filterData []byte) {
	_, n1, n2 := BloomFilter(str, filterData)
	filterData[n1] = 'y'
	filterData[n2] = 'y'
}
