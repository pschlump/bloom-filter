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
