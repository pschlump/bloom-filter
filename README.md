bloom-filter implementation in Go (golang)
==========================================

Bloom filter implementation in Go (golang)

Example:

```Go

	bloomFiler := NewBloomFilter(101)
	for _, s := range []string { "abc", "def", "ghi", "abc" } {
	likelyToHaveIt := bloomFilter.TestAndSet ( s )
		fmt.Printf ( "Likely to have seen %s before = %v\n", s, likelyToHaveIt )
	}

```

Functions:

`NewBloomFilter` returns a bloom filter with a lookup table size of `size`.  The table is stored efficiently in bits. In my opinion it is best to use a prime number for the size.

`String` prints out the bloom filter in a human readable format, sort of.

`Found` looks up the specified value, `str`, and returns true if it is likely to be found, false if not found.  Also the n1, n2 hash values are returned.

`Add` to marks the `str` value as seen in the bloom filter.

`TestAndSet` returns true if the `str` is like to have been seen, false otherwise.  It sets this value as having been seen.

Hash Functions:

Two hash functions are implemented:

```
	func Murmur(data []byte, seed uint32) (hash uint32) {
```

and

```
	func SuperFastHash(data []byte) (hash uint32) {
```


Alternatives
------------

You might want to look at: https://github.com/irfansharif/cfilter
These filters are usually better than Bloom filters.

